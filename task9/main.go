package main 
import "fmt"

func showBalance(n float64) {
	fmt.Printf("Ваш текущий баланс: %f", n)
}

func putMoney(n *float64, x float64) {
	if x <= 0 {
		fmt.Println("Сумма должна быть положительной")
		return
	}

	*n = *n + x

	fmt.Printf("На ваш счет зачислено: %f, Текущий баланс: %f", x, *n)

}

func withdrawMoney(n *float64, x float64) {
	
	if x <= 0 {
		fmt.Println("Сумма должна быть положительной")
		return
	}

	if *n == 0 {
		fmt.Println("Ваш Баланс нулевой, необходимо сначала внести деньги")
		return
	} 

	if *n - x < 0 {
		fmt.Printf("Недостаточно средств на счете, Текущий баланс: %f", *n)
		return
	} else {
		*n = *n - x 
		fmt.Printf("С вашего счета снято: %f, Текущий баланс: %f", x, *n)
	}

}

func main(){
	var n float64 
	fmt.Scan(&n)

	for true {
		var x string 
		fmt.Scan(&x)

		switch x {
			case "Баланс":
				showBalance(n)
			case "Положить":
				var amount float64 
				fmt.Scan(&amount)
				putMoney(&n, amount)
			case "Снять":
				var amount float64 
				fmt.Scan(&amount)
				withdrawMoney(&n, amount)
			default:
				fmt.Println("Неизвестная команда")

		}
	}

}