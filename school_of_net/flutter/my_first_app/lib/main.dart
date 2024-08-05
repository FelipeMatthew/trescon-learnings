import 'package:flutter/material.dart';

void main() => runApp(MyApp());

// Static app
class MyApp extends StatelessWidget
{
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "My app", // Titulo
      home: Scaffold( 
        appBar: AppBar(// Header
          title: Text('My app'),
        ),
        body: Center(
          child: Text('Olá from body part'),
        ), // Body
      ),
      
    );
  }
  
}