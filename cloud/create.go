package cloud

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// Resources are API's needed to execute a task
type Resources struct {
	CF     CFClient
	Logger Logger
}

// Create is the main logic tied to this lambda
func (c *cloud) Create(r Resources) func(ctx context.Context, cwEvent events.CloudWatchEvent) error {
	return func(ctx context.Context, cwEvent events.CloudWatchEvent) error {
		input := cloudformation.CreateStackInput{}
		// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/#CreateStackInput

		c.r.CF.CreateStack(&input)

		return nil
	}

}

// sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.GetRegion())}))
