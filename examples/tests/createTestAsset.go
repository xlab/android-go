// +build ignore

// This program creates a 1MB file whose bytes set to o%256, where o is the offset of the byte in the file, starting with 0.
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	var bytes = make([]byte, 1<<20)
	for i := range bytes {
		bytes[i] = byte(i)
	}
	var err = ioutil.WriteFile("assets/testAsset", bytes, 0666)
	if err != nil {
		log.Fatalln(err)
	}
}
