package main

import "fmt"

func main() {
	oppgave4a()
}

func oppgave4a() {
	//fakultet
	var f [6]string
	f[0] = "Helse-og idrettsfag \t\t "
	f[1] = "Humaniora og pedagogikk \t "
	f[2] = "Kunstfag \t\t\t\t\t  "
	f[3] = "Teknologi og realfag \t\t "
	f[4] = "Lærerutdanningen \t\t\t "
	f[5] = "Okonomi og samf.vitenskap \t "

	//antall
	var a [6]float64
	a[0] = 1829
	a[1] = 1525
	a[2] =  420
	a[3] = 2166
	a[4] = 1506
	a[5] = 3093

	//prosent
	var tot float64 = 10539

	fmt.Println("UiAs fakultet: \t\t\t\t Antall studenter: \t Sannsynlighet:")
	for i := 0; i < len(a); i ++ {
		fmt.Print(f[i])
		fmt.Print(a[i])
		fmt.Print("\t\t\t\t ")
		fmt.Println(a[i]/tot*100)
	}
}
