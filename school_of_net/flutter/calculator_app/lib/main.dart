import 'package:flutter/material.dart';

void main() => runApp(MyApp());

// Vai ser dinamico
class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  String result = "Resultado da soma: ";
  double num1 = 0.0;
  double num2 = 0.0;

  calculateValue() {
    setState(() {
      double cal = num1 + num2;
      this.result = "the sum result is: ${cal.toStringAsFixed(0)}";
    });
  }

  @override
  Widget build(BuildContext context) {
    // Component
    TextField numberOne = TextField(
      // input type from keyboard
      keyboardType: TextInputType.number,
      style: TextStyle(color: Colors.black),
      decoration: InputDecoration(
        labelText: 'Number one',
        labelStyle: TextStyle(
          color: Colors.deepPurpleAccent,
        ),
        enabledBorder: OutlineInputBorder(
            borderSide: BorderSide(color: Colors.deepPurpleAccent, width: 2.0)),
        focusedBorder: OutlineInputBorder(
            borderSide: BorderSide(color: Colors.deepPurpleAccent, width: 2.0)),
      ),
      onChanged: (value) => num1 = double.parse(value),
    );

    TextField numberTwo = TextField(
      // input type from keyboard
      keyboardType: TextInputType.number,
      style: TextStyle(color: Colors.black),
      decoration: InputDecoration(
        labelText: 'Number two',
        labelStyle: TextStyle(
          color: Colors.deepPurpleAccent,
        ),
        enabledBorder: OutlineInputBorder(
            borderSide: BorderSide(color: Colors.deepPurpleAccent, width: 2.0)),
        focusedBorder: OutlineInputBorder(
            borderSide: BorderSide(color: Colors.deepPurpleAccent, width: 2.0)),
      ),
      onChanged: (value) => num2 = double.parse(value),
    );

    ElevatedButton button = ElevatedButton(
      onPressed: calculateValue,
      child: Text('Calcular'),
      style: ElevatedButton.styleFrom(
        backgroundColor: Colors.deepPurpleAccent[400], // Cor de fundo roxo
        foregroundColor: Colors.white, // Cor do texto branco
      ),
    );

    Text result = Text(
      this.result,
      style: TextStyle(
        color: Colors.deepPurpleAccent[400],
      ),
    );

    Padding separator = Padding(padding: EdgeInsets.all(4.5));

    Column columns = Column(
      children: <Widget>[
        numberOne,
        separator,
        numberTwo,
        separator,
        SizedBox(
          child: button,
          width: double.infinity,
        ),
        separator,
        result,
      ],
    );

    return MaterialApp(
      title: 'Felipe Calculator',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Calculator'),
          centerTitle: true,
          backgroundColor: Colors.deepPurpleAccent[400],
          titleTextStyle: TextStyle(
            color: Colors.white,
            fontSize: 25.0,
          ),
        ),
        body: Padding(
          padding: EdgeInsets.all(15.0),
          child: columns,
        ),
      ),
    );
  }
}
