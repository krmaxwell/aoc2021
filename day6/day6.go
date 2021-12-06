package day6

func simulateDays(fish []int, days int) int {
	for day := 0; day < days; day++ {
		//fmt.Println(fish)
		fishLen := len(fish)
		for f := 0; f < fishLen; f++ {
			if fish[f] > 0 {
				fish[f]--
			} else if fish[f] == 0 {
				fish = append(fish, 8)
				fish[f] = 6
			}
		}
	}
	return len(fish)
}
