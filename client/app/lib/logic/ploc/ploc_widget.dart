import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class PLOCWidget<T extends ChangeNotifier> extends StatefulWidget {
  final Widget Function(BuildContext) builder;
  final Function(T) onModelReady;
  final T model;

  PLOCWidget({
    Key key,
    @required this.builder,
    this.onModelReady,
    this.model,
  }) : super(key: key);

  @override
  _PLOCWidgetState<T> createState() => _PLOCWidgetState();
}


class _PLOCWidgetState<T extends ChangeNotifier> extends State<PLOCWidget<T>> {

  T _model;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    _model = widget.model;

    if (widget.model != null) {
      widget.onModelReady(_model);
    }
  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      child: widget.builder(context),
      create: (BuildContext context) => _model,
    );
  }
}
