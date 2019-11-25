package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	rad "github.com/arma29/mid-rasp/radiation"

	"github.com/streadway/amqp"

	"github.com/arma29/mid-rasp/shared"
)

func main() {

	// Get Argument from command Line
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}
	ipContainer := os.Args[1]

	conn, err := amqp.Dial("amqp://guest:guest@" +
		ipContainer + ":" +
		strconv.Itoa(shared.RABBITMQ_PORT) + "/")
	shared.CheckError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	shared.CheckError(err)
	defer ch.Close()

	// declaração de filas , cria se não existir
	requestQueue, err := ch.QueueDeclare( // mesma fila de envio
		"request", false, false, false, false, nil)
	shared.CheckError(err)

	replyQueue, err := ch.QueueDeclare( // mesma fila de respostas
		"response", false, false, false, false, nil)
	shared.CheckError(err)

	// prepara o recebimento de mensagens do cliente
	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil)
	shared.CheckError(err)

	fmt.Println("Servidor pronto...")

	forever := make(chan bool) // travar
	fmt.Println("Time")
	go func() {
		for d := range msgsFromClient {

			// recebe request
			msgRequest := rad.Radiation{}
			err := json.Unmarshal(d.Body, &msgRequest)
			fmt.Printf("Estrutura Recebida: ")
			fmt.Println(msgRequest)
			shared.CheckError(err)

			// Medindo o tempo
			t1 := time.Now().UnixNano()
			t2 := msgRequest.Timestamp
			s := fmt.Sprintf("%d", t1-t2)
			fmt.Println(s)

			// processa request
			r := rad.IsRadiationDangerous(msgRequest.Value)
			if r {
				// prepara resposta
				replyMsg := rad.Validator{IsDangerous: r}
				replyMsgBytes, err := json.Marshal(replyMsg)
				shared.CheckError(err)

				// publica resposta
				err = ch.Publish("", replyQueue.Name, false, false,
					amqp.Publishing{ContentType: "text/plain", Body: replyMsgBytes})
				shared.CheckError(err)
			}
		}
	}()

	<-forever

}
