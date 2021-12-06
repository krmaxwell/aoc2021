package day6

func simulateDays(fish []int, days int) int64 {
	todaysFish := [9]int64{}

	// count existing fish by days remaining to reproduce
	for f := 0; f < len(fish); f++ {
		todaysFish[fish[f]]++
	}

	for day := 0; day < days; day++ {
		reproduction := todaysFish[0]
		for i := 0; i < 8; i++ {
			todaysFish[i] = todaysFish[i+1]
		}
		todaysFish[6] += reproduction
		todaysFish[8] = reproduction
	}
	var sum int64
	for i := 0; i < 9; i++ {
		sum += todaysFish[i]
	}
	return sum
}
