package main

func trailingZeroes(n int) int {
	cnt := 0
	for n/5 > 0 {
		cnt += n / 5
		n = n / 5
	}
	return cnt
}
