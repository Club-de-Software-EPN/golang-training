package main

import(
	"fmt"
	"time"
)

func showMyName(myName string){
    fmt.Println(myName)
}

func showMyLastname(myLastname string){
    fmt.Println(myLastname)
}

func main(){

    showMyName("Alejandro")
	go showMyLastname("Llanganate")
	showMyLastname("Valencia")
	
	// avoid doing this:
	time.Sleep(1*time.Second)
}