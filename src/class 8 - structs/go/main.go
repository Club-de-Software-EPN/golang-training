package main

import("fmt")

type ClubMember struct {
	id int64
	fullname string
	email string
	team string
}

func (member *ClubMember) toString() {
	fmt.Printf("Name --> %s\nEmail --> %s\nTeam --> %s",member.fullname, member.email, member.team)
}

func main() {
	securityMember := ClubMember{
		id: 0,
		fullname: "Alejandro Llanganate",
		email: "alealejandro@gmail.com",
		team: "CS",
	};

	fmt.Println(securityMember);
	securityMember.toString()
}