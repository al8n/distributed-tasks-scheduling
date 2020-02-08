import 'package:app/grpc/pb/tasks/tasks.pb.dart';
import 'package:app/services/grpc/tasks.services.dart';
import 'package:flutter/material.dart';

class IndexDesktopPage extends StatefulWidget {
  @override
  _IndexDesktopPageState createState() => _IndexDesktopPageState();
}

class _IndexDesktopPageState extends State<IndexDesktopPage> {

  GetAllTasksReply res;

  @override
  void initState() {

    res = new GetAllTasksReply();
    super.initState();
  }

  Future<void > _getAllTasks() async {
    var tasks = await TasksService.GetAllTasks();
    setState(() {

      print(tasks.tasks);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        width: 300,
        height: 400,
        decoration: BoxDecoration(
          color: Colors.blue
        ),
        child: MaterialButton(
          color: Colors.amber,
          onPressed: _getAllTasks,
          child: Text("获取任务"),
        ),
      ),
    );
  }
}
