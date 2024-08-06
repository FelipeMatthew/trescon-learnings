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
          title: Text('Magazine Luiza'),
          titleTextStyle: TextStyle(color: Colors.white),
          backgroundColor: Colors.blue,
        ),
        body: Center(child: Text('Home Magazine Luiza'),),
      ),
    );
  }
}
