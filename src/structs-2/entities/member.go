package entities

import "fmt"

type Member struct {
	Uid string
	Name string
	Lastname string
	Email string
	Status bool // true active ; false inactive
}

func (m *Member) ToAnInactiveMember(){
	m.Status = false
}

func (m Member) ToString(){
	var statusTag string
	if m.Status == true {
		statusTag = "Activo"
	}else{
		statusTag = "Inactivo"
	}
	fmt.Printf("---------- Club Member - %s --------\n",statusTag)
	fmt.Printf("%s %s (%s)\n", m.Name, m.Lastname, m.Email)
}