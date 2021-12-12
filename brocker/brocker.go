package main

import (
	"context"
	"log"
	"math/rand"
	"math"
	"net"
	"time"
	"sync"
	"strconv"


	pbInformante "inf343-tarea-3/protoBrokerInformantes"
	pbLeia "inf343-tarea-3/protoBrokerLeia"
	//pbFulcrum "inf343-tarea-3/protoServidorBroker"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/peer"
)

var adress = [...]string{"localhost:50061", "localhost:50062", "localhost:50063"}
//CAMBIAR A LAS DIRECCIONES DE LOS DISTINTOS SERVIDORES FULCRUM ej: dist14:puerto

const (
	portInformante = ":50051"
	portLeia = "50052"
)

type serverInformante struct {
	pbInformante.UnimplementedConnToBrokerServer
}
type serverLeia struct {
	pbLeia.UnimplementedConnToBrokerServer
}

func conexionLeia() {

	lis, err := net.Listen("tcp", portLeia)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := serverLeia{}
	grcpServer := grpc.NewServer()

	pbLeia.RegisterConnToBrokerServer(grcpServer, &s)
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

	pbInformante.RegisterConnToBrokerServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
