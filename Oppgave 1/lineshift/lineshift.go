package lineshift

import (
	"fmt"
	"strings"
	"../filutils"
)

//noinspection GoUnusedExportedFunction
func LineShift(filnavn string) {

	byteSlice := fileutils.FileToByteslice(filnavn)

	s := string(byteSlice[:])

	//Viser stringen som testes
	//fmt.Printf("%q", s)
	//fmt.Println()


	fmt.Print("Inneholder filen linjeskift: ")
	fmt.Println(strings.Contains(s, "\n"))

	fmt.Print("Inneholder filen carriage return: ")
	fmt.Println(strings.Contains(s, "\r"))
}

