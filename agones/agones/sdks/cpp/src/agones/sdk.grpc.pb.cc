// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: sdk.proto

#include "sdk.pb.h"
#include "sdk.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace agones {
namespace dev {
namespace sdk {

static const char* SDK_method_names[] = {
  "/agones.dev.sdk.SDK/Ready",
  "/agones.dev.sdk.SDK/Allocate",
  "/agones.dev.sdk.SDK/Shutdown",
  "/agones.dev.sdk.SDK/Health",
  "/agones.dev.sdk.SDK/GetGameServer",
  "/agones.dev.sdk.SDK/WatchGameServer",
  "/agones.dev.sdk.SDK/SetLabel",
  "/agones.dev.sdk.SDK/SetAnnotation",
  "/agones.dev.sdk.SDK/Reserve",
};

std::unique_ptr< SDK::Stub> SDK::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< SDK::Stub> stub(new SDK::Stub(channel));
  return stub;
}

SDK::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_Ready_(SDK_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Allocate_(SDK_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Shutdown_(SDK_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Health_(SDK_method_names[3], ::grpc::internal::RpcMethod::CLIENT_STREAMING, channel)
  , rpcmethod_GetGameServer_(SDK_method_names[4], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_WatchGameServer_(SDK_method_names[5], ::grpc::internal::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_SetLabel_(SDK_method_names[6], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SetAnnotation_(SDK_method_names[7], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Reserve_(SDK_method_names[8], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status SDK::Stub::Ready(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Ready_, context, request, response);
}

void SDK::Stub::experimental_async::Ready(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Ready_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::Ready(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Ready_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncReadyRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Ready_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncReadyRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Ready_, context, request, false);
}

::grpc::Status SDK::Stub::Allocate(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Allocate_, context, request, response);
}

void SDK::Stub::experimental_async::Allocate(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Allocate_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::Allocate(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Allocate_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncAllocateRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Allocate_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncAllocateRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Allocate_, context, request, false);
}

::grpc::Status SDK::Stub::Shutdown(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Shutdown_, context, request, response);
}

void SDK::Stub::experimental_async::Shutdown(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Shutdown_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::Shutdown(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Shutdown_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncShutdownRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Shutdown_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncShutdownRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Shutdown_, context, request, false);
}

::grpc::ClientWriter< ::agones::dev::sdk::Empty>* SDK::Stub::HealthRaw(::grpc::ClientContext* context, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::ClientWriterFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), rpcmethod_Health_, context, response);
}

void SDK::Stub::experimental_async::Health(::grpc::ClientContext* context, ::agones::dev::sdk::Empty* response, ::grpc::experimental::ClientWriteReactor< ::agones::dev::sdk::Empty>* reactor) {
  ::grpc::internal::ClientCallbackWriterFactory< ::agones::dev::sdk::Empty>::Create(stub_->channel_.get(), stub_->rpcmethod_Health_, context, response, reactor);
}

::grpc::ClientAsyncWriter< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncHealthRaw(::grpc::ClientContext* context, ::agones::dev::sdk::Empty* response, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncWriterFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Health_, context, response, true, tag);
}

::grpc::ClientAsyncWriter< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncHealthRaw(::grpc::ClientContext* context, ::agones::dev::sdk::Empty* response, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncWriterFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Health_, context, response, false, nullptr);
}

::grpc::Status SDK::Stub::GetGameServer(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::agones::dev::sdk::GameServer* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetGameServer_, context, request, response);
}

void SDK::Stub::experimental_async::GetGameServer(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::GameServer* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetGameServer_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::GetGameServer(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::GameServer* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetGameServer_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::GameServer>* SDK::Stub::AsyncGetGameServerRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::GameServer>::Create(channel_.get(), cq, rpcmethod_GetGameServer_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::GameServer>* SDK::Stub::PrepareAsyncGetGameServerRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::GameServer>::Create(channel_.get(), cq, rpcmethod_GetGameServer_, context, request, false);
}

::grpc::ClientReader< ::agones::dev::sdk::GameServer>* SDK::Stub::WatchGameServerRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request) {
  return ::grpc::internal::ClientReaderFactory< ::agones::dev::sdk::GameServer>::Create(channel_.get(), rpcmethod_WatchGameServer_, context, request);
}

void SDK::Stub::experimental_async::WatchGameServer(::grpc::ClientContext* context, ::agones::dev::sdk::Empty* request, ::grpc::experimental::ClientReadReactor< ::agones::dev::sdk::GameServer>* reactor) {
  ::grpc::internal::ClientCallbackReaderFactory< ::agones::dev::sdk::GameServer>::Create(stub_->channel_.get(), stub_->rpcmethod_WatchGameServer_, context, request, reactor);
}

::grpc::ClientAsyncReader< ::agones::dev::sdk::GameServer>* SDK::Stub::AsyncWatchGameServerRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::agones::dev::sdk::GameServer>::Create(channel_.get(), cq, rpcmethod_WatchGameServer_, context, request, true, tag);
}

::grpc::ClientAsyncReader< ::agones::dev::sdk::GameServer>* SDK::Stub::PrepareAsyncWatchGameServerRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Empty& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::agones::dev::sdk::GameServer>::Create(channel_.get(), cq, rpcmethod_WatchGameServer_, context, request, false, nullptr);
}

