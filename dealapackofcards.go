package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 1; i <= 4; i++ {
		fmt.Printf("Player %d: %d, %d, %d\n",
			i,
			deck[(i-1)*3],
			deck[(i-1)*3+1],
			deck[(i-1)*3+2],
		)
	}
}
