package advent

import (
	"fmt"
	"sync"
)

func Day18(part int) {
	input := getInput(18)

	program0 := createProgram(input, 0, "abcdefghijklmnopqrstuvwxyz")
	program1 := createProgram(input, 1, "abcdefghijklmnopqrstuvwxyz")
	*program1.Registers[int('p')-97] = 1

	ch0, ch1 := make(chan int, 100), make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go runDay18(program0, ch0, ch1)
	go runDay18(program1, ch1, ch0)

	wg.Wait()
}

func runDay18(p *Program, snd, rcv chan int) {
	var pos, sends int
	for {
		ins := p.Instructions[pos]
		switch ins.Type {
		case "snd":
			snd <- *ins.Value
			sends++
			fmt.Println(p.ID, "Total Sends:", sends)
		case "set":
			*ins.Value = *ins.Modifier
		case "add":
			*ins.Value += *ins.Modifier
		case "mul":
			*ins.Value *= *ins.Modifier
		case "mod":
			*ins.Value %= *ins.Modifier
		case "rcv":
			*ins.Value = <-rcv
		case "jgz":
			if *ins.Value > 0 {
				pos += *ins.Modifier
				continue
			}
		}
		pos++
	}
}
