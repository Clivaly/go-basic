package main

import "fmt"

func main() {
	var variavel1 int 
	var variavel2 *int 

	variavel1 = 10
	variavel2 = &variavel1
	
	fmt.Println(variavel2, variavel1)
	
	variavel1 = 100
	fmt.Println(variavel2, variavel1)
}
