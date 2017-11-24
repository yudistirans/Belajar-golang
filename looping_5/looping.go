package main

import "fmt"

func main(){
	for i:=1; i<5; i++{
		if i == 3{
			break
		}
		for j:=1; j<5; j++{
			fmt.Print("Matrik ","[",i,"]","[",j,"]")
			fmt.Println()
		}	
	}
}