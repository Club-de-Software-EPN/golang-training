package main

import(
	"fmt"
	"regexp"
)

func main(){

	// User input
	var userInput string
	fmt.Println("Enter your institutional email:")
	fmt.Scanln(&userInput)

	// Regex patterns
	validEpn := regexp.MustCompile(`^[a-z]{3,15}.[a-z]{3,15}[0-9]{0,2}(@epn.edu.ec)$`)
	validYachay := regexp.MustCompile(`^[a-z.]{4,}(@yachay.edu.ec)$`)
	validUsfq := regexp.MustCompile(`^[a-z0-9]{3,30}(@usfq.edu.ec)$`)

	// Switch with case condition
	switch {
	case validEpn.MatchString(userInput):
		fmt.Println("EPN system")
	case validYachay.MatchString(userInput):
		fmt.Println("Yachay system")
	case validUsfq.MatchString(userInput):
		fmt.Println("USFQ system")
	default:
		fmt.Println("Invalid user")
	}

	// Simple switch
	var number int64
	fmt.Print("Enter a number:")
	fmt.Scanln(&number)

	switch number {
	case 0:
		fmt.Println("You enter the number 0")
	case 3:
		fmt.Println("You enter the number 3")
	default:
		fmt.Println("You enter an other number")
	}

	// Switch with a condition
	switch number%2 {
	case 0:
		fmt.Println("Pair")
	default:
		fmt.Println("Odd")
	}

}