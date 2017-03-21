package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Av bit lengden
	fmt.Println("Antall * (Lengden på melding * sansynlighet) + (Lengden på ......")
	fmt.Print("Lengde på melding til 100 stk fra huffman bit lengde: ")
	tott := 100*(3*0.1518) + (3 * 0.0376) + (3 * 0.1717) + (2 * 0.1823) + (3 * 0.1487) + (2 * 0.3079)
	fmt.Print(tott)
	fmt.Print("\nTil binary:")
	fmt.Print(strconv.FormatInt(int64(tott), 2))
	fmt.Println("")
}

