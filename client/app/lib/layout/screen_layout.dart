import 'package:flutter/material.dart';

import 'device_screen_type.enum.dart';
import 'rs_builder.dart';
import 'size_information.dart';



class ScreenLayout extends StatelessWidget {

  final Widget mobile;
  final Widget tablet;
  final Widget desktop;
  final Widget webMobile;
  final Widget webTablet;
  final Widget webDesktop;

  const ScreenLayout({
    Key key,
    this.mobile,
    this.tablet,
    this.desktop,
    this.webMobile,
    this.webTablet,
    this.webDesktop,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return RSBuilder(
      builder: (BuildContext context, SizeInformation sizingInformation) {

        if (sizingInformation.deviceScreenType == DeviceScreenType.WebMobile) {
          return  webMobile != null ? webMobile : mobile;
        }

        if (sizingInformation.deviceScreenType == DeviceScreenType.WebTablet) {
          if (webTablet != null) {
            return webTablet;
          }

          return tablet != null ? tablet : mobile;
        }

        if (sizingInformation.deviceScreenType == DeviceScreenType.WebDesktop) {
          if (webDesktop != null) {
            return webDesktop;
          }

          return desktop != null ? desktop : mobile;
        }

        if (sizingInformation.deviceScreenType == DeviceScreenType.Tablet) {
          if (tablet != null) {
            return tablet;
          }
        }

        if (sizingInformation.deviceScreenType == DeviceScreenType.Desktop) {
          if (desktop != null) {
            return desktop;
          }
        }

        return mobile;
      },
    );
  }
}
