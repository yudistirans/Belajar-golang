package main

import "fmt"

func main(){
	a := 5
	b := 5
	c := 10
	d := 8

	fmt.Printf("a=%d,b=%d,c=%d,d=%d \n",a,b,c,d)
	fmt.Println("Apakah a == b ? ", a == b)
	fmt.Println("Apakah a != c ? ", a != c)
	fmt.Println("Apakah c < b ? ", c < b)
	fmt.Println("Apakah d <= c ? ", d <= c)
	fmt.Println("Apakah c > a ? ", c > a)
	fmt.Println("Apakah d >= b ? ", d >= b)
}

/*
Outputnya kode di atas
a = 5 , b = 5 , c = 10 , d = 8
Apakah a == b ?  true
Apakah a != c ?  true
Apakah c < b ?  false
Apakah d <= c ?  true
Apakah c > a ?  true
Apakah d >= b ?  true
*/