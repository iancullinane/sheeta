package cloud

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// Resources are API's needed to execute a task
type Resources struct {
	CF     CFClient
	Logger Logger
}

// Create is the main logic tied to this lambda
func Create(r Resources) func(ctx context.Context, cwEvent events.CloudWatchEvent) error {
	return func(ctx context.Context, cwEvent events.CloudWatchEvent) error {

		return nil
	}

}

// sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.GetRegion())}))
