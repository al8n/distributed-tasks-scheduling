import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart';

import 'size_information.dart';
import 'device_screen_type.enum.dart';


class RSBuilder extends StatelessWidget {
  final Widget Function(BuildContext context, SizeInformation sizeInformation) builder;


  const RSBuilder({
    Key key,
    @required this.builder,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    MediaQueryData mediaQueryData = MediaQuery.of(context);
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints boxConstraints) {
        SizeInformation sizeInformation = SizeInformation(
          orientation: mediaQueryData.orientation,
          deviceScreenType: getDeviceType(mediaQueryData),
          screenSize: mediaQueryData.size,
          localWidgetSize: Size(boxConstraints.maxWidth, boxConstraints.maxHeight),
        );

        return builder(context, sizeInformation);
      },
    );
  }
}

DeviceScreenType getDeviceType(MediaQueryData mediaQueryData) {

  // 获取高度和宽度中数值较小的作为设备宽度
  double deviceWidth = mediaQueryData.size.shortestSide;

  // 判断用户是否通过浏览器访问
  if (kIsWeb) {

    if (mediaQueryData.size.width >= 960) {
      return DeviceScreenType.WebDesktop;
    }

    if (mediaQueryData.size.width >= 600) {
      return DeviceScreenType.WebTablet;
    }

    return DeviceScreenType.WebMobile;
  }

  if (deviceWidth >= 960) {
    return DeviceScreenType.Desktop;
  }

  if (deviceWidth >= 600) {
    return DeviceScreenType.Tablet;
  }

  return DeviceScreenType.Mobile;
}