package main
import "fmt"
func split(sum int) (x, y int){
	var a int = sum + 4
	var b int = sum - 4
	fmt.Println(a + b)
	return
}

func main(){
	var sum int
	fmt.Scan(&sum)
	fmt.Println(split(sum))
}