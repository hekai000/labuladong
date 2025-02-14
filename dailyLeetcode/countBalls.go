package labuladong

func countBalls(lowLimit int, highLimit int) int {
	result := make(map[int]int)
	maxres := 0
	for i := lowLimit; i <= highLimit; i++ {
		index := 0
		j := i
		for j > 0 {

			index += j % 10
			j = j / 10
		}
		if _, ok := result[index]; ok {
			result[index] += 1
		} else {
			result[index] = 1
		}

	}
	for _, v := range result {
		if v > maxres {
			maxres = v
		}
	}
	return maxres

}
