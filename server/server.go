package server

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/PyMarcus/serialize/tools"
	"gopkg.in/yaml.v2"
	"github.com/pelletier/go-toml"
)

type Server struct{
	Host string 
	Port string 
}

var titles = [4]string{"name", "cpf", "age", "message"}

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
	var counter int = 0
	
	defer conn.Close()
	
	for{
		buffer := make([]byte, 512)
		
		n, _ := conn.Read(buffer)
		
		if n > 0{
			parts := strings.Split(strings.TrimSpace(string(buffer[:n])), "$")
			contentType := parts[0]
			body := parts[1]
			buffer = make([]byte, len(buffer))
			log.Println("Detected ", contentType)
						
			switch contentType {
			case "application/json":
				s.jsonResponse(conn, []byte(body), &data)
				break
			case "application/xml":
				s.xmlResponse(conn, []byte(body), &data)
				break
			case "application/csv":
				s.csvResponse(conn, []byte(body), &data)
				break
			case "application/yaml":
				s.yamlResponse(conn, []byte(body), &data)
				break
			case "application/toml":
				s.tomlResponse(conn, []byte(body), &data)
				break
			}
			conn.Write([]byte("OK"))
			counter++
			if counter >= 5{
				conn.Close()
				break
			}
							
		}	
	}
}

func (s *Server) jsonResponse(conn net.Conn, buffer []byte, data *tools.Data){
	log.Println("ORIGINAL FORMAT: ", string(buffer))
	
	receivedData := bytes.Trim(buffer, "\x00")
			
	err := json.Unmarshal(receivedData, &data)
	
	if err != nil {
		log.Println("Error while unmarshalling JSON:", err)
		return
	}
	
	s.print(data)
}

func (s *Server) xmlResponse(conn net.Conn, buffer []byte, data *tools.Data){
	log.Println("ORIGINAL FORMAT: ", string(buffer))
	
	receivedData := bytes.Trim(buffer, "\x00")
			
	err := xml.Unmarshal(receivedData, &data)
	
	if err != nil {
		log.Println("Error while unmarshalling XML:", err)
		return
	}
	
	s.print(data)
}

func (s *Server) csvResponse(conn net.Conn, buffer []byte, data *tools.Data){
	log.Println("ORIGINAL FORMAT: ", string(buffer))
		
	csv := strings.Split(string(buffer), ",")[len(titles):]
	
	data.Name = csv[0]
	data.Cpf = csv[1]
	data.Age = csv[2]
	data.Message = csv[3]
	 
	s.print(data)
}

func (s *Server) yamlResponse(conn net.Conn, buffer []byte, data *tools.Data){
	log.Println("ORIGINAL FORMAT: ", string(buffer))
		
	err := yaml.Unmarshal(buffer, &data)
	 
	if err != nil{
		log.Println(err)
	} 
	
	s.print(data)
}

func (s *Server) tomlResponse(conn net.Conn, buffer []byte, data *tools.Data){
	log.Println("ORIGINAL FORMAT: ", string(buffer))
				
	var decodedData tools.Data
	
	err := toml.Unmarshal([]byte(buffer), &decodedData)
	if err != nil {
		log.Println("Error while unmarshalling TOML:", err)
		return
	}

	data.Name = decodedData.Name
	data.Age = decodedData.Age
	data.Cpf = decodedData.Cpf
	data.Message = decodedData.Message

	s.print(data)
}

func (s *Server) print(data *tools.Data){
	log.Println(strings.Repeat("-", 20))
	log.Println("[Parsed data]")
	log.Println("Name: ", data.Name)
	log.Println("Age: ", data.Age)
	log.Println("CPF: ", data.Cpf)
	log.Println("Message: ", data.Message)
	log.Println(strings.Repeat("-", 20))
}
