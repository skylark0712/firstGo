// https://leetcode.cn/problems/count-collisions-on-a-road
package main

import "fmt"

func main() {
	fmt.Println(countCollisions("kklk"))
}

func countCollisions(directions string) int {
	n := len(directions)
	if n <= 1 {
		return 0
	}
	// stopped := make([]bool, n)
	// 不想写了，开摆
	total := 0
	return total
}
