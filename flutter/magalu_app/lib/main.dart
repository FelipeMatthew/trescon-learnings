import 'package:flutter/material.dart';

import 'package:magalu/main.dart';
import 'package:magalu/pages/home.dart';
import 'package:magalu/pages/shopping_cart.dart';
import 'package:magalu/pages/splash.dart';
import 'package:magalu/widgets/app_bar_search.dart';
import 'package:magalu/widgets/app_drawer_menu.dart';

// Para funcionar ferramenta de pesquisa precisa iniciar como stateless - estática, podendo assim dentro dela passar statefull;
void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Magalu clone app',
      home: Splash(),
    );
  }
}

// TODO: New File
