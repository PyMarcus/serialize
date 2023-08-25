package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/PyMarcus/serialize/tools"
)

type Server struct{
	Host string 
	Port string 
}

func (s *Server) Start(){
	bind := fmt.Sprintf("%s:%s", s.Host, s.Port)
	
	listener, err := net.Listen("tcp", bind)
	
	if err != nil{
		log.Fatal(err)
	}
	
	log.Println("Running on ", s.Host, ":" ,s.Port)
	
	for{
		conn, err := listener.Accept()
		if(tools.ThereIsError(err)){
			log.Println("Received connection from ", conn.RemoteAddr())
			
			go s.handleConnection(conn)	
		}
	}
}

func (s *Server) handleConnection(conn net.Conn){
	var data tools.Data

	defer conn.Close()
	
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	
	if err != nil{
		log.Println(err)
	}
	
	log.Println("ORIGINAL FORMAT: ", string(buffer))
	
	err = json.Unmarshal(buffer, &data)
	
	log.Println(strings.Repeat("-", 20))
	log.Println("[Parsed data]")
	log.Println("Name: ", data.Name)
	log.Println("Age: ", data.Age)
	log.Println("CPF: ", data.Cpf)
	log.Println("Message: ", data.Message)
	log.Println(strings.Repeat("-", 20))
	
	
}
