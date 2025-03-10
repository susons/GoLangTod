package main

import "fmt"

func main() {
	foo("Fpoooooo")

	s := bar("Baboooons")

	fmt.Println(s)

	s, n := dogYears("asdsadsda", 21)

	fmt.Println("Dog is !!!!!!   : ", s, n)

	fmt.Println("***************- Hands on 1 Functions -*****************")

	sum := sum(1, 23, 3, 545, 56, 657, 565)

	fmt.Println("Sum of numbers is: ", sum)

	fmt.Println("***************- Hands on 1 Functions -*****************")

	defer foo1()
	bar1()

}

func foo1() {
	fmt.Println("foo1:foo1: ")
}

func bar1() {
	fmt.Println("bar1:bar1: ")
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	n := 0

	for _, v := range ii {
		n += v
	}

	return n
}

func foo(s string) {
	fmt.Println("This is FOOO", s)
}

func bar(s string) string {
	fmt.Println("They call me Mr baboon", s)

	return s
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(age, "Dog is this old in god years"), age
}
