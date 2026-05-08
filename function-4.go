package main
import (
	"fmt"
)

func add(x, y int) int {
	fmt.Printf("%d \n", x + y)
	return x + y
}

func main(){
	fmt.Printf("Returned value %d ", add(5, 10))
}