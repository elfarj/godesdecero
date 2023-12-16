package main

import (
	"fmt"
	
	"github.com/elfar/godesdecero/variables"
)

func main()  {
	//fmt.Println("Hola Mundo")
 //variables.GetEnteros();
 //variables.RestoVariables();
 estado,texto := variables.ConvertirATexto(1524)
fmt.Println(estado);
fmt.Println(texto);
}