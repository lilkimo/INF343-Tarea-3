package main

import (
	"context"
	"log"
	"math/rand"
	//"math"
	"net"
	//"time"
	//"sync"
	//"strconv"


	pbInformante "inf343-tarea-3/protoBrokerInformantes"
	pbLeia "inf343-tarea-3/protoBrokerLeia"
	pbFulcrum "inf343-tarea-3/protoServidorBroker"
	//pbFulcrum "inf343-tarea-3/protoServidorBroker"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/peer"
)

var address = [...]string{"localhost:50061", "localhost:50062", "localhost:50063"}
//CAMBIAR A LAS DIRECCIONES DE LOS DISTINTOS SERVIDORES FULCRUM ej: dist14:puerto

const (
	portInformante = ":50051"
	portLeia = "50052"
)

type serverInformante struct {
	pbInformante.UnimplementedConnToBrokerFromInformanteServer
}
type serverLeia struct {
	pbLeia.UnimplementedConnToBrokerFromLeiaServer
}

func (s *serverLeia) GetNumberRebelds (ctx context.Context, in *pbLeia.MensajeToBroker) (*pbLeia.Respuesta, error) {

	var addressFulcrum string
	if (in.IpServidorFulcrum == "vacia") {
		addressFulcrum = address[rand.Intn(3)]
		conn, err := grpc.Dial(addressFulcrum, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("No se pudo conectar: %v", err)
		}
		defer conn.Close()
		c := pbFulcrum.NewConnToServidorFromBrokerClient(conn)
		r, err := c.LeiaGetNumberRebelds(context.Background(), &pbFulcrum.MensajeLeia{Comando: in.Comando, NombrePlaneta: in.NombrePlaneta, NombreCiudad: in.NombreCiudad})
	
		return &pbLeia.Respuesta{NumeroRebeldes: r.NumeroRebeldes, Vector: r.Vector, IpServidorFulcrum: r.IpServidorFulcrum}, nil

	} else {
		conn, err := grpc.Dial(in.IpServidorFulcrum, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("No se pudo conectar: %v", err)
		}
		defer conn.Close()
		c := pbFulcrum.NewConnToServidorFromBrokerClient(conn)
		r, err := c.LeiaGetNumberRebelds(context.Background(), &pbFulcrum.MensajeLeia{Comando: in.Comando, NombrePlaneta: in.NombrePlaneta, NombreCiudad: in.NombreCiudad})

		return &pbLeia.Respuesta{NumeroRebeldes: r.NumeroRebeldes, Vector: r.Vector, IpServidorFulcrum: r.IpServidorFulcrum}, nil
	}

}

func conexionLeia() {

	lis, err := net.Listen("tcp", portLeia)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := serverLeia{}
	grcpServer := grpc.NewServer()

	pbLeia.RegisterConnToBrokerFromLeiaServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	go conexionLeia()

	lis, err := net.Listen("tcp", portInformante)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := serverInformante{}
	grcpServer := grpc.NewServer()

	pbInformante.RegisterConnToBrokerFromInformanteServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
