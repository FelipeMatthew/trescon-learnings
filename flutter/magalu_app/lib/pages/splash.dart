import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:magalu/pages/home.dart';

class Splash extends StatefulWidget {
  const Splash({super.key});

  @override
  State<Splash> createState() => _SplashState();
}

class _SplashState extends State<Splash> {
  @override
  void initState() {
    super.initState();

    SystemChrome.setEnabledSystemUIMode(SystemUiMode.immersiveSticky);

    Future.delayed(const Duration(seconds: 5)).then((_) {
      // Close the splash screen
      Navigator.pushReplacement(
          context, MaterialPageRoute(builder: (context) => const HomeMyApp()));
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Container(
            color: Colors.white,
            child: Padding(
              padding: const EdgeInsets.all(20),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Image.asset('images/magalu-logo-1.png'),
                  Container(
                    height: 50,
                  ),
                  const CircularProgressIndicator(
                    color: Colors.blue,
                  ),
                  Container(
                    height: 50,
                  ),
                  const Text(
                    'Carregando...',
                    style: TextStyle(
                        color: Colors.blue,
                        fontSize: 20,
                        fontWeight: FontWeight.bold),
                  )
                ],
              ),
            )));
  }
}
