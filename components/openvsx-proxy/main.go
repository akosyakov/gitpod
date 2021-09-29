// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/allegro/bigcache"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	REQUEST_CACHE_KEY_HEADER = "X-Gitpod-Cache-Key"
	REQUEST_ID_HEADER        = "X-Gitpod-Request-ID"
	REQUEST_TIMESTAMP_HEADER = "X-Gitpod-Timestamp"
	REQID                    = "request_id"
)

var (
	cfg          *Config
	cacheManager *cache.Cache
)

func main() {
	if len(os.Args) != 2 {
		log.Panicf("Usage: %s </path/to/config.json>", os.Args[0])
	}

	var err error
	cfg, err = ReadConfig(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	log.WithField("config", string(cfg.ToJson())).Infof("starting OpenVSX proxy")

	startPromotheus(cfg)

	err = SetupCache()
	if err != nil {
		log.Panic(err)
	}

	upstream, err := url.Parse(cfg.URLUpstream)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(rw http.ResponseWriter, r *http.Request) {
			timestampStart := time.Now()
			r.Header.Set(REQUEST_TIMESTAMP_HEADER, strconv.FormatInt(timestampStart.UnixMilli(), 10))

			var (
				hitCacheRegular = false
				hitCacheBackup  = false
			)

			reqid := ""
			uuid, err := uuid.NewRandom()
			if err != nil {
				log.WithError(err).Warn("cannot generate a UUID")
				reqid = fmt.Sprintf("req%d", rand.Intn(999999))
			} else {
				reqid = uuid.String()
			}
			log.WithField(REQID, reqid).Infof("request %s %s", r.Method, r.URL)
			r.Header.Set(REQUEST_ID_HEADER, reqid)

			key, err := key(r)
			if err != nil {
				log.WithField(REQID, reqid).WithError(err).Error("cannot create cache key")
				r.Host = upstream.Host
				p.ServeHTTP(rw, r)
				finishLog(reqid, r.Method, r.URL, timestampStart, hitCacheRegular, hitCacheBackup)
				return
			}
			r.Header.Set(REQUEST_CACHE_KEY_HEADER, key)

			if cfg.CacheDurationRegular > 0 {
				cached, ok, err := ReadCache(key)
				if err != nil {
					log.WithField(REQID, reqid).WithError(err).Errorf("cannot read '%s' from cache", key)
					cacheMissBackupCounter.Inc()
				} else if !ok {
					log.WithField(REQID, reqid).Infof("cache has no entry for '%s'", key)
					cacheMissBackupCounter.Inc()
				} else {
					hitCacheBackup = true
					cacheHitBackupCounter.Inc()
					dateHeader := cached.Header.Get("Date")
					log.WithField(REQID, reqid).Infof("there is a cached value with date: %s", dateHeader)
					t, err := time.Parse("Mon, _2 Jan 2006 15:04:05 MST", dateHeader)
					if err != nil {
						log.WithField(REQID, reqid).WithError(err).Error("cannot parse date header of cached value")
					} else {
						maxOld := time.Now().Add(-time.Duration(cfg.CacheDurationRegular))
						if t.After(maxOld) {
							hitCacheRegular = true
							log.WithField(REQID, reqid).Infof("cached value is younger than %s - using cached value", cfg.CacheDurationRegular)
							for k, v := range cached.Header {
								for i, val := range v {
									if i == 0 {
										rw.Header().Set(k, val)
									} else {
										rw.Header().Add(k, val)
									}
								}
							}
							rw.Header().Set("X-Cache", "HIT")
							rw.WriteHeader(cached.StatusCode)
							rw.Write(cached.Body)
							finishLog(reqid, r.Method, r.URL, timestampStart, hitCacheRegular, hitCacheBackup)
							return
						} else {
							log.WithField(REQID, reqid).Infof("cached value is older than %s - ignoring cached value", cfg.CacheDurationRegular)
						}
					}
				}
			}

			r.Host = upstream.Host
			p.ServeHTTP(rw, r)
			finishLog(reqid, r.Method, r.URL, timestampStart, hitCacheRegular, hitCacheBackup)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(upstream)

	proxy.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, e error) {
		reqid := r.Header.Get(REQUEST_ID_HEADER)

		log.WithField(REQID, reqid).Warnf("proxy error for request %s: %v", r.URL, e)

		key := r.Header.Get(REQUEST_CACHE_KEY_HEADER)
		if key == "" {
			log.WithField(REQID, reqid).Error("cache key header is missing")
			rw.WriteHeader(http.StatusBadGateway)
			return
		}

		cached, ok, err := ReadCache(key)
		if err != nil {
			log.WithField(REQID, reqid).WithError(err).Errorf("cannot read '%s' from cache", key)
			rw.WriteHeader(http.StatusBadGateway)
			return
		}
		if !ok {
			log.WithField(REQID, reqid).Infof("cache has no entry for '%s'", key)
			rw.WriteHeader(http.StatusBadGateway)
			return
		}
		for k, v := range cached.Header {
			for i, val := range v {
				if i == 0 {
					rw.Header().Set(k, val)
				} else {
					rw.Header().Add(k, val)
				}
			}
		}
		rw.WriteHeader(cached.StatusCode)
		rw.Write(cached.Body)
		log.WithField(REQID, reqid).Info("used cached value due to a proxy error")
	}

	proxy.ModifyResponse = func(r *http.Response) error {

		reqid := r.Request.Header.Get(REQUEST_ID_HEADER)

		log.WithField(REQID, reqid).WithField("status", r.StatusCode).Info("getting response")
		statusCounter.WithLabelValues(strconv.Itoa(r.StatusCode)).Inc()

		key := r.Request.Header.Get(REQUEST_CACHE_KEY_HEADER)
		if key == "" {
			log.WithField(REQID, reqid).Error("cache key header is missing - sending response as is")
			return nil
		}

		if r.StatusCode >= 500 {
			// use cache if exists
			log.WithField(REQID, reqid).Warnf("trying to use cache value due to a status code of %d from upstream", r.StatusCode)
			cached, ok, err := ReadCache(key)
			if err != nil {
				log.WithField(REQID, reqid).WithError(err).Errorf("cannot read '%s' from cache", key)
				return nil
			}
			if !ok {
				log.WithField(REQID, reqid).Infof("cache has no entry for '%s'", key)
				return nil
			}
			r.Header = cached.Header
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(cached.Body))
			r.ContentLength = int64(len(cached.Body))
			r.StatusCode = cached.StatusCode
			log.WithField(REQID, reqid).Infof("used cache value due to a status code of %d from upstream", r.StatusCode)
			return nil
		}

		// no error (status code < 500)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithField(REQID, reqid).WithError(err).Error("error reading response body")
			return err
		}
		r.Body.Close()

		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") {
			log.WithField(REQID, reqid).Infof("replacing %d occurence(s) of '%s' in response body ...", strings.Count(string(body), cfg.URLUpstream), cfg.URLUpstream)
			bodyStr := strings.ReplaceAll(string(body), cfg.URLUpstream, cfg.URLLocal)
			body = []byte(bodyStr)
		} else {
			log.WithField(REQID, reqid).Infof("response is not JSON but '%s', skipping replacing '%s' in response body", contentType, cfg.URLUpstream)
		}

		cacheObj := &CacheObject{
			Header:     r.Header,
			Body:       body,
			StatusCode: r.StatusCode,
		}
		err = StoreCache(key, cacheObj)
		if err != nil {
			log.WithField(REQID, reqid).Errorf("error storing object to cache: %v", err)
		}
		log.WithField(REQID, reqid).Infof("stored value for key '%s' to cache", key)

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		r.ContentLength = int64(len(body))
		r.Header.Set("Content-Length", strconv.Itoa(len(body)))
		return nil
	}

	http.DefaultTransport.(*http.Transport).MaxIdleConns = cfg.MaxIdleConns
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = cfg.MaxIdleConnsPerHost

	http.HandleFunc("/", handler(proxy))
	log.Println("listening on port 8080 ...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func finishLog(reqid, method string, url *url.URL, start time.Time, hitCacheRegular, hitCacheBackup bool) {
	duration := time.Since(start)
	log.
		WithField(REQID, reqid).
		WithField("duration", duration).
		WithField("hit_cache_regular", hitCacheRegular).
		WithField("hit_cache_backup", hitCacheBackup).
		WithField("request", fmt.Sprintf("%s %s", method, url)).
		Info("request finished")
}

