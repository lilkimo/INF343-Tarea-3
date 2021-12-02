package main

import (
	"bufio"
	//"context"
	"fmt"
	"io/ioutil"

	"log"
	//"net"
	"os"
	"strings"
	//pbSvBk "INF343-Tarea-3\protoServidorBroker"
	//pbSvInf "INF343-Tarea-3\protoServidorInformante"
	//"google.golang.org/grpc"
)

const (
	port = ":50053"
)

var reg_planetas []string
var log_planetas []string

var comandos []string = []string{"AddCity", "UpdateNumber", "UpdateName", "DeleteCity"}

/*
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
*/

func valueInSlice(value string, list []string) bool {
	for _, b := range list {
		if b == value {
			return true
		}
	}
	return false
}

func command(comando string, planeta string, ciudad string, valor int32) {
	planetFile := fmt.Sprintf("servidores/%s.txt", planeta)
	planetLog := fmt.Sprintf("servidores/%s.log", planeta)
	var succ bool = false

	if valueInSlice(comando, comandos) {
		if comando == "AddCity" {
			if !valueInSlice(planeta, reg_planetas) {
				f, err := os.Create(planetFile)
				check(err)
				reg_planetas = append(reg_planetas, planeta)
				f.Close()
			}
			if city_exists(planeta, ciudad) {
				log.Printf("ciudad ya existe, no se puede agregar")
				return
			}
			if valor < 0 {
				valor = 0
			}
			AddCity(planeta, ciudad, valor)
			succ = true

		} else if comando == "UpdateName" {
			aux := strings.Split(ciudad, ":")
			if city_exists(planeta, aux[0]) {
				UpdateName(planeta, aux[0], aux[1])
				succ = true
			}

		} else {
			if city_exists(planeta, ciudad) {
				if comando == "UpdateNumber" {
					UpdateNumber(planeta, ciudad, valor)
					succ = true
				} else if comando == "DeleteCity" {
					DeleteCity(planeta, ciudad)
					succ = true
				}
			}
		}
	} else {
		log.Printf("comando invalido")
	}
	if succ == true {
		if !valueInSlice(planeta, log_planetas) {
			f, err := os.Create(planetLog)
			check(err)
			log_planetas = append(log_planetas, planeta)
			f.Close()
		}
		logPlanetario(comando, planeta, ciudad, valor)
	}
}

func get_city_data(planeta string, ciudad string) string {
	filename := fmt.Sprintf("servidores/%s.txt", planeta)
	var curr, curr_planet, curr_city string
	var num int32
	f, err := os.ReadFile(filename)
	check(err)

	scanner := bufio.NewScanner(strings.NewReader(string(f)))

	for scanner.Scan() {
		curr = scanner.Text()
		fmt.Sscanf(curr, "%s %s %d", &curr_planet, &curr_city, &num)
		if curr_planet == planeta && curr_city == ciudad {
			return fmt.Sprintf("%s %s %d\n", planeta, ciudad, num)
		}
	}
	return ""
}

func city_exists(planeta string, ciudad string) bool {
	filename := fmt.Sprintf("servidores/%s.txt", planeta)
	f, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	if !strings.Contains(string(f), ciudad) {
		log.Printf("Ciudad %s no existe.", ciudad)
		return false
	} else {
		return true
	}
}

func DeleteCity(planeta string, ciudad string) {

	filename := fmt.Sprintf("servidores/%s.txt", planeta)
	f, err := os.ReadFile(filename)
	check(err)

	city := get_city_data(planeta, ciudad)
	log.Printf(city)

	new_file := strings.Replace(string(f), city, "", 1)
	err = ioutil.WriteFile(filename, []byte(new_file), 0666)
	check(err)
}

func AddCity(planeta string, ciudad string, valor int32) {
	filename := fmt.Sprintf("servidores/%s.txt", planeta)

	f, err := os.OpenFile(filename, os.O_APPEND, 0600)
	check(err)
	str := fmt.Sprintf("%s %s %d\n", planeta, ciudad, valor)
	f.WriteString(str)
	f.Close()
	log.Printf(str)

}

func UpdateNumber(planeta string, ciudad string, nuevo_valor int32) {
	filename := fmt.Sprintf("servidores/%s.txt", planeta)
	ft, err := os.ReadFile(filename)
	check(err)

	city := get_city_data(planeta, ciudad)

	new := fmt.Sprintf("%s %s %d\n", planeta, ciudad, nuevo_valor)
	new_file := strings.Replace(string(ft), city, new, 1)

	err = ioutil.WriteFile(filename, []byte(new_file), 0666)
	check(err)
}

func UpdateName(planeta string, ciudad string, nuevo_valor string) {
	filename := fmt.Sprintf("servidores/%s.txt", planeta)
	f, err := os.ReadFile(filename)
	check(err)

	city := get_city_data(planeta, ciudad)
	new := strings.Replace(city, ciudad, nuevo_valor, 1)

	new_file := strings.Replace(string(f), city, new, 1)

	err = ioutil.WriteFile(filename, []byte(new_file), 0666)
	check(err)
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

	f, err := os.OpenFile(filename, os.O_APPEND, 0600)
	check(err)
	f.WriteString(str)
	f.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	command("AddCity", "coruscant", "temploJedi", 12)
	command("AddCity", "coruscant", "senado", 23)
	command("UpdateName", "coruscant", "senado:sewers", 0)
	command("DeleteCity", "coruscant", "temploJedi", 0)

	/*
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
	*/
}
