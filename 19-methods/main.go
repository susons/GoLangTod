package main

import "fmt"

type person struct {
	first string
}

type secretA struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("I am a person: ", p.first)
}

func (sa secretA) speak() {
	fmt.Println("I am a secret agent named: ", sa.first)
}

func main() {

	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}

	sa1 := secretA{
		person: person{
			first: "Smart",
		},
		ltk: true,
	}

	sa1.speak()

	p1.speak()
	p2.speak()

	saySomething(sa1)

}
