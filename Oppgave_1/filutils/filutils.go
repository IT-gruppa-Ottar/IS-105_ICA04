/**
Hentet fra faglærers repository
https://github.com/uia-worker/is105-ica03/tree/master/fileutils
 */
package fileutils

import (
	"io"
	"log"
	"os"
)

/**
Tar filnavnet og leser innholdet. Deretter returnerer det i en byteslice

@param filename av typen string
@return []byteslice av innholdet til filen
 */
//noinspection GoUnusedExportedFunction
func FileToByteslice(filename string) []byte {

	// Open file for lesing
	file, err := os.Open(filename)

	//Error melding
	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	return byteSlice
}