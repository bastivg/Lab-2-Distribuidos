package main

import (
	"bufio"
	"containerized-go-app/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

// Struct para la lectura de DATA.txt
type Node struct {
	ID       int
	DataNode int
	Estado   string
}

// Variables a utilizar
var id int = 1
var path = "DATA.txt"
var listaNombres = &proto.ListaNombres{}

type server struct {
	proto.UnimplementedMyServiceServer
}

func (s *server) SendONUMsg(ctx context.Context, in *proto.ConsultaPoblacion) (*proto.ListaNombres, error) {

	fmt.Printf("Solicitud de ONU recibida, mensaje enviado: %s\n", in.GetEstado())
	// Consultar los datos en DATA.txt
	nodos, err := leerArchivo(path)
	if err != nil {
		log.Fatalf("Error al leer el archivo: %v", err)
	}

	for _, nodo := range nodos {
		if nodo.Estado == "Infectado" && in.GetEstado() == "Infectado" {
			if nodo.DataNode == 1 {
				// Conectar con DatanNode 1 y preguntar por el Nombre y Apellido
				conn, err := grpc.Dial("10.6.46.113:50051", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("No se pudo conectar: %v", err)
				}
				defer conn.Close()
				c := proto.NewMyServiceClient(conn)
				r, err := c.SendOMSask(context.Background(), &proto.PedirDN{Id: int32(nodo.ID)})
				if err != nil {
					log.Fatalf("No se pudo enviar los datos: %v", err)
				}

				// Almacenar los datos en la lista a enviar
				nombreCompleto := &proto.NombreCompleto{
					Nombre:   r.GetNombre(),
					Apellido: r.GetApellido(),
				}

				listaNombres.Nombres = append(listaNombres.Nombres, nombreCompleto)

			} else if nodo.DataNode == 2 {
				// Conectar con DatanNode 2 y preguntar por el Nombre y Apellido
				conn, err := grpc.Dial("10.6.46.114:50052", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("No se pudo conectar: %v", err)
				}
				defer conn.Close()
				c := proto.NewMyServiceClient(conn)
				r, err := c.SendOMSask(context.Background(), &proto.PedirDN{Id: int32(nodo.ID)})
				if err != nil {
					log.Fatalf("No se pudo enviar los datos: %v", err)
				}

				// Almacenar los datos en la lista a enviar
				nombreCompleto := &proto.NombreCompleto{
					Nombre:   r.GetNombre(),
					Apellido: r.GetApellido(),
				}

				listaNombres.Nombres = append(listaNombres.Nombres, nombreCompleto)

			}
		} else if nodo.Estado == "Fallecido" && in.GetEstado() == "Fallecido" {
			if nodo.DataNode == 1 {
				// Conectar con DatanNode 1 y preguntar por el Nombre y Apellido
				conn, err := grpc.Dial("10.6.46.113:50051", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("No se pudo conectar: %v", err)
				}
				defer conn.Close()
				c := proto.NewMyServiceClient(conn)
				r, err := c.SendOMSask(context.Background(), &proto.PedirDN{Id: int32(nodo.ID)})
				if err != nil {
					log.Fatalf("No se pudo enviar los datos: %v", err)
				}

				// Almacenar los datos en la lista a enviar
				nombreCompleto := &proto.NombreCompleto{
					Nombre:   r.GetNombre(),
					Apellido: r.GetApellido(),
				}

				listaNombres.Nombres = append(listaNombres.Nombres, nombreCompleto)

			} else if nodo.DataNode == 2 {
				// Conectar con DatanNode 2 y preguntar por el Nombre y Apellido
				conn, err := grpc.Dial("10.6.46.114:50052", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("No se pudo conectar: %v", err)
				}
				defer conn.Close()
				c := proto.NewMyServiceClient(conn)
				r, err := c.SendOMSask(context.Background(), &proto.PedirDN{Id: int32(nodo.ID)})
				if err != nil {
					log.Fatalf("No se pudo enviar los datos: %v", err)
				}

				// Almacenar los datos en la lista a enviar
				nombreCompleto := &proto.NombreCompleto{
					Nombre:   r.GetNombre(),
					Apellido: r.GetApellido(),
				}

				listaNombres.Nombres = append(listaNombres.Nombres, nombreCompleto)

			}
		}
	}

	return listaNombres, nil
}

// Recibe la información de los continentes y almacena los datos en los datanodes
func (s *server) SendContinentMsg(ctx context.Context, in *proto.MensajeContinente) (*proto.RespuestaGenerica, error) {
	letra := in.GetApellido()
	fmt.Printf("Solicitud de Continente recibida, mensaje enviado: %s %s %s", in.GetNombre(), in.GetApellido(), in.GetEstado())
	if letra[0] <= 'M' {
		//conectar con datanode 1
		// Guardardo de datos en DATA.txt
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ID DataNode Estado
		if _, err := file.WriteString(fmt.Sprintf("%d %d %s \n", id, 1, in.GetEstado())); err != nil {
			log.Fatalf("error writing to file: %v", err)
		}
		// Conexión DataNode 1
		conn, err := grpc.Dial("10.6.46.113:50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("No se pudo conectar: %v", err)
		}
		defer conn.Close()
		c := proto.NewMyServiceClient(conn)
		// Envio de Datos a DataNode 1
		r, err := c.SendOMSdepositar(context.Background(), &proto.AlmacenarEnDN{Nombre: in.GetNombre(), Apellido: in.GetApellido(), Id: int32(id)})
		if err != nil {
			log.Fatalf("No se pudo enviar los datos: %v", err)
		}
		log.Printf("Se enviaron correctamente los datos en %s", r.GetMensaje())
		id++

	} else if letra[0] <= 'Z' {
		//DN2
		// Guardado de datos en DATA.txt
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ID DataNode Estado
		if _, err := file.WriteString(fmt.Sprintf("%d %d %s \n", id, 2, in.GetEstado())); err != nil {
			log.Fatalf("error writing to file: %v", err)
		}
		// Conexión DataNode 2
		conn, err := grpc.Dial("10.6.46.114:50052", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("No se pudo conectar: %v", err)
		}
		defer conn.Close()
		c := proto.NewMyServiceClient(conn)
		// Envio de Datos a DataNode 2
		r, err := c.SendOMSdepositar(context.Background(), &proto.AlmacenarEnDN{Nombre: in.GetNombre(), Apellido: in.GetApellido(), Id: int32(id)})
		if err != nil {
			log.Fatalf("No se pudo enviar los datos: %v", err)
		}
		log.Printf("Se enviaron correctamente los datos en %s", r.GetMensaje())
		id++
	}

	return &proto.RespuestaGenerica{Mensaje: "Ok"}, nil
}

func main() {
	// Creación de archiivos DATA.txt
	crearArchivo()

	lis, err := net.Listen("tcp", "0.0.0.0:50053")
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
func crearArchivo() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("Archivo creado exitosamente", path)
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func leerArchivo(nombreArchivo string) ([]Node, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var nodos []Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, " ")

		id, err := strconv.Atoi(partes[0])
		if err != nil {
			log.Fatalf("Error al convertir ID a entero: %v", err)
		}

		dataNode, err := strconv.Atoi(partes[1])
		if err != nil {
			log.Fatalf("Error al convertir DataNode a entero: %v", err)
		}

		nodos = append(nodos, Node{
			ID:       id,
			DataNode: dataNode,
			Estado:   partes[2],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nodos, nil
}
