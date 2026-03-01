package main

import "fmt"

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	IsAvailable bool
}

type Group struct {
	Name    string
	Admin   User
	Members []User
}

func main() {
	admin := User{1, "Hansen", "Jonatan", "hansenjonatann@gmail.com", true}
	user1 := User{2, "Wesly", "", "wesly@gmail.com", true}
	user2 := User{3, "Rubin", "", "rubin@gmail.com", true}

	members := []User{user1, user2}

	group := Group{"Gamer", admin, members}
	group.displayGroup()

}

func (group Group) displayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Admin : %s", group.Admin.FirstName)
	fmt.Println("")
	fmt.Println("Members :")
	for _, member := range group.Members {
		fmt.Println(member.FirstName)
	}
}
