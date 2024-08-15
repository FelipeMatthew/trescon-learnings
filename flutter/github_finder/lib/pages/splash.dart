import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:github_finder/pages/home.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();

    SystemChrome.setEnabledSystemUIMode(SystemUiMode.immersiveSticky);

    // Async
    _checkAuth().then((_) {
      // Close the splash screen
      Navigator.pushReplacement(
          context, MaterialPageRoute(builder: (context) => HomeScreen()));
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      body: Center(
          child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Image.asset(
            'images/github-logo.png',
            height: 400,
          ),
          const Text(
            'Bem vindo ao github finder',
            textAlign: TextAlign.center,
            style: TextStyle(
                color: Colors.white, fontWeight: FontWeight.bold, fontSize: 20),
          ),
          Container(
            height: 60,
          ),
          const CircularProgressIndicator(
            backgroundColor: Colors.white,
          ),
        ],
      )),
    );
  }

  // Verification user or login async method
  Future<void> _checkAuth() async {
    await Future.delayed(const Duration(seconds: 5));
  }
}
