package advent

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func DaySeven(part int) {
	c := http.Cookie{
		Domain: AdventURL,
		Name:   "session",
		Value:  "53616c7465645f5f09756315b18550b477d224cb3c4030558fa16f3972673051d6f8e0e58abf81d8891166118dbb85cb",
	}
	req, err := http.NewRequest("GET", "http://adventofcode.com/2017/day/7/input", nil)
	checkErr(err)
	req.AddCookie(&c)

	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	nodes := make(map[string]*Node)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		name, weight := words[0], parseWeight(words[1])
		parent, ok := nodes[name]
		if !ok {
			parent = &Node{
				Name:     name,
				Children: make(map[string]*Node),
			}
			nodes[name] = parent
		}
		parent.Weight = weight

		if len(words) < 4 {
			continue
		}

		words = words[3:]
		for _, w := range words {
			w = strings.Replace(w, ",", "", -1)
			node, ok := nodes[w]
			if !ok {
				node = &Node{
					Name:     w,
					Children: make(map[string]*Node),
				}
				nodes[w] = node
			}
			node.Parent = parent
			parent.Children[w] = node
		}
	}

	var top *Node
	for _, n := range nodes {
		if n.Parent == nil {
			top = n
		}
	}

	if part == 1 {
		fmt.Println(top.Name, top)
	}

	if part == 2 {
		top = nodes["bntzksk"]
		for name, node := range top.Children {
			fmt.Println(name, node.Weight, branchWeight(node))
			/*
				bntzksk 45118
				mvpqv 45110
				pthnz 45110
				xinxep 45110
				qnhvjec 45110
				znztzxd 45110
				silwwua 45110
			*/

			/*
				vmttcwe 2318 2615
				ukwlfcf 1818 2607
				zzpevgd 17 2607
			*/

		}
	}
}

func branchWeight(n *Node) int {
	w := n.Weight
	for _, n := range n.Children {
		w += branchWeight(n)
	}
	return w
}

func parseWeight(weight string) int {
	return toInt(weight[1 : len(weight)-1])
}

type Node struct {
	Name     string
	Weight   int
	Parent   *Node
	Children map[string]*Node
}

// {
// 	var (
// 		n   int
// 		err error
// 	)
// 	for {
// 		b := make([]byte, 16)

// 		n, err = resp.Body.Read(b)
// 		fmt.Println(n, string(b), err)
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }
