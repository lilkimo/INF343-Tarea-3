package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	"log"
	"net"
	"os"
	"strings"
	pbBroker "inf343-tarea-3/protoServidorBroker"
	pbInformante "inf343-tarea-3/protoServidorInformante"
	"google.golang.org/grpc"
)

type serverInformante struct {
	pbInformante.UnimplementedConnToServidorFromInformanteServer
}

type serverBroker struct {
	pbBroker.UnimplementedConnToServidorFromBrokerServer
}

const (
	port = ":50063"
	portInformante = ":50064"
)


var reg_planetas []string
var log_planetas []string

var comandos []string = []string{"AddCity", "UpdateNumber", "UpdateName", "DeleteCity"}

type vectorPlaneta struct {
	vector [] int32
	planeta string
}

var vectores [] vectorPlaneta

func (s *serverBroker) LeiaGetNumberRebelds(ctx context.Context, in *pbBroker.MensajeLeia) (*pbBroker.ServidorRespuestaLeia, error) {
	v := []int32{0,0,0}
	regciudad := get_city_data(in.NombrePlaneta, in.NombreCiudad)
	str := strings.SplitAfter(regciudad, " ")
	n , _ := strconv.Atoi(str[2])
	for _,vector := range vectores {
		if vector.planeta == in.NombrePlaneta {
			v = vector.vector
			break
		}
	}
	return &pbBroker.ServidorRespuestaLeia{NumeroRebeldes: int32(n), Vector: v, IpServidorFulcrum: "localhost"+port}, nil
}

func (s *serverInformante) Comando (ctx context.Context, in *pbInformante.MensajeToServidor) (*pbInformante.Respuesta, error) {
	vct := []int32{0}
	command(in.Comando, in.NombrePlaneta, in.NombreCiudad, in.NuevoValor)
	for _, vector := range vectores {
		if vector.planeta == in.NombrePlaneta {
			vct = vector.vector
		}
	}

	return &pbInformante.Respuesta{Vector: vct, IpServidorFulcrum: "localhost"+portInformante}, nil
}

func valueInSlice(value string, list []string) bool {
	for _, b := range list {
		if b == value {
			return true
		}
	}
	arr := []int32{0,0,0}
	vectores = append(vectores, vectorPlaneta{vector: arr, planeta: value})
	return false
}

func command(comando string, planeta string, ciudad string, valor string) {
	planetFile := fmt.Sprintf("servidores/%s.txt", planeta)
	planetLog := fmt.Sprintf("servidores/%s.log", planeta)
	var succ bool = false

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
		for _, vector := range vectores{
			if vector.planeta == planeta {
				vector.vector[0] += 1
				break
			}
		}
		valor2,_ := strconv.Atoi(valor)
		AddCity(planeta, ciudad, int32(valor2))
		succ = true

	} else if comando == "UpdateName" {
		if city_exists(planeta, ciudad) {
			UpdateName(planeta, ciudad, valor)
			succ = true
			for _, vector := range vectores{
				if vector.planeta == planeta {
					vector.vector[0] += 1
				}
			}
		}

	} else {
		if city_exists(planeta, ciudad) {
			if comando == "UpdateNumber" {
				valor2, _ := strconv.Atoi(valor)
				UpdateNumber(planeta, ciudad, int32(valor2))

				succ = true
				for _, vector := range vectores{
					if vector.planeta == planeta {
						vector.vector[0] += 1
					}
				}
			} else if comando == "DeleteCity" {
				DeleteCity(planeta, ciudad)
				succ = true
				for _, vector := range vectores{
					if vector.planeta == planeta {
						vector.vector[0] += 1
					}
				}
			}
		}
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

func logPlanetario(comando string, planeta string, ciudad string, valor string) {
	filename := fmt.Sprintf("servidores/%s.log", planeta)
	var str string

	if comando == "UpdateNumber" {
		valor2,_ := strconv.Atoi(valor)
		str = fmt.Sprintf("%s %s %s %d\n", comando, planeta, ciudad, int32(valor2))

	} else if comando == "AddCity" {
		valor2,_ := strconv.Atoi(valor)
		if valor2 >= 0 {
			str = fmt.Sprintf("%s %s %s %d\n", comando, planeta, ciudad, int32(valor2))
		} else {
			str = fmt.Sprintf("%s %s %s\n", comando, planeta, ciudad)
		}

	} else if comando == "UpdateName" {
		str = fmt.Sprintf("%s %s %s\n", comando, ciudad, valor)

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

func conexionInformante() {

	lis, err := net.Listen("tcp", portInformante)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := serverInformante{}
	grcpServer := grpc.NewServer()

	pbInformante.RegisterConnToServidorFromInformanteServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	/*command("AddCity", "coruscant", "temploJedi", 12)
	command("AddCity", "coruscant", "senado", 23)
	command("UpdateName", "coruscant", "senado:sewers", 0)
	command("DeleteCity", "coruscant", "temploJedi", 0)*/

	go conexionInformante()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := serverBroker{}
	grcpServer := grpc.NewServer()

	pbBroker.RegisterConnToServidorFromBrokerServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}