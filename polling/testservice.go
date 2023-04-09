package polling

import (
	"context"
	"go.temporal.io/sdk/activity"
)

type TestService struct {
	tryAttempts   int
	errorAttempts int
}

func NewTestService(errorAttempts int) TestService {
	return TestService{
		tryAttempts:   0,
		errorAttempts: errorAttempts,
	}
}

func (testService *TestService) GetServiceResult(ctx context.Context) (string, error) {

	activity.GetLogger(ctx).Info("GetServiceResult sending notification email as the process takes long time.")

	return "", nil
}
