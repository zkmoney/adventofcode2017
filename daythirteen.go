package advent

import (
	"fmt"
	"strings"
)

func DayThirteen(part int) {
	input := getInput(13)
	// input = []byte(`
	// 	0: 3
	// 	1: 2
	// 	4: 4
	// 	6: 4
	// `)

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	slice := make([]int, 128, 128)
	last := 0
	for _, line := range lines {
		strs := strings.Split(line, ":")
		if len(strs) < 2 {
			continue
		}
		depth := toInt(strs[0])
		slice[depth] = toInt(strs[1])
		if depth > last {
			last = depth
		}
	}
	slice = slice[:last+1]

	delay := 0
	for {
		_, caught := checkSeverity(slice, delay)
		if !caught {
			break
		}
		delay++
	}
	fmt.Println("Delay:", delay)
}

func checkSeverity(slice []int, delay int) (int, bool) {
	var (
		sev    int
		caught bool
	)
	for i := 0; i < len(slice); i++ {
		rng := slice[i]
		if rng == 0 {
			continue
		}
		cycle := (rng - 1) * 2
		cyclePos := (i + delay) % cycle
		// fmt.Println(i+delay, rng, cycle, cyclePos)
		if cyclePos == 0 {
			sev += i * rng
			caught = true
		}
	}
	fmt.Println("Delay:", delay, "Severity:", sev, "Caught:", caught)
	return sev, caught
}
