package main

import (
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
	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("Errror: %s", err)
	}

	defer f.Close()

	s := []byte("Hello nerds")

	_, err = f.Write(s)

	if err != nil {
		log.Fatalf("Errror: %s", err)
	}

}
