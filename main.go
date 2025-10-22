package main

import (
	"fmt"
	Search "minitask-3-go/lib"
)

func main()  {
	defer fmt.Println("Exit")
	users := []string{"Itsna", "Federus", "Yoga", "Fiki", "Sidik", "Ari"}
	loop := true

	for loop {
		var inputNama string
		fmt.Print("Enter the name you want to search for (0. exit): ")
		fmt.Scan(&inputNama)
		if (inputNama != "0"){
			result, err := Search.SearchPerson(users, &inputNama)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(result)
			}
		} else {
			loop = false
		}
	}
}