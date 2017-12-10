package advent

import (
	"bytes"
	"fmt"
)

const (
	StartGroup   = '{'
	EndGroup     = '}'
	StartGarbage = '<'
	EndGarbage   = '>'
	IgnoreChar   = '!'
)

func DayNine() {
	input := getInput(9)
	// lines := strings.Split(string(input), "\n")

	// input = []byte("{{{}}}")
	// input = []byte("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	// input = []byte("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	// input = []byte(`<{o"i!a,<{i<a>`)
	// input = []byte(`<!!!>>,`)

	depth := 0
	totalScore := 0

	inGarbage := false

	garbageChars := 0

	buf := bytes.NewBuffer(bytes.TrimSpace(input))
	for {
		r, _, err := buf.ReadRune()
		if err != nil {
			break
		}

		switch r {
		case IgnoreChar:
			_, _, err := buf.ReadRune()
			if err != nil {
				panic(err)
			}
		case StartGroup:
			if !inGarbage {
				depth++
			} else {
				garbageChars++
			}
		case EndGroup:
			if !inGarbage {
				totalScore += depth
				depth--
			} else {
				garbageChars++
			}
		case StartGarbage:
			if !inGarbage {
				inGarbage = true
			} else {
				garbageChars++
			}
		case EndGarbage:
			inGarbage = false
		default:
			if inGarbage {
				garbageChars++
			}
		}

		// fmt.Printf("Rune: %v  Depth: %v  Total: %v InGarbage: %v\n", r, depth, totalScore, inGarbage)
	}

	fmt.Println("Total", totalScore)
	fmt.Println("Num Garbage", garbageChars)
}
