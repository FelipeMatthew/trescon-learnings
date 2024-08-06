import 'package:flutter/material.dart';

import 'package:shopping_list/models/item.dart';

class ListScreen extends StatefulWidget {
  const ListScreen({super.key});

  @override
  State<ListScreen> createState() => _ListScreenState();
}

class _ListScreenState extends State<ListScreen> {
  List<ShoppingItem> shoppingItems = [
    ShoppingItem(title: 'Comprar feijão', isDone: true),
    ShoppingItem(title: 'Comprar Arroz'),
    ShoppingItem(title: 'Comprar cuscus'),
    ShoppingItem(title: 'Comprar mandioca'),
    ShoppingItem(title: 'Comprar jararaca'),
    ShoppingItem(title: 'Comprar ciclano'),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Lista de compras'),
        centerTitle: true,
        backgroundColor: Colors.redAccent,
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
    );
  }
}
