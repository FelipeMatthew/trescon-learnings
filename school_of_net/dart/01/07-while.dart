void main() {
  int age = 1;

  while (age < 18) {
    print("under age");
    print(age);
    age++;
  }

  do {
    print("under age");
    print(age);
    age++;
  }while (age < 18);
}