package main

import (
	"github.com/surajsub/temporal-event-example/activities"
	"github.com/surajsub/temporal-event-example/utils"
	"github.com/surajsub/temporal-event-example/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

// Workflow definition

func main() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	defer c.Close()

	// Create a new worker
	w := worker.New(c, utils.TemporalQueueName, worker.Options{})

	// Register workflows and activities
	w.RegisterWorkflow(workflows.ProvisionAWSResources)
	w.RegisterActivity(&activities.AWSImpl{})

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}
