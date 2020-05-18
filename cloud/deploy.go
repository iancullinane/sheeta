package cloud

import (
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
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
	dl, err := cm.r.S3.Download(buf, &input)
	if aerr, ok := err.(awserr.Error); ok {
		log.Printf("Error fetching key: %s", aerr)
	}

	sc := StackConfig{}
	if dl > 0 {
		err = yaml.Unmarshal(buf.Bytes(), &sc)
		if err != nil {
			panic(err)
		}
	}

	// Have to do this anyway to deal with slashes
	sc.Name = fmt.Sprintf("%s-%s", env, strings.Replace(stack, "/", "-", -1))
	log.Println(sc.Name)

	cr := cf.CreateStackInput{
		RoleARN:     aws.String(cm.cfg[cloudRoleKey]),
		StackName:   aws.String(sc.Name),
		TemplateURL: aws.String(templateURL),
		Capabilities: []*string{
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	for k, v := range sc.CloudConfig {
		cr.Parameters = append(cr.Parameters, &cf.Parameter{
			ParameterKey:   aws.String(k),
			ParameterValue: aws.String(v.(string)),
		})
	}

	for k, v := range sc.Tags {
		cr.Tags = append(cr.Tags, &cf.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

	_, err = cm.r.CF.CreateStack(&cr)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("Create Stack Request: %s", aerr)
	}

	return nil
}
