import 'package:flutter/material.dart';
import 'package:provider/provider.dart';



abstract class PLOCModelWidget<T> extends Widget {
  const PLOCModelWidget({Key key}) : super(key: key);

  @protected
  Widget build(BuildContext context, T model);

  @override
  _DataProviderElement<T> createElement() => _DataProviderElement<T>(this);

}

class _DataProviderElement<T> extends ComponentElement {
  _DataProviderElement(PLOCModelWidget widget) : super(widget);

  @override
  PLOCModelWidget get widget => super.widget;


  @override
  Widget build() => widget.build(this, Provider.of<T>(this));
}


abstract class PLOCModelWidget2<T, K> extends Widget {
  @protected
  Widget build(BuildContext context, T model, K model2);

  @override
  _DataProviderElement2<T, K> createElement() => _DataProviderElement2<T, K>(this);
}


class _DataProviderElement2<T, K> extends ComponentElement {
  _DataProviderElement2(PLOCModelWidget2 widget) : super(widget);

  @override
  PLOCModelWidget2 get widget => super.widget;


  @override
  Widget build() => widget.build(this, Provider.of<T>(this), Provider.of<K>(this));
}

abstract class PLOCModelWidget3<T, K, R> extends Widget {
  @protected
  Widget build(BuildContext context, T model, K model2, R model3);

  @override
  _DataProviderElement3<T, K, R> createElement() => _DataProviderElement3<T, K, R>(this);
}

class _DataProviderElement3<T, K, R> extends ComponentElement {
  _DataProviderElement3(PLOCModelWidget3 widget) : super(widget);

  @override
  PLOCModelWidget3 get widget => super.widget;


  @override
  Widget build() => widget.build(this, Provider.of<T>(this), Provider.of<K>(this), Provider.of<R>(this));
}

mixin PLOCModelAnimatedWidget3<T, K, R> implements PLOCModelWidget3, TickerProvider {}

abstract class PLOCModelWidget4<T, K, R, V> extends Widget {
  @protected
  Widget build(BuildContext context, T model, K model2, R model3, V model4);

  @override
  _DataProviderElement4<T, K, R, V> createElement() => _DataProviderElement4<T, K, R, V>(this);
}


class _DataProviderElement4<T, K, R, V> extends ComponentElement {
  _DataProviderElement4(PLOCModelWidget4 widget) : super(widget);

  @override
  PLOCModelWidget4 get widget => super.widget;


  @override
  Widget build() => widget.build(this, Provider.of<T>(this), Provider.of<K>(this), Provider.of<R>(this), Provider.of<V>(this));
}