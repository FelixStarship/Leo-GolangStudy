package main

import (
	myworkflow "github/FelixStarship/Leo-GolangStudy/2-temporalio/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	// 创建Temporal Worker
	w := worker.New(c, "my-task-queue", worker.Options{})

	// 注册工作流
	w.RegisterWorkflow(myworkflow.MyWorkflow)

	// 启动Temporal Worker
	if err := w.Run(worker.InterruptCh()); err != nil {
		panic(err)
	}
}
