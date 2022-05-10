package deploy

import (
	"io"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type CloudFormationDeployClient interface {
	CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
}

type S3Client interface {
	Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error)
}
