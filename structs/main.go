package main

import "fmt"

type contactinfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	// var obj person
	// obj.firstName="Yash"
	// obj.lastName="Tyagi"
	jim := person{
		firstName: "jim",
		lastName:  "party",
		contact: contactinfo{
			email:   "jim@gmail.com",
			zipcode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
