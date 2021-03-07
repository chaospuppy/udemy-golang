package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) cencus() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	(*p).lastName = newName
}

func main() {
	me := person{
		firstName: "Tim",
		lastName:  "Seagren",
		contactInfo: contactInfo{
			email: "example@invalid.com",
			zip:   11111,
		},
	}
	(&me).updateName("Lewis")
	me.cencus()
}
