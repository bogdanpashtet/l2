package main

import (
	"log"
	"net"
)
import "fmt"
import "bufio"
import "strings" // требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("Error in listening port")
	}

	// Открываем порт
	conn, _ := ln.Accept()
	defer conn.Close()

	// Запускаем цикл
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Println()
		fmt.Print("Message Received: ", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
