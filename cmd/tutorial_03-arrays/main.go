package main

import "fmt"

func main() {
	intArray := [...]int32{2, 4, 5}
	fmt.Println(intArray)

	// sliced
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)

	// Append
	intSlice = append(intSlice, 3, 5, 6, 7)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))
	fmt.Println(intSlice)

	// Map
	// make(type, len, cap)
	var intSlice3 []int32 = make([]int32, 4, 10)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam": 30}
	fmt.Println(myMap2["adam"])

	// For=
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		j++
		fmt.Println(j)
	}

	numbs := []int8{1,4,5,6,7}
	for index, value := range numbs {
		fmt.Printf("index:%d \n value: %d", index, value)
	}
}