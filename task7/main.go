package main 
import "fmt"

func main(){
	var n int; 
	fmt.Scan(&n)
	sum := 0;

	i :=1 
	for i <= n {
		if i % 2 == 0 && i % 3 == 0 {
			sum += i
		}
		i++
	}

	fmt.Println(sum)

}