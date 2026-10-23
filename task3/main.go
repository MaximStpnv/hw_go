package main 
import "fmt"

func main(){
	var n int; 
	sum := 0 
	fmt.Scan(&n)

	i := 1
	for i <= n {
		sum+=i
		i++
	}

	var avg float64 = float64(sum)/float64(n)

	fmt.Printf("Сумма: %d\nСреднее: %f\n", sum, avg)

}