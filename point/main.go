package main

import "github.com/01-edu/z01"

type point struct {
	x rune
	y rune
}

func setPoint(ptr *point) {
	ptr.x = 47
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	// var output [15]rune

	// one := 'b' - 'a'
	ten := 'k' - 'a'

	output := []rune{'x', ' ', '=', ' ', '0' + points.x/ten, '0' + points.x%ten, ',', ' ', 'y', ' ', '=', ' ', '0' + points.y/ten, '0' + points.y%ten, '\n'}

	// output[0] = 'x'
	// output[1] = ' '
	// output[2] = '='
	// output[3] = ' '

	// one := 'b' - 'a'
	// ten := 'j' - 'a' + one

	// output[4] = '0' + points.x/ten
	// output[5] = '0' + points.x%ten

	// output[6] = ','
	// output[7] = ' '
	// output[8] = 'y'
	// output[9] = ' '
	// output[10] = '='
	// output[11] = ' '

	// output[12] = '0' + points.y/ten
	// output[13] = '0' + points.y%ten

	// output[14] = '\n'

	for i := 0; i < 15; i++ {
		z01.PrintRune(output[i])
	}
}
