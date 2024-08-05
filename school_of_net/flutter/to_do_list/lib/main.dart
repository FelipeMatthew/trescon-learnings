import 'package:flutter/material.dart';

void main() => runApp(ListApp());

class ListApp extends StatelessWidget {
  const ListApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'List App',
      home: Scaffold(
        appBar: AppBar(
          title: Text('To do list'),
          backgroundColor: Colors.blueAccent,
          centerTitle: true,
        ),
        body: Center(
          child: Text('Hiii'),
        ),
      ),
    );
  }
}