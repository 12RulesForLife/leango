package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)
	fmt.Println("연산자를 입력하세요")
	line, _ = reader.ReadString('\n')
	op1 := strings.TrimSpace(line)

	fmt.Print("결과는 ... ")
	if op1 == "+" {
		fmt.Println(n1 + n2)
	} else if op1 == "-" {
		fmt.Println(n1 - n2)
	} else if op1 == "*" {
		fmt.Println(n1 * n2)
	} else if op1 == "/" {
		fmt.Println(n1 / n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}
