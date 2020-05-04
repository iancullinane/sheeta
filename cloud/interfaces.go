package cloud

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Logger defines package log functions
type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// CFClient exposes functions from github.com/aws/aws-sdk-go/service/cloudformation
type CFClient interface {
	CreateStack(input *cf.CreateStackInput) (*cf.CreateStackOutput, error)
}

// S3Client exposes functions from github.com/aws/aws-sdk-go/service/s3
type S3Client interface {
	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.GetObject
	GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error)
}
