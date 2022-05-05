package deploy

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

type deployCommands struct {
	cfClient CloudFormationDeployClient
	s3Client S3Client
}

func New(cfClient CloudFormationDeployClient, s3Client S3Client) *deployCommands {
	return &deployCommands{
		cfClient: cfClient,
		s3Client: s3Client,
	}
}

// StackConfig is used to generate the request
type StackConfig struct {
	Name        string                 `yaml:"name"`
	CloudConfig map[string]interface{} `yaml:"cloud-config"`
	Tags        map[string]string      `yaml:"tags"`
}

func (dc *deployCommands) Handler(data discordgo.ApplicationCommandInteractionData) string {

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(data.Options))
	for _, opt := range data.Options {
		optionMap[opt.Name] = opt
	}

	sc, err := dc.getStackConfig("sheeta-config-bucket", optionMap)
	if err != nil {
		return err.Error()
	}

	// "09052b4eb7aadd5864730f884542ed7405e30eab"

	if sc == nil {
		return "config does not exist in s3 for this template and environment"
	}

	cr := cloudformation.CreateStackInput{
		// RoleARN:     aws.String(cm.cfg[cloudRoleKey]),
		StackName:   aws.String(sc.Name),
		TemplateURL: aws.String("not sure"),
		Capabilities: []*string{
			// TODO::Move to config/flag
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	log.Printf("%#v", cr)

	buildCreateRequest(&cr, sc.CloudConfig, sc.Tags)

	_, err = dc.cfClient.CreateStack(&cr)
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("create Stack Request: %s", aerr).Error()
	} else {
		return err.Error()
	}

}

// getStackConfig gets a matching config from s3 matching the environment
// called when the user entered the `--env` flag
// func getStackConfig(env, stack, bucketName string, s3c S3Client) *StackConfig {
func (dc *deployCommands) getStackConfig(bucketName string, options map[string]*discordgo.ApplicationCommandInteractionDataOption) (*StackConfig, error) {

	// TODO::This is basically a backup sha, handle it a bit better
	sha := "0fbc9359e56b79932d06990aeff9524eafa631dc"
	// if v, ok := options["sha"]; ok {
	// 	sha = v.StringValue()
	// }

	tmpl, err := fixYamlSuffix(options["template"].StringValue())
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v/env/%v/%v", sha, options["env"].StringValue(), tmpl)

	input := s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	// This uses the S3 downloader
	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#Downloader
	buf := aws.NewWriteAtBuffer([]byte{})
	dl, err := dc.s3Client.Download(buf, &input)
	if aerr, ok := err.(awserr.Error); ok {
		return nil, fmt.Errorf("not found for key: %v\n%#v\naws err: %v", key, options, aerr.Error())
	}

	// Since the path and the stack name are derived from the key path,
	// we need to remove slashes for StackNameParameter compatibility
	var sc StackConfig
	if dl > 0 {
		log.Println("Stack config found in S3")
		err = yaml.Unmarshal(buf.Bytes(), &sc)
		if err != nil {
			// We don't need to fail here because not every stack has a config
			return nil, err
		}
		log.Printf("%#v", sc)
		return &sc, nil
	}

	return nil, nil
}

// TODO::Something clever aroud the fact the create and update share a similar
// interface for tags and params
func buildCreateRequest(cr *cloudformation.CreateStackInput, rcfg map[string]interface{}, tags map[string]string) {

	for k, v := range rcfg {
		cr.Parameters = append(cr.Parameters, &cloudformation.Parameter{
			ParameterKey:   aws.String(k),
			ParameterValue: aws.String(v.(string)),
		})
	}

	for k, v := range tags {
		cr.Tags = append(cr.Tags, &cloudformation.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

}

func fixYamlSuffix(s string) (string, error) {
	r := []rune(s)
	if !unicode.IsLetter(r[len(s)-1]) {
		return "", fmt.Errorf("template must end in a letter")
	}
	if !strings.HasSuffix(s, ".yaml") {
		s = s + ".yaml"
		return s, nil
	}
	return s, nil
}

// TODO::Something clever aroud the fact the create and update share a similar
// interface for tags and params
// func buildCreateRequest(cr *cf.CreateStackInput, rcfg map[string]interface{}, tags map[string]string) {

// 	for k, v := range rcfg {
// 		cr.Parameters = append(cr.Parameters, &cf.Parameter{
// 			ParameterKey:   aws.String(k),
// 			ParameterValue: aws.String(v.(string)),
// 		})
// 	}

// 	for k, v := range tags {
// 		cr.Tags = append(cr.Tags, &cf.Tag{
// 			Key:   aws.String(k),
// 			Value: aws.String(v),
// 		})
// 	}

// }
