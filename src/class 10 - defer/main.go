package main

import(
	"fmt"
)

func showName(name string){
	fmt.Println(name)
}

func showLastname(lastname string){
	fmt.Println(lastname)
}

func showAge(age int64){
	fmt.Println(age)
}

func showExample1(){
	fmt.Println("-------- Example 1 --------")
	showName("Ola")
	showLastname("Bini")
	showAge(35)
}

func showExample2(){
	fmt.Println("-------- Example 2 - defer on name --------")
	defer showName("Ola")
	showLastname("Bini")
	showAge(35)
}

func showExample3(){
	fmt.Println("-------- Example 3 - defer on name and lastname --------")
	defer showName("Ola")
	defer showLastname("Bini")
	showAge(35)
}

func showExample4(){
	defer fmt.Println("-------- Example 4 - defer on everything --------")
	defer showName("Ola")
	defer showLastname("Bini")
	defer showAge(35)
}

func main(){
	showExample1()
	showExample2()
	showExample3()
	showExample4()
}