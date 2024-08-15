import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:github_finder/models/user.dart';
import 'package:http/http.dart' as http;

class ResultsScreen extends StatelessWidget {
  final String username;

  const ResultsScreen({super.key, required this.username});

  // HTTP REQUEST! SZ
  Future<User> fetchUser() async {
    var url = Uri.https('api.github.com', '/users/$username');
    final response = await http.get(url);

    if (response.statusCode == 200) {
      // Converter o body string para map {}
      final jsonBody = json.decode(response.body);
      return User.fromJson(jsonBody);
    } else {
      throw Exception('Something got wrong');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: const Text('Request result'),
        centerTitle: true,
        backgroundColor: Colors.black,
        titleTextStyle: const TextStyle(
            color: Colors.white, fontSize: 25, fontWeight: FontWeight.bold),
        iconTheme: const IconThemeData(color: Colors.white),
      ),
      body: Center(
        child: FutureBuilder<User>(
            future: fetchUser(),
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const CircularProgressIndicator(
                  backgroundColor: Colors.white,
                );
              } else if (snapshot.hasError) {
                return Text(
                  "Error: ${snapshot.error}",
                  style: const TextStyle(color: Colors.white),
                );
              } else if (snapshot.hasData) {
                // Valida se tem dados
                final user = snapshot.data!;
                return Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      ClipRRect(
                        borderRadius: BorderRadius.circular(100),
                        child: Image.network(
                          user.avatar_url,
                          height: 200,
                        ),
                      ),
                      Container(
                        height: 20,
                      ),
                      Text(
                        user.name,
                        style: const TextStyle(
                            fontSize: 16,
                            fontWeight: FontWeight.bold,
                            color: Colors.white),
                      ),
                      Container(
                        height: 20,
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: [
                          Container(
                            padding: const EdgeInsets.all(10),
                            child: Text(
                              "Seguidores: ${user.followers}",
                              style: const TextStyle(
                                  fontSize: 16, color: Colors.white),
                            ),
                          ),
                          Text(
                            "Repositórios: ${user.public_repos}",
                            style: const TextStyle(
                                fontSize: 16, color: Colors.white),
                          ),
                        ],
                      ),
                      Container(
                        height: 20,
                      ),
                      ElevatedButton(
                          onPressed: () {
                            Navigator.pop(context);
                          },
                          child: const Text(
                            "Procurar novo perfil",
                            style: TextStyle(color: Colors.black),
                          ))
                    ],
                  ),
                );
              } else {
                return const Text(
                  "No data found",
                  style: TextStyle(color: Colors.white),
                );
              }
            }),
      ),
    );
  }
}
