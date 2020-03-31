package main

import "fmt"

func main() {
	var a []int
	PrintLenCap(a)

	a1 := make([]int, 0, 8)
	PrintLenCap(a1)

	a1 = append(a1, 2)
	PrintLenCap(a1)

	b := make([]int, 2, 3)
	b1 := append(b, 1)
	fmt.Printf("%p %p\n", b, b1)
	fmt.Println(b, b1)

	b1 = append(b, 2)
	fmt.Printf("%p %p\n", b, b1)
	fmt.Println(b, b1)

	b1 = append(b1, 3)
	fmt.Printf("%p %p\n", b, b1)
	fmt.Println(b, b1)

	c := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", &c)
	c1, _ := RemoveFront(c)
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p %p\n", c, c1)
	fmt.Println(c, c1)
}

//PrintLenCap of a Slice
func PrintLenCap(a []int) {
	fmt.Printf("len() = %d / ", len(a))
	fmt.Printf("cap() = %d\n", cap(a))
}

// RemoveFront of a Slice
func RemoveFront(s []int) ([]int, int) {
	fmt.Printf("%p\n", &s)
	front := s[0]
	s = s[1:]
	return s, front
}
