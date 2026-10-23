package main
import "fmt"


func printTable(n int ){
	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d", i, j, i*j)
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}

}

func main(){
	var n int; 
	fmt.Scan(&n)
	printTable(n)
}
