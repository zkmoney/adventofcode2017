package advent

import (
	"fmt"
	"strings"
)

func DayFive(part int) {
	input := getInput(5)
	lines := strings.Split(string(input), "\n")

	var instrux []int
	for _, line := range lines[:len(lines)-1] {
		instrux = append(instrux, toInt(line))
	}

	var num, idx int
	for {
		num++

		inc := 1
		ins := instrux[idx]
		if part == 2 && ins >= 3 {
			inc = -1
		}
		instrux[idx] += inc

		idx += ins
		if idx < 0 || idx >= len(instrux) {
			break
		}
	}
	fmt.Println("Num steps", num)
}
