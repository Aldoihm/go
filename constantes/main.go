package main

import "fmt"

// Constantes a nivel de paquete y en forma de agrupaciones
// Las constantes a diferencia de las variables, no necesitan ser usadas forzosamente en el programa para que compile
const (
	os     = "linux"
	domain = "ed.team"
)

// Constante con iota para las secuencias, comenzando desde cero, ejemplo
const (
	Ene = iota + 1
	Feb
	Mar
	Abr
	May
	Jun
)

func main() {
	/*1ra forma para declarar constantes
	const os string = "linux"*/

	/*Segunda forma para declarar constantes
	const os, domain string = "linux", "ed.team"*/

	//No se puede usar el operador de variable corta

	/*Tercera forma
	const os, domain = "linux", 1*/

	fmt.Println(os, domain)
	fmt.Println(May)
}
