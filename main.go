package main

import (
	"fmt"
	"strings"
)

func searchPerson(users []string, name string) []string {
	var result []string
	var y int = 0
	for x := range users {
		if strings.Contains(users[x], name) {
			result = append(result, users[x])
			y++
		}
	}
	return result
}

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
	fmt.Println(searchPerson(users, "Clemen"))
	fmt.Println(searchPerson(users, "clemen"))
	fmt.Println(searchPerson(users, "G"))
}
