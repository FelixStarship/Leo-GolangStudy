package main

import (
	"github.com/temporalio/samples-go/polling/frequent"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort:  "10.9.1.1:16345",
		Namespace: "upgradecenter-foo",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, frequent.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(frequent.FrequentPolling)

	activities := &frequent.PollingActivities{}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
