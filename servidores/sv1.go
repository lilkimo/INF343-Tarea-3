package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("zarkoPichula.txt", os.O_APPEND|os.O_CREATE, 0666)

	f.WriteString("pichula")
	f.Close()
}
