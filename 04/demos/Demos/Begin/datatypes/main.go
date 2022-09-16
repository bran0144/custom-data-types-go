package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Wilson")
	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.FullName())
}
