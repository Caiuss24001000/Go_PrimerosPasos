package main 

import "fmt"

func main(){
	
	slice := []int{1,2,3,4,5,6,7,8,9}
	
	/* 
		Agrega en la ultima posición del arreglo el dato ademas si la capacidad 
		es superada duplica su capacidad
	*/
	//slice = append(slice, 2)		

	/*Creando arreglos con la función make*/
	copia := make([]int, len(slice), cap(slice)*2)		//make(Type the array, leng, capacity)

	fmt.Println("The capacity of slice is:", cap(slice))
	fmt.Println("The capacity of copy is:", cap(copia))

	/* 
		Usando la función copy
		Solo copea la cantidad minima de los arreglos (considera el mas chico)

		Para solucionar la deficiencia de que solo se copen unos datos se puede pasar
		la función len() para obtener la longitud del arreglo a copear
	*/
	copy(copia, slice)			//copy(copy, source)

	fmt.Println("slice:\t", slice)
	fmt.Println("copy:\t", copia)

	copia = append(copia, 5)		//Agrega en la ultima posición del arreglo 
	fmt.Println(copia,"copy con capacity\t", cap(copia))
	fmt.Println(slice,"slice con capacity\t", cap(slice))
}