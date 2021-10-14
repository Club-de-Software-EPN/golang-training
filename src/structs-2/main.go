package main

import (
	"github.com/alejandrollanganate/golang-training/entities"
)

func main() {
	member1 := entities.Member{
		Uid: "0124ASD",
		Name: "Alejandro",
		Lastname: "Llanganate",
		Email: "alejo@gmail.com",
		Status: true,
	}
	member1.ToString()
	member1.ToAnInactiveMember()
	member1.ToString()

	team := entities.Team{
		Tid: "ASD",
		Name: "CyberSecurity",
		Leader: member1,
	}
	team.ToString()
	team.AddMember(
		entities.Member{
			Uid: "asdsad",
			Name: "Juanjo",
			Lastname: "Jaramillo",
			Email: "juanjo@gmail.com",
			Status: true,
		},
	)
	team.ToString()
}