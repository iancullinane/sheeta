package cloud

import (
	"io"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bwmarrin/discordgo"
)

//go:generate mockgen -package cloud -source=interface.go -destination interface_test.go

// Logger defines package log functions
type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// CFClient exposes functions from github.com/aws/aws-sdk-go/service/cloudformation
type CFClient interface {
	CreateStack(input *cf.CreateStackInput) (*cf.CreateStackOutput, error)
	UpdateStack(input *cf.UpdateStackInput) (*cf.UpdateStackOutput, error)
}

// S3Client exposes functions from github.com/aws/aws-sdk-go/service/s3
type S3Client interface {
	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.GetObject
	// GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error)
	Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error)
}

type Msngr interface {
	SendErrorToUser(s *discordgo.Session, err error, channelID string, content string)
	SendSuccessToUser(s *discordgo.Session, channelID string, content string)
}
