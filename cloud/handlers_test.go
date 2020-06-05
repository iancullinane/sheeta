package cloud

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
)

func TestUnit_SubmitEvent(t *testing.T) {

	ready := discordgo.Ready{
		User: &discordgo.User{
			ID: "Current User",
		},
	}

	sesh := &discordgo.Session{
		State: &discordgo.State{
			Ready: ready,
		},
	}

	type Resources struct {
		mockBot *MockBot
		mockS3  *MockS3Client
		mockCfn *MockCFClient
		logger  *MockLogger
	}

	createTestResources := func(t *testing.T) (Resources, func()) {
		ctrl := gomock.NewController(t)

		r := Resources{
			mockBot: NewMockBot(ctrl),
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
		Setup   func(r Resources)
	}{
		{
			Name: "success - message with app mention",
			Session: &discordgo.Session{
				State: &discordgo.State{
					Ready: ready,
				},
			},
			Message: &discordgo.MessageCreate{
				&discordgo.Message{
					Content:   "value cloud deploy lies",
					ChannelID: "test_channel",
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
			},
			Setup: func(r Resources) {
				r.mockBot.EXPECT().SendErrorToUser(sesh, gomock.Any(), gomock.Any(), gomock.Any())
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

			c := NewCloud(services, map[string]string{})
			c.GenerateCLI()
			c.DeployHandler(tc.Session, tc.Message)

		})
	}
}
