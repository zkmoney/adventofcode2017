package advent

import (
	"fmt"
	"strings"
)

func Day19(part int) {
	input := getInput(19)
	lines := strings.Split(string(input), "\n")

	d := Diagram{Grid: make(mapGrid)}
	// h := len(lines)
	w := 0
	for y, line := range lines {
		for x, s := range line {
			d.Mark(x, y, string(s))
		}
		if len(line) > w {
			w = len(line)
		}
	}

	var (
		x, y int
	)
	// Get starting values
	for x = 0; x < w; x++ {
		s := d.At(x, y)
		if s == "|" {
			break
		}
	}

	var chars string
	var total int
	dir := 2 // 2 is down
	for {
		s := d.At(x, y)
		if s == "|" {
			// if dir != 2 && dir != 0 {
			// 	panic("HOW DID THIS HAPPEN - Y")
			// }
		} else if s == "-" {
			// if dir != 1 && dir != 3 {
			// 	panic("HOW DID THIS HAPPEN - X")
			// }
		} else if s == "+" {
			oldDir := dir
			switch {
			case d.At(x, y+1) == "|" && dir != 0:
				dir = 2
			case d.At(x, y-1) == "|" && dir != 2:
				dir = 0
			case d.At(x+1, y) == "-" && dir != 3:
				dir = 1
			case d.At(x-1, y) == "-" && dir != 1:
				dir = 3
			default:
				panic("Nooooooo!")
			}
			fmt.Println("Switching!", x, y, oldDir, dir)
		} else if s == " " {
			fmt.Println("FOUND SPACE")
			break
		} else {
			fmt.Println("NEW CHAR", s)
			chars = fmt.Sprintf("%s%s", chars, s)
		}

		if dir == 2 {
			y++
		} else if dir == 0 {
			y--
		} else if dir == 1 {
			x++
		} else if dir == 3 {
			x--
		}

		total++
		// fmt.Println("Up next:", x, y, dir, d.At(x, y))
	}

	fmt.Println("Order", chars)
	fmt.Println("Total", total)

	// for y := 0; y < h; y++ {
	// 	for x := 0; x < w; x++ {
	// 		fmt.Print(d.At(x, y))
	// 	}
	// 	fmt.Print("\n")
	// }
}

type Diagram struct {
	Grid
}

func (d *Diagram) Traverse() {

}
