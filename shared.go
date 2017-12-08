package advent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const AdventURL = "http://adventofcode.com"

func getInput(day int) []byte {
	c := http.Cookie{
		Domain: AdventURL,
		Name:   "session",
		Value:  "53616c7465645f5f09756315b18550b477d224cb3c4030558fa16f3972673051d6f8e0e58abf81d8891166118dbb85cb",
	}
	urlStr := fmt.Sprintf("http://adventofcode.com/2017/day/%d/input", day)
	req, err := http.NewRequest("GET", urlStr, nil)
	checkErr(err)
	req.AddCookie(&c)

	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return b
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func addString(tot int, s string) int {
	i, _ := strconv.Atoi(s)
	return tot + i
}

func checkErr(err error) {
	if err == nil {
		return
	}
	panic(err)
}
