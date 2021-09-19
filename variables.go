//Paquete main necesario para cualquier programa 
package main

// Se importan los paquetes necesarios
import (
    "fmt" 
    "bufio" 
    "os"
)

func main(){
	var x int 
	x = 25
	y := 10
	z := 200

	cadena := "Hola soy una cadena"

	fmt.Println(cadena)    //Imprime una cadena en pantalla con salto de linea
	
    //Imprime un numero con un formato especifico(parecido a C)
    fmt.Printf("El valor de x es %d\n", x)
	fmt.Printf("El valor de y es %d\n", y)
	fmt.Printf("El valor de z es %d\n", z)
    
    //En este caso no es necesario especificar el tipo de dato ya que solo con colocar la variable se puede imprimir, separandolo por una coma 
    fmt.Println("El valor de x es", x,y,z, cadena)

    cadena = "Se modifico"
    fmt.Printf("Usando el parametro v que seria %v\n", cadena)

    //Pedir datos por consola con scanf
    var edad string
    fmt.Println("Coloca una edad: ")
    fmt.Scanf("%v\n", &edad)
    fmt.Printf("Mi edad es %s\n", edad)

    //Pedir datos por consola con Reader
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ingresa un nombre: ")

    //Termina cuando ingresa un caracter especifico 
    nombre, err := reader.ReadString('\n')  //Las comillas simples identifican un caracter
    
    if(err != nil){
        fmt.Println(err)
    }else {
        fmt.Println("Hola "+nombre)
    }

    //Compracion de cual es el mayor de 3 numeros opción 1
    if x<y {
        if y>z {
            fmt.Println("y es el numero mas grande", y)
        } else {
            fmt.Println("z es el numero mas grande", z)
        }
        
    } else if x>z {
        fmt.Println("x es el numero mas grande", x)
    } else {
        fmt.Println("z es el numero mas grande", z)
    }

    //Compracion de cual es el mayor de 3 numeros opción 2
    if x<y && z<y {
        fmt.Println("y es el numero mas grande", y)
    } else if z<x {
        fmt.Println("x es el numero mas grande", x)
    } else {
        fmt.Println("z es el numero mas grande", z)
    }
}