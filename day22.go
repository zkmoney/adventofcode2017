package advent

import (
	"fmt"
	"strings"
)

func Day22(part int) {
	input := getInput(22)
	// input = []byte(`..#
	// #..
	// ...`)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	padding := 5
	dim := (padding * 2) + len(lines)

	n := NewNetwork(dim, dim)
	for row, line := range lines {
		line = strings.TrimSpace(line)
		for i, s := range line {
			n.Mark(i+padding, row+padding, string(s))
		}
	}

	x := dim / 2
	y := dim / 2
	dir := 0
	infected := 0

	for i := 0; i < 10000000; i++ {
		cur := n.At(x, y)
		if cur == "#" {
			n.Mark(x, y, "F")
			dir++
		} else if cur == "F" {
			n.Mark(x, y, ".")
			dir += 2
		} else if cur == "W" {
			n.Mark(x, y, "#")
			infected++
		} else {
			n.Mark(x, y, "W")
			dir--
		}

		if dir < 0 {
			dir = 3
		} else if dir > 3 {
			dir = dir % 4
		}

		switch dir {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}
	n.Print()

	fmt.Println("Infected", infected)
}

type Network struct {
	Grid
	Width  int
	Height int
}

func NewNetwork(w, h int) *Network {
	n := Network{
		Grid:   NewMapGrid(),
		Width:  w,
		Height: h,
	}
	return &n
}

func (n *Network) At(x, y int) string {
	// if x < 0 || x > n.Width || y < 0 || y > n.Height {
	// 	return " "
	// }
	return n.Grid.At(x, y)
}

func (n *Network) Print() {
	for y := 0; y < n.Height; y++ {
		for x := 0; x < n.Width; x++ {
			s := n.At(x, y)
			if s == "" {
				s = "."
			}
			fmt.Print(s)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

type Grid interface {
	At(x, y int) string
	Mark(x, y int, s string)
}

type mapGrid map[string]string

func NewMapGrid() Grid {
	return make(mapGrid)
}

func (mg mapGrid) Mark(x, y int, s string) {
	mg[mg.key(x, y)] = s
}

func (mg mapGrid) At(x, y int) string {
	return mg[mg.key(x, y)]
}

func (mg mapGrid) key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

type sliceGrid [][]string

func NewGrid(w, h int) Grid {
	var g sliceGrid
	g = make([][]string, h)
	for i := 0; i < h; i++ {
		g[i] = make([]string, w)
	}
	return &g
}

func (g *sliceGrid) Mark(x, y int, r string) {
	(*g)[y][x] = r
}

func (g *sliceGrid) At(x, y int) string {
	return (*g)[y][x]
}
