package main

import (
	"fmt"
)

func preimageSizeFZF(K int) int {
	step := 0
	for step < K {
		step = step*5 + 1
	}

	for K > 0 {
		step = (step - 1) / 5
		if K/step == 5 {
			return 0
		}

		K %= step
	}

	return 5
}

func main() {
	fmt.Printf("%d\n", preimageSizeFZF(6))
}
