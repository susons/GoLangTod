package main

import "fmt"

func main() {
	a := make(map[string]int)

	a["lucas"] = 28
	a["steph"] = 37
	a["george"] = 78

	fmt.Println(a)

	for k, v := range a {
		fmt.Println("key: ", k, " value: ", v)
	}

	delete(a, "lucas")

	v, ok := a["banana"]

	if ok {
		fmt.Println("Value exist!!!", v)

	} else {
		fmt.Println("Value doesnt exist?????")
	}

	fmt.Println(a)
}
