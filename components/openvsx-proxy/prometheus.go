// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package main

import (
	"net/http"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cacheHitBackupCounter  prometheus.Counter
	cacheMissBackupCounter prometheus.Counter
	statusCounter          *prometheus.CounterVec
)

func startPromotheus(cfg *Config) {
	reg := prometheus.NewRegistry()

	if cfg.PrometheusAddr != "" {
		reg.MustRegister(
			prometheus.NewGoCollector(),
			prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		)

		handler := http.NewServeMux()
		handler.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

		go func() {
			err := http.ListenAndServe(cfg.PrometheusAddr, handler)
			if err != nil {
				log.WithError(err).Error("Prometheus metrics server failed")
			}
		}()
		log.WithField("addr", cfg.PrometheusAddr).Info("started Prometheus metrics server")
	}

	createMetrics()
	collectors := []prometheus.Collector{
		cacheHitBackupCounter,
		cacheMissBackupCounter,
		statusCounter,
	}
	for _, c := range collectors {
		err := reg.Register(c)
		if err != nil {
			log.WithError(err).Error("register Prometheus metric failed")
		}
	}
}

func createMetrics() {
	namespace := "gitpod"
	subsystem := "openvsx_proxy"
	cacheHitBackupCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "cache_hit_backup_total",
		Help:      "The total amount of hits of the backup cache.",
	})
	cacheMissBackupCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "cache_miss_backup_total",
		Help:      "The total amount of misses of the backup cache.",
	})
	statusCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "status_total",
		Help:      "The total amount of responses by status.",
	}, []string{"code"})
}
