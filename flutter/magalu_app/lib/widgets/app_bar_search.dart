import 'package:flutter/material.dart';

class AppSearchBar extends SearchDelegate<String> {
  final products = <String>[
    'Micro-ondas',
    'Fogão',
    'Televisão',
    'Armário',
    'Lava e seca',
    'Iphone',
    'Pia de marmore'
    'Macbook',
  ];

  final recentSearch = <String>[
    'Macbook',
    'Iphone',
    'Televisão',
    'Micro-ondas',
  ];

  @override
  List<Widget>? buildActions(BuildContext context) {
    return <Widget>[
      IconButton(
          onPressed: () {
            query = ''; // Limpa resultado de pesquisa
            close(context, ''); // Fecha a pagina aberta e volta para ultima
          },
          icon: const Icon(Icons.cancel))
    ];
  }

  @override
  Widget? buildLeading(BuildContext context) {
    // Ação de voltar
    return IconButton(
        onPressed: () {
          query = '';
          close(context, '');
        },
        icon: AnimatedIcon(
          icon: AnimatedIcons.menu_arrow,
          progress: transitionAnimation,
        ));
  }

  @override
  Widget buildResults(BuildContext context) {
    // Fazer a parte da pesquisa
    return Center(
      child: Text('Resultado para pesquisa: $query'),
    );
  }

  @override
  Widget buildSuggestions(BuildContext context) {
    final results = query.isEmpty ? recentSearch : products;

    // Ultimos itens pesquisados
    return ListView.builder(
        itemCount: results.length,
        itemBuilder: (context, index) {
          return ListTile(
            onTap: () {
              // TODO
              
              query = '';
              close(context, '');
            },
            leading: const Icon(Icons.youtube_searched_for),
            title: Text(results[index]),
          );
        });
  }
}

// * Query serve para pegar o que foi escrito na ferramenta de pesquisa automaticamente.