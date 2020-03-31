package main

import "fmt"

func main() {
	var x int = 32
	switch 30 {
	case x:
		fmt.Println("x = 30")
	case x - 1:
		fmt.Println("x = 30 - 1")
	case x - 2:
		fmt.Println("x = 30 - 2")
	}

}
