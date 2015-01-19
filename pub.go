package main

import (
	"log"

	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func main() {
	// TODO: read from config and handle flags
	broker := "tcp://128.199.132.229:60000"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientId("sample-pub")

	cl := mqtt.NewClient(opts)
	_, err := cl.Start()
	if err != nil {
		log.Fatal(err)
	}

	// publish a message
	m := mqtt.NewMessage([]byte{68, 55, 1})
	<-cl.PublishMessage("vspark", m)
	log.Println("published message")
}
