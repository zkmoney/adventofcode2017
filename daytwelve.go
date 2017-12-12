package advent

import (
	"fmt"
	"strings"
)

type Pipe struct {
	Val   int
	Pipes []*Pipe
}

func DayTwelve(part int) {
	input := getInput(12)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	pipes := make([]*Pipe, len(lines))
	for _, line := range lines {
		segs := strings.Split(line, "<->")
		idx := toInt(segs[0])
		if pipes[idx] == nil {
			pipes[idx] = &Pipe{Val: idx}
		}

		if len(segs) < 2 {
			continue
		}

		for _, s := range strings.Split(segs[1], ",") {
			num := toInt(s)
			if pipes[num] == nil {
				pipes[num] = &Pipe{Val: num}
			}
			pipes[idx].Pipes = append(pipes[idx].Pipes, pipes[num])
		}
	}

	var (
		// start = 0
		groupCount = 0
		seen       = make([]bool, len(pipes))
	)
	for _, pipe := range pipes {
		if seen[pipe.Val] {
			continue
		}
		groupCount++
		traversePipes(pipe, seen)
	}
	fmt.Println(groupCount)

	// var count int
	// for _, b := range seen {
	// 	if b {
	// 		count++
	// 	}
	// }
}

func traversePipes(pipe *Pipe, seen []bool) {
	if seen[pipe.Val] == true {
		return
	}
	seen[pipe.Val] = true
	for _, p := range pipe.Pipes {
		traversePipes(p, seen)
	}
}
