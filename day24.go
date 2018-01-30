package advent

import (
	"fmt"
	"sort"
	"strings"
)

var day24Input = `14/42
2/3
6/44
4/10
23/49
35/39
46/46
5/29
13/20
33/9
24/50
0/30
9/10
41/44
35/50
44/50
5/11
21/24
7/39
46/31
38/38
22/26
8/9
16/4
23/39
26/5
40/40
29/29
5/20
3/32
42/11
16/14
27/49
36/20
18/39
49/41
16/6
24/46
44/48
36/4
6/6
13/6
42/12
29/41
39/39
9/3
30/2
25/20
15/6
15/23
28/40
8/7
26/23
48/10
28/28
2/13
48/1`

func Day24(part int) {
	day24Input = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`
	lines := strings.Split(strings.TrimSpace(day24Input), "\n")

	var mags []*Magnet
	for _, line := range lines {
		strs := strings.Split(strings.TrimSpace(line), "/")
		a := toInt(strs[0])
		b := toInt(strs[1])
		if a < b && a != 0 {
			tmp := a
			a = b
			b = tmp
		}
		mag := Magnet{A: a, B: b}
		mags = append(mags, &mag)
	}

	sort.Slice(mags, func(i, j int) bool {
		if mags[i].A == 0 {
			return true
		} else if mags[j].A == 0 {
			return false
		}
		return mags[i].A > mags[j].A
	})

	start := mags[0]
	bridge, rest := mags[0:1], mags[1:]
	findBridge(0, start, false, bridge, rest)

	fmt.Println("TOTAL DONE:", day24Max)
	fmt.Println("TOTAL LONGEST:", day24Longest)
	fmt.Println("ALL:", day24Maxes)

	var maxLen int
	for k := range day24All {
		if k > maxLen {
			maxLen = k
		}
	}

	var max int
	for _, val := range day24All[maxLen] {
		fmt.Println(val)
		// if val > max {
		// 	max = val
		// }
	}
	fmt.Println("MAX", maxLen, max)
}

var (
	day24Max     int
	day24Longest int
	day24Maxes   = make(map[int]int)
	day24All     = make(map[int][]*Magnet)
)

func findBridge(total int, start *Magnet, useA bool, bridge, rest []*Magnet) int {
	total += start.A + start.B
	if total > day24Max {
		day24Max = total
	}

	bridgeLen := len(bridge)
	if bridgeLen > day24Longest {
		day24Longest = bridgeLen
	}
	if bmax := day24Maxes[bridgeLen]; total > bmax {
		day24Maxes[bridgeLen] = total
	}

	day24All[bridgeLen] = append(day24All[bridgeLen], bridge...)

	var port int
	if useA {
		port = start.A
	} else {
		port = start.B
	}

	lenRest := len(rest)

	// var max int
	for idx := 0; idx < lenRest; idx++ {
		mag := rest[idx]
		if mag.A == port || mag.B == port {
			var a bool
			if mag.B == port {
				a = true
			}

			rest := newMagnets(rest, idx)
			newBridge := append(bridge, mag)
			findBridge(total, mag, a, newBridge, rest)
		}
	}

	return total
}

func newMagnets(mags []*Magnet, idx int) []*Magnet {
	var nmags []*Magnet
	for i, m := range mags {
		if i == idx {
			continue
		}
		nmags = append(nmags, m)
	}
	return nmags
}

type Connection struct {
	Prev *int
	Next *int
}

type Magnet struct {
	A int
	B int
}
