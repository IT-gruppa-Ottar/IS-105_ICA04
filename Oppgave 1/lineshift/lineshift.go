package lineshift

import (
	"fmt"
	"strings"
)

//noinspection GoUnusedExportedFunction
func LineShift(filnavn []byte) {
	s := string(filnavn[:])

	//Viser stringen som testes
	//fmt.Printf("%q", s)
	//fmt.Println()


	fmt.Print("Inneholder filen linjeskift: ")
	fmt.Println(strings.Contains(s, "\n"))

	fmt.Print("Inneholder filen carriage return: ")
	fmt.Println(strings.Contains(s, "\r"))
}

