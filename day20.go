package advent

import (
	"fmt"
	"math"
	"strings"
)

func Day20(part int) {
	input := getInput(20)

	var parts []*Particle
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		fs := strings.Fields(line)
		p := parseIt(fs, 0)
		v := parseIt(fs, 1)
		a := parseIt(fs, 2)
		part := Particle{
			X: &Forces{Pos: toInt(p[0]), Vel: toInt(v[0]), Acc: toInt(a[0])},
			Y: &Forces{Pos: toInt(p[1]), Vel: toInt(v[1]), Acc: toInt(a[1])},
			Z: &Forces{Pos: toInt(p[2]), Vel: toInt(v[2]), Acc: toInt(a[2])},
		}
		parts = append(parts, &part)
	}

	if part == 1 {
		min := parts[0].ManhattanAccel()
		minIdx := 0
		for i, p := range parts {
			if p.ManhattanAccel() < min {
				min = p.ManhattanAccel()
				minIdx = i
			}
		}
		fmt.Println("Min", minIdx)
		return
	}

	fmt.Println("Total Particles:", len(parts))

	runs := 100
	for i := 0; i < runs; i++ {
		m := make(map[string]*Particle)
		for _, p := range parts {
			if p.Destroyed {
				continue
			}

			key := p.PositionString()
			if v := m[key]; v != nil {
				v.Destroyed = true
				p.Destroyed = true
			} else {
				m[key] = p
			}
		}

		var destroyed int
		for _, p := range parts {
			if p.Destroyed {
				destroyed++
			}
		}

		fmt.Println("Total Left:", len(parts)-destroyed)

		for _, p := range parts {
			if p.Destroyed {
				continue
			}
			p.Update()
		}
	}
}

func parseIt(fs []string, idx int) []string {
	return strings.Split(fs[idx][3:strings.Index(fs[idx], ">")], ",")
}

type Particle struct {
	X, Y, Z   *Forces
	Destroyed bool
}

func (p *Particle) Update() {
	p.X.Vel += p.X.Acc
	p.Y.Vel += p.Y.Acc
	p.Z.Vel += p.Z.Acc

	p.X.Pos += p.X.Vel
	p.Y.Pos += p.Y.Vel
	p.Z.Pos += p.Z.Vel
}

func (p *Particle) ManhattanDistance() float64 {
	return math.Abs(float64(p.X.Pos)) + math.Abs(float64(p.Y.Pos)) + math.Abs(float64(p.Z.Pos))
}

func (p *Particle) ManhattanAccel() float64 {
	return math.Abs(float64(p.X.Acc)) + math.Abs(float64(p.Y.Acc)) + math.Abs(float64(p.Z.Acc))
}

func (p *Particle) PositionString() string {
	return fmt.Sprintf("%d:%d:%d", p.X.Pos, p.Y.Pos, p.Z.Pos)
}

type Forces struct {
	Pos int
	Vel int
	Acc int
}
