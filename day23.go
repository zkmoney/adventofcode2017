package advent

import (
	"fmt"
	"strings"
)

var instruxMap = map[string]int{
	"snd": 0,
	"set": 1,
	"add": 2,
	"mul": 3,
	"mod": 4,
	"rcv": 5,
	"jgz": 6,
}

type Program struct {
	ID                 int
	Instructions       []*Instruction
	Registers          []*int
	InstructionPointer int
}

type Instruction struct {
	Type     string
	Value    *int
	Modifier *int
}

func createProgram(input []byte, pid int, regSet string) *Program {
	p := Program{
		ID:        pid,
		Registers: make([]*int, len(regSet)),
	}
	for i := 0; i < len(p.Registers); i++ {
		p.Registers[i] = new(int)
	}

	inputs := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, in := range inputs {
		parts := strings.Split(strings.TrimSpace(in), " ")
		ins := Instruction{Type: parts[0]}

		v := parts[1]
		if strings.ContainsAny(v, regSet) {
			ins.Value = p.Registers[int(byte(v[0]))-97]
		} else {
			i := toInt(v)
			ins.Value = &i
		}

		if len(parts) > 2 {
			v := parts[2]
			if strings.ContainsAny(v, regSet) {
				ins.Modifier = p.Registers[int(byte(v[0]))-97]
			} else {
				i := toInt(v)
				ins.Modifier = &i
			}
		}
		p.Instructions = append(p.Instructions, &ins)
	}

	return &p
}

func Day23(part int) {
	input := getInput(23)

	p := createProgram(input, 0, "abcdefgh")
	*p.Registers[0] = 1

	runProgram23(p)
}

func runProgram23(p *Program) {
	var pos, muls int
	i := 0
	for i < 100 {
		if pos < 0 || pos >= len(p.Instructions) {
			break
		}
		ins := p.Instructions[pos]
		switch ins.Type {
		case "set":
			*ins.Value = *ins.Modifier
		case "sub":
			*ins.Value -= *ins.Modifier
		case "mul":
			*ins.Value *= *ins.Modifier
			muls++
		case "jnz":
			if *ins.Value != 0 {
				pos += *ins.Modifier
				continue
			}
		}

		fmt.Println(i, pos, ins)
		p.PrintRegisters()
		// if i%100000 == 0 {

		// }

		pos++

		if p.Register(7) != 0 {
			fmt.Println("NOT ZERO!", p.Register(7))
			break
		}
		i++
	}
	fmt.Println("muls", muls)
}

func (p *Program) PrintRegisters() {
	for i, v := range p.Registers {
		fmt.Printf("%s: %-9d ", []byte{byte(i + 97)}, *v)
	}
	fmt.Print("\n")
}

func (p *Program) Register(idx int) int {
	return *p.Registers[idx]
}
