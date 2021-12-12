package day11

type Octopus struct {
	Energy int
	Flash  bool
}

type School [][]Octopus

func ConstructOctopi(field []string) School {
	octopi := make(School, len(field))
	for i := range octopi {
		octopi[i] = make([]Octopus, len(field[0]))
		for j := range field[i] {
			energy := int(field[i][j] - '0')
			octopi[i][j] = Octopus{Flash: false, Energy: energy}
		}
	}
	return octopi
}

func (s School) ProcessStep() int {
	for i := range s {
		for j := range s[i] {
			s.Energize(i, j)
		}
	}
	count := 0
	for i := range s {
		for j := range s[i] {
			if s[i][j].Flash {
				s[i][j].Energy = 0
				s[i][j].Flash = false
				count++
			}
		}
	}
	return count
}

func (s School) Energize(i int, j int) {
	if i >= 0 && j >= 0 && i < len(s) && j < len(s[0]) && s[i][j].Energy < 10 {
		s[i][j].Energy++
		if s[i][j].Energy > 9 {
			s[i][j].Flash = true
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}
					s.Energize(i+di, j+dj)
				}
			}
		}
	}
}

func (o *Octopus) ValidateFlash() {
	if o.Energy > 9 {
		o.Flash = true
	}
}
