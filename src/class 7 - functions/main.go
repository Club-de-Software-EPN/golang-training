package main

import (
	"fmt"
	"time"
)

// Functions that take parameters

func combineWords1(word1 string, word2 string) string { // combineWords1() is equivalent to combineWords2()
	return fmt.Sprintf("1. %s %s\n", word1, word2)
}

func combineWords2(word1, word2 string) string {
	return fmt.Sprintf("2. %s %s\n", word1, word2)
}

// Multiple results

func combineWords3(word1, word2 string) (int, string) {
	return 3, fmt.Sprintf("%s %s\n", word1, word2)
}

func combineWords4(word1, word2 string) (number int, text string) { // named return values
	number = 4
	text = fmt.Sprintf("%s %s\n", word1, word2)
	return number, text
}

// Functions without arguments

func getActualDate(dateFormat string) string{
	// time package uses this date formate as a reference: Mon Jan 2 15:04:05 MST 2006
	currentTime := time.Now()
	return  currentTime.Format(dateFormat)
}

// fun main must have no arguments
func main(){
	fmt.Println(combineWords1("Alejandro", "Llanganate")) // 1. Alejandro Llanganate
	fmt.Println(combineWords2("Software", "Club")) // 2. Software Club
	fmt.Println(combineWords3("Cyber Security", "Team")) // 3 Cyber Security Team
	fmt.Println(combineWords4("Ola", "Bini")) // 4 Ola bini
	fmt.Println(getActualDate("2006-01-02")) // 2021-07-01
	fmt.Println(getActualDate("02/01/2006")) //01-07-2021
	fmt.Println(getActualDate("Day: Mon Month: Jan Date: 02/01/2006 Time: 15h:04min")) //01-07-2021
}

