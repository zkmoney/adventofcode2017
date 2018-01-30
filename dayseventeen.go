package advent

import "fmt"

func DaySeventeen(part int) {
	input := 371

	spins := 50000000
	pos := 0
	buf := make([]int, spins+1)

	prev := 0

	for len := 1; len < spins; len++ {
		pos = (input + pos) % len

		// for z := len - 1; z > pos; z-- {
		// 	buf[z+1] = buf[z]
		// }
		pos++
		buf[pos] = len

		if len > 1 && prev != buf[1] {
			prev = buf[1]
			fmt.Println(prev)
		}
		// fmt.Println(buf[0:10])
	}

	fmt.Println(buf[0:10])

	// for i, v := range buf {
	// 	if v == 2017 {
	// 		fmt.Println(buf[i+1])
	// 		return
	// 	}
	// }
}
