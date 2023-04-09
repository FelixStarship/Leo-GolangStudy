package workflow

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
	"time"
)

func MyWorkflow(ctx workflow.Context) error {

	// 定义定时器
	future := workflow.NewTimer(ctx, time.Hour*24) // 在每天的9点启动工作流

	// 在这里编写你要执行的代码
	fmt.Println("任务开始时间已到达！")
	return nil
}
