package main

import (
	"./filemode"
	"os"
)

/**
Importerer "filemode" mappen og kjøer FileInfo() med fil hentet igjennom os.Args[]
 */
func main(){
	filename := os.Args[1]
	filemode.FileInfo(filename)
}