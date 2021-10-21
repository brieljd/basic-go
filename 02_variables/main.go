package main

import "fmt"

func main() {
	//primer manera de definir variables
	var nombre, apellido string
	nombre = "Gabriel"
	apellido = "Jimenez"
	//segunda manera de definir variables
	apellido2 := "Díaz"
	//tercera manera de definir variables
	usuario1, usuario2 := "valeria", "santo"

	fmt.Println(nombre, apellido, apellido2)
	fmt.Println(usuario1, usuario2)
}
