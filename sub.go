package main

import (
	"log"

	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/audreylim/vspark"
)

func main() {
	// TODO: read from config and handle flags
	broker := "tcp://128.199.132.229:60000"

	errChan := make(chan error)
	go func() {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(broker)
		opts.SetClientId("sample-sub")
		opts.SetCleanSession(true)

		cl := mqtt.NewClient(opts)
		_, err := cl.Start()
		if err != nil {
			errChan <- err
		}

		// subscribe to a topic
		// we are setting QOS to zero in this case
		tf, err := mqtt.NewTopicFilter("vspark", byte(mqtt.QOS_ZERO))
		if err != nil {
			errChan <- err
		}

		_, err = cl.StartSubscription(handleMessage, tf)
		if err != nil {
			errChan <- err
		}
	}()

	// ping spark to retrieve its IP address from Spark Cloud
	// err = vspark.PingSpark()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("waiting for messages...")
	select {
	case e := <-errChan:
		log.Fatal(e)
	}
}

func handleMessage(client *mqtt.MqttClient, msg mqtt.Message) {
	p := msg.Payload()
	log.Printf("received message: %s", p)

	val := p[2]
	pin := string(p[0:2])

	err := vspark.PinMode(pin, "OUTPUT")
	if err != nil {
		log.Fatal(err)
	}

	err = vspark.DigitalWrite(pin, val)
	if err != nil {
		log.Fatal(err)
	}
}
