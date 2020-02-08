import 'package:flutter/material.dart';
import 'package:sailor/sailor.dart';

class SailorSizeTransition extends CustomSailorTransition {
  @override
  Widget buildTransition(BuildContext context, Animation<double> animation, Animation<double> secondaryAnimation, Widget child) {
    // TODO: implement buildTransition
    return SizeTransition(
      sizeFactor: animation,
      child: child,
    );
  }
}

class SailorSlideTransition extends CustomSailorTransition {
  @override
  Widget buildTransition(BuildContext context, Animation<double> animation, Animation<double> secondaryAnimation, Widget child) {
    // TODO: implement buildTransition
    return SlideTransition(

      child: child,
      position: Tween(
        begin: Offset(0, .2),
        end: Offset(0, 0),
      ).animate(animation),
    );
  }
}


class SailorFadeTransition extends CustomSailorTransition {
  @override
  Widget buildTransition(BuildContext context, Animation<double> animation, Animation<double> secondaryAnimation, Widget child) {
    // TODO: implement buildTransition
    return FadeTransition(
      opacity: animation,
      child: child,
    );
  }
}

class SailorScaleTransition extends CustomSailorTransition {
  @override
  Widget buildTransition(BuildContext context, Animation<double> animation, Animation<double> secondaryAnimation, Widget child) {
    // TODO: implement buildTransition
    return ScaleTransition(
      scale: Tween<double>(
        begin: 0.0,
        end: 1.0,
      ).animate(
        CurvedAnimation(
          parent: animation,
          curve: Curves.ease,
        ),
      ),
      child: child,
    );
  }
}