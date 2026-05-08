package main

import "fmt"


func calc(x string, y, z int) (string, int){
	return x, y + z
}
func main(){
	var x int
	var y int
	fmt.Scan(&x, &y)
	fmt.Printf("Got: %d %d", x, y)
}