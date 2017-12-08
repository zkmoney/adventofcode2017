package main

import (
	"fmt"
	"math"
)

func main() {
	// target := 325489
	// count := 1
	// for i := 0; ; i++ {
	// 	perSide := i * 2
	// 	count += (4 * perSide)
	// 	fmt.Println(perSide, count)
	// 	if count > target || i > 10 {
	// 		fmt.Println("Found it", math.Sqrt(float64(count)))
	// 		break
	// 	}
	// }
	dayThree()
}

func dayThree() {
	target := 325489
	sideLen := int(math.Floor(math.Sqrt(float64(target))))
	if sideLen%2 == 0 {
		sideLen++ // Account for being in the middle of the longer side
	}
	remain := target - ((sideLen - 2) * (sideLen - 2))
	remain = remain % (sideLen - 1)
	fmt.Println(sideLen - 1 - remain)

	// rad := 1
	// dim := (rad * 2) + 1
	// grid := make([][]int, dim)
	// for i := 0; i < dim; i++ {
	// 	grid[i] = make([]int, dim)
	// }

	// for i := 0; i < dim*dim; i++ {
	// 	fmt.Println()
	// }
}
