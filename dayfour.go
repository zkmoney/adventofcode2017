package advent

import (
	"fmt"
	"sort"
	"strings"
)

func DayFour(part int) {
	input := getInput(4)
	lines := strings.Split(string(input), "\n")

	var total int
	for _, line := range lines {
		var (
			ss    []string
			found bool
		)

		words := strings.Split(line, " ")
		if len(words) < 2 {
			continue
		}
		for _, w := range words {
			if w == "" {
				continue
			}

			if part == 2 {
				w = sortChars(w)
			}

			if strSliceContains(ss, w) {
				found = true
				break
			}
			ss = append(ss, w)
		}
		if !found {
			total++
		}
	}
	fmt.Println(len(lines))
	fmt.Println(total)
}

func strSliceContains(ss []string, s string) bool {
	for _, str := range ss {
		if s == str {
			return true
		}
	}
	return false
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortChars(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
