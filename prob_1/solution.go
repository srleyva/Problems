package main

import "fmt"

func main() {
	fmt.Println(numWays(4))
	fmt.Println(numWaysX(12, []int{1, 3, 5, 6}))
}

func numWays(n int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	ways[1] = 1
	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n]
}

func numWaysX(n int, x []int) int {
	y := 0
	for _, i := range x {
		if n >= i {
			y += numWays(n - i)
		}
	}
	return y
}
