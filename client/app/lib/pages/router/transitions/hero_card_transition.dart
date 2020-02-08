import 'package:flutter/material.dart';


class CardDetailsPageRoute extends PageRouteBuilder {

  final Widget enterPage;
  final Color backgroundColor;
  final Duration duration;

  CardDetailsPageRoute({
    @required this.enterPage,
    @required this.backgroundColor,
    this.duration = const Duration(milliseconds: 800),
  }) : super(
      transitionDuration: duration,
      pageBuilder: (context, animation, secondaryAnimation) => enterPage,
      transitionsBuilder: (context, animation, secondaryAnimation, child) {
        Animation<double> fadeIn = Tween<double>(begin: 0, end: 1).animate(CurvedAnimation(curve: Interval(.7, 1), parent: animation));
        Animation<double> fadeOut =
        Tween<double>(begin: 0, end: 1).animate(CurvedAnimation(curve: Interval(0, .2), parent: animation));

        return Stack(children: <Widget>[
          FadeTransition(
            opacity: fadeOut,
            child: Container(color: backgroundColor),
          ),
          FadeTransition(
            opacity: fadeIn,
            child: child,
          ),
        ]);
      });
}