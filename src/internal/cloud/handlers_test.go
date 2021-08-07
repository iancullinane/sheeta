package cloud

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
)

func TestUnit_SubmitEvent(t *testing.T) {

	config := map[string]string{
		"bucketName": "bucket_name",
		"cloudRole":  "arn:aws:iam::123456789:role/chat-ops-role",
		"region":     "us-east-2",
	}

	msg := &discordgo.MessageCreate{
		&discordgo.Message{
			Content: "sheeta cloud deploy --env dev --stack test-stack",
			Author: &discordgo.User{
				ID: "CurrentUser",
			},
			Mentions: []*discordgo.User{
				&discordgo.User{
					ID:       "SheetaID",
					Username: "sheeta",
				},
			},
		},
	}

	configInput := s3.GetObjectInput{
		Bucket: aws.String("bucket_name"),
		Key:    aws.String(fmt.Sprintf("env/%s/%s.yaml", "dev", "test-stack")),
	}

	cfnInput := cloudformation.CreateStackInput{
		Capabilities: []*string{aws.String("CAPABILITY_AUTO_EXPAND"), aws.String("CAPABILITY_NAMED_IAM")},
		RoleARN:      aws.String("arn:aws:iam::123456789:role/chat-ops-role"),
		StackName:    aws.String("dev-test-stack"),
		TemplateURL:  aws.String("https://bucket_name.s3-us-east-2.amazonaws.com/templates/test-stack.yaml"),
	}

	sesh := &discordgo.Session{
		State:                  discordgo.NewState(),
		Ratelimiter:            discordgo.NewRatelimiter(),
		StateEnabled:           true,
		Compress:               true,
		ShouldReconnectOnError: true,
		ShardID:                0,
		ShardCount:             1,
		MaxRestRetries:         3,
		Client:                 &http.Client{Timeout: (20 * time.Second)},
		UserAgent:              "DiscordBot",
		LastHeartbeatAck:       time.Now().UTC(),
	}

	sesh.State.User = &discordgo.User{
		ID:       "SheetaID",
		Username: "sheeta",
	}

	type Resources struct {
		mockS3  *MockS3Client
		mockCfn *MockCFClient
		logger  *MockLogger
	}

	createTestResources := func(t *testing.T) (Resources, func()) {
		ctrl := gomock.NewController(t)

		r := Resources{
			mockS3:  NewMockS3Client(ctrl),
			mockCfn: NewMockCFClient(ctrl),
			logger:  NewMockLogger(ctrl),
		}

		return r, ctrl.Finish
	}

	for _, tc := range []struct {
		Name    string
		Session *discordgo.Session
		Message *discordgo.MessageCreate
		// DownloadInput
		Setup func(r Resources)
	}{
		{
			Name:    "deploy - no existing config",
			Session: sesh,
			Message: msg,
			Setup: func(r Resources) {
				r.mockS3.EXPECT().Download(gomock.Any(), &configInput).Return(*aws.Int64(0), nil)
				r.mockCfn.EXPECT().CreateStack(&cfnInput).Return(nil, nil)
			},
		},
	} {

		t.Run(tc.Name, func(t *testing.T) {
			r, stop := createTestResources(t)
			defer stop()

			tc.Setup(r)

			services := Services{
				S3: r.mockS3,
				CF: r.mockCfn,
			}

			c := NewCloud(services, config)
			c.Handler(tc.Session, tc.Message)

		})
	}
}
