import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  _MyApp createState() => _MyApp();
}

class _MyApp extends State<MyApp> {
  int totalClicks = 0;

  void _increment(){
    setState(() {
      totalClicks++;
    });
  } 
  void _decrement() => totalClicks--;

  @override
  Widget build(BuildContext context) {
    TextStyle textStyled = TextStyle(color: Colors.white, fontSize: 25);

    return MaterialApp(
      title: 'My app',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Statefull app'),
          titleTextStyle: textStyled,
          centerTitle: true,
          backgroundColor: Colors.greenAccent,
        ),
        body: Center(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Text(
                'Total de cliques: $totalClicks',
                style: TextStyle(color: Colors.greenAccent),
              ),
              ElevatedButton(
                onPressed: _increment,
                child: Text('Clique aqui'),
                style: TextButton.styleFrom(
                  foregroundColor: Colors.greenAccent, // Cor do texto
                ),
              ),
            ],
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: _increment,
          child: Icon(Icons.add),
          backgroundColor: Colors.greenAccent,
        ),
      ),
    );

    
  }
}
