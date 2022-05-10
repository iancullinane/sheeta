package exp

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/src/internal/bot"
)

type expCommands struct{}

func New() *expCommands {
	return &expCommands{}
}

func (exp *expCommands) Handler(data discordgo.ApplicationCommandInteractionData, ctl bot.Controller) string {

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		ctl.MsgCh <- fmt.Sprint(i)
	}

	for {
		select {
		case msg := <-ctl.MsgCh:
			fmt.Println("received:", msg)
		case <-ctl.Done:
			return "done"
		}
	}
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
