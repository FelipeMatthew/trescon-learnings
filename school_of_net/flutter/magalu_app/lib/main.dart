import 'package:flutter/material.dart';

import 'package:magalu_app/widgets/app_bar_search.dart';
import 'package:magalu_app/widgets/app_drawer_menu.dart';

// Para funcionar ferramenta de pesquisa precisa iniciar como stateless - estática, podendo assim dentro dela passar statefull;
void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      title: 'Magalu clone app',
      home: HomeMyApp(),
    );
  }
}

// TODO: New File
class HomeMyApp extends StatefulWidget {
  const HomeMyApp({super.key});

  @override
  State<HomeMyApp> createState() => _HomeMyAppState();
}

class _HomeMyAppState extends State<HomeMyApp> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // * Menu drawer
      drawer: const Drawer(child: MenuDrawer()),
      appBar: AppBar(
        title: const Text(
          'Magazine Luiza',
        ),
        iconTheme: const IconThemeData(color: Colors.white),
        titleTextStyle: const TextStyle(
            color: Colors.white, fontSize: 20.0, fontWeight: FontWeight.bold),
        backgroundColor: Colors.blue,
        actions: <Widget>[
          IconButton(onPressed: () {}, icon: const Icon(Icons.shopping_cart)),
          IconButton(
              onPressed: () {
                showSearch(context: context, delegate: AppSearchBar());
              },
              icon: const Icon(Icons.search)),
        ],
      ),
      body: const Center(
        child: Text('Home Magazine Luiza'),
      ),
    );
  }
}



// TODO Create scafold file 
 
