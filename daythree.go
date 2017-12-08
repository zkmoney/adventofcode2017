package advent

import (
	"fmt"
	"math"
)

const dayThreeTarget = 325489

func DayThree(part int) {
	if part == 1 {
		sideLen := int(math.Floor(math.Sqrt(float64(dayThreeTarget))))
		if sideLen%2 == 0 {
			sideLen++ // Account for being in the middle of the longer side
		}
		remain := dayThreeTarget - ((sideLen - 2) * (sideLen - 2))
		remain = remain % (sideLen - 1)
		fmt.Println(sideLen - 1 - remain)
		return
	}

	fmt.Println("Value:", dayThreePartTwo())
}

func dayThreePartTwo() int {
	dim := 31
	grid := make([][]int, dim)
	for i := 0; i < dim; i++ {
		grid[i] = make([]int, dim)
	}

	startPos := int(math.Floor(float64(dim) / 2))

	x := startPos
	y := startPos
	grid[y][x] = 1

	for i := 1; i < 5; i++ {
		j := i * 2

		x++
		if val := fillSpot(grid, x, y); val > dayThreeTarget {
			return val
		}
		for dir := 0; dir < 4; dir++ {
			for b := 0; b < j; b++ {
				switch dir {
				case 0:
					if b == 0 {
						continue
					}
					y--
				case 1:
					x--
				case 2:
					y++
				case 3:
					x++
				}
				if val := fillSpot(grid, x, y); val > dayThreeTarget {
					return val
				}
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	return 0
}

func fillSpot(grid [][]int, x, y int) int {
	// fmt.Println(x, y)
	// val := step
	// if step == 0 {
	// 	val = -1
	// }
	// // grid[y][x] = val
	// step++
	fmt.Println(grid[y+1][x], grid[y+1][x+1], grid[y][x+1], grid[y-1][x+1], grid[y-1][x], grid[y-1][x-1], grid[y][x-1], grid[y+1][x-1])
	val := grid[y+1][x] + grid[y+1][x+1] + grid[y][x+1] + grid[y-1][x+1] + grid[y-1][x] + grid[y-1][x-1] + grid[y][x-1] + grid[y+1][x-1]
	grid[y][x] = val
	return val
}
