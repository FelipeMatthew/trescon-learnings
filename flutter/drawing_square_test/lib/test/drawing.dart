import 'dart:typed_data';
import 'dart:ui' as ui;
import 'package:drawing_square_test/test/result.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter/services.dart';

class DesenhandoQuadrado extends StatefulWidget {
  const DesenhandoQuadrado({super.key});

  @override
  State<DesenhandoQuadrado> createState() => _DesenhandoQuadradoState();
}

class _DesenhandoQuadradoState extends State<DesenhandoQuadrado> {
  final GlobalKey _globalKey = GlobalKey();
  final List<Offset> _pontos = [];
  ui.Image? _image;

  @override
  void initState() {
    super.initState();
    _loadImage('assets/tubo.png');
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
    if (_image != null) {
      // Calcule as dimensões e posição da imagem
      final double imageWidth = _image!.width.toDouble();
      final double imageHeight = _image!.height.toDouble();
      final double left = (MediaQuery.of(context).size.width - imageWidth) / 2;
      final double top = (MediaQuery.of(context).size.height - imageHeight) / 2;

      // Verifique se o ponto está dentro dos limites da imagem
      if (ponto.dx >= left &&
          ponto.dx <= left + imageWidth &&
          ponto.dy >= top &&
          ponto.dy <= top + imageHeight) {
        setState(() {
          _pontos.add(ponto);
        });
      }
    }
  }

  Future<void> _navegarParaImagemEditada() async {
    final Uint8List? imageData = await _capturarImagem();
    if (imageData != null) {
      Navigator.of(context).push(
        MaterialPageRoute(
          builder: (context) => Resultado(imageData: imageData),
        ),
      );
    }
  }

  Future<Uint8List?> _capturarImagem() async {
    try {
      RenderRepaintBoundary boundary = _globalKey.currentContext!
          .findRenderObject() as RenderRepaintBoundary;
      ui.Image image = await boundary.toImage();
      ByteData? byteData =
          await image.toByteData(format: ui.ImageByteFormat.png);
      return byteData?.buffer.asUint8List();
    } catch (e) {
      print(e);
      return null;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          GestureDetector(
            onTapDown: (TapDownDetails detalhes) {
              RenderBox box = context.findRenderObject() as RenderBox;
              Offset localPosition = box.globalToLocal(detalhes.globalPosition);
              _adicionarQuadrado(localPosition);
            },
            child: RepaintBoundary(
              key: _globalKey,
              child: CustomPaintContainer(pontos: _pontos, image: _image),
            ),
          ),
          Positioned(
            bottom: 20,
            left: 20,
            right: 20,
            child: ElevatedButton(
              onPressed: _navegarParaImagemEditada,
              child: const Text('Ver Imagem'),
            ),
          ),
        ],
      ),
    );
  }
}

class CustomPaintContainer extends StatelessWidget {
  final List<Offset> pontos;
  final ui.Image? image;

  const CustomPaintContainer(
      {required this.pontos, required this.image, Key? key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return CustomPaint(
      size: Size.infinite,
      painter: QuadradoPainter(pontos: pontos, image: image),
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
    return true;
  }
}
