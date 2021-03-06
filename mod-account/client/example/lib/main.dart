
import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:intl/intl.dart';
import 'package:mod_account/modules/account_module.dart';
import 'package:mod_account/core/i18n/mod_account_localization.dart';

void main() => runApp(ModularApp(module: AppModule()));

class AppModule extends MainModule {
  // here will be any class you want to inject into your project (eg bloc, dependency)
  @override
  List<Bind> get binds => [];

  // here will be the routes of your module
  @override
  List<Router> get routers => [
        Router(
          "/",
          module: AccountModule("/")
        ),
      ];

// add your main widget here
  @override
  Widget get bootstrap => App();
}

class App extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      initialRoute: "/",
      onGenerateRoute: Modular.generateRoute,
      localizationsDelegates: [
        ModAccountLocalizationsDelegate(Locale('en')),
      ],
    );
  }
}