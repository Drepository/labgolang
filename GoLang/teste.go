package main

import "fmt"

func main() {
	sequencia := []int{5, 10, 15, 21}
	resultado := 55
	testasequencia := soma(sequencia)

	fmt.Println("")

	if testasequencia == resultado {
		fmt.Print("PASSOU...")
	} else {
		fmt.Print("FALHOU...")
	}

	fmt.Println("Valor alcançado:", testasequencia, "esperado:", resultado)

	fmt.Println("")

}

func soma(numeros []int) int {
	somanumeros := 0
	for _, numero := range numeros {
		somanumeros += numero
	}
	return somanumeros
}
