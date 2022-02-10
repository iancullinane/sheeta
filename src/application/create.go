package application

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/src/internal/services"
)

func CreateSlashCommands(ssmStore *ssm.SSM) error {

	dToken, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/token"))
	if err != nil {
		panic(err)
	}

	d, err := discordgo.New("Bot " + *dToken.Parameter.Value)
	if err != nil {
		panic(err)
	}

	d.Open()
	defer d.Close()

	for _, v := range commands {

		cmds, _ := d.ApplicationCommands("703973863335264286", "")
		log.Printf("%#v", cmds)
		for _, v := range cmds {
			err := d.ApplicationCommandDelete("703973863335264286", "", v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}

		log.Println("Add " + v.Name)
		_, err := d.ApplicationCommandCreate("703973863335264286", "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	// whs, _ := s.ChannelWebhooks("703965708165447734")
	// for _, v := range whs {
	// 	log.Printf("%#v", v)
	// }

	// for _, v := range commands {
	// 	_, err := s.WebhookCreate("703965708165447734", "basic-command", "")
	// 	if err != nil {
	// 		log.Panicf("Cannot create '%v' webhook: %v", v.Name, err)
	// 	}
	// }

	return nil
}
