package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("zarkoPichula", os.O_APPEND, 0666)

	f.WriteString("pichula")
	f.Close()
}
