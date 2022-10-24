package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	// Arrays
	var array [5]int
	array[0] = 10
	array[1] = 100
	array[2] = 1000
	fmt.Println(array)
	
	array2 := [5]string{"um", "dois", "tres", "quatro"}
	fmt.Println(array2)
	
	array3 := [...]int{1, 2, 3, 4, 5} // baseado no tamanho dos indices que passou, estatico
	fmt.Println(array3)
	
	// Slices
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = append(slice, 19)
	fmt.Println(slice)
	
	slice2 := array2[1:3]
	fmt.Println(slice2)

	fmt.Println("----------")
	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 11)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 14)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade

}
