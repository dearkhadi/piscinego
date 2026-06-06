package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 1, 2, 3, 4, 3, 4, 2, 3}
	unmatch := piscine.Unmatch(a)
	fmt.Println(unmatch)
}
