package advent

import (
	"fmt"
	"strings"
)

type DanceMove struct {
	Type byte
	A    int
	B    int
}

func DaySixteen(part int) {
	input := getInput(16)

	dance := make([]int, 16)
	for i := 0; i < 16; i++ {
		dance[i] = i + 97
	}

	var moves []*DanceMove
	for _, m := range strings.Split(strings.TrimSpace(string(input)), ",") {
		move := DanceMove{Type: m[0]}
		rest := m[1:]
		switch move.Type {
		case 's':
			end := 16 - toInt(rest)
			move.A = end
		case 'x':
			ps := strings.Split(rest, "/")
			a, b := ps[0], ps[1]
			move.A = toInt(a)
			move.B = toInt(b)
		case 'p':
			ps := strings.Split(rest, "/")
			a, b := []byte(ps[0]), []byte(ps[1])
			move.A = int(a[0])
			move.B = int(b[0])
		}

		moves = append(moves, &move)
	}

	var (
		key   string
		found bool

		seen = make(map[string]struct{})
	)

	times := 1000000000
	cycle := 60
	for i := 0; i < times%cycle; i++ {
		for _, move := range moves {
			switch move.Type {
			case 's':
				dance = append(dance[move.A:], dance[0:move.A]...)
			case 'x':
				swapEls(dance, move.A, move.B)
			case 'p':
				swapEls(dance, findPos(dance, move.A), findPos(dance, move.B))
			}
		}
		key = toString(dance)
		if _, found = seen[key]; found {
			fmt.Println("Cycle len", i)
			return
		}
		seen[key] = struct{}{}
	}

	fmt.Println(key)
}

func toString(in []int) string {
	var out []byte
	for _, i := range in {
		out = append(out, byte(i))
	}
	return string(out)
}

func findPos(arr []int, test int) int {
	for i, v := range arr {
		if v == test {
			return i
		}
	}
	return -1
}

func swapEls(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
