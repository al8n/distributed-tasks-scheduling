import 'package:app/grpc/grpc.dart';
import 'package:app/grpc/pb/tasks/tasks.pb.dart';
import 'package:app/grpc/pb/tasks/tasks.pbgrpc.dart';

class  TasksService {
  static Future<GetAllTasksReply> GetAllTasks() async {
    var client = TasksClient(GrpcClientSingleton().client);
    return await client.getAllTasks(GetAllTasksRequest());
  }
}
