package ejercicios
import ("strconv")

func ConvNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "hubo un error: " + err.Error()
	}

	if num > 100 {
		return num, "es mayor de 100"
	} else {
		return num, "es menor de 100"
	}
}