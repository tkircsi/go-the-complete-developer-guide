package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//  Ways for define a Struct
	// #1
	// harry := person{"Harry", "Potter", contactInfo{"potter@gmail.com", 1234}}
	// fmt.Println(harry)

	// #2
	// sherlock := person{firstName: "Scherlock", lastName: "Holmes"}
	// fmt.Println(sherlock)

	// var poirot person
	// fmt.Println(poirot)
	// fmt.Printf("%+v\n", poirot)

	// poirot.firstName = "Herkules"
	// poirot.lastName = "Poirot"
	// poirot.contactInfo.email = "poirot@gmail.com"
	// fmt.Printf("%+v\n", poirot)

	bud := person{
		firstName: "Bud",
		lastName:  "Spencer",
		contactInfo: contactInfo{
			email:   "bud@mail.hu",
			zipCode: 1234,
		},
	}
	bud.print()
	bud.updateName("Terence")
	bud.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
