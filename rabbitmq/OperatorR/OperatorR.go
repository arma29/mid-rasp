package main

import (
	"github.com/streadway/amqp"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/arma29/mid-mid-perf/shared"
	"github.com/arma29/mid-mid-perf/application"
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
		strconv.Itoa(shared.RABBITMQ_PORT) +"/")
	shared.CheckError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	shared.CheckError(err)
	defer ch.Close()

	// declaração de filas , cria se não existir
	requestQueue, err := ch.QueueDeclare( // mesma fila de envio
		"request", false, false, false, false, nil, )
	shared.CheckError(err)

	replyQueue, err := ch.QueueDeclare( // mesma fila de respostas
		"response", false, false, false, false, nil, )
	shared.CheckError(err)

	// prepara o recebimento de mensagens do cliente
	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil, )
	shared.CheckError(err)

	fmt.Println("Servidor pronto...")

	forever := make(chan bool) // travar 
	go func(){
		for d := range msgsFromClient {

			// recebe request
			msgRequest := shared.Request{}
			err := json.Unmarshal(d.Body, &msgRequest)
			shared.CheckError(err)
	
			// processa request
			r := application.CalcFibonacci(msgRequest.Req)

			// prepara resposta
			replyMsg := shared.Response{Res: r}
			replyMsgBytes, err := json.Marshal(replyMsg)
			shared.CheckError(err)
	
			// publica resposta
			err = ch.Publish("", replyQueue.Name, false, false,
				amqp.Publishing{ContentType: "text/plain", Body: replyMsgBytes,})
			shared.CheckError(err)
		}
	}()

	<- forever
	
}