package exp

import (
	"github.com/bwmarrin/discordgo"
)

type expCommands struct{}

func New() *expCommands {
	return &expCommands{}
}

func (exp *expCommands) Handler(i *discordgo.Interaction, d *discordgo.Session) {

	d.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			// Note: this isn't documented, but you can use that if you want to.
			// This flag just allows you to create messages visible only for the caller of the command
			// (user who triggered the command)
			Flags:   1 << 6,
			Content: "Surprise!",
		},
	})
}

// func f1(quit chan bool) {
// 	go func() {
// 		for {
// 			println("f1 is working...")
// 			time.Sleep(1 * time.Second)

// 			select {
// 			case <-quit:
// 				fmt.Println("stopping")
// 				return
// 			default:
// 			}
// 		}
// 	}()
// }
