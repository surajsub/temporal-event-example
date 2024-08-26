package main

import (
	"context"
	"encoding/json"
	"github.com/surajsub/temporal-event-example/utils"
	"github.com/surajsub/temporal-event-example/workflows"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"go.temporal.io/sdk/client"
)

const BUILD = "provision-workflow-"

// ProvisionMessage is the structure of the message we expect to receive
type ProvisionMessage struct {
	ResourceType string            `json:"resource_type"`
	Params       map[string]string `json:"params"`
	Operation    string            `json:"operation,omitempty"`
}

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	log.Println("Connected to NATS...")

	// Subscribe to NATS topic
	_, err = nc.Subscribe(utils.NATS_AWS_PROVISION_QUEUE, func(m *nats.Msg) {
		log.Println("Received a message from NATS")

		// Unmarshal the message into the ProvisionMessage struct
		var provisionMsg ProvisionMessage
		if err := json.Unmarshal(m.Data, &provisionMsg); err != nil {
			log.Printf("Error unmarshaling NATS message: %v", err)
			return
		}

		log.Printf("Message received: %+v\n", provisionMsg)

		// Start the Temporal workflow
		if err := StartProvisioningWorkflow(provisionMsg); err != nil {
			log.Printf("Error starting provisioning workflow: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to NATS: %v", err)
	}

	log.Println("Listening for NATS messages...")

	// Keep the connection alive
	select {}
}

func StartProvisioningWorkflow(provisionMsg ProvisionMessage) error {
	// Create a Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		return err
	}
	defer c.Close()

	log.Println("Temporal client created successfully")

	if provisionMsg.Operation == "destroy" {
		log.Printf("Doing a destroy Operatio")
	} else {
		log.Println("We are doing a build")
	}

	// Workflow options
	options := client.StartWorkflowOptions{
		ID:        BUILD + provisionMsg.ResourceType + "-" + time.Now().Format("20060102150405"),
		TaskQueue: "provision-task-queue",
	}

	// Start the workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, workflows.ProvisionAWSResources, provisionMsg.ResourceType, provisionMsg.Params)
	if err != nil {
		log.Printf("Error executing workflow: %v", err)
		return err
	}

	log.Printf("Started workflow with WorkflowID: %s, RunID: %s", we.GetID(), we.GetRunID())

	return nil
}
