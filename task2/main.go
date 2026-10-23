package main
import "fmt"


func isPositiv(n int) bool {
	if n >= 0 {
		return true 
	} else {
		return false 
	}
}

func isEven(n int) bool {
	if n % 2 == 0 {
		return true 
	} else {
		return false 
	}
}

func isThreeFiveDivision(n int) bool {
	if n % 3 == 0 && n % 5 == 0 {
		return true 
	} else {
		return false 
	}
}



func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Число равно нулю")

	} else {
		if isPositiv(n) {
			fmt.Println("Число положительное")
		} else {
			fmt.Println("Число отрицательное")
		}

		if isEven(n) {
			fmt.Println("Число четное")
		} else {
			fmt.Println("Число нечетное")
		}

		if isThreeFiveDivision(n) {
			fmt.Println("Число делится на 3 и 5 без остатка")
		} else {
			fmt.Println("Число не делится на 3 и 5 без остатка")
		}

	}

}
