package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("zarkoPichula.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)

	f.WriteString("pichula")
	f.Close()
}
