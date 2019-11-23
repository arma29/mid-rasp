package shared

import (
	"fmt"
	"os"
)

const SAMPLE_SIZE = 10000 // Warm-up: 30% , Post: 10%
const RABBITMQ_PORT = 5672

const QUEUE_SERVER_HOST = "localhost"
const QUEUE_SERVER_PORT = 7574

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}

// same as fib parameter
type Request struct {
	Req int32
}

// same as fib response
type Response struct {
	Res int32
}
