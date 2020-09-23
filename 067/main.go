package main

import (
	"fmt"
)

func main() {
	fmt.Println("10-5:", subtract(10, 5))
	fmt.Println("20385-187:", subtract(20385, 187))
}

func subtract(x, y int) int {
	return x - y
}
