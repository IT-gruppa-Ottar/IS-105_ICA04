package main

import (
	"./filemode"
	"os"
)


func main(){
	filename := os.Args[1]
	filemode.FileInfo(filename)
}