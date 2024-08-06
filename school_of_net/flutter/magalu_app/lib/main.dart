import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Magalu clone app',
      home: Scaffold(
        appBar: AppBar(
          title: const Text(
            'Magazine Luiza',
          ),
          titleTextStyle: const TextStyle(color: Colors.white, fontSize: 22.0),
          backgroundColor: Colors.blue,
          actions: <Widget>[IconButton(onPressed: onPressed, icon: icon)],
        ),
        body: Center(
          child: Text('Home Magazine Luiza'),
        ),
      ),
    );
  }
}
