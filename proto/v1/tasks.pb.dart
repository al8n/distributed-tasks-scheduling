///
//  Generated code. Do not modify.
//  source: tasks.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Task extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Task', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..aOS(1, 'name')
    ..aOS(2, 'command')
    ..aOS(3, 'expression')
    ..hasRequiredFields = false
  ;

  Task._() : super();
  factory Task() => create();
  factory Task.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Task.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Task clone() => Task()..mergeFromMessage(this);
  Task copyWith(void Function(Task) updates) => super.copyWith((message) => updates(message as Task));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Task create() => Task._();
  Task createEmptyInstance() => create();
  static $pb.PbList<Task> createRepeated() => $pb.PbList<Task>();
  @$core.pragma('dart2js:noInline')
  static Task getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Task>(create);
  static Task _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get command => $_getSZ(1);
  @$pb.TagNumber(2)
  set command($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCommand() => $_has(1);
  @$pb.TagNumber(2)
  void clearCommand() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get expression => $_getSZ(2);
  @$pb.TagNumber(3)
  set expression($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasExpression() => $_has(2);
  @$pb.TagNumber(3)
  void clearExpression() => clearField(3);
}

class SaveOneTaskRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SaveOneTaskRequest', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..aOS(1, 'token')
    ..aOM<Task>(2, 'task', subBuilder: Task.create)
    ..aOS(3, 'operator')
    ..hasRequiredFields = false
  ;

  SaveOneTaskRequest._() : super();
  factory SaveOneTaskRequest() => create();
  factory SaveOneTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SaveOneTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  SaveOneTaskRequest clone() => SaveOneTaskRequest()..mergeFromMessage(this);
  SaveOneTaskRequest copyWith(void Function(SaveOneTaskRequest) updates) => super.copyWith((message) => updates(message as SaveOneTaskRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SaveOneTaskRequest create() => SaveOneTaskRequest._();
  SaveOneTaskRequest createEmptyInstance() => create();
  static $pb.PbList<SaveOneTaskRequest> createRepeated() => $pb.PbList<SaveOneTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static SaveOneTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SaveOneTaskRequest>(create);
  static SaveOneTaskRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  Task get task => $_getN(1);
  @$pb.TagNumber(2)
  set task(Task v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTask() => $_has(1);
  @$pb.TagNumber(2)
  void clearTask() => clearField(2);
  @$pb.TagNumber(2)
  Task ensureTask() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.String get operator => $_getSZ(2);
  @$pb.TagNumber(3)
  set operator($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasOperator() => $_has(2);
  @$pb.TagNumber(3)
  void clearOperator() => clearField(3);
}

class SaveOneTaskReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SaveOneTaskReply', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..a<$core.int>(1, 'errorNumber', $pb.PbFieldType.O3)
    ..aOS(2, 'msg')
    ..aOM<Task>(3, 'task', subBuilder: Task.create)
    ..hasRequiredFields = false
  ;

  SaveOneTaskReply._() : super();
  factory SaveOneTaskReply() => create();
  factory SaveOneTaskReply.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SaveOneTaskReply.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  SaveOneTaskReply clone() => SaveOneTaskReply()..mergeFromMessage(this);
  SaveOneTaskReply copyWith(void Function(SaveOneTaskReply) updates) => super.copyWith((message) => updates(message as SaveOneTaskReply));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SaveOneTaskReply create() => SaveOneTaskReply._();
  SaveOneTaskReply createEmptyInstance() => create();
  static $pb.PbList<SaveOneTaskReply> createRepeated() => $pb.PbList<SaveOneTaskReply>();
  @$core.pragma('dart2js:noInline')
  static SaveOneTaskReply getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SaveOneTaskReply>(create);
  static SaveOneTaskReply _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get errorNumber => $_getIZ(0);
  @$pb.TagNumber(1)
  set errorNumber($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasErrorNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearErrorNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);

  @$pb.TagNumber(3)
  Task get task => $_getN(2);
  @$pb.TagNumber(3)
  set task(Task v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTask() => $_has(2);
  @$pb.TagNumber(3)
  void clearTask() => clearField(3);
  @$pb.TagNumber(3)
  Task ensureTask() => $_ensure(2);
}

class DeleteOneTaskRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DeleteOneTaskRequest', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..aOS(1, 'token')
    ..aOS(2, 'taskKey', protoName: 'taskKey')
    ..hasRequiredFields = false
  ;

  DeleteOneTaskRequest._() : super();
  factory DeleteOneTaskRequest() => create();
  factory DeleteOneTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteOneTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DeleteOneTaskRequest clone() => DeleteOneTaskRequest()..mergeFromMessage(this);
  DeleteOneTaskRequest copyWith(void Function(DeleteOneTaskRequest) updates) => super.copyWith((message) => updates(message as DeleteOneTaskRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteOneTaskRequest create() => DeleteOneTaskRequest._();
  DeleteOneTaskRequest createEmptyInstance() => create();
  static $pb.PbList<DeleteOneTaskRequest> createRepeated() => $pb.PbList<DeleteOneTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static DeleteOneTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteOneTaskRequest>(create);
  static DeleteOneTaskRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get taskKey => $_getSZ(1);
  @$pb.TagNumber(2)
  set taskKey($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTaskKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearTaskKey() => clearField(2);
}

class DeleteOneTaskReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DeleteOneTaskReply', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..a<$core.int>(1, 'errorNumber', $pb.PbFieldType.O3)
    ..aOS(2, 'msg')
    ..hasRequiredFields = false
  ;

  DeleteOneTaskReply._() : super();
  factory DeleteOneTaskReply() => create();
  factory DeleteOneTaskReply.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteOneTaskReply.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DeleteOneTaskReply clone() => DeleteOneTaskReply()..mergeFromMessage(this);
  DeleteOneTaskReply copyWith(void Function(DeleteOneTaskReply) updates) => super.copyWith((message) => updates(message as DeleteOneTaskReply));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteOneTaskReply create() => DeleteOneTaskReply._();
  DeleteOneTaskReply createEmptyInstance() => create();
  static $pb.PbList<DeleteOneTaskReply> createRepeated() => $pb.PbList<DeleteOneTaskReply>();
  @$core.pragma('dart2js:noInline')
  static DeleteOneTaskReply getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteOneTaskReply>(create);
  static DeleteOneTaskReply _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get errorNumber => $_getIZ(0);
  @$pb.TagNumber(1)
  set errorNumber($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasErrorNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearErrorNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);
}

class KillOneTaskRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('KillOneTaskRequest', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..aOS(1, 'token')
    ..aOS(2, 'taskKey', protoName: 'taskKey')
    ..hasRequiredFields = false
  ;

  KillOneTaskRequest._() : super();
  factory KillOneTaskRequest() => create();
  factory KillOneTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KillOneTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  KillOneTaskRequest clone() => KillOneTaskRequest()..mergeFromMessage(this);
  KillOneTaskRequest copyWith(void Function(KillOneTaskRequest) updates) => super.copyWith((message) => updates(message as KillOneTaskRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static KillOneTaskRequest create() => KillOneTaskRequest._();
  KillOneTaskRequest createEmptyInstance() => create();
  static $pb.PbList<KillOneTaskRequest> createRepeated() => $pb.PbList<KillOneTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static KillOneTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KillOneTaskRequest>(create);
  static KillOneTaskRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get taskKey => $_getSZ(1);
  @$pb.TagNumber(2)
  set taskKey($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTaskKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearTaskKey() => clearField(2);
}

class KillOneTaskReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('KillOneTaskReply', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..a<$core.int>(1, 'errorNumber', $pb.PbFieldType.O3)
    ..aOS(2, 'msg')
    ..hasRequiredFields = false
  ;

  KillOneTaskReply._() : super();
  factory KillOneTaskReply() => create();
  factory KillOneTaskReply.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KillOneTaskReply.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  KillOneTaskReply clone() => KillOneTaskReply()..mergeFromMessage(this);
  KillOneTaskReply copyWith(void Function(KillOneTaskReply) updates) => super.copyWith((message) => updates(message as KillOneTaskReply));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static KillOneTaskReply create() => KillOneTaskReply._();
  KillOneTaskReply createEmptyInstance() => create();
  static $pb.PbList<KillOneTaskReply> createRepeated() => $pb.PbList<KillOneTaskReply>();
  @$core.pragma('dart2js:noInline')
  static KillOneTaskReply getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KillOneTaskReply>(create);
  static KillOneTaskReply _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get errorNumber => $_getIZ(0);
  @$pb.TagNumber(1)
  set errorNumber($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasErrorNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearErrorNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);
}

class GetOneTaskRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetOneTaskRequest', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..aOS(1, 'token')
    ..aOS(2, 'taskKey', protoName: 'taskKey')
    ..hasRequiredFields = false
  ;

  GetOneTaskRequest._() : super();
  factory GetOneTaskRequest() => create();
  factory GetOneTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetOneTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetOneTaskRequest clone() => GetOneTaskRequest()..mergeFromMessage(this);
  GetOneTaskRequest copyWith(void Function(GetOneTaskRequest) updates) => super.copyWith((message) => updates(message as GetOneTaskRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetOneTaskRequest create() => GetOneTaskRequest._();
  GetOneTaskRequest createEmptyInstance() => create();
  static $pb.PbList<GetOneTaskRequest> createRepeated() => $pb.PbList<GetOneTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static GetOneTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetOneTaskRequest>(create);
  static GetOneTaskRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get taskKey => $_getSZ(1);
  @$pb.TagNumber(2)
  set taskKey($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTaskKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearTaskKey() => clearField(2);
}

class GetOneTaskReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetOneTaskReply', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..a<$core.int>(1, 'errorNumber', $pb.PbFieldType.O3)
    ..aOS(2, 'msg')
    ..aOM<Task>(3, 'task', subBuilder: Task.create)
    ..hasRequiredFields = false
  ;

  GetOneTaskReply._() : super();
  factory GetOneTaskReply() => create();
  factory GetOneTaskReply.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetOneTaskReply.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetOneTaskReply clone() => GetOneTaskReply()..mergeFromMessage(this);
  GetOneTaskReply copyWith(void Function(GetOneTaskReply) updates) => super.copyWith((message) => updates(message as GetOneTaskReply));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetOneTaskReply create() => GetOneTaskReply._();
  GetOneTaskReply createEmptyInstance() => create();
  static $pb.PbList<GetOneTaskReply> createRepeated() => $pb.PbList<GetOneTaskReply>();
  @$core.pragma('dart2js:noInline')
  static GetOneTaskReply getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetOneTaskReply>(create);
  static GetOneTaskReply _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get errorNumber => $_getIZ(0);
  @$pb.TagNumber(1)
  set errorNumber($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasErrorNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearErrorNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);

  @$pb.TagNumber(3)
  Task get task => $_getN(2);
  @$pb.TagNumber(3)
  set task(Task v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTask() => $_has(2);
  @$pb.TagNumber(3)
  void clearTask() => clearField(3);
  @$pb.TagNumber(3)
  Task ensureTask() => $_ensure(2);
}

class GetAllTasksRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetAllTasksRequest', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..aOS(1, 'token')
    ..hasRequiredFields = false
  ;

  GetAllTasksRequest._() : super();
  factory GetAllTasksRequest() => create();
  factory GetAllTasksRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetAllTasksRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetAllTasksRequest clone() => GetAllTasksRequest()..mergeFromMessage(this);
  GetAllTasksRequest copyWith(void Function(GetAllTasksRequest) updates) => super.copyWith((message) => updates(message as GetAllTasksRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetAllTasksRequest create() => GetAllTasksRequest._();
  GetAllTasksRequest createEmptyInstance() => create();
  static $pb.PbList<GetAllTasksRequest> createRepeated() => $pb.PbList<GetAllTasksRequest>();
  @$core.pragma('dart2js:noInline')
  static GetAllTasksRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetAllTasksRequest>(create);
  static GetAllTasksRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);
}

class GetAllTasksReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetAllTasksReply', package: const $pb.PackageName('taskspb'), createEmptyInstance: create)
    ..a<$core.int>(1, 'errorNumber', $pb.PbFieldType.O3)
    ..aOS(2, 'msg')
    ..pc<Task>(3, 'tasks', $pb.PbFieldType.PM, subBuilder: Task.create)
    ..hasRequiredFields = false
  ;

  GetAllTasksReply._() : super();
  factory GetAllTasksReply() => create();
  factory GetAllTasksReply.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetAllTasksReply.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetAllTasksReply clone() => GetAllTasksReply()..mergeFromMessage(this);
  GetAllTasksReply copyWith(void Function(GetAllTasksReply) updates) => super.copyWith((message) => updates(message as GetAllTasksReply));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetAllTasksReply create() => GetAllTasksReply._();
  GetAllTasksReply createEmptyInstance() => create();
  static $pb.PbList<GetAllTasksReply> createRepeated() => $pb.PbList<GetAllTasksReply>();
  @$core.pragma('dart2js:noInline')
  static GetAllTasksReply getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetAllTasksReply>(create);
  static GetAllTasksReply _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get errorNumber => $_getIZ(0);
  @$pb.TagNumber(1)
  set errorNumber($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasErrorNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearErrorNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<Task> get tasks => $_getList(2);
}

