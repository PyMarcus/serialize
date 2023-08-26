package tools

import(
	"log"
	"net"
)


type Data struct{
	Name     string  `json:"name" xml:"name" yaml:"name" toml:"name"`   
	Cpf      string  `json:"cpf" xml:"cpf" yaml:"cpf" toml:"cpf"` 
	Age      string  `json:"age" xml:"age" yaml:"age" toml:"age"`
	Message  string  `json:"message" xml:"message" yaml:"message" toml:"message"`
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