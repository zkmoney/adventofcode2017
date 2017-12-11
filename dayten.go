package advent

import (
	"bytes"
	"fmt"
	"strings"
)

func DayTen(part int) {
	input := getInput(10)
	var lengths []int
	if part == 1 {
		for _, str := range strings.Split(strings.TrimSpace(string(input)), ",") {
			lengths = append(lengths, toInt(str))
		}
	} else {
		for _, b := range bytes.TrimSpace(input) {
			lengths = append(lengths, int(b))
		}
		lengths = append(lengths, 17, 31, 73, 47, 23)
	}

	// lengths = []int{3, 4, 1, 5}
	// listLen := 5

	fmt.Println(lengths)

	listLen := 256
	list := make([]int, listLen)
	for i := 0; i < listLen; i++ {
		list[i] = i
	}

	pos := 0
	numRounds := 64
	for round := 0; round < numRounds; round++ {
		for skipSize, length := range lengths {
			swapElements(list, pos, length)
			trueSkip := round*len(lengths) + skipSize
			pos = (pos + length + trueSkip) % listLen
		}
	}

	// fmt.Println(list)
	if part == 1 {
		fmt.Println(list[0], list[1], list[0]*list[1])
	} else if part == 2 {
		var final string
		for i := 0; i < 16; i++ {
			var val int
			val = list[i*16+0] ^ list[i*16+1] ^ list[i*16+2] ^ list[i*16+3] ^ list[i*16+4] ^ list[i*16+5] ^ list[i*16+6] ^ list[i*16+7] ^ list[i*16+8] ^ list[i*16+9] ^ list[i*16+10] ^ list[i*16+11] ^ list[i*16+12] ^ list[i*16+13] ^ list[i*16+14] ^ list[i*16+15]
			fmt.Printf("%d - %02x\n", val, val)
			final = fmt.Sprintf("%s%02x", final, val)
		}
		fmt.Println(final)
	}
}

func swapElements(arr []int, base int, size int) {
	for i := 0; i < (size/2)+(size%2); i++ {
		a := (base + i) % len(arr)
		b := (base + size - 1 - i) % len(arr)

		// fmt.Println(a, b)
		tmp := arr[a]
		arr[a] = arr[b]
		arr[b] = tmp
	}
	// fmt.Println("Base", base, "Size", size)
	// fmt.Println("Array", arr)
}
