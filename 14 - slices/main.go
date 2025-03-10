package main

import "fmt"

func main() {

	fmt.Println("***************- Hands on 1 -*****************")
	a := [5]int{}

	for i := 0; i < 5; i++ {
		a[i] = i * 3
		fmt.Println("Value: ", a[i])
	}

	for _, v := range a {
		fmt.Println("Range value: ", v)
	}

	fmt.Println("***************- Hands on 2 -*****************")

	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range b {
		fmt.Printf("%v \t %T \t %v\n", i, v, v)
	}

	fmt.Println("***************- Hands on 3 -*****************")

	// ● [42 43 44 45 46]
	// ● [47 48 49 50 51]
	// ● [44 45 46 47 48]
	// ● [43 44 45 46 47

	fmt.Println(b[:5])
	fmt.Println(b[5:])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])

	fmt.Println("***************- Hands on 4 -*****************")

	b = append(b, 52)
	fmt.Println(b)

	b = append(b, 53, 54, 55)
	fmt.Println(b)

	c := []int{56, 57, 58, 59, 60}

	b = append(b, c...)

	fmt.Println(b)

	d := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	d = append(d[:3], d[6:]...)

	fmt.Println(d)

	fmt.Println("***************- Hands on 5 -*****************")

	x1 := make([]string, 0, 50)

	x1 = append(x1, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println("X1: ", x1)

	for i := 0; i < len(x1); i++ {
		fmt.Printf("%v: \t%v\n", i+1, x1[i])
	}

	fmt.Println("***************- Hands on 6 -*****************")

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	xp := [][]string{jb, jm}

	for i, v := range xp {
		fmt.Println(i, v)

		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
