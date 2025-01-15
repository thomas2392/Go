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
	fmt.Println(brad)
	fmt.Printf("%+v", brad)
}
