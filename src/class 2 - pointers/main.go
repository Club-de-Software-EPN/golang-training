package main

import(
	"fmt"
)

func main(){

	// Pointers
	fmt.Println("--- Pointers ---")

	var pointer1, pointer2 *int
	fmt.Printf("pointer1: %p, pointer2: %p\n", pointer1, pointer2) // 0x0 0x0

	number := 2
	fmt.Println("number value:", number) // 2 - variable value
	fmt.Println("number address:", &number) // addr - variable address

	// same addresses and pointer values
	pointer1 = &number
	pointer2 = pointer1 
	fmt.Printf("number address: %p, pointer1: %p, pointer2: %p\n",&number, pointer1, pointer2) // addr == addr == addr
	fmt.Printf("number value: %d, pointer1 value: %d, pointer2 value: %d\n", number, *pointer1, *pointer2) // 2 2 2

	*pointer1 = 105 //  * operator denotes the pointer's value
	fmt.Printf("number address: %p, pointer1: %p, pointer2: %p\n",&number, pointer1, pointer2) // 105 105 105

	// modify one pointer entity
	pointer2 = nil // 0x0 
	fmt.Printf("number value: %d, pointer1 value: %d, pointer2 value: %p\n", &number, pointer1, pointer2) // addr addr 0

}