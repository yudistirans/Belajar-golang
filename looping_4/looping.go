package main

import "fmt"

func main(){
	for i:=1; i<=30; i++{
		if i%3 != 0 {
			continue
		}

		if i > 25{
			break
		}

		fmt.Print(i, " ")
	}
	fmt.Println()
}

// outputnya adalah
// 3 6 9 12 15 18 21 24