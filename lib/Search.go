package Search

import (
	"fmt"
	"strings"
)

func SearchPerson(users []string, name *string) (result []string, err error) {
	defer func (){
		r := recover()
		if r != nil {
			err = fmt.Errorf("%s not found", r)
		}
	}()
	found := false
	for x := range len(users) {
		if (strings.EqualFold(users[x], *name)){
			result = append(result, users[x])
			found = true
		}
	}
	if !found {
		panic(*name)
	} 
	return result, err
}