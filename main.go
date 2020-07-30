package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/arshabbir/logger/logger"
	//"github.com/arshabbir/logger/logger"
)

func main() {

	//Initialize the logger with log path & size of the log files

	fmt.Println("Starting the  main module...")

	log := logger.NewLogger(".\\logs", 2*1024, false)

	wait := make(chan int)

	go writeLog(log, 1)

	go writeLog(log, 2)

	<-wait

	return

}

func writeLog(log logger.Log, id int) {

	var logMesg string
	for {
		randnum := rand.Intn(400000)

		if randnum%5 == 0 {
			logMesg = fmt.Sprintf("%d,,%s", id, "this is a INFO  message")
			log.LogString(logger.INFO, http.StatusOK, logMesg)
		}
		logMesg = fmt.Sprintf("%d,,%s", id, "this is a ERROR message")
		log.LogString(logger.ERROR, http.StatusOK, logMesg)
		time.Sleep(time.Millisecond * 1)
	}

}
