package entities

import (
	"fmt"
	"log"
)

type Team struct {
	Tid string
	Name string
	Leader Member
	members []Member	
}

func (t *Team) AddMember(newMember Member){
	if(newMember.Status == true){
		t.members = append(t.members, newMember)
	}else{
		log.Fatalln("Imposible to add a inactive member to a team")
	}
}

func (t *Team) ToString(){
	fmt.Printf("*********** %s Team *************\n", t.Name)
	fmt.Println("--> Leader:")
	t.Leader.ToString()
	fmt.Println("--> Total members",len(t.members))
}