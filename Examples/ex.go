package main
import "fmt"

func main() {
		var a,b int
		

		fmt.Println("Enter a value")
		fmt.Scan(&a)

		fmt.Println("Enter b value")
		fmt.Scan(&b)

		sum := a + b

		fmt.Println("sum is:",sum)
}