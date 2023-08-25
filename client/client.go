package client

import (
	"fmt"
	"log"
	"net"
	"encoding/json"
	
	"github.com/PyMarcus/serialize/tools"
)

type Client struct{
	ServerIp   string
	ServerPort string
}

func (c *Client) SendData(name, cpf, age, message *string){

	log.Printf("To send [name: %s, cpf: %s, age: %s, message: %s]\n", *name, *cpf, *age, *message)

	d := &tools.Data{
		Name: *name,
		Cpf: *cpf,
		Age: *age,
		Message: *message,
	}

	bind := fmt.Sprintf("%s:%s", c.ServerIp, c.ServerPort)
	
	conn, err := net.Dial("tcp", bind)
	defer conn.Close()	
	
	tools.ThereIsError(err)
	
	json, err := c.convertToJson(d)
			
	tools.ThereIsError(err)
	if tools.WriteToServer(conn, json){
		log.Println("JSON was been sended!")
	}else{
		log.Println("Fail to send JSON")
	}
}

func (c *Client) convertToJson(d *tools.Data) ([]byte, error){
	jsonData, err := json.Marshal(d)
	
	if err != nil{
		log.Panicf("Fail to convert to json %v", err)
		return nil, err
	}
	
	return jsonData, nil
} 
