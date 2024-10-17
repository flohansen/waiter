package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	timeout := flag.String("timeout", "10s", "The duration for the task to run")
	flag.Parse()

	t, err := time.ParseDuration(*timeout)
	if err != nil {
		log.Fatalf("could not parse timeout duration: %s", err)
	}

	wait := time.After(t)
	timeLeft := t

	for {
		select {
		case <-wait:
			return
		case <-time.After(time.Second):
			timeLeft = timeLeft - time.Second
			log.Printf("time left: %s", timeLeft)
		}
	}
}
