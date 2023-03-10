package main

import (
	"fmt"
	"time"
)

func main() {

	mychannel1 := make(chan time.Time)

	fmt.Println("Starting periodic task")
	go func() {
		for dh := range time.NewTicker(3 * time.Second).C {
			fmt.Println("Running periodic task at 3 interval time")
			// reading data from db periodicly
			mychannel1 <- dh
		}
	}()

	go func() {
		for {
			recieveMsg := <-mychannel1
			// calling handlers and do update db
			fmt.Println("Received message: ", recieveMsg)
		}
	}()

	time.Sleep(1 * time.Minute)
}
