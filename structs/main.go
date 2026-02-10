package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}
type contactinfo struct {
	email   string
	zipcode int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.firstName = "Allen"
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
	var allen person
	var rmp person
	fmt.Printf("%+v", rmp)
	fmt.Println()
	allen = person{"Allen", "Das", contactinfo{"msaswata@gmail.com", 721607}}
	fmt.Println(allen.firstName)
	fmt.Println(allen.lastName)
	fmt.Println(allen.contact.email)
	fmt.Println(allen.contact.zipcode)
	allen.print()
	alex.print()
	fmt.Println()
	fmt.Println()

	allen.updateName("Sam").print()
	fmt.Println()
	fmt.Println()
	allen.updateNamebyreference("Sonu")
	allen.print()

}
func (p person) updateName(newFirstName string) person {
	p.firstName = newFirstName
	return p
}
func (p *person) updateNamebyreference(newFirstName string) {
	p.firstName = newFirstName

}
func (p person) print() {
	fmt.Printf("%+v", p)
}
