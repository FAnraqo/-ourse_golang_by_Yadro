package main

import (
	"fmt"
)

func main() {
	mas := []int{}
	flag := false
	for !flag {
		var tmp int
		fmt.Scan(&tmp)
		if tmp != 0 {
			mas = append(mas, tmp)
		} else {
			flag = true
		}
	}

	var k int
	fmt.Scan(&k)

	tmp_k := k
	tmp := make([]int, len(mas), cap(mas))
	copy(tmp, mas)

	max := 0
	for tmp_k != 0 {
		max = 0
		max_i := -1
		for i := 0; i < len(tmp); i++ {
			if tmp[i] > max {
				max = tmp[i]
				max_i = i
			}
		}
		tmp = append(tmp[0:(max_i)], tmp[(max_i+1):]...)
		tmp_k -= 1
	}

	fmt.Print(mas, k, "-й максимум:", max, "\n")
}
