package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m["moneypenny_jenny"] = []string{`intelligence`, `literature`, `computer science`}
	m["no_dr"] = []string{`cats`, `ice cream`, `sunsets`}
	m["flame_in_jar"] = []string{"flames", "cats", "jenots"}

	delete(m, "flame_in_jar")

	fmt.Println("M: ", m)

	for key, value := range m {
		fmt.Printf("key:\t%v, value:\t%v\n", key, value)

		for i, value2 := range value {
			fmt.Printf("index:\t%v, value2:\t%v\n", i, value2)
		}
	}

}

// key value
// `bond_james` `shaken, not stirred`, `martinis`, `fast cars`
// `moneypenny_jenny` `intelligence`, `literature`, `computer science`
// `no_dr` `cats`, `ice cream`, `sunsets`
