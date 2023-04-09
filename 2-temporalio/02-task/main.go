package main

import (
	"context"
	"github.com/google/uuid"
	myworkflow "github/FelixStarship/Leo-GolangStudy/2-temporalio/workflow"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {

	c, err := client.Dial(client.Options{
		HostPort:  "10.9.1.1:16345",
		Namespace: "starship-dev",
	})

	if err != nil {
		panic(err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "timer_" + uuid.New().String(),
		TaskQueue: "my-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, myworkflow.MyWorkflow)
	if err != nil {
		panic(err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
