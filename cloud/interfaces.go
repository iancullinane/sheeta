package cloud

import cf "github.com/aws/aws-sdk-go/service/cloudformation"

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type CFClient interface {
	CreateStack(input *cf.CreateStackInput) (*cf.CreateStackOutput, error)
}
