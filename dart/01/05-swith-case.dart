void main() {
  String action = "open";

  switch (action) {
    case "open":
      print("the door is open");
      break;
    case "close":
      print("the door is close");
      break;
    case "locked":
      print("the door is locked");
      break;
    default: 
      print("if nothing of this happens im gona run this text");
  }
}