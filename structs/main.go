package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.firstName = "Allen"
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
	var allen person
	var rmp person
	fmt.Printf("%+v", rmp)
	allen = person{"Allen", "Das"}
	fmt.Println(allen)
}
