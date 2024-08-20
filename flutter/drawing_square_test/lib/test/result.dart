import 'dart:typed_data';

import 'package:flutter/material.dart';

class Resultado extends StatelessWidget {
  final Uint8List imageData;
  const Resultado({super.key, required this.imageData});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Imagem editada'),),
      body: Center(
        child: Image.memory(imageData),
      ),
    );
  }
}
