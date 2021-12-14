package main

import (
	"fmt"
	//"math/rand"
	//"strconv"
	//"time"

	"context"
	"log"

	pbBroker "inf343-tarea-3/protoBrokerInformantes"
	pbFulcrum "inf343-tarea-3/protoServidorInformante"
	"google.golang.org/grpc"
)

const (
	addressBroker = "localhost:50051"
)
var addressFulcrum = [...]string{"localhost:50062"}
type data struct {
	comandos []string
	reloj []int32
	servidor string
}

func main() {

	direccionToId := make(map[string]int)
	direccionToId["localhost:50062"] = 0
	direccionToId["localhost:50064"] = 1
	direccionToId["localhost:50066"] = 2
	
	conexiones := make([]pbFulcrum.ConnToServidorFromInformanteClient,3)


	conn, err := grpc.Dial(addressBroker, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pbBroker.NewConnToBrokerFromInformanteClient(conn)
	ctx := context.Background()
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
	}
	
	for i, addr := range addressFulcrum {
		conn2, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("No se pudo conectar: %v", err)
		}
		defer conn2.Close()
		conexiones[i] = pbFulcrum.NewConnToServidorFromInformanteClient(conn2)
	}
	

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
		arg3 = "0"
	} else {
		strcmd = fmt.Sprintf("%s %s %s %s", comando, arg1, arg2, arg3)
	}
	
	val, ok := informacion[arg1];
	var cFulcrum pbFulcrum.ConnToServidorFromInformanteClient
	if ok {
		switch val.servidor {
			case addressFulcrum[0]:
				cFulcrum = conexiones[0]
				break
			/*case addressFulcrum[1]:
				cFulcrum = conexiones[1]
				break
			case addressFulcrum[2]:
				cFulcrum = conexiones[2]
				break*/
		}
		r, err := cFulcrum.Comando(ctx, &pbFulcrum.MensajeToServidor{Comando: comando, NombrePlaneta: arg1, NombreCiudad: arg2, NuevoValor: arg3})
		if err != nil {
			log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
		}
		reloj := r.GetVector()
		ip := r.GetIpServidorFulcrum()
		
		if (reloj[direccionToId[ip]] >= val.reloj[direccionToId[ip]]){
			val.reloj = reloj
		} else {
			fmt.Println("Inconsistencia encontrada")
		}
		val.servidor = ip
		val.comandos = append(val.comandos, strcmd)
		// Si ya existía el registro actualizo reloj y ip (¿Quizás hacer esa weá de
		// agarrar el máximo valor de los componentes del viejo reloj y el nuevo?)
		// y appendeo a los comandos el nuevo comando ejecutado.
	} else {
		// Si la weá no existe la chanta así tal cual.
		fmt.Println("a")
		rS, err := c.ObtenerDireccion(ctx, &pbBroker.MensajeToBrokerFromInformante{IpServidorFulcrum: "vacia"})
		ip := rS.Direccion
		fmt.Println("b")
		if err != nil {
			log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
		}
		switch ip {
		case addressFulcrum[0]:
			cFulcrum = conexiones[0]
			break
		/*case addressFulcrum[1]:
			cFulcrum = conexiones[1]
			break
		case addressFulcrum[2]:
			cFulcrum = conexiones[2]
			break*/
		}
		r, err := cFulcrum.Comando(ctx, &pbFulcrum.MensajeToServidor{Comando: comando, NombrePlaneta: arg1, NombreCiudad: arg2, NuevoValor: arg3})
		if err != nil {
			log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
		}
		reloj := r.GetVector()
		ip = r.GetIpServidorFulcrum()
		val = data{comandos: []string{strcmd}, reloj: reloj, servidor: ip}
	}
	informacion[arg1] = val
}