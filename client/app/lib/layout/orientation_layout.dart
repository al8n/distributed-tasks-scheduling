import 'package:flutter/material.dart';

class OrientationLayout extends StatelessWidget {

  final Widget Function(BuildContext) landscape;
  final Widget Function(BuildContext) portrait;

  const OrientationLayout({
    Key key,
    this.landscape,
    @required this.portrait,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {

    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints constraints) {
        Orientation orientation = MediaQuery.of(context).orientation;

        // 如果当前设备为横放并且传入了横放时的UI组件
        if (orientation == Orientation.landscape) {
          if (landscape != null) {
            return landscape(context);
          }
        }

        // 默认返回竖放状态的传入的组件
        return portrait(context);
      },
    );
  }
}