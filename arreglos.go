package main

import "fmt"

func main(){

	//Declaracion de dos tipos de arreglos 
	var arreglo [10]int 		//Arreglo vacio
	numeros := [3]int {1,2}		//Arreglo inicializado

	fmt.Println(arreglo)		//Impresion de arreglo
	fmt.Println(numeros)		//Impresion de arreglo
	fmt.Println(len(arreglo))	//Impresion del tamaño del arreglo

	//Recorriendo todo el arreglo
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = i*2 		//Definiendo valor al arreglo
		fmt.Println("El dato es ", arreglo[i])
	}

	//Declaración de una matriz 3x2
	var matriz [3][2]int 

	//Asignamos el valor de la posición (0,1)
	matriz[0][1] = len(arreglo)

	fmt.Println(matriz)			//Impresión de la matriz
}