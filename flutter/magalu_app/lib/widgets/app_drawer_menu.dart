import 'package:flutter/material.dart';
import 'package:magalu/main.dart';
import 'package:magalu/pages/bag.dart';
import 'package:magalu/pages/department.dart';
import 'package:magalu/pages/favorite.dart';
import 'package:magalu/pages/home.dart';
import 'package:magalu/pages/my_account.dart';

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
            page: const HomeMyApp(),
            context: context,
            icon: const Icon(
              Icons.home,
              color: Colors.blue,
            ),
            text: 'Home'),
        _itemDrawer(
            page: const Department(),
            context: context,
            icon: const Icon(Icons.list),
            text: 'Departamentos',
            badge: ''),
        const Divider(thickness: 2),
        _itemDrawer(
            page: const Favorite(),
            context: context,
            icon: const Icon(Icons.favorite),
            text: 'Favoritos'),
        _itemDrawer(
            page: const Bag(),
            context: context,
            icon: const Icon(Icons.shopping_bag),
            text: 'Sacola'),
        _itemDrawer(
            page: const MyAccount(),
            context: context,
            icon: const Icon(Icons.account_circle),
            text: 'Minha conta'),
        const Divider(thickness: 2),
        _itemDrawer(
            page: const HomeMyApp(),
            context: context,
            icon: const Icon(Icons.logout),
            text: 'Sair'),
      ],
    );
  }

  Widget _itemDrawer(
      {required context,
      required page,
      required Icon icon,
      required String text,
      String badge = ''}) {
    return ListTile(
      leading: IconTheme(
        data: const IconThemeData(),
        child: icon,
      ),
      title: Text(text),
      trailing: Container(
        decoration: BoxDecoration(borderRadius: BorderRadius.circular(100)),
        child: badge != ''
            ? Padding(
                padding: const EdgeInsets.fromLTRB(8, 2, 8, 2),
                child: Text(
                  badge,
                  style: const TextStyle(backgroundColor: Colors.blue),
                ),
              )
            : const Icon(Icons.arrow_forward_ios, size: 10),
      ),
      onTap: () {
        Navigator.push(context, MaterialPageRoute(builder: (context) => page));
      },
    );
  }
}
