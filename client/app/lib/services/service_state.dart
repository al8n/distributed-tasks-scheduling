import 'package:flutter/material.dart';
import 'package:app/services/sp/sp_service.dart';
import 'package:app/services/locator.dart';

mixin SPServiceMixin {
  @protected
  final SPService spService = locator<SPService>();
}