package variables

import(
	"fmt"
	"time"
	"strconv"
)
var Nombre string;
var Estado bool;
var Sueldo float32;
var Fecha time.Time;

func RestoVariables(){
	Nombre = "Elfar";
	Estado = true;
	Sueldo = 11.562;
	Fecha = time.Now();
	fmt.Println("Nombre = ",Nombre);
	fmt.Println("Estado = ",Estado);
	fmt.Println("Sueldo = ",Sueldo);
	fmt.Println("Fecha = ",Fecha);
}

func ConvertirATexto(numero int) (bool,string){
	var texto string;
	texto = strconv.Itoa(numero);
    return true,texto;
}