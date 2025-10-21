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
	loop := true

	for loop {
		var inputNama string
		fmt.Print("Enter the name you want to search for (0. exit): ")
		fmt.Scan(&inputNama)
		if (inputNama != "0"){
			result := searchPerson(users, inputNama)
			fmt.Println(result)
		} else {
			loop = false
		}
	}
}