package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type agent struct {
	person
	ltk bool
}

func main() {
	sa1 := agent{
		person: person{
			first: "james",
			last:  "bond",
			age:   40,
		},
		ltk: true,
	}

	p1 := person{
		first: "first",
		last:  "last",
		age:   21,
	}

	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "batman",
		last:  "coooook",
		age:   28,
	}

	fmt.Println(p2)
	fmt.Printf("%T \t %#v \t %v\n", p1, p1, p1)
	fmt.Printf("%T \t %#v \t %v\n", sa1, sa1, sa1)
}
