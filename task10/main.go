package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("N должно быть больше нуля")
		return
	}

	sum := 0
	max := 0
	min := 0
	even_num := 0
	odd_num := 0

	i := 0
	for i < n {
		var x int
		fmt.Scan(&x)

		sum += x

		if i == 0 {
			max = x
			min = x
		} else {
			if x > max {
				max = x
			}

			if x < min {
				min = x
			}
		}

		if x % 2 == 0 {
			even_num++
		} else {
			odd_num++
		}

		i++
	}

	var avg float64 = float64(sum) / float64(n)

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее: %.2f\n", avg)
	fmt.Printf("Максимум: %d\n", max)
	fmt.Printf("Минимум: %d\n", min)
	fmt.Printf("Четных чисел: %d\n", even_num)
	fmt.Printf("Нечетных чисел: %d\n", odd_num)
}
