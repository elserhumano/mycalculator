package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println("Suma de los dos operadores matematicamente: ", operador1+operador2)
		return operador1 + operador2

	case "-":
		fmt.Println("Resta de los dos operadores matematicamente: ", operador1-operador2)
		return operador1 - operador2

	case "*":
		fmt.Println("Multiplicacion de los dos operadores matematicamente: ", operador1*operador2)
		return operador1 * operador2

	case "/":
		fmt.Println("Division de los dos operadores matematicamente: ", operador1/operador2)
		return operador1 / operador2

	default:
		fmt.Println("Operacion no soportada")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
