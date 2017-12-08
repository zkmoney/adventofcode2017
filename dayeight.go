package advent

import (
	"fmt"
	"strings"
)

func DayEight() {
	input := getInput(8)
	lines := strings.Split(string(input), "\n")

	var (
		instrux   []*Instrux
		registers = make(map[string]*Register)
	)
	for _, line := range lines {
		ws := strings.Split(line, " ")
		if len(ws) < 6 {
			continue
		}

		regName := ws[0]
		reg, ok := registers[regName]
		if !ok {
			reg = &Register{Name: regName}
			registers[regName] = reg
		}

		condRegName := ws[4]
		condReg, ok := registers[condRegName]
		if !ok {
			condReg = &Register{Name: condRegName}
			registers[condRegName] = condReg
		}

		ins := Instrux{
			Register: reg,
			Type:     ws[1],
			Amount:   toInt(ws[2]),
			Cond: &Cond{
				Register: condReg,
				Op:       ws[5],
				Val:      toInt(ws[6]),
			},
		}
		instrux = append(instrux, &ins)
	}

	var maxForRealz int

	for _, ins := range instrux {
		var (
			exec bool

			b = ins.Cond.Register
		)

		switch ins.Cond.Op {
		case OpGT:
			if b.Val > ins.Cond.Val {
				exec = true
			}
		case OpGTE:
			if b.Val >= ins.Cond.Val {
				exec = true
			}
		case OpLT:
			if b.Val < ins.Cond.Val {
				exec = true
			}
		case OpLTE:
			if b.Val <= ins.Cond.Val {
				exec = true
			}
		case OpEq:
			if b.Val == ins.Cond.Val {
				exec = true
			}
		case OpNE:
			if b.Val != ins.Cond.Val {
				exec = true
			}
		}

		if !exec {
			continue
		}

		switch ins.Type {
		case Inc:
			ins.Register.Val += ins.Amount
		case Dec:
			ins.Register.Val -= ins.Amount
		}

		if ins.Register.Val > maxForRealz {
			maxForRealz = ins.Register.Val
		}
	}

	var max int
	for _, reg := range registers {
		if reg.Val > max {
			max = reg.Val
		}
	}

	fmt.Println("Max at end", max)
	fmt.Println("Max at anytime", maxForRealz)
}

const (
	Inc = "inc"
	Dec = "dec"

	OpGT  = ">"
	OpGTE = ">="
	OpLT  = "<"
	OpLTE = "<="
	OpEq  = "=="
	OpNE  = "!="
)

type Register struct {
	Name string
	Val  int
}

type Instrux struct {
	Register *Register
	Type     string
	Amount   int

	Cond *Cond
}

type Cond struct {
	Register *Register
	Op       string
	Val      int
}
