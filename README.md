# Go_PrimerosPasos
Se listan pequeños ejemplos y ejercicios del lenguaje Go  

1. [DOWNLOAD GO](#download-go)
2. [Installation in Linux](#installation-in-linux)
3. [Creating Go Workspace](#creating-go-workspace)
    * [Optional directory](#optional-directory)
4. [Global variables](#glonal-variables)
5. [Creating Hello Word](#hello-world)

### __DOWNLOAD GO__  
Visita [https://golang.org/dl/][DOWNLOADGO] para descargar GO desde su pagina oficial y selecciona el adecuado para tu sistema operativo  


### __Installation in Linux__  
Extraer el archivo que se descargo previamente en el directorio `/usr/local`. Se recomienda borrar el contanido de la dirección `/usr/local/go` en caso de una instalación previa  

**Example**  

```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz
```  

La versión de Go que se usa en este comando esta actualizada para el 2021, si su versión es diferente a la de este ejemplo cambiela por la suya  

Nota: Los comandos podrían necesitar permisos root o sudo, en caso de error usarlos independientemente  


### __Creating Go Workspace__  
Ahora crearemos nuestro espacio de trabajo, normalmente este espacio usa dos directorios principales los cuales son `src` y `bin`  

Normalmente nuestro `Go_workspace` se ubica en `$HOME` por lo que nuestro arbol seria:

- HOME
    - Go_workspace

**Example**  

`mkdir ~/Go_workspace` o `mkdir $HOME/Go_workspace` cualquiera de las dos instrucciones deberia de funcionar  

Ya que tenemos nuestro directorio de trabajo ahora es crear nuestros directorios secundarios `src` y `bin`  

- Go_workspace
    - src: el cual contiene los archivos fuente de Go
    - bin: el cual contiene los ejecutables y las herramientas de Go

**Example**  

```bash
cd ~/Go_workspace
mkdir bin src
```  

#### __Optional directory__  
Si se trabaja con github en varios proyectos es recomendable usar un directorio `github.com` para guardarlos  

Tambien podemos crear un un directorio con el usuario con el que se este trabajando un proyecto determinado 

- src
    - github.com(optional) 
        - NombreDeUsuario(optional) 

**Example**  

```bash
cd ~/Go_workspace/src
mkdir github.com/NombreDeUsuario
```  

### __Global variables__  
Se agregaran las variables `$GOPATH` y `$PATH`, si bien la variable `$GOPATH` ya no es tan usada, algunas herramientas aun la usan ademas de ser una buena practica el incluirla  

Para agregar las variables globales en nuestra shell debemos de saber en que archivo se deben de agregar, si nunca has modificado la configuración de tu shell seria en `~/.profile` pero en caso de que si se hubiera cambiado por una zsh o alguna otra deberia de ser en `.zshrc` o `.bashrc` segun corresponda  

Abrimos el archivo `~/.profile` con su editor favorito en este caso usaremos nano  

```bash
nano ~/.profile
```  

Añadiendo `$GOPATH` y `$PATH` al archivo 
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin 
```  
Esto nos permitira ejecutar cualquier programa que compilemos o se descarge dentro de nuestro sistema  

Actualizamos la shell para cargar las variables globales en el sistema  

```bash
. ~/.profile
```

Para verificar que nuestro PATH este actualizado usaremos el comando `echo` el cual nos listara las direcciones que contiene el PATH, al final deberia de estar la dirección /usr/local/go/bin  

```bash
echo $PATH
```

Para verificar la instalación usamos el comando `go version` el cual nos listaria la version que tenemos instalada 

```bash 
go version
```  

`Output >> go version go1.17.2 linux/amd64`  

### __Creating Hello World__  

Ya que tenemos nuestro espacio de trabajo(Workspace), las variables de entorno estan actualizadas, ahora hagamos nuestro primer hola mundo  

En nuestro directorio home abrimos un archivo con nuestro editor nano  

```bash
nano hola.go
```  

Escribimos el siguiente codigo en nuestro archivo `hola.go`  

```go
package main

import "fmt"

func main(){
    fmt.Println("Hola mundo")
}
```  

Tambien puede descargar el archivo `hola.go` dentro de este repositorio con `wget`  

```bash 
wget https://raw.githubusercontent.com/Caiuss24001000/Go_PrimerosPasos/main/hola.go
```  

Ahora para correr usamos la siguiente instrucción

```bash 
go run hola.go
```
`Output >> Hola Mundo`  

<!-- Enlaces requeridos por el Readme -->
[DOWNLOADGO]: https://golang.org/dl/