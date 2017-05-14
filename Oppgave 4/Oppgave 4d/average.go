package main

import (
	"fmt"
	"strconv"
)

/**
Regner ut gjennomsnittslengden på en melding til 100 tilfeldig valgt
 */
func main() {
	// Av bit lengden
	fmt.Println("Antall * (Lengden på melding * sansynlighet) + (Lengden på ......")
	fmt.Print("Lengde på melding til 100 stk fra huffman bit lengde: ")


	temp := (3*0.1518) + (3 * 0.0376) + (3 * 0.1717) + (2 * 0.1823) + (3 * 0.1487) + (2 * 0.3079)
	total := temp * 100
	fmt.Print(total)
	fmt.Print("\nTil binary: ")
	fmt.Print(strconv.FormatInt(int64(total), 2))
	fmt.Println("")
}