package main

import (
	"fmt"
	"os"
	
	"github.com/PyMarcus/serialize/client"
	"github.com/PyMarcus/serialize/server"
)

func main(){

	var name     string 
	var cpf      string 
	var age      string 
	var message  string
	
	args := os.Args[1:] 

	if len(args) == 0 {
		fmt.Println("Run server or client")
		return
	}
	
	c := client.Client{
		ServerIp: "localhost",
		ServerPort: "8000",
	}
	
	
	if args[0] == "client"{
		fmt.Print("Name: ")
		fmt.Scan(&name)

		fmt.Print("CPF: ")
		fmt.Scan(&cpf)

		fmt.Print("Age: ")
		fmt.Scan(&age)

		fmt.Print("Message: ")
		fmt.Scan(&message)
		c.SendData(&name, &cpf, &age, &message)
	}else{
		server.RunServer()
	}
	
}