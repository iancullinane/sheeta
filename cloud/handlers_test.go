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
					Content: "cmd deploy something something",
					Author: &discordgo.User{
						ID: "CurrentUser",
					},
					Mentions: []*discordgo.User{
						&discordgo.User{
							ID: "SheetaID",
						},
					},
				},
			},
			Setup: func(r Resources) {
				// r.userStorage.EXPECT().GetGitHubKey("User1").Return("API123", nil)
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
			c.DeployHandler(tc.Session, tc.Message)

		})
	}
}
