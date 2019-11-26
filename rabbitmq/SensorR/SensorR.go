package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"

	rad "github.com/arma29/mid-rasp/radiation"
	queue "github.com/arma29/mid-rasp/sensorqueue"
	"github.com/arma29/mid-rasp/shared"
)

var (
	rabbitConn    *amqp.Connection
	rabbitChannel *amqp.Channel

	requestQueue amqp.Queue
	replyQueue   amqp.Queue

	msgsFromServer <-chan amqp.Delivery

	amqpURI string
	err     error
)

func main() {
	// Get Argument from command Line
	if len(os.Args) != 4 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	user := os.Args[1]
	password := os.Args[2]
	ipContainer := os.Args[3]

	amqpURI = "amqp://" + user + ":" + password + "@" +
		ipContainer + ":" +
		strconv.Itoa(shared.RABBITMQ_PORT) + "/"

	// Prepara as estruturas do middleware
	initRabbit()
	// Fechamento de canais deve ser no escopo main
	defer rabbitConn.Close()
	defer rabbitChannel.Close()

	// loop de escutar
	go waitMsgs(msgsFromServer)

	// inicia fila de armazenamento do sensor
	queue.InitQueue()
	// loop de enviar
	for i := 0; i < 100; i++ {

		queue.Enqueue(rad.Radiation{Value: rad.GenerateRadiationValue(),
			Timestamp: time.Now().UnixNano()})

		if !rabbitConn.IsClosed() {
			parseQueue()
		} else {
			fmt.Println("Sem conexão")
			// tenta conexão mais uma vez
			if hasReconnected() {

				initRabbit()
				// Fechar canais deve ser no escopo main
				defer rabbitConn.Close()
				defer rabbitChannel.Close()

				go waitMsgs(msgsFromServer)
			}
		}
		// Garantir taxa máxima
		time.Sleep(shared.REAL_TIME)
	}

}

func initRabbit() {
	rabbitConn = initConn()
	rabbitChannel = initCh()
	initUtils()
}

func initConn() *amqp.Connection {
	// cria conexão
	conn, err := amqp.Dial(amqpURI)
	shared.CheckError(err)
	return conn
}

func initCh() *amqp.Channel {
	// cria o canal
	ch, err := rabbitConn.Channel()
	shared.CheckError(err)
	return ch
}

func initUtils() {
	// declara  filas, cria se não existir
	requestQueue, err = rabbitChannel.QueueDeclare( // fila de envio
		"request", false, false, false, false, nil)
	shared.CheckError(err)

	replyQueue, err = rabbitChannel.QueueDeclare( // fila de respostas
		"response", false, false, false, false, nil)
	shared.CheckError(err)

	// cria consumidor <-> fila de respostas -> async
	msgsFromServer, err = rabbitChannel.Consume(replyQueue.Name, "", true, false,
		false, false, nil)
	shared.CheckError(err)
}

func waitMsgs(ch <-chan amqp.Delivery) {
	for res := range ch {
		msgReply := rad.Validator{}
		err := json.Unmarshal([]byte(res.Body), &msgReply)
		shared.CheckError(err)
		fmt.Printf("Falha registrada em: ")
		fmt.Println(time.Unix(0, msgReply.Timestamp))

		// Acender o Led GPIO lib
	}
}

func parseQueue() {
	for !queue.Empty() {

		// prepara request
		msgRequest := queue.Peek()
		msgRequestBytes, err := json.Marshal(msgRequest)
		shared.CheckError(err)
		fmt.Printf("Estrutura Enviada: ")
		fmt.Println(msgRequest)

		// publica request <-> fila de envio
		err = rabbitChannel.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		shared.CheckError(err)

		// remove da fila
		queue.Dequeue()
	}
}

func hasReconnected() bool {
	_, err := amqp.Dial(amqpURI)
	if err == nil {
		return true
	}
	return false
}
