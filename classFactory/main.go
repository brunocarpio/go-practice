package main

import (
	"classFactory/appliances"
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter the preferred appliance")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")

	var myType int

	_, err := fmt.Scan(&myType)
	if err != nil {
		fmt.Println(errors.New("Error while scanning type option"))
	}

	myAppliance, err := appliances.CreateAppliance(myType)
	if err != nil {
		fmt.Println(err)
	} else {
		myAppliance.Start()
		fmt.Printf("myAppliance.GetPurpose(): %v\n", myAppliance.GetPurpose())
	}
}
