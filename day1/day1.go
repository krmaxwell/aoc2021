package day1

func SonarSweep(data []int) int {
	decreaseCount := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			decreaseCount++
		}
	}
	return decreaseCount
}
