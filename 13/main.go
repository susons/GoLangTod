package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Init func(): !")
}

func main() {
	fmt.Println("start: !")

	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("Numbers: x=%v and y=%v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println(" <4")
	} else if x > 6 && y > 6 {
		fmt.Println("< 6")
	} else {
		fmt.Println("This is else")
	}

	switch {
	case x < 4 && y < 4:
		fmt.Println("'Switch'  <4")
	case x > 6 && y > 6:
		fmt.Println("'Switch' < 6")
	default:
		fmt.Println("'Switch' This is else")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("numbers: %v \n", i)
	}

	i := 20

	for i > 0 {
		fmt.Printf("for loop: %v \n", i)
		i--
	}

	i = 5

	for {
		if i == 10 {
			fmt.Printf("Break was a break \n")
			break
		}

		i++
	}

	for i = 0; i < 20; i++ {
		if i%2 != 0 {
			fmt.Printf("not round number! %v\n", i)
		}
	}

	fmt.Println("***************FOR*****************")

	xi := []int{21, 32, 453, 4545, 454, 123, 234}

	for i1, v1 := range xi {
		fmt.Printf("this is a number index %v nad this is value %v.\n", i1, v1)
	}

	fmt.Println("***************MAP*****************")
	m := map[string]int{
		"james": 42,
		"money": 12,
	}

	for key, value := range m {
		fmt.Printf("key: %v, value: %v", key, value)
	}

	fmt.Println("***************COMMA, COMMA OK IDIOM*****************")

	age := m["james"]

	q := m["Q"]

	money := m["money"]

	fmt.Printf("Age: %v, Q: %v, money: %v \n", age, q, money)

	if v, ok := m["Q"]; !ok {
		fmt.Println("Q doeant exist", v)
	}

	if v, ok := m["james"]; ok {
		fmt.Println("james does exist Booom", v)
	}

	fmt.Println("***************STATEMENT, STATEMENT IDIOM*****************")

	for i := 0; i < 20; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("X is equals to 3, x == 3")
		}
	}
}
