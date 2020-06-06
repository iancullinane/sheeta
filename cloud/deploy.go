package cloud

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (cm *cloud) Deploy(s Services, req *cli.Context) error {

	// The req here is the cli app context, these would be keys you set via
	// `--env` in discord
	env := req.String("env")
	stack := req.String("stack")

	log.Println(env)
	log.Println(stack)

	stackTemplateURL := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/templates/%s.yaml",
		cm.cfg[bucketNameKey],
		cm.cfg[regionKey],
		stack,
	)

	sc := getStackConfig(env, stack, cm.cfg[bucketNameKey], cm.s.S3)

	cr := cf.CreateStackInput{
		RoleARN:     aws.String(cm.cfg[cloudRoleKey]),
		StackName:   aws.String(sc.Name),
		TemplateURL: aws.String(stackTemplateURL),
		Capabilities: []*string{
			// TODO::Move to config/flag
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	buildCreateRequest(&cr, sc.CloudConfig, sc.Tags)
	_, err := cm.s.CF.CreateStack(&cr)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("Create Stack Request: %s", aerr)
	}

	return nil
}
