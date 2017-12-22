package advent

import (
	"fmt"
	"math/bits"
)

type BigGrid [128][128]int

func DayFourteen(part int) {
	key := "vbqugkhl"
	extraInputs := []int{17, 31, 73, 47, 23}

	var (
		totalUsed int
		grid      BigGrid
	)

	for i := 0; i < 128; i++ {
		inputBytes := []byte(fmt.Sprintf("%s-%d", key, i))
		var inputs []int
		for _, b := range inputBytes {
			inputs = append(inputs, int(b))
		}
		inputs = append(inputs, extraInputs...)

		_, hexInts := knotHash(inputs)

		// Part 1
		var used int
		for j, hexInt := range hexInts {
			ones := bits.OnesCount(uint(hexInt))
			binStr := fmt.Sprintf("%08b", hexInt)
			// hexStr := fmt.Sprintf("%02x", hexInt)
			// fmt.Printf("%s - %s - %d\n", hexStr, binStr, ones)
			used += ones

			for k, s := range binStr {
				grid[i][j*8+k] = toInt(string(s))
			}
		}
		// fmt.Println(prettyHex(hexInts), " - ", used)
		totalUsed += used

		// Part 2
		// var visited [128][128]bool
	}

	fmt.Println("Total Used: ", totalUsed)

	regions := grid.CountRegions()
	// printGrid(grid)
	fmt.Println(regions)
}

func (g *BigGrid) CountRegions() int {
	g.Print()
	num := 2
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if g.At(x, y) < 2 {
				num = g.CheckRegion(x, y, num)
			}
		}
	}
	g.Print()
	return num - 2
}

func (g *BigGrid) CheckRegion(x, y, cur int) int {
	if g.At(x, y) != 1 {
		return cur
	}

	reg := cur
	g.Mark(x, y, reg)

	if g.At(x-1, y) == 1 {
		g.CheckRegion(x-1, y, reg)
	}
	if g.At(x, y-1) == 1 {
		g.CheckRegion(x, y-1, reg)
	}
	if g.At(x+1, y) == 1 {
		g.CheckRegion(x+1, y, reg)
	}
	if g.At(x, y+1) == 1 {
		g.CheckRegion(x, y+1, reg)
	}

	cur++
	return cur

	// var reg int
	// if adj := g.At(x-1, y); adj > 1 {
	// 	reg = adj
	// } else if adj := g.At(x, y-1); adj > 1 {
	// 	reg = adj
	// } else {
	// 	reg = cur
	// 	cur++
	// }
}

func (g *BigGrid) Mark(x, y, num int) {
	g[y][x] = num
}

func (g *BigGrid) At(x, y int) int {
	if x < 0 || x > 127 || y < 0 || y > 127 {
		return 0
	}
	return g[y][x]
}

func (g *BigGrid) Print() {
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			var c interface{}
			switch v := g.At(x, y); v {
			case 0:
				c = " "
			case 1:
				c = "X"
			default:
				c = v % 10
			}
			fmt.Print(c)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
