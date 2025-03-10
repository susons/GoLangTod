package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	favIc []string
}

func main() {
	p1 := person{
		first: "1",
		last:  "last",
		favIc: []string{"banana", "candies"},
	}

	p2 := person{
		first: "Jennuy",
		last:  "Wooodhouse",
		favIc: []string{"icecream", "cola"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIc)
	fmt.Println(p2.favIc)

	for i, v := range p1.favIc {
		fmt.Println(i, ": ", p1.first, " Favorite foood is: ", v)
	}
	for i, v := range p2.favIc {
		fmt.Println(i, ": ", p2.first, " Favorite foood is: ", v)
	}

	fmt.Println("***************- Hands on 2 -*****************")
	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for _, v := range m {
		fmt.Println(v)

		for _, v2 := range v.favIc {
			fmt.Println("The fav for ", v.first, " is ", v2)
		}
	}

	fmt.Println("***************- Hands on 3 -*****************")

	v1 := vehicle{
		engine: engine{electric: true},
		make:   "ford",
		model:  "mustang",
		doors:  4,
		color:  "blue",
	}
	v2 := vehicle{
		engine: engine{electric: false},
		make:   "honda",
		model:  "vasabi",
		doors:  2,
		color:  "red",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println("***************- Hands on 4 -*****************")

	v3 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:   "weqweeqeq",
		friends: map[string]int{"catston": 12, "booongaaas": 30},
		favDrinks: []string{
			"bananas", "carrots",
		},
	}

	for k, v := range v3.friends {
		fmt.Println(p1.first, " - friends - ", k, v)
	}

	for _, v := range v3.favDrinks {
		fmt.Println(p1.first, " - fav drinks - ", v)
	}

}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}
