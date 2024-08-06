import 'package:flutter/material.dart';

import 'package:shopping_list/models/item.dart';
import 'package:shopping_list/ui/add_item.dart';

class ListScreen extends StatefulWidget {
  const ListScreen({super.key});

  @override
  State<ListScreen> createState() => _ListScreenState();
}

class _ListScreenState extends State<ListScreen> {

  void _addItem() {
      showDialog(
        context: context, 
        builder: (BuildContext context) {
          return AddItem();
        },
      );
    }

  List<ShoppingItem> shoppingItems = [
    ShoppingItem(title: 'Comprar feijão', isDone: true),
    ShoppingItem(title: 'Comprar Arroz'),
    ShoppingItem(title: 'Comprar cuscus'),
    ShoppingItem(title: 'Comprar mandioca'),
    ShoppingItem(title: 'Comprar jararaca'),
    ShoppingItem(title: 'Comprar ciclano'),
    ShoppingItem(title: 'Comprar pipoca'),
  ];

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
        itemCount: shoppingItems.length,
        itemBuilder: (context, index) {
          final item = shoppingItems[index];

          return ListTile(
            onTap: () {},
            leading: CircleAvatar(
              backgroundColor: Colors.redAccent,
              child: IconTheme(
                child: Icon(Icons.done),
                data: IconThemeData(color: Colors.white),
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