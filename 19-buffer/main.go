package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))

	return err
}

func main() {
	p := person{
		first: "locomotive",
	}

	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)

	fmt.Println(b.String())

	fmt.Println("***************- Hands on 1 anonymous func -*****************")

	foo()
	func(s string) {
		fmt.Printf("This is anonymous with param: %v! \n", s)
	}("test")

	fmt.Println("***************- Hands on 2 func expression -*****************")

	x := func(s string) {
		fmt.Printf("This is a function expression: %v\n", s)
	}

	x("Bazookas")

	fmt.Println("***************- Hands on 3 return func -*****************")

	y := bar()
	y()
	fmt.Println(y())

	fmt.Println("***************- Hands on 4 func callbacks -*****************")

	
}

func foo() {
	fmt.Println("Foo ran!: ")
}

func bar() func() int {
	return func() int {
		return 45
	}
}
