package dynamic_programming

func BruteMaxValue(size int, prices []int) (value int) {
	if size == 0 {
		return 0
	} else {
		value = 0
		for i := 1; i <= size; i++ {
			value = max(value, prices[i]+BruteMaxValue(size-i, prices))
		}
		return
	}
}

func MemoizedMaxValue(size int, prices []int) int {
	memo := make([]int, len(prices))
	return memoizedMaxValue(size, prices, memo)
}

func memoizedMaxValue(size int, prices []int, memo []int) int {
	if memo[size] > 0 {
		return memo[size]
	}
	if size == 0 {
		return 0
	} else {
		value := 0
		for i := 1; i <= size; i++ {
			value = max(value, prices[i]+memoizedMaxValue(size-i, prices, memo))
		}
		memo[size] = value
		return value
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
