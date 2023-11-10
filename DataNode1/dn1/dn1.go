package main

import (
	"bufio"
	"containerized-go-app/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedMyServiceServer
}

func (s *server) SendOMSask(ctx context.Context, in *proto.PedirDN) (*proto.NombreCompleto, error) {
	var apellido, nombre string
	var err error
	log.Println(in.GetId())
	apellido, nombre, err = BuscarPorID(int(in.GetId()))

	if err != nil {
		// Manejar el error, como devolver un error gRPC
		return nil, status.Error(codes.NotFound, "ID no encontrado1")
	}
	// Crear y devolver una respuesta gRPC con el nombre y apellido
	response := &proto.NombreCompleto{
		Nombre:   nombre,
		Apellido: apellido,
	}
	return response, nil
}

func (s *server) SendOMSdepositar(ctx context.Context, in *proto.AlmacenarEnDN) (*proto.RespuestaGenerica, error) {
	var err int
	go func() {
		err = EscribirArchivo(int(in.GetId()), in.GetNombre(), in.GetApellido())
	}()
	if err == 0 {
		return &proto.RespuestaGenerica{Mensaje: "Se guardo con exito"}, nil
	} else {
		return &proto.RespuestaGenerica{Mensaje: "Fracaso"}, nil
	}
}

func main() {

	// Creacion del servidor, editar el segundo valor con la ip y puerto a usar
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Funciones Utiles

func EscribirArchivo(id int, nombre string, apellido string) int {
	// Abre el archivo en modo escritura. Si el archivo no existe, lo crea.
	file, err := os.OpenFile("DATA.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return 1 // Devuelve un código de error
	}
	defer file.Close() // Aseguramos que el archivo se cierre al finalizar la función
	// Concatena ID, nombre y apellido en una cadena
	str := fmt.Sprintf("%d %s %s\n", id, nombre, apellido)
	// Escribe la cadena en el archivo
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return 1 // Devuelve un código de error
	}
	fmt.Printf("Se guardó %s %s en DATA.txt\n", nombre, apellido)
	return 0 // Éxito
}

func BuscarPorID(id int) (string, string, error) {
	file, err := os.Open("DATA.txt")
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 3 {
			currentID := parts[0]
			if currentID == fmt.Sprintf("%d", id) {
				return parts[1], parts[2], nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return "", "", err
	}
	return "", "", fmt.Errorf("ID no encontrado")
}
