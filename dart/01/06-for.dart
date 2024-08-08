void main() {
  List myList = [1,2,3,4,5,6,7,8,9,10, "anything", 49.60, false, true];

  for (int i = 0; i < 10; i++) {
    print(i);
  }
  for (int i = 0; i < myList.length; i++) {
    print(myList[i]);
  }
  // Range or forEach
  for (var list in myList) {
    print(list);
  }
}