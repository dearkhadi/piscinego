package main

import (
	"fmt"
	"os"
)

func main() {
	myArg := os.Args[1:]

	for _, arg := range myArg {
		switch arg {
		case "01":
			fmt.Println("Alert!!!")
			return
		case "galaxy":
			fmt.Println("Alert!!!")
			return
		case "galaxy 01":
			fmt.Println("Alert!!!")
			return
		}
	}
}
