package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

type count int

func loginfo(s fmt.Stringer) {
	log.Println("Log from stringer!: ", s)
}

func (c count) String() string {
	return fmt.Sprint("This will be a number: ", strconv.Itoa(int(c)))
}

func (b book) String() string {
	return fmt.Sprint("The title of book is ", b.title)
}

func main() {
	b := book{
		title: "harry potter",
	}

	var n count = 42

	loginfo(b)
	log.Println(n)
}
