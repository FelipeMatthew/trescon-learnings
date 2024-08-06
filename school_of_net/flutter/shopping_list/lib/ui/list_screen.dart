import 'package:flutter/material.dart';

import 'package:shopping_list/models/item.dart';
import 'package:shopping_list/ui/add_item.dart';

class ListScreen extends StatefulWidget {
  const ListScreen({super.key});

  @override
  State<ListScreen> createState() => _ListScreenState();
}

class _ListScreenState extends State<ListScreen> {
  void _addItem() async {
    final item = await showDialog<ShoppingItem>(
      context: context,
      builder: (BuildContext context) {
        return AddItem();
      },
    );

    if (item != null) {
      setState(() {
        myItems.add(item);
      });
    }
  }

  List<ShoppingItem> myItems = [];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Lista de compras'),
        centerTitle: true,
        backgroundColor: Colors.redAccent,
        titleTextStyle: const TextStyle(color: Colors.white, fontSize: 25.0),
      ),
      body: ListView.separated(
        separatorBuilder: (context, index) => const Divider(
          color: Colors.redAccent,
        ),
        itemCount: myItems.length,
        itemBuilder: (context, index) {
          final item = myItems[index];

          return ListTile(
            onTap: () {
              setState(() {
                myItems[index].isDone = !myItems[index].isDone;
              });
            },
            onLongPress: () {}, // Segurar
            leading: CircleAvatar(
              backgroundColor: Colors.redAccent,
              child: IconTheme(
                data: const IconThemeData(color: Colors.white),
                child: Icon(item.isDone ? Icons.done_all : Icons.done),
              ),
            ),
            title: Text(
              item.title,
              style: const TextStyle(color: Colors.redAccent),
            ),
          );
        },
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.redAccent,
        child: const Icon(
          Icons.add,
          color: Colors.white,
        ),
        onPressed: _addItem,
      ),
    );
  }
}
