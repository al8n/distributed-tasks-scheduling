import 'package:get_it/get_it.dart';
import 'package:app/services/sp/sp_service.dart';

GetIt locator = GetIt.instance;

Future setupLocator() async {
  SPService sp = await SPService.getInstance();
  locator.registerSingleton<SPService>(sp);
}