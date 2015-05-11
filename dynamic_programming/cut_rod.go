package dynamic_programming

import "math"

func BruteMaxValue(size int, prices []int) (value int) {
	if size == 0 {
		return 0
	} else {
		value = math.MinInt64
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
		value := math.MinInt64
		for i := 1; i <= size; i++ {
			value = max(value, prices[i]+memoizedMaxValue(size-i, prices, memo))
		}
		memo[size] = value
		return value
	}
}

func BottomUpMaxValue(size int, prices []int) int {
	rs := make([]int, len(prices))
	rs[0] = 0
	for i := 1; i <= size; i++ {
		value := math.MinInt64
		for j := 1; j <= i; j++ {
			value = max(value, prices[j]+rs[i-j])
		}
		rs[i] = value
	}
	return rs[size]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
