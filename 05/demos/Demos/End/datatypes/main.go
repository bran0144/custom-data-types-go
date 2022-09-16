package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	name1 := Name{First: "", Last: ""}
	//name2 := Name{First: "James", Last: "Wilson"}

	//ssn := organization.NewSocialSecurityNumber("123-45-6789")
	//eu := organization.NewEuropeanUnionIdentifier("12345", "France")
	//eu2 := organization.NewEuropeanUnionIdentifier("12345", "France")

	if name1.Equals(Name{}) {
		println("We match")
	}
}

type Name struct {
	First string
	Last string
	Middle []string
}

func (n Name) Equals(otherName Name) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

type OtherName struct {
	First string
	Last string
}