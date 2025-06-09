package utils

import "strings"

func SearchPerson(users []string, name string) []string {
	var result []string
	for x := range users {
		if strings.Contains(strings.ToLower(users[x]), strings.ToLower(name)) {
			result = append(result, users[x])
		}
	}
	return result
}
