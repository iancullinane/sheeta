package deploy

import (
	"encoding/json"
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
	Name        string `yaml:"name"`
	Template    *string
	Result      string
	CloudConfig map[string]string `yaml:"cloud-config"`
	Tags        map[string]string `yaml:"tags"`
}

// 09052b4eb7aadd5864730f884542ed7405e30eab
// asg/auto-scaling-group

func (dc *deployCommands) Handler(data discordgo.ApplicationCommandInteractionData) string {

	// optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(data.Options))
	optionMap := make(map[string]string, len(data.Options))
	for _, opt := range data.Options {
		optionMap[opt.Name] = opt.StringValue()
	}

	log.Println(prettyPrint(optionMap))

	sc, err := dc.getStackConfig("sheeta-config-bucket", optionMap)
	if err != nil {
		return err.Error()
	}

	// "09052b4eb7aadd5864730f884542ed7405e30eab"

	cr := cloudformation.CreateStackInput{
		// RoleARN:     aws.String(cm.cfg[cloudRoleKey]),
		StackName:    aws.String(sc.Name),
		TemplateBody: sc.Template,
		Capabilities: []*string{
			// TODO::Move to config/flag
			aws.String("CAPABILITY_AUTO_EXPAND"),
			aws.String("CAPABILITY_NAMED_IAM"),
		},
	}

	buildCreateRequest(&cr, sc.CloudConfig, sc.Tags)

	resp, err := dc.cfClient.CreateStack(&cr)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return fmt.Sprintf("aws CreateStack err: %v\n%v", err.Error(), prettyPrint(sc))
			}
		}
		return fmt.Sprint(err.Error())
	}

	return *resp.StackId

}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (dc *deployCommands) getFileFromBucket(bucket, key string) ([]byte, error) {
	// This uses the S3 downloader
	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#Downloader
	input := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := dc.s3Client.Download(buf, &input)
	if aerr, ok := err.(awserr.Error); ok {
		return nil, fmt.Errorf("not found for key: %v\naws err: %v", key, aerr.Error())
	}

	return buf.Bytes(), nil
}

// getStackConfig gets a matching config from s3 matching the environment
// called when the user entered the `--env` flag
// func getStackConfig(env, stack, bucketName string, s3c S3Client) *StackConfig {
func (dc *deployCommands) getStackConfig(bucketName string, options map[string]string) (StackConfig, error) {

	// TODO::This is basically a backup sha, handle it a bit better
	sc := StackConfig{}
	sha := "0fbc9359e56b79932d06990aeff9524eafa631dc" // cfg sha
	if _, ok := options["sha"]; !ok {
		options["sha"] = "09052b4eb7aadd5864730f884542ed7405e30eab" // cfn sha
	}
	if _, ok := options["env"]; !ok {
		options["env"] = "dev" // cfn sha
	}
	// if v, ok := options["sha"]; ok {
	// 	sha = v.StringValue()
	// }

	configKey := fmt.Sprintf("%v/env/%v/%v", sha, "dev", options["env-config"])
	log.Println("Getting config from: " + configKey)
	configFile, err := dc.getFileFromBucket(bucketName, configKey)
	if err != nil {
		sc.Result = fmt.Sprintf("get config failed: %v", err.Error())
		return sc, err
	}
	if len(configFile) > 0 {
		log.Println("Stack config found in S3")
		err = yaml.Unmarshal(configFile, &sc)
		if err != nil {
			// We don't need to fail here because not every stack has a config
			log.Printf("that one err %#v", err)
		}
		log.Printf("%#v", sc)
	}

	tmplKey := fmt.Sprintf("%v/templates/%v", options["sha"], options["template"])
	log.Println("Getting template from: " + tmplKey)

	cfnFile, err := dc.getFileFromBucket("sheeta-cfn-bucket", tmplKey)
	if err != nil {
		sc.Result = fmt.Sprintf("get cfn failed: %v", err.Error())
		return sc, err
	}

	if len(cfnFile) > 0 {
		log.Println("Stack template found in S3")
		str := string(cfnFile)
		sc.Template = aws.String(str)
	} else {
		fmt.Println("none found")
	}

	return sc, nil
}

// TODO::Something clever aroud the fact the create and update share a similar
// interface for tags and params
func buildCreateRequest(cr *cloudformation.CreateStackInput, rcfg map[string]string, tags map[string]string) {

	for k, v := range rcfg {
		cr.Parameters = append(cr.Parameters, &cloudformation.Parameter{
			ParameterKey:   aws.String(k),
			ParameterValue: aws.String(v),
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
