package frequent

import (
	"context"
	"time"

	"github.com/temporalio/samples-go/polling"
)

type PollingActivities struct {
	TestService  *polling.TestService
	PollInterval time.Duration
}

func (a *PollingActivities) DoPoll(ctx context.Context) (string, error) {

	//timer1 := time.NewTimer(time.Second * 3)
	timer2 := time.NewTimer(time.Minute * 500)

	i := 0
EXIT:
	for {
		select {

		case <-ctx.Done():
			return "", ctx.Err()

		case <-time.NewTimer(time.Minute * 3).C:
			i++
			a.TestService.GetServiceResult(ctx)
			if i == 3 {
				//break EXIT
			}
		case <-timer2.C:
			break EXIT
		}
		//activity.RecordHeartbeat(ctx)

	}

	return "", nil
}
