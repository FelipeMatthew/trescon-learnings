import 'package:flutter/material.dart';

class ListScreen extends StatefulWidget {
  const ListScreen({super.key});

  @override
  State<ListScreen> createState() => _ListScreenState();
}

class _ListScreenState extends State<ListScreen> {
  List<String> shoppingItems = ['Comprar feijão', 'Comprar arroz'];

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
              shoppingItems[index],
              style: const TextStyle(color: Colors.redAccent),
            ),
          );
        },
      ),
    );
  }
}
