package cloud

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (cm *cloud) Update(r Resources, req *cli.Context) error {

	env := req.String("env")
	stack := req.String("stack")

	templateURL := fmt.Sprintf("https://craft-cf-bucket.s3-us-east-2.amazonaws.com/templates/%s.yaml", stack)
	// envConfig := fmt.Sprintf("https://craft-cf-bucket.s3-us-east-2.amazonaws.com/env/%s/%s.yaml", env, stack)

	input := s3.GetObjectInput{
		Bucket: aws.String(cm.cfg[bucketNameKey]),
		Key:    aws.String(fmt.Sprintf("env/%s/%s.yaml", env, stack)),
	}

	buf := aws.NewWriteAtBuffer([]byte{})
	n, err := cm.r.S3.Download(buf, &input)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("Download error: %s", aerr)
	}

	fmt.Printf("Downloaded %v bytes %v\n", len(buf.Bytes()), n)

	var sc *StackConfig
	err = yaml.Unmarshal(buf.Bytes(), &sc)
	if err != nil {
		panic(err)
	}

	cr := cloudformation.UpdateStackInput{
		RoleARN:     aws.String("arn:aws:iam::346096930733:role/chat-ops-role"),
		StackName:   aws.String(sc.Name),
		TemplateURL: aws.String(templateURL),
		Capabilities: []*string{
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	_, err = cm.r.CF.UpdateStack(&cr)
	if aerr, ok := err.(awserr.Error); ok {
		log.Println(aerr)
		return fmt.Errorf("Create Stack Request: %s", aerr)
	}

	log.Println("Deplou func")

	return nil
}
