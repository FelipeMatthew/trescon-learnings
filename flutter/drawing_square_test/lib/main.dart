import 'package:drawing_square_test/test/drawing.dart';
import 'package:flutter/material.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Desenhando quadrado projeto'),
        ),
        body: const DesenhandoQuadrado(),
      ),
    );
  }
}
