package main

import "fmt"

func main() {
	fruit := "🍌"
	//Variable de tipo puntero
	var p *string
	p = &fruit
	//Se usa el & ampersan para decir que quiero ver la direccion de la variable
	fmt.Printf("Variable Normal: Tipo: %T, Valor: %v, Dirección: %v \n", fruit, fruit, &fruit)
	//* Hace referencia al valor que se almacena en donde apunta el puntero
	fmt.Printf("Variable Puntero: Tipo %T, Valor: %v, Desreferenciación: %v \n", p, p, *p)
	*p = "🍍"
	fmt.Printf("Variable Normal: Tipo: %T, Valor: %v, Dirección: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Variable Puntero: Tipo %T, Valor: %v, Desreferenciación: %v \n", p, p, *p)
}
