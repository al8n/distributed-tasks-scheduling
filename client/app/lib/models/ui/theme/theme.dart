import 'package:app/constant/theme.dart';
import 'package:app/models/ui/ploc_model.dart';
import 'package:app/services/service_state.dart';
import 'package:flutter/material.dart';


class ThemeModel extends PlocBaseModel with SPServiceMixin {
  MaterialColor get themeColor => colors[_themeIndex];

  ThemeModel({@required themeIndex, @required isDark}) {
    _isDark = isDark;
    _themeIndex = themeIndex;
    currentTheme = getThemeData(colors[themeIndex], isDark);
  }

  int _themeIndex;

  int get themeIndex => _themeIndex;

  bool _isDark;

  bool get isDark => _isDark;

  ThemeData _currentTheme;

  ThemeData get currentTheme => _currentTheme;

  set currentTheme(ThemeData themeData) {
    _currentTheme = themeData;
    notifyListeners();
  }

  void setMode(bool tb) {
    if (tb != _isDark) {
      _isDark = tb;
      spService.darkMode = _isDark;
      _currentTheme = getThemeData(colors[_themeIndex], _isDark);
      notifyListeners();
    }
  }

  void setThemeIndex(int index) {
    if (index != _themeIndex) {
      _themeIndex = index;
      spService.themeIndex = _themeIndex;
      _currentTheme = getThemeData(colors[_themeIndex], _isDark);
      notifyListeners();
    }
  }

  final List<MaterialColor> colors = [
    Colors.lightBlue,
    Colors.blue,
    Colors.indigo,
    Colors.red,
    Colors.pink,
    Colors.purple,
    Colors.deepPurple,
    Colors.lime,
    Colors.lightGreen,
    Colors.green,
    Colors.teal,
    Colors.cyan,
    Colors.yellow,
    Colors.amber,
    Colors.orange,
    Colors.deepOrange,
  ];
}