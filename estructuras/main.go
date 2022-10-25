package main

import "fmt"

func main() {
	//las estructuras dejan almacenar una colección de campos, como las clases en otros lenguajes de programación
	type course struct {
		Name     string
		Profesor string
		Country  string
	}

	db := course{
		Name:     "Bases de Datos",
		Profesor: "Aldo",
		Country:  "México",
	}
	git := course{"Git", "Beto", "Bolivia"}
	css := course{Profesor: "Albaro"}            //si no asgino valor a los campos, go deja valor a 0 dependiendo del tipo de dato
	fmt.Printf("Bases de datos: %+v\n", db.Name) //v= valor, +=nombre del campo
	fmt.Printf("Git: %+v\n", git)
	fmt.Printf("CSS: %+v\n", css)

	//Para acceder a cada uno de los campos usamos "." punto

	//PUNTEROS A ESTRUCTURAS
	p := &db
	fmt.Printf("Puntero: %+v\n", p)
	//cambiar valores desde el puntero
	(*p).Profesor = "Alexis"
	p.Country = "Colombia"
	fmt.Printf("Base d Datos: %+v\n", db)
}
