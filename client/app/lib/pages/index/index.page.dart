import 'package:app/layout/orientation_layout.dart';
import 'package:app/layout/screen_layout.dart';
import 'package:app/pages/index/index.destop.page.dart';
import 'package:app/pages/index/index.mobile.page.dart';
import 'package:flutter/material.dart';


class IndexPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ScreenLayout(
      desktop: OrientationLayout(
        portrait: (BuildContext context) => IndexDesktopPage(),
      ),
      mobile: OrientationLayout(
        portrait: (BuildContext context) => IndexDesktopPage(),
      ),
    );
  }
}
