package main

import (
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"github.com/surajsub/temporal-event-example/utils"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {

	paramData, resource, err := utils.ParseInputFlags()
	if err != nil {
		fmt.Println("Validation error:", err)
		os.Exit(1)
	}
	fmt.Println("All parameters validated successfully.")

	// Create the message
	msg := models.ProvisionMessage{
		ResourceType: resource,
		Params:       paramData,
	}

	// Marshal the message to JSON

	log.Printf("The data being send to the the systems is %v\n", msg.Params)
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error marshaling message to JSON: %v", err)
	}

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	// Publish the message
	err = nc.Publish(utils.NATS_AWS_PROVISION_QUEUE, data)
	if err != nil {
		log.Fatalf("Error publishing message to NATS: %v", err)
	}

	log.Println("Message sent successfully!")
}
