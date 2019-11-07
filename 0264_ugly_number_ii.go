package main

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		min := min(2*dp[i2], 3*dp[i3], 5*dp[i5])
		if min == 2*dp[i2] {
			i2++
		}
		if min == 3*dp[i3] {
			i3++
		}
		if min == 5*dp[i5] {
			i5++
		}
		dp[i] = min
	}

	return dp[n-1]
}

func min(a, b, c int) int {
	m := a
	if m > b {
		m = b
	}

	if m > c {
		m = c
	}

	return m
}
