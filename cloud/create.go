package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/urfave/cli/v2"

	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// Create is the main logic tied to this lambda
// func (c *cloud) Create(r Resources) func(ctx context.Context, cwEvent events.CloudWatchEvent) error {
// 	return func(ctx context.Context, cwEvent events.CloudWatchEvent) error {
// 		input := cloudformation.CreateStackInput{
// 			// Parameters []*Parameter `type:"list"`
// 		}
// 		// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/#CreateStackInput

// 		c.r.CF.CreateStack(&input)

// 		return nil
// 	}

// }

func (cm *cloud) Deploy(r Resources, req *cli.Context) error {

	cr := cloudformation.CreateStackInput{
		RoleARN:     aws.String("arn:aws:iam::346096930733:role/cf-role-CF-role"),
		StackName:   aws.String(req.String("stack-name")),
		TemplateURL: aws.String(req.String("template-name")),
	}

	_, err := cm.r.CF.CreateStack(&cr)
	if err != nil {
		return err
	}

	return nil
}

// sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.GetRegion())}))
// var (
// 	cr = cloudformation.CreateStackInput{
// 		// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
// 		// role that AWS CloudFormation assumes to create the stack. AWS CloudFormation
// 		// uses the role's credentials to make calls on your behalf. AWS CloudFormation
// 		// always uses this role for all future operations on the stack. As long as
// 		// users have permission to operate on the stack, AWS CloudFormation uses this
// 		// role even if the users don't have permission to pass it. Ensure that the
// 		// role grants least privilege.
// 		//
// 		// If you don't specify a value, AWS CloudFormation uses the role that was previously
// 		// associated with the stack. If no role is available, AWS CloudFormation uses
// 		// a temporary session that is generated from your user credentials.
// 		// Capabilities []*string `type:"list"`
// 		RoleARN:   aws.String("arn:aws:iam::346096930733:role/cf-role-CF-role"),
// 		StackName: aws.String("stack name"),
// 		// Tags: []*cloudformation.Tag{
// 		// 	&cloudformation.Tag{
// 		// 		Key:   aws.String("Name"),
// 		// 		Value: aws.String("TestName"),
// 		// 	},
// 		// },

// 		// Location of file containing the template body. The URL must point to a template
// 		// (max size: 460,800 bytes) that is located in an Amazon S3 bucket. For more
// 		// information, go to the Template Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
// 		// in the AWS CloudFormation User Guide.
// 		//
// 		// Conditional: You must specify either the TemplateBody or the TemplateURL
// 		// parameter, but not both.
// 		TemplateURL: aws.String("someseurl"),
// 	}
// )
