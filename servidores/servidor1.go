package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	//pbSvBk "INF343-Tarea-3\protoServidorBroker"
	//pbSvInf "INF343-Tarea-3\protoServidorInformante"

	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

var reg_planetas []string
var log_planetas []string

type DataNodeServer struct {
	pb.UnimplementedNameDataServiceServer
}

func (s *DataNodeServer) RegistrarJugadas(ctx context.Context, in *pb.JugadaToData) (*pb.RespuestaJugada, error) {
	log.Printf("Input - pl: %d | cd: %d | va: \n", in.IdJugador, in.Etapa)
	var jgs []int32
	guardarJugada(in.IdJugador, in.Jugada, in.Etapa)
	jgs = obtenerJugada(in.IdJugador, in.Etapa)
	return &pb.RespuestaJugada{Jugadas: jgs, Cantidad: int32(len(jgs))}, nil
}

func valueInSlice(value string, list []string) bool {
	for _, b := range list {
		if b == value {
			return true
		}
	}
	return false
}

func registroPlanetario(planeta string, ciudad string, valor int32) {
	filename := fmt.Sprintf("servidores/%s.txt", planeta)
	str := fmt.Sprintf("%s %s %d\n", planeta, ciudad, valor)

	if valueInSlice(planeta, reg_planetas) {
		//do nothing
	} else {
		_ = append(reg_planetas, planeta)
		f, err := os.Create(filename)
		check(err)
		f.Close()
	}
	f, err := os.OpenFile(filename, os.O_APPEND, 0600)
	check(err)
	f.WriteString(str)
	f.Close()
}

func logPlanetario(comando string, planeta string, ciudad string, valor int32) {
	filename := fmt.Sprintf("servidores/%s.log", planeta)
	var str string

	if comando == "UpdateNumber" {
		str = fmt.Sprintf("%s %s %s %d\n", comando, planeta, ciudad, valor)

	} else if comando == "AddCity" {
		if valor >= 0 {
			str = fmt.Sprintf("%s %s %s %d\n", comando, planeta, ciudad, valor)
		} else {
			str = fmt.Sprintf("%s %s %s\n", comando, planeta, ciudad)
		}

	} else if comando == "UpdateName" {
		split := strings.Split(ciudad, ":")
		str = fmt.Sprintf("%s %s %s\n", comando, split[0], split[1])

	} else if comando == "DeleteCity" {
		str = fmt.Sprintf("%s %s %s\n", comando, planeta, ciudad)
	}

	if valueInSlice(planeta, log_planetas) {
		//do nothing
	} else {
		_ = append(log_planetas, planeta)
		f, err := os.Create(filename)
		check(err)
		f.Close()
	}
	f, err := os.OpenFile(filename, os.O_APPEND, 0600)
	check(err)
	f.WriteString(str)
	f.Close()
}

func obtenerJugada(idJugador int32, etapa int32) []int32 {
	filename := fmt.Sprintf("dataNode/jugador_%d__etapa_%d.txt", idJugador, etapa)

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	var num int32
	var str string
	var pl []int32
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		str = scanner.Text()
		fmt.Sscanf(str, "%d", &num)
		pl = append(pl, num)

	}
	return pl
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fatal Error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNameDataServiceServer(s, &DataNodeServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal Error: %v", err)
	}
}
