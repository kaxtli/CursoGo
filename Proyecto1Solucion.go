package main

import "fmt"

func main() {
	var dolar, euro float32
	var pesos float32

	fmt.Println("TransforDivisas v0.1")
	fmt.Println("Ingrese los datos solicitados a continuación.")
	fmt.Print("Precio del Euro (mxn)> ")
	fmt.Scan(&euro)
	fmt.Print("Precio del Dólar (mxn)> ")
	fmt.Scan(&dolar)

	for {
		fmt.Print("Dinero(mxn) -1 para salir>")
		fmt.Scan(&pesos)

		if pesos < 0 {
			break
		}

		dolares := pesos / dolar
		euros := pesos / euro

		fmt.Printf("MXN: %.2f | EUR: %.2f | USD: %.2f\n\n", pesos, euros, dolares)
	}
}