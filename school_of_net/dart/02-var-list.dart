void main() {
  List groceryList = [1,2,3,4,5,6,7,8,9,10, "anything", 49.60, false, true];
  List user = ["Felipe", 20, 87.60];

  Map dog = {
    "name": "felipe",
    "age": 30,
    "weight": 87.0,
    39: "trinta e nove",
  };

  print(groceryList[4]);
  print(groceryList.length);

  print(user);
  print(user.length);

  print(dog["name"]);
  }