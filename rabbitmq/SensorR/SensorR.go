package sensorR

// import (
// 	"github.com/streadway/amqp"
// 	"encoding/json"
// 	"fmt"
// 	"time"
// 	"os"
// 	"strconv"
	
// 	"github.com/arma29/mid-mid-perf/shared"
// )

// func main() {
// 	// Get Argument from command Line
// 	if len(os.Args) != 3 {
// 		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
// 		os.Exit(1)
// 	}

// 	ipContainer := os.Args[1]

// 	// conecta ao servidor de mensageria
// 	conn, err := amqp.Dial("amqp://guest:guest@" + 
// 		ipContainer + ":" + 
// 		strconv.Itoa(shared.RABBITMQ_PORT) +"/")
// 	shared.CheckError(err)
// 	defer conn.Close()

// 	// cria o canal
// 	ch, err := conn.Channel()
// 	shared.CheckError(err)
// 	defer ch.Close()

// 	// declara  filas, cria se não existir
// 	requestQueue, err := ch.QueueDeclare( // fila de envio
// 		"request", false, 	false, 	false, 	false, 	nil, )	
// 	shared.CheckError(err)

// 	replyQueue, err := ch.QueueDeclare( // fila de respostas	
// 		"response", false, false, false, false, nil, )
// 	shared.CheckError(err)

// 	// cria consumidor <-> fila de respostas -> async
// 	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
// 		false, false, nil, )
// 	shared.CheckError(err)

// 	number, _ := strconv.Atoi(os.Args[2])
// 	fmt.Println("Fibonacci,Answer,Time")
// 	for i := 0; i<shared.SAMPLE_SIZE; i++{

// 		t1 := time.Now()

// 		// prepara request
// 		msgRequest := shared.Request{Req: int32(number)} // Fibonacci 5
// 		msgRequestBytes,err := json.Marshal(msgRequest)
// 		shared.CheckError(err)

// 		// publica request <-> fila de envio 
// 		err = ch.Publish("", requestQueue.Name, false, false,
// 			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes,})
// 		shared.CheckError(err)

// 		// recebe resposta em bytes
// 		x := <- msgsFromServer

// 		// deserializa
// 		msgReply := shared.Response{}
// 		err = json.Unmarshal([]byte(x.Body), &msgReply)
// 		shared.CheckError(err)

// 		t2 := time.Now()
// 		xtime := float64(t2.Sub(t1).Nanoseconds()) / 1000000
// 		s := fmt.Sprintf("%d,%d,%f", number, msgReply.Res, xtime)
// 		fmt.Println(s)
// 	}
// }