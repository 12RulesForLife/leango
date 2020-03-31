package main

import "fmt"

func main() {
	a := 4
	b := 2
	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)
	fmt.Printf("%v\n", a^b)
	fmt.Println(a >> 1)
	fmt.Println(a << 1)

	if 3 > 4 {
		fmt.Println("true")
	}
	fmt.Println("false")

	a = 5
	if a == 3 || a == 4 {
		fmt.Println("a is 3 or 4")
	} else {
		fmt.Println("a is another")
	}
}
