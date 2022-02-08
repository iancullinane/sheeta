package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func GetParameterDecrypted(svc ssmiface.SSMAPI, name *string) (*ssm.GetParameterOutput, error) {
	results, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           name,
		WithDecryption: aws.Bool(true),
	})

	return results, err
}
