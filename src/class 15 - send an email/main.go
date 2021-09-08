package main

import (
	"log"
	"os"
	rsc "github.com/alejandrollanganate/golang-training/resources"
	"github.com/joho/godotenv"
)

func getCredentials() (email string, password string){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
  	userEmail := os.Getenv("EMAIL")
  	userPassword := os.Getenv("PASSWORD")
	return userEmail, userPassword
}

func main(){

	// In order to get email and password from .env file
	email, password := getCredentials()


	rsc.CaptureEntireScreenshot("img1")
	rsc.CaptureEntireScreenshot("img2")

	pathsImages := []string{"0_img1.png","0_img2.png"}


	rsc.SendEmail(
		email,
		password,
		email,
		"Club de Software EPN - Cybersecuriry TEAM",
		pathsImages,
	)
}
