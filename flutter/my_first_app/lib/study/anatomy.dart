import 'package:flutter/material.dart';

void main() => runApp(WidgetApp());

class WidgetApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {

    Text text = Text('Exemplo');
    TextField entry = TextField();

    return MaterialApp(
      title: 'Widgets app',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Widgets App'),
          titleTextStyle: TextStyle(
            color: Colors.white,
            fontSize: 25,
          ),
          centerTitle: true,
          backgroundColor: Colors.deepPurple,
        ),
        body: Padding(
          padding: EdgeInsets.all(12.0),
          child: Column(
            children: <Widget>[
              text,
              entry,
              Text(
                'Exemplo 2', 
                style: TextStyle(color: Colors.deepPurple)
              ),
              Row(
                children: <Widget>[
                  Text('01'),
                  Text('02'),
                ],
              ),
              ElevatedButton(
                onPressed: () {},
                child: Text('Clique-me'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
