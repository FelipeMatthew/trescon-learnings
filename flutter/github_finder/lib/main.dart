import 'package:flutter/material.dart';
import 'package:github_finder/pages/home.dart';

// TODO: Splash screen
void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Github finder',
      home: HomeScreen(),
    );
  }
}
