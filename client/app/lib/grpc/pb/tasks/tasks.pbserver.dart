///
//  Generated code. Do not modify.
//  source: tasks.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'tasks.pb.dart' as $0;
import 'tasks.pbjson.dart';

export 'tasks.pb.dart';

abstract class TasksServiceBase extends $pb.GeneratedService {
  $async.Future<$0.OneTaskReply> saveOneTask($pb.ServerContext ctx, $0.SaveOneTaskRequest request);
  $async.Future<$0.OneTaskReply> deleteOneTask($pb.ServerContext ctx, $0.DeleteOneTaskRequest request);
  $async.Future<$0.GetOneTaskReply> getOneTask($pb.ServerContext ctx, $0.GetOneTaskRequest request);
  $async.Future<$0.GetAllTasksReply> getAllTasks($pb.ServerContext ctx, $0.GetAllTasksRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'SaveOneTask': return $0.SaveOneTaskRequest();
      case 'DeleteOneTask': return $0.DeleteOneTaskRequest();
      case 'GetOneTask': return $0.GetOneTaskRequest();
      case 'GetAllTasks': return $0.GetAllTasksRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'SaveOneTask': return this.saveOneTask(ctx, request);
      case 'DeleteOneTask': return this.deleteOneTask(ctx, request);
      case 'GetOneTask': return this.getOneTask(ctx, request);
      case 'GetAllTasks': return this.getAllTasks(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TasksServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TasksServiceBase$messageJson;
}

