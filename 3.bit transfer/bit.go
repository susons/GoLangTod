package main

import "fmt"

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	_ = iota // c0 == 0
	a
	b
	c
	d
	e
	f
)

type ByteSize float64

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("%d \t %b\n", KB, KB)
	fmt.Printf("%d \t %b\n", MB, MB)
	fmt.Printf("%d \t %b\n", GB, GB)
	fmt.Printf("%d \t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}
