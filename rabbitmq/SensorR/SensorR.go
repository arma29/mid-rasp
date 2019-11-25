package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"

	rad "github.com/arma29/mid-rasp/radiation"
	"github.com/arma29/mid-rasp/shared"
)

func main() {
	// Get Argument from command Line
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	ipContainer := os.Args[1]

	// conecta ao servidor de mensageria
	conn, err := amqp.Dial("amqp://guest:guest@" +
		ipContainer + ":" +
		strconv.Itoa(shared.RABBITMQ_PORT) + "/")
	shared.CheckError(err)
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.CheckError(err)
	defer ch.Close()

	// declara  filas, cria se não existir
	requestQueue, err := ch.QueueDeclare( // fila de envio
		"request", false, false, false, false, nil)
	shared.CheckError(err)

	replyQueue, err := ch.QueueDeclare( // fila de respostas
		"response", false, false, false, false, nil)
	shared.CheckError(err)

	// cria consumidor <-> fila de respostas -> async
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
		false, false, nil)
	shared.CheckError(err)

	// enviar
	for {
		// prepara request
		msgRequest := rad.Radiation{Value: rad.GenerateRadiationValue(), Timestamp: time.Now().UnixNano()}
		msgRequestBytes, err := json.Marshal(msgRequest)
		fmt.Printf("Estrutura Enviada: ")
		fmt.Println(msgRequest)
		shared.CheckError(err)

		// publica request <-> fila de envio
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		shared.CheckError(err)

		// Aguarda resposta
		listenChannel(msgsFromServer)

		// Garantir taxa máxima
		time.Sleep(shared.REAL_TIME)
	}
}

// Escutar o canal e dar um timeout a ele
func listenChannel(ch <-chan amqp.Delivery) {
	select {
	case res := <-ch:
		fmt.Println("Deu tempo: ")
		msgReply := rad.Validator{}
		err := json.Unmarshal([]byte(res.Body), &msgReply)
		shared.CheckError(err)

		//Acender o led de alguma forma
	case <-time.After(shared.WAIT_TIME):
		fmt.Println("Não deu tempo")
	}
}
