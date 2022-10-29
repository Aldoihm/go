package main

import (
	"fmt"

	"github.com/Aldoihm/go/greet"
)

func main() {
	fmt.Println(greet.Italian())

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
*/