::grpc::Status SDK::Stub::SetLabel(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue& request, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SetLabel_, context, request, response);
}

void SDK::Stub::experimental_async::SetLabel(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_SetLabel_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::SetLabel(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_SetLabel_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncSetLabelRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_SetLabel_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncSetLabelRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_SetLabel_, context, request, false);
}

::grpc::Status SDK::Stub::SetAnnotation(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue& request, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SetAnnotation_, context, request, response);
}

void SDK::Stub::experimental_async::SetAnnotation(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_SetAnnotation_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::SetAnnotation(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_SetAnnotation_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncSetAnnotationRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_SetAnnotation_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncSetAnnotationRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::KeyValue& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_SetAnnotation_, context, request, false);
}

::grpc::Status SDK::Stub::Reserve(::grpc::ClientContext* context, const ::agones::dev::sdk::Duration& request, ::agones::dev::sdk::Empty* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Reserve_, context, request, response);
}

void SDK::Stub::experimental_async::Reserve(::grpc::ClientContext* context, const ::agones::dev::sdk::Duration* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Reserve_, context, request, response, std::move(f));
}

void SDK::Stub::experimental_async::Reserve(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::agones::dev::sdk::Empty* response, std::function<void(::grpc::Status)> f) {
  return ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Reserve_, context, request, response, std::move(f));
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::AsyncReserveRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Duration& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Reserve_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::agones::dev::sdk::Empty>* SDK::Stub::PrepareAsyncReserveRaw(::grpc::ClientContext* context, const ::agones::dev::sdk::Duration& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::agones::dev::sdk::Empty>::Create(channel_.get(), cq, rpcmethod_Reserve_, context, request, false);
}

SDK::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::Empty, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::Ready), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::Empty, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::Allocate), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::Empty, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::Shutdown), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[3],
      ::grpc::internal::RpcMethod::CLIENT_STREAMING,
      new ::grpc::internal::ClientStreamingHandler< SDK::Service, ::agones::dev::sdk::Empty, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::Health), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::Empty, ::agones::dev::sdk::GameServer>(
          std::mem_fn(&SDK::Service::GetGameServer), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[5],
      ::grpc::internal::RpcMethod::SERVER_STREAMING,
      new ::grpc::internal::ServerStreamingHandler< SDK::Service, ::agones::dev::sdk::Empty, ::agones::dev::sdk::GameServer>(
          std::mem_fn(&SDK::Service::WatchGameServer), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[6],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::KeyValue, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::SetLabel), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[7],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::KeyValue, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::SetAnnotation), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SDK_method_names[8],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SDK::Service, ::agones::dev::sdk::Duration, ::agones::dev::sdk::Empty>(
          std::mem_fn(&SDK::Service::Reserve), this)));
}

SDK::Service::~Service() {
}

::grpc::Status SDK::Service::Ready(::grpc::ServerContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::Allocate(::grpc::ServerContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::Shutdown(::grpc::ServerContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::Health(::grpc::ServerContext* context, ::grpc::ServerReader< ::agones::dev::sdk::Empty>* reader, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) reader;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::GetGameServer(::grpc::ServerContext* context, const ::agones::dev::sdk::Empty* request, ::agones::dev::sdk::GameServer* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::WatchGameServer(::grpc::ServerContext* context, const ::agones::dev::sdk::Empty* request, ::grpc::ServerWriter< ::agones::dev::sdk::GameServer>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::SetLabel(::grpc::ServerContext* context, const ::agones::dev::sdk::KeyValue* request, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::SetAnnotation(::grpc::ServerContext* context, const ::agones::dev::sdk::KeyValue* request, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SDK::Service::Reserve(::grpc::ServerContext* context, const ::agones::dev::sdk::Duration* request, ::agones::dev::sdk::Empty* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace agones
}  // namespace dev
}  // namespace sdk
