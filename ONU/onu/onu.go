package main

import (
	"containerized-go-app/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var state string

func main() {

	i := 1
	n := 2
	conn, err := grpc.Dial("10.6.46.139:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	c := proto.NewMyServiceClient(conn)
	// Implementar una consulta por pantalla para determinar si quiere ver a los Infectados o Fallecidos
	for i < n {
		var input string
		log.Println("Consulta de Estados de la ONU:")
		log.Println("1.- Infectados")
		log.Println("2.- Fallecido")
		log.Println("Seleccione una opción (1 o 2): ")
		fmt.Scanln(&input)
		if input == "1" {
			state = "Infectado"
		} else if input == "2" {
			state = "Fallecido"
		}
		// Comunicación con OMS

		r, err := c.SendONUMsg(context.Background(), &proto.ConsultaPoblacion{Estado: state})
		if err != nil {
			log.Fatalf("No se pudo enviar los datos: %v", err)
		}
		log.Printf("Lista de nombres con estado %s", state)
		for _, NombreCompleto := range r.GetNombres() {
			log.Printf("%v %v", NombreCompleto.GetNombre(), NombreCompleto.GetApellido())
		}

		log.Println("¿Quiere continuar? (S/N)")
		fmt.Scanln(&input)
		if input == "N" {
			i = 3
		}
	}
}
