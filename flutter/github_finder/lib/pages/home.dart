import 'package:flutter/material.dart';
import 'package:github_finder/pages/results.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        resizeToAvoidBottomInset:
            false, // Vai evitar erro do input ficar por de trás do teclado
        body: Column(
          children: [
            Container(
                width: double.maxFinite,
                decoration: const BoxDecoration(
                    color: Colors.black,
                    borderRadius:
                        BorderRadius.only(bottomLeft: Radius.circular(100))),
                child: Padding(
                  padding: const EdgeInsets.all(70),
                  child: Image.asset(
                    'images/github-logo.png',
                    height: 200,
                  ),
                )),
            Padding(
              padding: const EdgeInsets.all(40),
              child: Column(children: [
                const Text(
                  'Github Finder',
                  style: TextStyle(
                      color: Colors.black,
                      fontSize: 30,
                      fontWeight: FontWeight.bold),
                ),
                const Text('Insira o usuário que deseja encontrar',
                    style: TextStyle(
                      color: Colors.black,
                      fontSize: 18,
                    )),
                Container(
                  height: 30,
                ),
                const TextField(
                  decoration: InputDecoration(
                      border: OutlineInputBorder(
                          borderSide:
                              BorderSide(color: Colors.black, width: 1.0)),
                      focusedBorder: OutlineInputBorder(
                          borderSide:
                              BorderSide(color: Colors.black, width: 1.0)),
                      hintText: 'Username',
                      labelText: 'Username',
                      labelStyle: TextStyle(color: Colors.black)),
                ),
                Container(
                  height: 30,
                ),
                SizedBox(
                    width: double.maxFinite,
                    child: ElevatedButton(
                        onPressed: () {
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => const ResultsScreen()));
                        },
                        style: const ButtonStyle(
                            padding: WidgetStatePropertyAll<EdgeInsetsGeometry>(
                                EdgeInsets.fromLTRB(40, 10, 40, 10)),
                            backgroundColor:
                                WidgetStatePropertyAll<Color>(Colors.black)),
                        child: const Text(
                          'Buscar',
                          style: TextStyle(color: Colors.white, fontSize: 20),
                        )))
              ]),
            )
          ],
        ));
  }
}
