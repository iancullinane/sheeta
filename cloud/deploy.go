package cloud

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (cm *cloud) Deploy(r Resources, req *cli.Context) error {

	env := req.String("env")
	stack := req.String("stack")

	// TODO::Move to config
	templateURL := fmt.Sprintf("https://craft-cf-bucket.s3-us-east-2.amazonaws.com/templates/%s.yaml", stack)

	input := s3.GetObjectInput{
		Bucket: aws.String(cm.cfg[bucketNameKey]),
		Key:    aws.String(fmt.Sprintf("env/%s/%s.yaml", env, stack)),
	}

	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := cm.r.S3.Download(buf, &input)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("Download error: %s", aerr)
	}

	var sc *StackConfig
	err = yaml.Unmarshal(buf.Bytes(), &sc)
	if err != nil {
		panic(err)
	}

	cr := cloudformation.CreateStackInput{
		RoleARN:     aws.String(cm.cfg["cloud-role"]),
		StackName:   aws.String(sc.Name),
		TemplateURL: aws.String(templateURL),
		Capabilities: []*string{
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	_, err = cm.r.CF.CreateStack(&cr)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("Create Stack Request: %s", aerr)
	}

	return nil
}
