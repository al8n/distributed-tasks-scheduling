import 'package:shared_preferences/shared_preferences.dart';

class SPService {
  static SPService _instance;
  static SharedPreferences _sp;

  static Future<SPService> getInstance() async {
    if (_instance == null) {
      _instance = SPService();
    }

    if (_sp == null) {
      _sp = await SharedPreferences.getInstance();
    }

    return _instance;
  }

  /// Theme
  static const String DarkModeKey = 'mode';
  static const String ThemeIndexKey = 'theme';

  bool get darkMode => _getFromDevice(DarkModeKey) ?? false;

  set darkMode(bool value) => saveToDevice(DarkModeKey, value);

  int get themeIndex => _getFromDevice(ThemeIndexKey) ?? 0;

  set themeIndex(int value) => saveToDevice(ThemeIndexKey, value);

  /// utils
  dynamic _getFromDevice(String key) {
    var value = _sp.get(key);
    print('(TRACE) LocalStorageService:_getFromDisk. key: $key value: $value');
    return value;
  }

  void saveToDevice<T>(String key, T content) {
    print('(TRACE) LocalStorageService:_saveStringToDisk. key: $key value: $content');

    if (content is String) {
      _sp.setString(key, content);
    }

    if (content is bool) {
      _sp.setBool(key, content);
    }

    if (content is int) {
      _sp.setInt(key, content);
    }

    if (content is double) {
      _sp.setDouble(key, content);
    }

    if (content is List<String>) {
      _sp.setStringList(key, content);
    }
  }
}
