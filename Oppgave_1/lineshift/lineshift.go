package lineshift

import (
	"fmt"
	"strings"
	"../filutils"
)

/**
Tar filen og kjører "strings.Contains" på "\n" og "\r".
Retrunerer true eller false

@param filnavn
 */
//noinspection GoUnusedExportedFunction
func LineShift(filnavn string) {
	byteSlice := fileutils.FileToByteslice(filnavn)

	s := string(byteSlice[:])

	fmt.Print("Inneholder filen linjeskift: ")
	fmt.Println(strings.Contains(s, "\n"))

	fmt.Print("Inneholder filen carriage return: ")
	fmt.Println(strings.Contains(s, "\r"))
}