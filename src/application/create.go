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

	for _, v := range commands {

		cmds, _ := d.ApplicationCommands("703973863335264286", "")
		log.Printf("Deleting %s", v.Name)
		for _, v := range cmds {
			err := d.ApplicationCommandDelete("703973863335264286", "", v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}

		log.Println("Add " + v.Name)
		rsp, err := d.ApplicationCommandCreate("703973863335264286", "703669646502395954", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		log.Printf("%#v", rsp)
	}

	return nil
}

func DeleteSlashCommands(ssmStore *ssm.SSM) error {

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
		log.Printf("Deleting %s", v.Name)
		for _, v := range cmds {
			err := d.ApplicationCommandDelete("703973863335264286", "", v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	return nil
}
