package main

import "fmt"

func main(){
	a := true
	b := true
	c := false

	//operator &&
	fmt.Println("true && true ? ", a && b)
	fmt.Println("true && false ? ", a && c)
	fmt.Println("false && true ? ", c && a)

	//operator ||
	fmt.Println("true || true ? ", a || b)
	fmt.Println("true || false ? ", a || c)
	fmt.Println("false || true ? ", c || a)
	
	//operator !
	fmt.Println("!true ? ", !a)
	fmt.Println("!false ?", !c)

}

/*
true && true ?  true
true && false ?  false
false && true ?  false
true || true ?  true
true || false ?  true
false || true ?  true
!true ?  false
!false ? true
*/