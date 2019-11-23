package shared

import (
	"fmt"
	"os"
	"time"
)

const SAMPLE_SIZE = 10000 // Warm-up: 30% , Post: 10%
const RABBITMQ_PORT = 5672

const QUEUE_SERVER_HOST = "localhost"
const QUEUE_SERVER_PORT = 7574

const REAL_TIME = 10 * time.Second
const EXPERIMENT_TIEM = 0

// CheckError is
func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
