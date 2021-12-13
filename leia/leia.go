package main

import (
	"fmt"
)

type data struct {
	cantRebeldes int
	reloj []int
	servidor string
}

func main() {
	// {ciudad: (cantRebeldes, reloj, servidor)}
	informacion := make(map[string]data)
	
	var comando string;
	var arg1 string;
	var arg2 string;

	for {
		fmt.Scanf("%s %s %s\n", &comando, &arg1, &arg2)
		if (comando != "GetNumberRebelds") || (arg1 == "") || (arg2 == "") {
			fmt.Println("Entrada inválida, intente nuevamente.")
		} else {
			break
		}
	}
	cantRebeldes, reloj, ip := ObtenerCantRebeldes(arg1, arg2)
	// Asumo que dos ciudades, aunque estén en diferentes planetas),
	// no pueden tener el mismo nombre.
	val, ok := informacion[arg2]
	if ok {
		// Aquí habría que aplicar Monotonic Reads, ni idea de cómo la verdad xd.
		// Me imagino que hay que revisar el reloj o weás así no sé nada xuxetumare.
		val.cantRebeldes = cantRebeldes
		val.reloj = reloj
		val.servidor = ip
	} else {
		// Si la weá no existe la chanta así tal cual. 
		val = data{cantRebeldes: cantRebeldes, reloj: reloj, servidor: ip}
	}
	informacion[arg2] = val
}

func ObtenerCantRebeldes(nombrePlaneta string, nombreCiudad string) (int, []int, string) {
	return 3, make([]int, 3), "localhost"
}