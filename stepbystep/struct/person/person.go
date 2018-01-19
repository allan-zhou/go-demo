package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upperPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var person11 Person
	person11.firstName = "jialiang"
	person11.lastName = "zhou"
	upperPerson(&person11)
	fmt.Printf("%s full name is: %s %s \n\n", "person11", person11.firstName, person11.lastName)

	person12 := Person{"jialiang","zhou"}
	upperPerson(&person12)
	fmt.Printf("%s full name is: %s %s \n\n", "person12", person12.firstName, person12.lastName)

	var person13 Person = Person{"jialiang", "zhou"}
	upperPerson(&person13)
	fmt.Printf("%s full name is: %s %s \n\n", "person13", person13.firstName, person13.lastName)

	person21 := new(Person)
	person21.firstName = "jialiang"
	(*person21).lastName = "zhou"   //合法
	upperPerson(person21)
	fmt.Printf("%s full name is: %s %s \n\n", "person21", person21.firstName, person21.lastName)

	person22 := &Person{"jialiang", "zhou"}
	upperPerson(person22)
	fmt.Printf("%s full name is: %s %s \n\n", "person22", person22.firstName, person22.lastName)

}
