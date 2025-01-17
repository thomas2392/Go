package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	brad := person{
		firstName: "Brad",
		lastName:  "Pitt",
		contact: contactInfo{
			email:   "brad@gmail.com",
			zipCode: 94000,
		},
	}
	bradPointer := &brad
	bradPointer.updateName("Jhon")
	brad.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
