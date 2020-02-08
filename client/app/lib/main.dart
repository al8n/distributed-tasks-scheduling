import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:sailor/sailor.dart';
import 'package:flutter_villains/villains/villains.dart';
import 'package:app/pages/router/router.dart';
import 'package:app/services/locator.dart';
import 'package:app/services/sp/sp_service.dart';
import 'package:app/models/ui/theme/theme.dart';
import 'package:app/pages/index/index.page.dart';

Future<void> main() async {
  try {
    WidgetsFlutterBinding.ensureInitialized();
    await setupLocator();
    SPService _sp = locator<SPService>();
    bool mode = _sp.darkMode;
    int themeIndex = _sp.themeIndex;
    Routes.createRoutes();

    runApp(App(
      mode:  mode,
      themeIndex: themeIndex,
    ));
  } catch (e) {
    print(e);
    print('Locator setup has failed');
  }
}



class App extends StatefulWidget {
  final bool mode;
  final int themeIndex;

  App({
    @required this.mode,
    @required this.themeIndex,
  });

  @override
  _AppState createState() => _AppState();
}

class _AppState extends State<App> {
  final PageController controller = PageController();
  final GlobalKey<ScaffoldState> scaffoldStateKey = GlobalKey<ScaffoldState>();
  final GlobalKey drawerKey = GlobalKey(debugLabel: "drawer key");

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider.value(
          value: ThemeModel(themeIndex: widget.themeIndex, isDark: widget.mode),
        ),
      ],
      child: Consumer(
        builder: (BuildContext context, ThemeModel tm, Widget _){
          return MaterialApp(
            title: 'Flutter Demo',
            theme: tm.currentTheme,
            debugShowCheckedModeBanner: false,
            home: IndexPage(),
            navigatorObservers: [
              SailorLoggingObserver(),
              VillainTransitionObserver(),
            ],
            onGenerateRoute: Routes.sailor.generator(),
            navigatorKey: Routes.sailor.navigatorKey,
          );
        },
      ),
    );
  }
}
