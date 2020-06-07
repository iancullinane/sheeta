package cloud

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (cm *cloud) Update(s Services, req *cli.Context) error {

	// These values are derived via the cli library
	env := req.String("env")
	stack := req.String("stack")

	sc := getStackConfig(env, stack, cm.cfg[bucketNameKey], cm.s.S3)

	stackTemplateURL := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/templates/%s.yaml",
		cm.cfg[bucketNameKey],
		cm.cfg[regionKey],
		stack,
	)

	ur := cf.UpdateStackInput{
		RoleARN:     aws.String(cm.cfg[cloudRoleKey]),
		StackName:   aws.String(sc.Name),
		TemplateURL: aws.String(stackTemplateURL),
		Capabilities: []*string{
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	buildUpdateRequest(&ur, sc.CloudConfig, sc.Tags)

	_, err := cm.s.CF.UpdateStack(&ur)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("Create Stack Request: %s", aerr)
	}

	return nil
}
