import 'package:flutter/material.dart';

import 'device_screen_type.enum.dart';



class SizeInformation {
  // 横放还是竖放
  final Orientation orientation;
  // 设备类型
  final DeviceScreenType deviceScreenType;
  // 设备屏幕的尺寸
  final Size screenSize;
  // 当前组件的尺寸
  final Size localWidgetSize;

  SizeInformation({
    this.orientation,
    this.deviceScreenType,
    this.screenSize,
    this.localWidgetSize,
  });

  @override
  String toString() {
    return 'Orientation: $orientation DeviceType:$deviceScreenType ScreenSize:$screenSize LocalWidgetSize:$localWidgetSize';
  }
}
