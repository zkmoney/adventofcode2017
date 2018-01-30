package advent

import "fmt"

func DayFifteen(part int) {
	inputA := 277
	factorA := 16807

	inputB := 349
	factorB := 48271

	div := 2147483647
	mask := 65535

	if part == 1 {
		var count int
		for i := 0; i < 40000000; i++ {
			inputA = (inputA * factorA) % div
			inputB = (inputB * factorB) % div

			if (inputA & mask) == (inputB & mask) {
				count++
			}
		}
		fmt.Println(count)
		return
	}

	gen := func(in, factor, mod int) int {
		for {
			in = (in * factor) % div
			if in%mod == 0 {
				return in
			}
		}
	}

	var count int
	for i := 0; i < 5000000; i++ {
		inputA = gen(inputA, factorA, 4)
		inputB = gen(inputB, factorB, 8)

		if (inputA & mask) == (inputB & mask) {
			count++
		}
	}
	fmt.Println(count)
}
