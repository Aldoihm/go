package main

import "fmt"

func main() {
	/*	if condition {

		} else if condition {

		} else {

		}
	*/ //sin parentesis
	/* 	emoji := "๐"
	   	if emoji == "๐ต" {
	   		fmt.Println("Es un cactus")
	   	} else if emoji == "๐" {
	   		fmt.Println("Es una carita")
	   	} else {
	   		fmt.Printf("El emoji es: %v\n", emoji)
	   	} */

	//Declarando la variable y al mismo tiempo usando if
	//solo se podrรก usar la variable dentro de la estructura de if, por fuera ya no
	/* 	if variable; condition{

	   	}else if condition{

	   	}else{

	   	} */

	if emoji := "๐"; emoji == "๐ต" {
		fmt.Println("Es un cactus")
	} else if emoji == "๐" {
		fmt.Println("Es una carita")
	} else {
		fmt.Printf("El emoji es: %v\n", emoji)
	}
}
