package dynamic_programming

func BruteMaxValue(size int, prices []int) (value int) {
	if size == 0 {
		return 0
	} else {
		value = 0
		for i := 1; i <= size; i++ {
			rest := BruteMaxValue(size-i, prices)
			val := prices[i] + rest
			value = max(value, val)
		}
		return
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
