package cloud

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

func containsUser(s []*discordgo.User, e string) bool {
	for _, a := range s {
		if a.Username == e {
			return true
		}
	}
	return false
}

// getStackConfig gets a matching config from s3 matching the environment
// called when the user entered the `--env` flag
func getStackConfig(env, stack, bucketName string, s3c S3Client) *StackConfig {

	input := s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fmt.Sprintf("env/%s/%s.yaml", env, stack)),
	}

	// This uses the S3 downloader
	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#Downloader
	buf := aws.NewWriteAtBuffer([]byte{})
	dl, err := s3c.Download(buf, &input)
	if aerr, ok := err.(awserr.Error); ok {
		log.Printf("Error fetching key: %s", aerr)
	}

	// Since the path and the stack name are derived from the key path,
	// we need to remove slashes for StackNameParameter compatibility
	sc := StackConfig{}
	if dl > 0 {
		err = yaml.Unmarshal(buf.Bytes(), &sc)
		if err != nil {
			// We don't need to fail here because not every stack has a config
			log.Println(err)
			return nil
		}

		// use the name from the config
		n := sc.Name
		sc.Name = fmt.Sprintf("%s-%s", n, env)
		return &sc
	}

	sc.Name = fmt.Sprintf("%s-%s", env, strings.Replace(stack, "/", "-", -1))

	return &sc
}

// TODO::Something clever aroud the fact the create and update share a similar
// interface for tags and params
func buildCreateRequest(cr *cf.CreateStackInput, rcfg map[string]interface{}, tags map[string]string) {

	for k, v := range rcfg {
		cr.Parameters = append(cr.Parameters, &cf.Parameter{
			ParameterKey:   aws.String(k),
			ParameterValue: aws.String(v.(string)),
		})
	}

	for k, v := range tags {
		cr.Tags = append(cr.Tags, &cf.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

	log.Printf("%#v", cr)

}

func buildUpdateRequest(cr *cf.UpdateStackInput, rcfg map[string]interface{}, tags map[string]string) {

	for k, v := range rcfg {
		cr.Parameters = append(cr.Parameters, &cf.Parameter{
			ParameterKey:   aws.String(k),
			ParameterValue: aws.String(v.(string)),
		})
	}

	for k, v := range tags {
		cr.Tags = append(cr.Tags, &cf.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

}
