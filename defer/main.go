package main

import (
	"fmt"
	"os"
)

func main() {
	//-------------------------PARTE 1--------------------------------------------
	//defer = deferir, aplazar la ejecución hasta el final de la función
	//Crea una pila donde se mete cada tarea que tiene la palabra defer y al final va sacando 1 x 1, recordar que en un pila, lo primero que entra es lo último que sale y lo último que sale, es lo primero que sale
	/* defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	a := 5
	defer fmt.Println("Defer:", a)
	a = 10
	fmt.Println(a) */

	//---------------------------PARTE2------------------------------------------
	//USUS:
	//Limpiar recursos, cerrar archivos, cerrar conexiones de red o cerrar controladores de bases de datos
	file, err := os.Create("hello.txt") //Crea un archivo y retorno 2 opciones
	//1.- Un puntero que señala a un archivo File
	//2.- Un error
	if err != nil {
		fmt.Printf("Ocurrió un error al crear un archivo: %v \n", err)
		return
	}
	defer file.Close()
	_, er := file.Write([]byte("Hello Gophers")) //Escribe en el archivo y devuelve 2 resultados
	//1.- número de bytes que fueron escritos
	//2.- un error
	if er != nil {
		fmt.Printf("Ocurrio un error al escribir el archivo: %v \n", er)
		return
	}

}
