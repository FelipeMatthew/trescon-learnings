void main() {
  hello("felipe ");
  hello("felipe ", lastName: "azevedo");
  helloOrdenado("Felipe ", "Matthew ", 30);
  helloOrdenado("Felipe ");
}

// Opcional parameter
void hello(String name, { String lastName = "matthew" }) {
  print(name + lastName);
}

// Optional ordenado nao precisa passr o nome mas sim o valor na ordem
void helloOrdenado(String name, [String lastName = "da silva ", int age = 0]) {
   print(name + lastName + age.toString());
}