package day13

import "fmt"

type Paper [][]int

type Point struct {
	X, Y int
}

func MakePaper(sizeX, sizeY int) Paper {
	p := make(Paper, sizeY)
	//fmt.Printf("Creating paper with %d rows and %d columns\n", sizeY, sizeX)
	for r := range p {
		p[r] = make([]int, sizeX)
	}
	return p
}

func (p Paper) SetDots(dots []Point) {
	for _, dot := range dots {
		p[dot.Y][dot.X] = 1
	}
}

func (p Paper) Fold(x, y int) Paper {
	var newP Paper
	if y > 0 {
		newP = MakePaper(len(p[0]), y)
		for i := range newP {
			for j := range newP[0] {
				newP[i][j] = p[len(p)-1-i][j] | p[i][j]
			}
		}
	} else if x > 0 {
		newP = MakePaper(x, len(p))
		for i := range newP {
			for j := range newP[0] {
				//fmt.Printf("Combining row %d col %d and row %d col %d\n", i, len(p[0])-1-j, i, j)
				newP[i][j] = p[i][j] | p[i][len(p[0])-1-j]
				//fmt.Printf("New value is %d\n", newP[i][j])
			}
		}
	} else {
		return p
	}

	return newP
}

func (p Paper) CheckDot(x, y int) int {
	return p[y][x]
}

func (p Paper) CountDots() int {
	sum := 0
	for i := range p {
		for j := range p[0] {
			sum += p[i][j]
		}
	}
	return sum
}

func (p Paper) Print() {
	for i := range p {
		for j := range p[0] {
			if p[i][j] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
