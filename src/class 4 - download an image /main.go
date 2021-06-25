package main

import(
	"fmt"
	"os/exec"
)

func main(){

		// Download an image
		const img string = "logo-EPN-blanco.png"
		const url string = "https://www.epn.edu.ec/wp-content/uploads/2019/02/"+img
	
		output, err := exec.Command("ls", img).CombinedOutput() // returns a binary and an standard error
	
		if(err != nil){ // the logo was not download before so it doesn't exist in the actual directory
			output, err := exec.Command("wget", url).CombinedOutput() // we download EPN's logo
			if(err == nil){
				fmt.Println(string(output))
			}
		}else{
			fmt.Println("You have previously downloaded", string(output))
		}
}