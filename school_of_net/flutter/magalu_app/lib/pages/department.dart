import 'package:flutter/material.dart';

class Department extends StatelessWidget {
  const Department({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        iconTheme: const IconThemeData(color: Colors.white),
        title: const Text('Departamentos'),
        backgroundColor: Colors.blue,
        titleTextStyle: const TextStyle(
            color: Colors.white, fontSize: 20.0, fontWeight: FontWeight.bold),
      ),
      body: const Center(
        child: Text('Departamentos'),
      ),
    );
  }
}
