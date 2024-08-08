import 'dart:math';

import 'package:flutter/material.dart';

class ListScreen extends StatefulWidget {
  const ListScreen({super.key});

  @override
  State<ListScreen> createState() => _ListScreenState();
}

class _ListScreenState extends State<ListScreen> {

  void _addToDo() {
    setState(() {
      int random = new Random().nextInt(100);
      todo.add('Task $random');
    });
  }

  List<String> todo = [
    'To do 1',
    'To do 2',
    'To do 3',
    'To do 4',
    'To do 5',
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Dynamic list'),
        backgroundColor: Colors.pinkAccent,
        centerTitle: true,
      ),
      body: ListView.separated(
        separatorBuilder: (context, index) => Divider(
          color: Colors.pinkAccent,
        ),
        itemCount: todo.length,
        itemBuilder: (BuildContext context, index) {
          return ListTile(
            onTap: () {},
            leading: const CircleAvatar(
              child: IconTheme(
              child: Icon(Icons.ac_unit),
              data: IconThemeData(
                color: Colors.white,
              ), // Icon color
            ),
            backgroundColor: Colors.pinkAccent,
            ),
            title: Text(
              todo[index],
              style: TextStyle(color: Colors.pinkAccent),
            ),
          );
        },
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.pinkAccent,
        child: Icon(Icons.add),
        onPressed: _addToDo,
      ),
    );
  }
}
