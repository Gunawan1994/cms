// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var core_v1_auth_auth_pb = require('../../../core/v1/auth/auth_pb.js');
var global_meta_meta_pb = require('../../../global/meta/meta_pb.js');

function serialize_LoginRequest(arg) {
  if (!(arg instanceof core_v1_auth_auth_pb.LoginRequest)) {
    throw new Error('Expected argument of type LoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_LoginRequest(buffer_arg) {
  return core_v1_auth_auth_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RegisterRequest(arg) {
  if (!(arg instanceof core_v1_auth_auth_pb.RegisterRequest)) {
    throw new Error('Expected argument of type RegisterRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RegisterRequest(buffer_arg) {
  return core_v1_auth_auth_pb.RegisterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UserResponse(arg) {
  if (!(arg instanceof core_v1_auth_auth_pb.UserResponse)) {
    throw new Error('Expected argument of type UserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UserResponse(buffer_arg) {
  return core_v1_auth_auth_pb.UserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthServiceService = exports.AuthServiceService = {
  // rpc VerifyToken(AuthVerifyToken) returns (UserResponse);
loginUser: {
    path: '/AuthService/LoginUser',
    requestStream: false,
    responseStream: false,
    requestType: core_v1_auth_auth_pb.LoginRequest,
    responseType: core_v1_auth_auth_pb.UserResponse,
    requestSerialize: serialize_LoginRequest,
    requestDeserialize: deserialize_LoginRequest,
    responseSerialize: serialize_UserResponse,
    responseDeserialize: deserialize_UserResponse,
  },
  registerUser: {
    path: '/AuthService/RegisterUser',
    requestStream: false,
    responseStream: false,
    requestType: core_v1_auth_auth_pb.RegisterRequest,
    responseType: core_v1_auth_auth_pb.UserResponse,
    requestSerialize: serialize_RegisterRequest,
    requestDeserialize: deserialize_RegisterRequest,
    responseSerialize: serialize_UserResponse,
    responseDeserialize: deserialize_UserResponse,
  },
};

exports.AuthServiceClient = grpc.makeGenericClientConstructor(AuthServiceService, 'AuthService');
