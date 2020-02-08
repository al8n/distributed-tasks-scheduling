import 'package:app/models/type/ploc_status.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter/foundation.dart';

abstract class PlocBaseModel extends ChangeNotifier {

  PlocStatus _status = PlocStatus.idle;

  PlocStatus get status => _status;

  void setBusy() {
    _status = PlocStatus.busy;
    notifyListeners();
  }

  void setIdle() {
    _status = PlocStatus.idle;
    notifyListeners();
  }

  void initialise() {}
}