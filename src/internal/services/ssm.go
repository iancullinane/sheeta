package services

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func GetParameter(svc ssmiface.SSMAPI, name *string) (*ssm.GetParameterOutput, error) {
	results, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: name,
	})

	return results, err
}
