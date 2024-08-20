import 'package:flutter/material.dart';
import 'dart:ui' as ui;
import 'package:flutter/services.dart';

class DesenhandoQuadrado extends StatefulWidget {
  const DesenhandoQuadrado({super.key});

  @override
  State<DesenhandoQuadrado> createState() => _DesenhandoQuadradoState();
}

class _DesenhandoQuadradoState extends State<DesenhandoQuadrado> {
  final List<Offset> _pontos = [];
  ui.Image? _image;

  @override
  void initState() {
    super.initState();
    _loadImage('assets/tubo.jpg');
  }

  Future<void> _loadImage(String path) async {
    final ByteData data = await rootBundle.load(path);
    final Uint8List bytes = data.buffer.asUint8List();
    final ui.Image image = await decodeImageFromList(bytes);
    setState(() {
      _image = image;
    });
  }

  void _adicionarQuadrado(Offset ponto) {
    setState(() {
      _pontos.add(ponto);
    });
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTapDown: (TapDownDetails detalhes) {
        RenderBox box = context.findRenderObject() as RenderBox;
        Offset localPosition = box.globalToLocal(detalhes.globalPosition);
        _adicionarQuadrado(localPosition);
      },
      child: CustomPaintContainer(pontos: _pontos, image: _image),
    );
  }
}

class CustomPaintContainer extends StatefulWidget {
  final List<Offset> pontos;
  final ui.Image? image;

  const CustomPaintContainer(
      {required this.pontos, required this.image, Key? key})
      : super(key: key);

  @override
  State<CustomPaintContainer> createState() => _CustomPaintContainerState();
}

class _CustomPaintContainerState extends State<CustomPaintContainer> {
  @override
  Widget build(BuildContext context) {
    return CustomPaint(
      size: Size.infinite,
      painter: QuadradoPainter(pontos: widget.pontos, image: widget.image),
    );
  }
}

class QuadradoPainter extends CustomPainter {
  final List<Offset> pontos;
  final ui.Image? image;

  QuadradoPainter({required this.pontos, this.image});

  @override
  void paint(Canvas canvas, Size size) {
    if (image != null) {
      final double imageWidth = image!.width.toDouble();
      final double imageHeight = image!.height.toDouble();
      final Offset center = Offset(
        (size.width - imageWidth) / 2,
        (size.height - imageHeight) / 2,
      );
      canvas.drawImage(image!, center, Paint());

      Paint paint = Paint()
        ..color = Colors.blue
        ..strokeWidth = 2.0
        ..style = PaintingStyle.stroke;

      for (Offset ponto in pontos) {
        canvas.drawRect(
            Rect.fromLTWH(ponto.dx - 12.5, ponto.dy - 12.5, 25, 25), paint);
      }
    }
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true; // Force repaint every time.
  }
}
