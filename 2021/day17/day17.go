package day17

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

type probe struct {
	x, y                 int
	xVelocity, yVelocity int
}

func (p *probe) step() {
	p.stepX()
	p.stepY()
}

func (p *probe) stepY() {
	p.y += p.yVelocity
}

func (p *probe) stepX() {
	p.x += p.xVelocity
}

func (p *probe) updateVelocity() {
	if p.xVelocity > 0 {
		p.xVelocity--
	} else if p.xVelocity < 0 {
		p.xVelocity++
	}

	p.yVelocity--
}

func (p *probe) inTarget(target targetArea) bool {
	return target.x1 <= p.x && p.y <= target.x2 && target.y1 <= p.y && p.y <= target.y2
}

func (p *probe) outOfTarget(target targetArea) bool {
	return (p.x == 0 && (p.y < target.x1 || p.y > target.x2)) || (p.y < target.y1)
}

type targetArea struct {
	x1, x2 int
	y1, y2 int
}

func partA(in string) int {
	inputParts := strings.Split(in, " ")
	target := targetArea{}
	target.x1, target.x2 = extract(inputParts[2])
	target.y1, target.y2 = extract(inputParts[3])
	maxY := 0

	for x := 0; x <= target.x2; x++ {
		for y := target.y1; y <= -target.y1; y++ {
			p := probe{
				xVelocity: x,
				yVelocity: y,
			}
			maxHeight := 0

			for {
				p.step()
				if p.y > maxHeight {
					maxHeight = p.y
				}
				if p.inTarget(target) {
					if maxHeight > maxY {
						maxY = maxHeight
					}
					break
				} else if p.outOfTarget(target) {
					break
				}

				p.updateVelocity()
			}
		}
	}

	return maxY
}

func partB(in string) int {
	inputParts := strings.Split(in, " ")
	target := targetArea{}
	target.x1, target.x2 = extract(inputParts[2])
	target.y1, target.y2 = extract(inputParts[3])
	counter := 0

	for x := 0; x <= target.x2; x++ {
		for y := target.y1; y <= -1*target.y1; y++ {
			p := probe{
				xVelocity: x,
				yVelocity: y,
			}

			for {
				p.step()
				if p.inTarget(target) {
					counter++
					break
				} else if p.outOfTarget(target) {
					break
				}

				p.updateVelocity()
			}
		}
	}

	counter = 0

	for x := 0; x <= target.x2; x++ {
		for y := target.y1; y <= -target.y1; y++ {
			p := probe{
				xVelocity: x,
				yVelocity: y,
			}
			for {
				p.step()
				if target.x1 <= p.x && p.x <= target.x2 && target.y1 <= p.y && p.y <= target.y2 {
					counter++
					break
				} else if (p.x == 0 && (p.y < target.x1 || p.y > target.x2)) || (p.y < target.y1) {
					break
				}
				p.updateVelocity()
			}
		}
	}

	return counter
}

func extract(s string) (int, int) {
	s = s[2:] // remove: x=
	if s[len(s)-1] == ',' {
		s = s[:len(s)-1]
	}

	xParts := strings.Split(s, "..")
	a := strToInt(xParts[0])
	b := strToInt(xParts[1])
	return a, b
}

func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
