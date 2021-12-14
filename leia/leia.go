package main

import (
	"fmt"
	//"math/rand"
	//"strconv"
	//"time"

	"context"
	"log"

	pb "inf343-tarea-3/protoBrokerLeia"
	"google.golang.org/grpc"
)



const (
	address = "dist16:50052"
)

type data struct {
	cantRebeldes string
	reloj [] int32
	servidor string
}
func main() {
	direccionToId := make(map[string]int)
	direccionToId["dist13:50061"] = 0
	direccionToId["dist14:50063"] = 1
	direccionToId["dist15:50065"] = 2
	
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewConnToBrokerFromLeiaClient(conn)

	ctx := context.Background()
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
	}
	// {ciudad: (cantRebeldes, reloj, servidor)}
	informacion := make(map[string]data)
	
	var comando string;
	var arg1 string;
	var arg2 string;

	fmt.Println("Presione ENTER sin ingresar un comando pasa salir")
	for {
		for {
			fmt.Print("Ingrese comando: ")
			fmt.Scanf("%s %s %s\n", &comando, &arg1, &arg2)
			if (comando == "") && (arg1 == "") && (arg2 == "") {
				return
			} else if (comando != "GetNumberRebelds") || (arg1 == "") || (arg2 == "") {
				fmt.Println("Entrada inválida, intente nuevamente.")
			} else {
				break
			}
		}

		// Asumo que dos ciudades, aunque estén en diferentes planetas),
		// no pueden tener el mismo nombre.
		val, ok := informacion[arg1]
		if ok {
			// Aquí habría que aplicar Monotonic Reads, ni idea de cómo la verdad xd.
			// Me imagino que hay que revisar el reloj o weás así no sé nada xuxetumare.
			rS, err := c.GetNumberRebelds(ctx, &pb.MensajeToBrokerFromLeia{Comando: comando, NombrePlaneta: arg1, NombreCiudad: arg2, IpServidorFulcrum: val.servidor})
			if err != nil {
				log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
			}
			cantRebeldes := rS.GetNumeroRebeldes()
			reloj := rS.GetVector()
			ip := rS.GetIpServidorFulcrum()

			if (val.reloj[direccionToId[ip]] <= reloj[direccionToId[ip]]){
				val.cantRebeldes = cantRebeldes
				fmt.Printf("Cantidad de rebeldes %s\n", cantRebeldes)
				val.reloj = reloj
				val.servidor = ip
			} else {
				fmt.Println("Inconsistencia encontrada")
			}
		} else {
			// Si la weá no existe la chanta así tal cual. 
			rS, err := c.GetNumberRebelds(ctx, &pb.MensajeToBrokerFromLeia{Comando: comando, NombrePlaneta: arg1, NombreCiudad: arg2, IpServidorFulcrum: "vacia"})
			if err != nil {
				log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
			}
			fmt.Println(rS.GetNumeroRebeldes())
			cantRebeldes := rS.GetNumeroRebeldes()
			reloj := rS.GetVector()
			ip := rS.GetIpServidorFulcrum()

			val.cantRebeldes = cantRebeldes
			fmt.Printf("Cantidad de rebeldes %s\n", cantRebeldes)
			val.reloj = reloj
			val.servidor = ip
		}
		informacion[arg2] = val
	}
}


func ObtenerCantRebeldes(nombrePlaneta string, nombreCiudad string) (int, []int, string) {
	return 3, make([]int, 3), "localhost"
}