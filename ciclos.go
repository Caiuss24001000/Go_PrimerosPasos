package main

import "fmt"

func main(){
	fin := 10
	//Ciclo for comun, imprime 10 veces 
	for i := 0; i < fin; i++ {
		fmt.Println("Hola desde i con ", i+1)
	}

	//Ciclo for alterno, se salta el numero 2 
	i := 0
	for {
		if i == 2 {
			i++ 
			continue //Continua con el ciclo(un tipo de break)
		}

		fmt.Println("Valor de i", i)
		i++

		if i>fin{
			break
		}
	}
}