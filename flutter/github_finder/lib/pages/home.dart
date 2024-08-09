import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        height: 400,
        width: double.maxFinite,
        color: Colors.black,
        child: const Icon(
          Icons.access_alarm,
          color: Colors.white,
        ),
      ),
    );
  }
}
