import 'package:flutter/material.dart';

class DesenhandoQuadrado extends StatefulWidget {
  const DesenhandoQuadrado({super.key});

  @override
  State<DesenhandoQuadrado> createState() => _DesenhandoQuadradoState();
}

class _DesenhandoQuadradoState extends State<DesenhandoQuadrado> {
  final List<Offset> _pontos = []; // Vai retornar uma lista de posições

  // Armazena no  _pontos os valores e posições
  void _adicionarQuadrado(Offset ponto) {
    setState(() {
      _pontos.add(ponto);
    });
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTapDown: (TapDownDetails detalhes) {
        // Pegar os detalhes do clique
        // Converte a posição para coordenadas
        RenderBox box = context.findRenderObject();
        Offset localPosition = box.globalToLocal(detalhes.globalPosition);
        _adicionarQuadrado(localPosition);
      },
      child: CustomPaint(
        size: Size.infinite,
        painter: QuadradoPainter(pontos: _pontos),
        child: Center(
          child: Image.asset(),
        ),
      ),
    );
  }
}

class QuadradoPainter extends CustomPainter {
  final List<Offset> pontos;

  QuadradoPainter({required this.pontos});

//
  @override
  void paint(Canvas canvas, Size size) {
    Paint paint = Paint()
      ..color = Colors.blue
      ..strokeWidth = 3.0
      ..style = PaintingStyle.stroke;

    // Desenha o quadrado no ponto pintado
    for (Offset ponto in pontos) {
      canvas.drawRect(
          Rect.fromLTWH(ponto.dx - 25, ponto.dy - 25, 50, 50), paint);
    }
  }

  @override
  bool shouldRepaint(QuadradoPainter oldDelegate) {
    return true;
  }
}
