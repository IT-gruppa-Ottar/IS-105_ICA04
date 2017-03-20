package main

import (
	"./filutils"
	"./lineshift"
	//"fmt"
	//"strings"
	//"bytes"
	//"fmt"
	//"bytes"
)


func main() {
	fileutils.FileToByteslice("text1.txt")
	fileutils.FileToByteslice("text2.txt")



	println("-----1b------")

	lineshift.LineShiftCheck()


	//Alternative måter jeg prøvde å løse det på. /Fungerer ikke)
	//fmt.Println(bytes.Contains([]byte(byteSlice), []byte("10")))

	//fmt.Println(bytes.ContainsAny(fileutils.FileToByteslice("text1.txt", "10")


}
