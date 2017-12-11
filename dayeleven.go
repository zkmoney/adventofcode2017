package advent

import (
	"fmt"
	"strings"
)

func DayEleven(part int) {
	input := getInput(11)
	steps := strings.Split(strings.TrimSpace(string(input)), ",")
	fmt.Println(string(input))

	var max int
	updateMax := func(x, y int) {
		if x > max {
			max = x
		}
		if y > max {
			max = y
		}
	}

	var x, y int
	for _, step := range steps {
		switch step {
		case "n":
			y++
		case "ne":
			y++
			x++
		case "nw":
			y++
			x--
		case "s":
			y--
		case "sw":
			y--
			x--
		case "se":
			y--
			x++
		}
		updateMax(x, y)
	}
	fmt.Println("X", x)
	fmt.Println("Y", y)
	fmt.Println("Max", max)

}
