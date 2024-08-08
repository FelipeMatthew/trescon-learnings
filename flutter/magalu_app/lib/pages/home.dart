import 'package:flutter/material.dart';

import 'package:magalu/pages/shopping_cart.dart';
import 'package:magalu/widgets/app_bar_search.dart';
import 'package:magalu/widgets/app_drawer_menu.dart';

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
        drawer: const Drawer(
          width: 250,
          child: MenuDrawer(),
        ),
        appBar: AppBar(
          title: const Text(
            'Magazine Luiza',
          ),
          iconTheme: const IconThemeData(color: Colors.white),
          titleTextStyle: const TextStyle(
              color: Colors.white, fontSize: 20.0, fontWeight: FontWeight.bold),
          backgroundColor: Colors.blue,
          actions: <Widget>[
            IconButton(
                onPressed: () {
                  Navigator.push(
                      context,
                      MaterialPageRoute(
                          builder: (context) => const ShoppingCart()));
                },
                icon: const Icon(Icons.shopping_cart)),
            IconButton(
                onPressed: () {
                  showSearch(context: context, delegate: AppSearchBar());
                },
                icon: const Icon(Icons.search)),
          ],
        ),
        body: RefreshIndicator(
          onRefresh: _refresh,
          child: const Text('to'),
        ));
  }

  Future<void> _refresh() async {
    await Future.delayed(const Duration(seconds: 2));

    setState(() {
      Future.value();
    });
    return Future.value();
  }
}



// TODO Create scafold file 
 

