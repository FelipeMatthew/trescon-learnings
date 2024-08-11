import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:github_finder/models/user.dart';
import 'package:http/http.dart' as http;

class ResultsScreen extends StatelessWidget {
  const ResultsScreen({super.key});

  // HTTP REQUEST! SZ
  Future<User> fetchUser() async {
    var url = Uri.https('api.github.com', '/users/felipematthew');
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
        child: Text('Hii'),
      ),
    );
  }
}
