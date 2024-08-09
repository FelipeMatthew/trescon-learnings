import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
          width: double.maxFinite,
          decoration: const BoxDecoration(
              color: Colors.black,
              borderRadius:
                  BorderRadius.only(bottomLeft: Radius.circular(100))),
          child: Padding(
            padding: const EdgeInsets.all(70),
            child: Image.asset(
              'images/github-logo.png',
              height: 200,
            ),
          )),
    );
  }
}
