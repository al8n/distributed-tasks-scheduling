///
//  Generated code. Do not modify.
//  source: tasks.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'tasks.pb.dart' as $0;
export 'tasks.pb.dart';

class TasksClient extends $grpc.Client {
  static final _$saveOneTask =
      $grpc.ClientMethod<$0.SaveOneTaskRequest, $0.SaveOneTaskReply>(
          '/taskspb.Tasks/SaveOneTask',
          ($0.SaveOneTaskRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.SaveOneTaskReply.fromBuffer(value));
  static final _$deleteOneTask =
      $grpc.ClientMethod<$0.DeleteOneTaskRequest, $0.DeleteOneTaskReply>(
          '/taskspb.Tasks/DeleteOneTask',
          ($0.DeleteOneTaskRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.DeleteOneTaskReply.fromBuffer(value));
  static final _$killOneTask =
      $grpc.ClientMethod<$0.KillOneTaskRequest, $0.KillOneTaskReply>(
          '/taskspb.Tasks/KillOneTask',
          ($0.KillOneTaskRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.KillOneTaskReply.fromBuffer(value));
  static final _$getOneTask =
      $grpc.ClientMethod<$0.GetOneTaskRequest, $0.GetOneTaskReply>(
          '/taskspb.Tasks/GetOneTask',
          ($0.GetOneTaskRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetOneTaskReply.fromBuffer(value));
  static final _$getAllTasks =
      $grpc.ClientMethod<$0.GetAllTasksRequest, $0.GetAllTasksReply>(
          '/taskspb.Tasks/GetAllTasks',
          ($0.GetAllTasksRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetAllTasksReply.fromBuffer(value));

  TasksClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.SaveOneTaskReply> saveOneTask(
      $0.SaveOneTaskRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$saveOneTask, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.DeleteOneTaskReply> deleteOneTask(
      $0.DeleteOneTaskRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$deleteOneTask, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.KillOneTaskReply> killOneTask(
      $0.KillOneTaskRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$killOneTask, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.GetOneTaskReply> getOneTask(
      $0.GetOneTaskRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getOneTask, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.GetAllTasksReply> getAllTasks(
      $0.GetAllTasksRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getAllTasks, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class TasksServiceBase extends $grpc.Service {
  $core.String get $name => 'taskspb.Tasks';

  TasksServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.SaveOneTaskRequest, $0.SaveOneTaskReply>(
        'SaveOneTask',
        saveOneTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.SaveOneTaskRequest.fromBuffer(value),
        ($0.SaveOneTaskReply value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.DeleteOneTaskRequest, $0.DeleteOneTaskReply>(
            'DeleteOneTask',
            deleteOneTask_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.DeleteOneTaskRequest.fromBuffer(value),
            ($0.DeleteOneTaskReply value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.KillOneTaskRequest, $0.KillOneTaskReply>(
        'KillOneTask',
        killOneTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.KillOneTaskRequest.fromBuffer(value),
        ($0.KillOneTaskReply value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetOneTaskRequest, $0.GetOneTaskReply>(
        'GetOneTask',
        getOneTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetOneTaskRequest.fromBuffer(value),
        ($0.GetOneTaskReply value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetAllTasksRequest, $0.GetAllTasksReply>(
        'GetAllTasks',
        getAllTasks_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.GetAllTasksRequest.fromBuffer(value),
        ($0.GetAllTasksReply value) => value.writeToBuffer()));
  }

  $async.Future<$0.SaveOneTaskReply> saveOneTask_Pre($grpc.ServiceCall call,
      $async.Future<$0.SaveOneTaskRequest> request) async {
    return saveOneTask(call, await request);
  }

  $async.Future<$0.DeleteOneTaskReply> deleteOneTask_Pre($grpc.ServiceCall call,
      $async.Future<$0.DeleteOneTaskRequest> request) async {
    return deleteOneTask(call, await request);
  }

  $async.Future<$0.KillOneTaskReply> killOneTask_Pre($grpc.ServiceCall call,
      $async.Future<$0.KillOneTaskRequest> request) async {
    return killOneTask(call, await request);
  }

  $async.Future<$0.GetOneTaskReply> getOneTask_Pre($grpc.ServiceCall call,
      $async.Future<$0.GetOneTaskRequest> request) async {
    return getOneTask(call, await request);
  }

  $async.Future<$0.GetAllTasksReply> getAllTasks_Pre($grpc.ServiceCall call,
      $async.Future<$0.GetAllTasksRequest> request) async {
    return getAllTasks(call, await request);
  }

  $async.Future<$0.SaveOneTaskReply> saveOneTask(
      $grpc.ServiceCall call, $0.SaveOneTaskRequest request);
  $async.Future<$0.DeleteOneTaskReply> deleteOneTask(
      $grpc.ServiceCall call, $0.DeleteOneTaskRequest request);
  $async.Future<$0.KillOneTaskReply> killOneTask(
      $grpc.ServiceCall call, $0.KillOneTaskRequest request);
  $async.Future<$0.GetOneTaskReply> getOneTask(
      $grpc.ServiceCall call, $0.GetOneTaskRequest request);
  $async.Future<$0.GetAllTasksReply> getAllTasks(
      $grpc.ServiceCall call, $0.GetAllTasksRequest request);
}
