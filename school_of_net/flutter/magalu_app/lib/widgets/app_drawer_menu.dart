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
        _itemDrawer(
            icon: const Icon(
              Icons.home,
              color: Colors.blue,
            ),
            text: 'Home'),
        _itemDrawer(
            icon: const Icon(Icons.list), text: 'Departamentos', badge: ''),
        const Divider(thickness: 2),
        _itemDrawer(icon: const Icon(Icons.favorite), text: 'Favoritos'),
        _itemDrawer(icon: const Icon(Icons.shopping_bag), text: 'Sacola'),
        _itemDrawer(
            icon: const Icon(Icons.account_circle), text: 'Minha conta'),
        const Divider(thickness: 2),
        _itemDrawer(icon: const Icon(Icons.logout), text: 'Sair'),
      ],
    );
  }

  Widget _itemDrawer(
      {required Icon icon, required String text, String badge = ''}) {
    return ListTile(
      leading: IconTheme(
        data: IconThemeData(),
        child: icon,
      ),
      title: Text(text),
      trailing: Container(
        decoration: BoxDecoration( borderRadius: BorderRadius.circular(100)),
        child: badge != ''
          ? Padding(
              padding: EdgeInsets.fromLTRB(8, 2, 8, 2),
              child: Text(
                badge,
                style: TextStyle(backgroundColor: Colors.blue),
              ),
            )
          : Icon(Icons.arrow_forward_ios, size: 10),
      ),
      onTap: () {},
    );
  }
}