func key(r *http.Request) (string, error) {
	bh, err := bodyHash(r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s %d %s", r.Method, r.URL, r.ContentLength, bh), nil
}

func bodyHash(r *http.Request) (string, error) {
	if r == nil || r.Body == nil {
		return "", xerrors.Errorf("request or body is nil")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return hash(body), nil
}

func hash(v []byte) string {
	h := sha1.New()
	h.Write(v)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

type CacheObject struct {
	Header     http.Header
	Body       []byte
	StatusCode int
}

func (cacheObj *CacheObject) ToJson() ([]byte, error) {
	b, err := json.Marshal(cacheObj)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert cache object to JSON: %v", err)
	}
	return b, nil
}

func (cacheObj *CacheObject) FromJson(v []byte) error {
	err := json.Unmarshal(v, cacheObj)
	if err != nil {
		return xerrors.Errorf("cannot convert cache object from JSON: %v", err)
	}
	return nil
}

func SetupCache() error {
	bigcacheClient, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(cfg.CacheDurationBackup)))
	if err != nil {
		return err
	}
	bigcacheStore := store.NewBigcache(bigcacheClient, nil)
	cacheManager = cache.New(bigcacheStore)
	return nil
}

func StoreCache(key string, obj *CacheObject) error {
	b, err := obj.ToJson()
	if err != nil {
		return err
	}
	return cacheManager.Set(key, b, nil)
}

func ReadCache(key string) (obj *CacheObject, ok bool, err error) {
	b, err := cacheManager.Get(key)
	if err == bigcache.ErrEntryNotFound {
		return nil, false, nil
	} else if err != nil || b == nil {
		return
	}
	obj = &CacheObject{}
	if err = obj.FromJson(b.([]byte)); err != nil {
		return
	}
	ok = true
	return
}
