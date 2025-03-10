package main

import "fmt"

func main() {
	fmt.Println("***************- Array -*****************")

	fmt.Println("brrrrrr datatypes start")
	a := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Printf("%T\n", a)

	fmt.Println("***************- Slices -*****************")

	s := []string{"hello potatoes", "Bazookas"}
	fmt.Println(s)

	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	for index, value := range xs {
		fmt.Printf("index: \t%v, value: \t%v\n", index, value)
	}

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	fmt.Println("***************- Slices Append -*****************")

	xi := []int{42, 43, 44}

	fmt.Println(xi)
	xi = append(xi, 0, 9, 8, 7)
	fmt.Printf("xi is now: %v\n", xi)

	fmt.Println("***************- Slices Slice Slice -*****************")

	xi = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("slice start: ", xi)

	//inclusive:exclusive
	fmt.Printf("xi (inclusive:exclusive) - %#v\n", xi[0:4])

	// [exclusive]
	fmt.Printf("xi :exclusive - %#v\n", xi[:7])

	//inclusive:
	fmt.Printf("xi inclusive - %#v\n", xi[4:])

	//everything
	fmt.Printf("xi inclusive - %#v\n", xi[:])

	fmt.Println("***************- Slices delete -*****************")

	xi = append(xi[:4], xi[5:]...)

	fmt.Printf("xi inclusive - %#v\n", xi)

	fmt.Println("***************- Slices MAKE -*****************")

	xi = make([]int, 0, 10)

	fmt.Println(xi)
	fmt.Println("length: ", len(xi))
	fmt.Println("cap: ", cap(xi))

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("length: ", len(xi))
	fmt.Println("cap: ", cap(xi))

	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println("length: ", len(xi))
	fmt.Println("cap: ", cap(xi))

	fmt.Println("***************- Multi dimension Slices -*****************")

	jb := []string{"James", "bond", "martini", "chocolate"}
	jm := []string{"jenny", "aluminium", "coffe", "ice cream"}
	tm := []string{"Tom", "Busiwalker", "sprite", "Vaffles"}

	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println(tm)

	xxs := [][]string{jb, jm, tm}
	fmt.Println("3D slice: ", xxs)

	fmt.Println("***************- Slices memory implications -*****************")

	a1 := []int{0, 1, 2, 3, 4, 5}
	// b1 := a1

	b1 := make([]int, 6)
	copy(b1, a1)

	fmt.Println("A1\t", a1)
	fmt.Println("B1\t", b1)

	a1[0] = 7

	fmt.Println("A1\t", a1)
	fmt.Println("B1\t", b1)
}
