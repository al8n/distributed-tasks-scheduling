///
//  Generated code. Do not modify.
//  source: tasks.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Task$json = const {
  '1': 'Task',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'command', '3': 2, '4': 1, '5': 9, '10': 'command'},
    const {'1': 'expression', '3': 3, '4': 1, '5': 9, '10': 'expression'},
  ],
};

const SaveOneTaskRequest$json = const {
  '1': 'SaveOneTaskRequest',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'task', '3': 2, '4': 1, '5': 11, '6': '.taskspb.Task', '10': 'task'},
    const {'1': 'operator', '3': 3, '4': 1, '5': 9, '10': 'operator'},
  ],
};

const DeleteOneTaskRequest$json = const {
  '1': 'DeleteOneTaskRequest',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'taskKey', '3': 2, '4': 1, '5': 9, '10': 'taskKey'},
  ],
};

const OneTaskReply$json = const {
  '1': 'OneTaskReply',
  '2': const [
    const {'1': 'error_number', '3': 1, '4': 1, '5': 5, '10': 'errorNumber'},
    const {'1': 'msg', '3': 2, '4': 1, '5': 9, '10': 'msg'},
  ],
};

const GetOneTaskRequest$json = const {
  '1': 'GetOneTaskRequest',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'taskKey', '3': 2, '4': 1, '5': 9, '10': 'taskKey'},
  ],
};

const GetOneTaskReply$json = const {
  '1': 'GetOneTaskReply',
  '2': const [
    const {'1': 'error_number', '3': 1, '4': 1, '5': 5, '10': 'errorNumber'},
    const {'1': 'msg', '3': 2, '4': 1, '5': 9, '10': 'msg'},
    const {'1': 'task', '3': 3, '4': 1, '5': 11, '6': '.taskspb.Task', '10': 'task'},
  ],
};

const GetAllTasksRequest$json = const {
  '1': 'GetAllTasksRequest',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

const GetAllTasksReply$json = const {
  '1': 'GetAllTasksReply',
  '2': const [
    const {'1': 'error_number', '3': 1, '4': 1, '5': 5, '10': 'errorNumber'},
    const {'1': 'msg', '3': 2, '4': 1, '5': 9, '10': 'msg'},
    const {'1': 'tasks', '3': 3, '4': 3, '5': 11, '6': '.taskspb.Task', '10': 'tasks'},
  ],
};

const TasksServiceBase$json = const {
  '1': 'Tasks',
  '2': const [
    const {'1': 'SaveOneTask', '2': '.taskspb.SaveOneTaskRequest', '3': '.taskspb.OneTaskReply', '4': const {}},
    const {'1': 'DeleteOneTask', '2': '.taskspb.DeleteOneTaskRequest', '3': '.taskspb.OneTaskReply', '4': const {}},
    const {'1': 'GetOneTask', '2': '.taskspb.GetOneTaskRequest', '3': '.taskspb.GetOneTaskReply', '4': const {}},
    const {'1': 'GetAllTasks', '2': '.taskspb.GetAllTasksRequest', '3': '.taskspb.GetAllTasksReply', '4': const {}},
  ],
};

const TasksServiceBase$messageJson = const {
  '.taskspb.SaveOneTaskRequest': SaveOneTaskRequest$json,
  '.taskspb.Task': Task$json,
  '.taskspb.OneTaskReply': OneTaskReply$json,
  '.taskspb.DeleteOneTaskRequest': DeleteOneTaskRequest$json,
  '.taskspb.GetOneTaskRequest': GetOneTaskRequest$json,
  '.taskspb.GetOneTaskReply': GetOneTaskReply$json,
  '.taskspb.GetAllTasksRequest': GetAllTasksRequest$json,
  '.taskspb.GetAllTasksReply': GetAllTasksReply$json,
};

