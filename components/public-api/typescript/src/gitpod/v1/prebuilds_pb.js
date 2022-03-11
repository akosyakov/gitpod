// source: gitpod/v1/prebuilds.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

var google_rpc_status_pb = require('../../google/rpc/status_pb.js');
goog.object.extend(proto, google_rpc_status_pb);
goog.exportSymbol('proto.gitpod.v1.GetRunningPrebuildRequest', null, global);
goog.exportSymbol('proto.gitpod.v1.GetRunningPrebuildResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gitpod.v1.GetRunningPrebuildRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gitpod.v1.GetRunningPrebuildRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.gitpod.v1.GetRunningPrebuildRequest.displayName = 'proto.gitpod.v1.GetRunningPrebuildRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gitpod.v1.GetRunningPrebuildResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gitpod.v1.GetRunningPrebuildResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.gitpod.v1.GetRunningPrebuildResponse.displayName = 'proto.gitpod.v1.GetRunningPrebuildResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gitpod.v1.GetRunningPrebuildRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.gitpod.v1.GetRunningPrebuildRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gitpod.v1.GetRunningPrebuildRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gitpod.v1.GetRunningPrebuildRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    contextUrl: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gitpod.v1.GetRunningPrebuildRequest}
 */
proto.gitpod.v1.GetRunningPrebuildRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gitpod.v1.GetRunningPrebuildRequest;
  return proto.gitpod.v1.GetRunningPrebuildRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gitpod.v1.GetRunningPrebuildRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gitpod.v1.GetRunningPrebuildRequest}
 */
proto.gitpod.v1.GetRunningPrebuildRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setContextUrl(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gitpod.v1.GetRunningPrebuildRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gitpod.v1.GetRunningPrebuildRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gitpod.v1.GetRunningPrebuildRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gitpod.v1.GetRunningPrebuildRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContextUrl();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string context_url = 1;
 * @return {string}
 */
proto.gitpod.v1.GetRunningPrebuildRequest.prototype.getContextUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.gitpod.v1.GetRunningPrebuildRequest} returns this
 */
proto.gitpod.v1.GetRunningPrebuildRequest.prototype.setContextUrl = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.gitpod.v1.GetRunningPrebuildResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gitpod.v1.GetRunningPrebuildResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gitpod.v1.GetRunningPrebuildResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    responseStatus: (f = msg.getResponseStatus()) && google_rpc_status_pb.Status.toObject(includeInstance, f),
    prebuildId: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gitpod.v1.GetRunningPrebuildResponse}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gitpod.v1.GetRunningPrebuildResponse;
  return proto.gitpod.v1.GetRunningPrebuildResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gitpod.v1.GetRunningPrebuildResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gitpod.v1.GetRunningPrebuildResponse}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_rpc_status_pb.Status;
      reader.readMessage(value,google_rpc_status_pb.Status.deserializeBinaryFromReader);
      msg.setResponseStatus(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrebuildId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gitpod.v1.GetRunningPrebuildResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gitpod.v1.GetRunningPrebuildResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gitpod.v1.GetRunningPrebuildResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponseStatus();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_rpc_status_pb.Status.serializeBinaryToWriter
    );
  }
  f = message.getPrebuildId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional google.rpc.Status response_status = 1;
 * @return {?proto.google.rpc.Status}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.getResponseStatus = function() {
  return /** @type{?proto.google.rpc.Status} */ (
    jspb.Message.getWrapperField(this, google_rpc_status_pb.Status, 1));
};


/**
 * @param {?proto.google.rpc.Status|undefined} value
 * @return {!proto.gitpod.v1.GetRunningPrebuildResponse} returns this
*/
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.setResponseStatus = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.gitpod.v1.GetRunningPrebuildResponse} returns this
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.clearResponseStatus = function() {
  return this.setResponseStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.hasResponseStatus = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string prebuild_id = 2;
 * @return {string}
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.getPrebuildId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.gitpod.v1.GetRunningPrebuildResponse} returns this
 */
proto.gitpod.v1.GetRunningPrebuildResponse.prototype.setPrebuildId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.gitpod.v1);
