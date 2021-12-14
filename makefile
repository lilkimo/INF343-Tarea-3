sv1: 
	go run servidores/servidor1.go

sv2:
        go run servidores/servidor2.go

sv3:
        go run servidores/servidor3.go

almirante: 
	go run informantes/almirante.go

broker:
        go run brocker/brocker.go

leia: 
        go run informantes/leia.go

tano:
        go run informantes/tano.go