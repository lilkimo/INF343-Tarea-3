package main

import (
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("zarkoPichula.txt", []byte("pichula"), os.ModeAppend)

	if err != nil {
		panic(err)
	}
}
