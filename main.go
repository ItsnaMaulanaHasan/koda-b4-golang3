package main

import (
	"fmt"
	Search "minitask-3-go/lib"
)

func main()  {
	users := []string{"Itsna", "Federus", "Yoga", "Fiki", "Sidik", "Ari"}
	loop := true

	for loop {
		var inputNama string
		fmt.Print("Enter the name you want to search for (0. exit): ")
		fmt.Scan(&inputNama)
		if (inputNama != "0"){
			result := Search.SearchPerson(users, inputNama)
			fmt.Println(result)
		} else {
			loop = false
		}
	}
}