import 'package:flutter/material.dart';

class ResultsScreen extends StatelessWidget {
  const ResultsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: const Text('Request result'),
        centerTitle: true,
        backgroundColor: Colors.black,
        titleTextStyle: const TextStyle(
            color: Colors.white, fontSize: 25, fontWeight: FontWeight.bold),
        iconTheme: const IconThemeData(color: Colors.white),
      ),
      body: Center(
        child: Text('Hii'),
      ),
    );
  }
}
