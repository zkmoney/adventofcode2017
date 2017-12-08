package advent

import (
	"fmt"
	"strings"
)

func DaySix(part int) {
	input := getInput(6)
	lines := strings.Split(string(input), "\t")

	var banks []int
	for _, line := range lines {
		banks = append(banks, toInt(strings.TrimSpace(line)))
	}

	var (
		numBanks      = len(banks)
		numBits  uint = 4 // Hardcode for now
	)

	var hashes = make(map[uint]int)

	var (
		total     int
		foundHash uint
	)
	for {
		num, idx := maxInt(banks)
		banks[idx] = 0
		for i := 0; i < num; i++ {
			banks[(idx+1+i)%numBanks]++
		}

		total++
		var hash uint
		for i, n := range banks {
			hash += (uint(n) << (numBits * uint(i)))
		}
		fmt.Println(hash)

		if _, ok := hashes[hash]; ok {
			foundHash = hash
			break
		}

		hashes[hash] = total
	}

	fmt.Println("Num trips", total)
	fmt.Println("Between", total-hashes[foundHash])
}

func maxInt(ints []int) (int, int) {
	var max, maxIdx int
	for idx, i := range ints {
		if i > max {
			max = i
			maxIdx = idx
		}
	}
	return max, maxIdx
}
