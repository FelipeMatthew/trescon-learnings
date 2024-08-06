import 'package:flutter/material.dart';

class MenuDrawer extends StatelessWidget {
  const MenuDrawer({super.key});

  @override
  Widget build(BuildContext context) {
    return ListView(
      children: <Widget>[
        SizedBox(
          height: 120,
          child: DrawerHeader(
              decoration: const BoxDecoration(color: Colors.blue),
              child: Row(
                children: [
                  Container(
                    height: 60,
                    width: 60,
                    decoration: BoxDecoration(
                        borderRadius:
                            const BorderRadius.all(Radius.circular(50)),
                        border: Border.all(color: Colors.white, width: 4.0)),
                    child: const CircleAvatar(
                      backgroundColor: Colors.blue,
                      foregroundColor: Colors.white,
                      child: Icon(Icons.people),
                    ),
                  ),
                  const Padding(
                    padding: EdgeInsets.all(15),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text('Olá Felipe',
                            style: TextStyle(
                                color: Color.fromRGBO(255, 255, 255, 0.88),
                                fontWeight: FontWeight.bold,
                                fontSize: 20)),
                        Text('Nível avançado',
                            style: TextStyle(
                              color: Colors.white,
                            )),
                      ],
                    ),
                  )
                ],
              )),
        ),
        // Options
        _itemDrawer(Icon(Icons.home, color: Colors.blue,), 'Home'),
        _itemDrawer(Icon(Icons.list), 'Departamentos'),
        _itemDrawer(Icon(Icons.favorite), 'Favoritos'),
        _itemDrawer(Icon(Icons.shopping_bag), 'Sacola'),
        Divider(thickness: 2),
        _itemDrawer(Icon(Icons.account_circle), 'Minha conta'),
        _itemDrawer(Icon(Icons.logout), 'Sair'),
      ],
    );
  }

  Widget _itemDrawer(Icon icon, String text ) {
    return ListTile(
      leading: IconTheme(
        data: IconThemeData(),
        child: icon,
      ),
      title: Text(text),
      onTap: () {},
    );
  }
}
