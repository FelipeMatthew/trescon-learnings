import 'package:flutter/material.dart';
import 'package:yolo_v8_obj_indentifier/camera.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
        title: 'YOLO OBJECT IDENTIFIER', home: YoloVideo());
  }
}
 