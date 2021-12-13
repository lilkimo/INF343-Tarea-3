package main

import (
	"fmt"
)

type data struct {
	comandos []string
	reloj []int
	servidor string
}

func main() {
	// {planeta: (reloj, servidor)}
	informacion := make(map[string]data)
	
	var comando string
	var arg1 string
	var arg2 string
	var arg3 string

	for {
		fmt.Println("dou")
		fmt.Scanf("%s %s %s %s\n", &comando, &arg1, &arg2, &arg3)
		if (arg1 == "") || (arg2 == "") {
			fmt.Println("Entrada inválida, intente nuevamente. (1)")
			continue
		}
		if (comando == "AddCity") || (comando == "UpdateName") || (comando == "UpdateNumber") {
			if arg3 == "" {
				fmt.Println("Entrada inválida, intente nuevamente. (2)")
				continue
			}
		} else if comando != "DeleteCity" {
			fmt.Println("Entrada inválida, intente nuevamente. (3)")
			continue
		}
		break
	}
	var strcmd string
	if comando == "DeleteCity" {
		strcmd = fmt.Sprintf("%s %s %s", comando, arg1, arg2)
	} else {
		strcmd = fmt.Sprintf("%s %s %s %s", comando, arg1, arg2, arg3)
	}

	ip := ObtenerFulcrum(comando, arg1, arg2, arg3)
	reloj := ActualizarRegistros(ip, comando, arg1, arg2, arg3)

	val, ok := informacion[arg1];
	if ok {
		// Si ya existía el registro actualizo reloj y ip (¿Quizás hacer esa weá de
		// agarrar el máximo valor de los componentes del viejo reloj y el nuevo?)
		// y appendeo a los comandos el nuevo comando ejecutado.
		val.reloj = reloj
		val.servidor = ip
		val.comandos = append(val.comandos, strcmd)
	} else {
		// Si la weá no existe la chanta así tal cual. 
		val = data{comandos: []string{strcmd}, reloj: reloj, servidor: ip}
	}
	informacion[arg1] = val
}

func ObtenerFulcrum(comando string, arg1 string, arg2 string, arg3 string) string {
	return "localhost"
}

func ActualizarRegistros(ip string, comando string, arg1 string, arg2 string, arg3 string) []int {
	return make([]int, 3)
}