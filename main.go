package main

import (
	"fmt"
	"strings"
)

func searchPerson(users []string, name string) string {
	for x := range len(users) {
		if (strings.EqualFold(users[x], name)){
			return users[x]
		}
	}
	return "Name not found"
}

func main()  {
	users := []string{"Itsna", "Federus", "Yoga", "Fiki", "Sidik", "Ari"}
	var inputNama string
	fmt.Print("Enter the name you want to search for: ")
	fmt.Scan(&inputNama)
	result := searchPerson(users, inputNama)
	fmt.Println(result)
}