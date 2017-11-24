package main

import "fmt"

func main(){
	var a = 8
	switch{
	case a == 10 :
		{
			fmt.Println("Perfect")
			fmt.Println("Amazing")
		}
	case a == 9 : 
		{
			fmt.Println("Awesome")
		}
	default :
		fmt.Println("Not bad")
	}	
}
