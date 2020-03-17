package something

import (
	"fmt"
	"strings"
)

func SayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}

func Multiply(a, b int) int {
	return a * b
}

func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

func LenAndUpper_nakedReturn(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func SuperAdd(numbers ...int) int {
	defer fmt.Println("return the sum of numbers")
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func CanIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func CanIDrink2(age int) bool {
	switch {
	case age > 18:
		return false
	case age <= 18:
		return true
	}
	return false
}

func CanIDrink3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 50:
		return true
	}
	return false
}
