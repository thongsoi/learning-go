package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Wattanapong",
		LastName:  "Thongsoi",
	}
	fmt.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}
