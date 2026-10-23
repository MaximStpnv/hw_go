package main

import "fmt"

func main() {
	var n int

	positive_num := 0
	negative_num := 0
	zeroes_num := 0

	fmt.Scan(&n)

	i := 0

	for i < n {
		var x int
		fmt.Scan(&x)

		if x > 0 {
			positive_num++
		} else if x < 0 {
			negative_num++
		} else if x == 0 {
			zeroes_num++
		}

		i++
	}

	fmt.Printf(
		"Положительные числа: %d\nОтрицательные числа: %d\nНули: %d\n",
		positive_num,
		negative_num,
		zeroes_num)

}
