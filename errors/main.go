package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errorCrewNotFound = errors.New("No member found in crew.")

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 9,
	"Raul":  10,
}

type findError struct {
	name, server, msg string
}

func (e findError) Error() string {
	return e.msg
}

func findSC(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	// if v, ok := scMapping[name]; !ok {
	// 	return -1, errorCrewNotFound
	// } else {
	// 	return v, nil
	// }

	// if v, ok := scMapping[name]; !ok {
	// 	return -1, errorCrewNotFound
	// } else {
	// 	return v, nil
	// }

	// if v, ok := scMapping[name]; !ok {
	// 	return -1, fmt.Errorf("Crew member %s could not be found on server '%s'", name, server)
	// } else {
	// 	return v, nil
	// }

	if v, ok := scMapping[name]; !ok {
		return -1, findError{name, server, "Crew member not found."}
	} else {
		return v, nil
	}

}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("A panic recovered", err)
	// 	}
	// }()

	// if sc, err := findSC("Bob", "Server1"); err != nil {
	// 	fmt.Println("Error ocurred by searching for clearance level.", err)
	// } else {
	// 	fmt.Println("Clearance level found:", sc)
	// }

	// _, err := findSC("Bob", "Server1")
	// if err == errorCrewNotFound {
	// 	fmt.Println("Confirmed error is crew not found!")
	// }

	if sc, err := findSC("Bob", "Server1"); err != nil {
		fmt.Println("Error occurred while searching.", err)
		if v, ok := err.(findError); ok {
			fmt.Println("Server name:", v.server)
			fmt.Println("Crew member name:", v.name)
		}
	} else {
		fmt.Println("Crew member has security clearance", sc)
	}
}
