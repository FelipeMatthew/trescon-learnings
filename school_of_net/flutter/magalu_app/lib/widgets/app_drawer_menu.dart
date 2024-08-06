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
        )
      ],
    );
  }
}
