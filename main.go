package main

import (
	"fmt"
	//"runtime"
	"github.com/elfar/godesdecero/ejercicios"
	//"github.com/elfar/godesdecero/variables"
)

func main()  {
	//fmt.Println("Hola Mundo")
 //variables.GetEnteros();
 //variables.RestoVariables();
/*
 estado,texto := variables.ConvertirATexto(1524)
fmt.Println(estado);
fmt.Println(texto);*/

// if so := runtime.GOOS; so == "Linux." || so == "so x." {
// 	fmt.Println("Esto no es windows, es ",so)
// }else{
// 	fmt.Println("Esto es windows")
// }

// switch so := runtime.GOOS; so {
// case "linux":
// 	fmt.Println("Esto es linux")
// case "darwid":
// 	fmt.Println("Esto es darwind")
// default: 
// 	fmt.Printf("%s \n ",so)	
// }

numero,texto := ejercicios.ConvNumerico("200")
fmt.Println("numero ", numero)
fmt.Println("numero ", texto)
}