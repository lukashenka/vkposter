package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/lukashenka/vkposter/config"
	vp "github.com/lukashenka/vkposter/vp"
)

func init() {

}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	p := vp.InitProcessing()
	go p.Start()

	<-stop
	log.Println("Shutting down the server...")
	p.Stop()

	log.Println("Server gracefully stopped")

}
