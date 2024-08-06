import 'package:flutter/material.dart';
import 'package:shopping_list/models/item.dart';

class AddItem extends StatelessWidget {
  // Pega valor de um textfield
  final itemC = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text('Adicionar item'),
      content: TextField(
        autofocus: true,
        controller: itemC,
      ),
      actions: <Widget>[
        ElevatedButton(
          onPressed: Navigator.of(context).pop, // Vai fechar o modal
          style: ElevatedButton.styleFrom(foregroundColor: Colors.redAccent),
          child: const Text('Cancelar'),
        ),
        ElevatedButton(
          style: ElevatedButton.styleFrom(foregroundColor: Colors.redAccent),
          child: const Text('Adicionar'),
          onPressed: () {
            final item = ShoppingItem(title: itemC.value.text); // Pega o texto exato
            itemC.clear();

            Navigator.of(context).pop(item);
          },
        ),
      ],
    );
  }
}
