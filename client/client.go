package client

import (
	"encoding/json"
	"encoding/xml"
	"bytes"
	"fmt"
	"log"
	"net"


	"github.com/PyMarcus/serialize/tools"
	"gopkg.in/yaml.v2"
	"github.com/pelletier/go-toml"

)

// Client sends data: json, xml, csv, yaml, toml
type Client struct {
	ServerIp   string
	ServerPort string
}

func (c *Client) SendData(name, cpf, age, message *string) {

	log.Printf("To send [name: %s, cpf: %s, age: %s, message: %s]\n", *name, *cpf, *age, *message)

	headers := make(map[string]string, 0)

	d := &tools.Data{
		Name:    *name,
		Cpf:     *cpf,
		Age:     *age,
		Message: *message,
	}

	bind := fmt.Sprintf("%s:%s", c.ServerIp, c.ServerPort)

	conn, err := net.Dial("tcp", bind)
	defer conn.Close()

	tools.ThereIsError(err)
	
	c.sendJson(headers, conn, d)
	
	buffer := make([]byte, 512)
	n, _ := conn.Read(buffer)
	if string(buffer[:n]) == "OK"{
		c.sendXml(headers, conn, d)
	}
	
	buffer = make([]byte, 512)
	n, _ = conn.Read(buffer)
	if string(buffer[:n]) == "OK"{
		c.sendCsv(headers, conn, d)
	}
	
	buffer = make([]byte, 512)
	n, _ = conn.Read(buffer)
	if string(buffer[:n]) == "OK"{
		c.sendYaml(headers, conn, d)
	}
	
	buffer = make([]byte, 512)
	n, _ = conn.Read(buffer)
	if string(buffer[:n]) == "OK"{
		c.sendToml(headers, conn, d)
	}
}

func (c *Client) sendJson(headers map[string]string, conn net.Conn, d *tools.Data) {
    jsonRequest, err := c.convertToJson(d)
    headers["Content-Type"] = "application/json"

    if err != nil {
        log.Panicln(err)
    }

    tools.WriteToServer(conn,
        []byte(fmt.Sprintf("%s$%s\n", headers["Content-Type"],
            jsonRequest)))
}

func (c *Client) sendXml(headers map[string]string, conn net.Conn, d *tools.Data) {
    xmlRequest, err := c.convertToXml(d)
    headers["Content-Type"] = "application/xml"

    if err != nil {
        log.Panicln(err)
    }

    tools.WriteToServer(conn,
        []byte(fmt.Sprintf("%s$%s\n", headers["Content-Type"],
		bytes.Trim(xmlRequest, "\x00") )))
}

func (c *Client) sendCsv(headers map[string]string, conn net.Conn, d *tools.Data) {
    csvRequest, err := c.convertToCsv(d)
    headers["Content-Type"] = "application/csv"

    if err != nil {
        log.Panicln(err)
    }

    tools.WriteToServer(conn,
        []byte(fmt.Sprintf("%s$%s\n", headers["Content-Type"],
            csvRequest)))
}

func (c *Client) sendYaml(headers map[string]string, conn net.Conn, d *tools.Data) {
    yamlRequest, err := c.convertToYaml(d)
    headers["Content-Type"] = "application/yaml"

    if err != nil {
        log.Panicln(err)
    }

    tools.WriteToServer(conn,
        []byte(fmt.Sprintf("%s$%s\n", headers["Content-Type"],
		yamlRequest)))
}

func (c *Client) sendToml(headers map[string]string, conn net.Conn, d *tools.Data) {
    tomlRequest, err := c.convertToToml(d)
    headers["Content-Type"] = "application/toml"

    if err != nil {
        log.Panicln(err)
    }

    tools.WriteToServer(conn,
        []byte(fmt.Sprintf("%s$%s\n", headers["Content-Type"],
		tomlRequest)))
}

func (c *Client) convertToJson(d *tools.Data) ([]byte, error) {
	jsonData, err := json.Marshal(d)

	if err != nil {
		log.Panicf("Fail to convert to json %v", err)
		return nil, err
	}

	return jsonData, nil
}

func (c *Client) convertToXml(d *tools.Data) ([]byte, error){
	xml, err := xml.MarshalIndent(d, "", "  ")
	
	if err != nil{
		fmt.Println("Error:", err)
		return nil, err 
	}
	return xml, nil
}

func (c *Client) convertToCsv(d *tools.Data) ([]byte, error){
	return []byte(fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s,%s,%s", "name", "cpf", "age", "message",
		 d.Name, d.Cpf, d.Age, d.Message,
	)), nil
}

func (c *Client) convertToYaml(d *tools.Data) ([]byte, error){
	yaml, err := yaml.Marshal(d)
	
	if err != nil{
		fmt.Println("Error:", err)
		return nil, err 
	}
	return yaml, nil
}

func (c *Client) convertToToml(d *tools.Data) ([]byte, error){
	toml, err := toml.Marshal(d)
	
	if err != nil{
		fmt.Println("Error:", err)
		return nil, err 
	}
	return toml, nil
}
