package main 
import "fmt"

func main(){
	var n, max, min int 
	fmt.Scan(&n)
	i := 0;

	for i < n {
		var x int; 
		fmt.Scan(&x)

		if (i == 0){
			max = x;
			min = x; 
		} else { 
			
			if (x > max){
				max = x
			}

			if (x < min){
				min = x
				
			}
		}
		i++
	}

	fmt.Printf("Максимум: %d\nМинимум: %d\n", max, min)



}