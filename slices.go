package main

import "fmt"

func main(){
	
	//Propiedades de un slice(3 datos)
	//
	//Puntero al arreglo que esta referenciando
	//Longuitud del arreglo al que apunta 
	//Capacidad 
	
	var slice []int 
	slice_nv := []int {1,2,3}

	if slice == nil {			//Si es nula o no esta inicilizado el slice
		fmt.Println("La matriz esta vacia")
	} else {
		fmt.Println(slice)
		fmt.Println(slice_nv)
	}

	//Convertir un arreglo a slice 
	arreglo := [3]int {1,2,3}
	//esto se conoce como slicing
	slice = arreglo[:]			//Se define la posicion a imprimir [inicial:final]

	for i := 0; i < len(arreglo)+1; i++ {	//Se agrega +1 para tomar el ultimo elemento del arreglo
		slice = arreglo[:i]
		fmt.Println("Se convirtio un arreglo a un slice: ", slice)

		slice = arreglo[i:]
		fmt.Println("Se convirtio un arreglo a un slice(invertido): ", slice)
	}

}