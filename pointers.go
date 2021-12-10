package main

import "fmt"

func main(){
	/*
		1.- Es una dirección de memoria
		2.- En lugar de obtener el valor obtenemos la dirección
		3.- x, y apuntan a la misma dirección de memoria (=> as123d => 6)
		4.- x altera el valor en memoria => as123d => 6
		5.- y quiere obtener el valor y seria => 6

		Declarate pointer
		*T => *int, *string, *Struct
		
		Default value is zero => nil
	 */
	
	var x,y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)


}