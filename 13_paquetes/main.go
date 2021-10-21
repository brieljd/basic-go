package main

import (
	"fmt"

	"../13_paquetes/saludar"

	depedida "../13_paquetes/despedida"
)

func main() {

	saludar.Saludar("Gabriel")
	saludar.MeVes = "Esto es un string asignado desde el main a la variable en el paquete"
	fmt.Println(saludar.MeVes)

	nombre := "Gabriel"
	depedida.Despedirse(nombre)
}
