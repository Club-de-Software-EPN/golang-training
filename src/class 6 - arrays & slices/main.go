package main

import(
	"fmt"
)

func main(){

	// Arrays
	var careers = [2]string{"software engineering","computer science"}
	var language = [...]string{"g","o","l","a","n","g"} 

	fmt.Printf("Array: %s, Type: %T\n",careers, careers)
	fmt.Printf("Array: %s, Type: %T\n",language, language)

	// Fill an array
	numbers := [5]int{}
	fmt.Println("Before: ", numbers)
	numbers[0] = 10
	numbers[4] = 50
	fmt.Println("After: ", numbers)

	// Slice & append()
	simpleSlice := make([]int, 2, 5) // len can't be larger than cap in make
	fmt.Printf("Original simple slice: %d. Type: %T\n", simpleSlice, simpleSlice)
	newSlice := append(simpleSlice, 10)
	fmt.Println("Original simple slice after append(simpleSlice):", simpleSlice)
	fmt.Println("New simple slice after append(simpleSlice):", newSlice)

	// len(), cap()
	fmt.Printf("Length -> simpleSlice: %d, Capacity -> simpleSlice: %d\n", len(simpleSlice), cap(simpleSlice))
	fmt.Printf("Length -> newSlice: %d, Capacity -> newSlice: %d\n", len(newSlice), cap(newSlice))
	
	// Slice literals
	var sliceLiteral = []string{"a", "e", "i", "o", "u"}
	fmt.Printf("Slice literal: %s. Type: %T\n", sliceLiteral, sliceLiteral)

	// Slices from an array
	array := [5]float64{12.5, 10.8, 0.5, 7.6, 14.1}
	slice1 := array[2:] 
	slice2 := array[:1]
	slice3 := array[len(array)-1:]
	slice4 := array[1:4]

	fmt.Println("Slice 1:", slice1) // [0.5, 7.6, 14.1]
	fmt.Println("Slice 2:", slice2) // [12.5]
	fmt.Println("Slice 3:", slice3) // [14.1]
	fmt.Println("Slice 4:", slice4) // [10.8, 0.5, 7.6]

	// Iterate an array
	fmt.Println("Vowels:", sliceLiteral)
	for index, element := range sliceLiteral { // index and vowels
		fmt.Printf("Index: %d, Element: %s\n", index, element)
	}
	for index := range sliceLiteral { // only the index
		fmt.Printf("Index: %d\n", index)
	}
	for index := range sliceLiteral { // only vowels
		fmt.Printf("Vowel: %s\n", sliceLiteral[index])
	}
}