package main

import(
	"fmt"
)

func main(){

	// User input
	var name string
	fmt.Println("- ¿What's your name?")
	fmt.Scanln(&name)
	fmt.Println("- Welcome",name)

	// Simple For
	var html string
	
	for i := 1 ; i <= 100 ; i++ {
		if i%13 == 0 {
			html += fmt.Sprintf("<li> %d. Item </li>\n", i) // multiples of thirteen
		}
	}
	fmt.Println(html)

	// Avoid infinite loops
        /* 
	for {
		fmt.Printf("Infinity\n")  
	  } 
	*/

	// While loop simulation
	var lastname string = ""
	for true {
		if lastname == "llanganate"{ 
			break 
		}else{
			fmt.Println("- ¿What's your lastname?")
			fmt.Scanln(&lastname)
		}
	}
	fmt.Printf("- Ok, Hi %s %s \n", name, lastname)

	// Factorial of 10
	var i int = 10
	factorial := 1
	for i > 0 {
		factorial *= i
		i--
	}
	fmt.Printf("!10 = %d\n", factorial) // 3628800
}
