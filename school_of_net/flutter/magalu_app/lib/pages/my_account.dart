import 'package:flutter/material.dart';

class MyAccount extends StatelessWidget {
  const MyAccount({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        iconTheme: const IconThemeData(color: Colors.white),
        title: const Text('Minha conta'),
        backgroundColor: Colors.blue,
        titleTextStyle: const TextStyle(
            color: Colors.white, fontSize: 20.0, fontWeight: FontWeight.bold),
      ),
      body: const Center(
        child: Text('Minha conta'),
      ),
    );
  }
}
