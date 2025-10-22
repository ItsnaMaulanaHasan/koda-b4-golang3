package Search

import "strings"

func SearchPerson(users []string, name string) []string {
	var result []string
	for x := range len(users) {
		if (strings.EqualFold(users[x], name)){
			result = append(result, users[x])
			return result
		}
	}
	return result
}