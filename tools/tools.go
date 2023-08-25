package tools

import(
	"log"
	"net"
)


type Data struct{
	Name     string  `json:"name"`   
	Cpf      string  `json:"cpf"` 
	Age      string  `json:"age"`
	Message  string  `json:"message"`
}

func ThereIsError(err any) bool{
	if err != nil{
		log.Panic(err)
		return false
	}
	
	return true
}

func WriteToServer(conn net.Conn, content []byte) bool {
	_, err := conn.Write(content)
	return ThereIsError(err)
}