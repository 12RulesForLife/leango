package main 
import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d단\n",i)
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n",i,j, i*j)
		}
		fmt.Println()
	} 
}