package main

import (
	"encoding/json" // the file will have a json format
	"fmt"
	"os" // Create files
	"time"
	"io/ioutil"
)

// Using OS
func saveDataOnJSONFile(data ClubMember, fileName string){
	// Create file
	membersFile, err := os.Create(fmt.Sprintf("%s.json", fileName))
	// Close file
	defer membersFile.Close()
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "		") // data, prefix, indent string
	if err != nil {
		panic(err)
	}
	membersFile.WriteString(string(jsonData))
}

func saveDataOnJSONFileUsingIoutil(data ClubMember, fileName string){
	jsonData, err := json.MarshalIndent(data, "", "		") // data, prefix, indent string
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s.json", fileName), jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

type ClubMember struct {
	Id int64 `json:"idEstudiante"`
	Name string `json:"nombre"`
	Email string `json:"correo"`
	Created time.Time
}

func main(){
	// data
	member1 := ClubMember{
		Id: 0,
		Name: "Alejandro Llanganate",
		Email: "example@gmail.com",
		Created: time.Now(),
	}
	member2 := ClubMember{
		Id: 1,
		Name: "Linus Torvalds",
		Email: "lntvs@gmail.com",
		Created: time.Now(),
	}

	saveDataOnJSONFile(member1, "exampleUsingOS")
	saveDataOnJSONFileUsingIoutil(member2, "exampleUsingIoutil")
}