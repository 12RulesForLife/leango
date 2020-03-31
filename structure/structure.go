package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func main() {
	p1 := Person{"개똥이", 10}
	p2 := Person{name: "말순이"}
	p3 := Person{name: "개말순", 0}
	fmt.Println(p1, p2)
}
