package day1

func SonarSweepNaive(data []int) int {
	decreaseCount := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			decreaseCount++
		}
	}
	return decreaseCount
}

func SonarSweepSlidingSum(data []int) int {
	var sums []int
	var newSum int
	for i := 0; i < len(data)-2; i++ {
		newSum = data[i] + data[i+1] + data[i+2]
		sums = append(sums, newSum)
	}
	return SonarSweepNaive(sums)
}
