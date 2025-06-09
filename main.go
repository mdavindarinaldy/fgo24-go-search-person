package main

import (
	"fgo24-go-search-person/utils"
	"fmt"
)

func main() {
	users := []string{
		"Leanne Graham",
		"Ervin Howell",
		"Clementine Bauch",
		"Patricia Lebsack",
		"Chelsey Dietrich",
		"Mrs. Dennis Schulist",
		"Kurtis Weissnat",
		"Nicholas Runolfsdottir V",
		"Glena Reichert",
		"Clementina DuBuque",
	}
	fmt.Println(utils.SearchPerson(users, "Clemen"))
	fmt.Println(utils.SearchPerson(users, "clemen"))
	fmt.Println(utils.SearchPerson(users, "G"))
}
