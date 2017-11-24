package main

import "fmt"

func main(){
	a := 1
	for{
		fmt.Print(a, " ")
		a++
		if a == 11{
			break
		}
	}
	fmt.Println()
}

// outputnya adalah
// 1 2 3 4 5 6 7 8 9 10