package main

import (
	"fmt"

	"github.com/Aldoihm/go/greet"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(greet.Italian())
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())

}

/* Cómo usar los paquetes que nosotros creamos o paquetes de terceros: con MODULOS
Los Modulos nos ayudan administrar las dependencias de nuestros paquetes y controlar las versiones de estos paquetes

Podemos tener diferentes módulos en nuestro repositorio
Pero se recomienda que tengamos uno solo en la raíz de nuestro proyecto

go mod init: crea un nuevo modulo de go
Se recomienda alojarlo en la ruta donde se encuentra nuestro código fuente

Ejemplo:
go mod init github.com/Aldoihm/go

Intenté varias maneras.. pero la de arriba es la correcta
NOTAS IMPORTANTES PORQUE TUVE VARIOS ERRORES

1.- La ruta del módulo: github.com/Aldoihm/go/greet es independiente a la ruta donde se crea el archivo go.mod, son 2 cosas distintas, porfavor no confundir
*/

/* DESCARGAR UN PAQUETES DE TERCEROS TAMBIÉN LES DICEN DEPENDENCIAS
link donde se encontraba alojado el repositorio
https://pkg.go.dev/rsc.io/quote


RESUELTO:
	1.-Ejecuta el comando go get que se muestra abajo
	2.-Ejecuta el comando go build para generear el ejecutable
	3.-Ejecutamos con ./go (go fue porque me genero el archivo go.exe))

go get rsc.io/quote:
	1.-se descarga e instala el paquete : para ver donde se descarga seguir la siguiente lista de istrucciones
	2.-se modifico el archivo go.mod agregando en el archivo el paquete que se desargó
	3.-se crea el archivo go.sum: lleva el control de las dependencias directas e indirectas con las que se está trabajando
	Dependencias directas: el paquete que nosotros descargamos (quote)
	Dependencias indirectas: los paquetes que necesitaba el paquete que descargamos (sample, text)
	4.- Tabien el archivo go.sum lleva el control de las versiones de los paquetes a través de un hash, que es un id pero el teacher les dice hash
	Con el fin de asegurar que no se han modificado los paquetes de forma accidental o forma maliciosa
	NOTA: Después de ejecutar este comando, nos aparecerá las funciones después de que coloquemos el punto. de las nuevas dependencias

go build: genera un archivo ejecutable o binario
	Este comando se ejecuta en el directorio donde esté el main.go de nuestra app
*/

/* Ver donde se descarga nuestros paquetes de terceros
go env: muesta go la lista de todas las variables que usa go para trabajar
GOPATH: Es una variable
		Nos muestra el directorio donde se creo el gopath: "Raíz del proyecto"
		Dentro del gopath, es donde se clonan nuestras dependencias
		1.- En mi caso fue en: C:\Users\aldo_\go
		2.- Ir a esa ruta y deberia encontrarse la carpeta pkg, entrar en ella
		3.- Debería estar la carpeta mod (modulos), aqui se descargan nuestras dependencias, entramos
		4.- Listamos por fechas con el comando ll -t
		5.- Encontramos la carpeta que se descargo: src.io. Es el dominio donde se encuentra el repositorio que clonamos
		6.- Ingresar a la carpeta y listamos y
		7.- Deberíamos encontrar el paquete que estamos usando y su versión en mi caso : quote@v1.5.2


*/
