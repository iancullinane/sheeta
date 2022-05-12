package server

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/bwmarrin/discordgo"
)

type serverCommands struct {
	ec2Client EC2InstanceClient
}

func New(ec2Client EC2InstanceClient) *serverCommands {
	return &serverCommands{
		ec2Client: ec2Client,
	}
}

func respond(text string, i *discordgo.Interaction, d *discordgo.Session) {
	d.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			// Note: this isn't documented, but you can use that if you want to.
			// This flag just allows you to create messages visible only for the caller of the command
			// (user who triggered the command)
			Flags:   1 << 6, // ephemeral! https://discord.com/developers/docs/resources/channel#message-object-message-flags
			Content: text,
		},
	})
}

func (h *serverCommands) Handler(i *discordgo.Interaction, d *discordgo.Session) {

	cmd := i.ApplicationCommandData()

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(cmd.Options))
	for _, opt := range cmd.Options {
		optionMap[opt.Name] = opt
	}

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:sheeta"),
				Values: []*string{
					aws.String("server-sorrow"),
				},
			},
		},
	}

	describeResult, err := h.ec2Client.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				respond(aerr.Error(), i, d)
				return
			}
		} else {
			respond(aerr.Error(), i, d)
			return
		}
	}

	if v, ok := optionMap["start-server"]; ok {
		if v.Value == true {
			resp, err := h.startInstance(*describeResult.Reservations[0].Instances[0].InstanceId)
			if err != nil {
				log.Println(err)
			}
			respond(*resp.StartingInstances[0].InstanceId+" starting", i, d)
			return
		} else {
			resp, err := h.stopInstance(*describeResult.Reservations[0].Instances[0].InstanceId)
			if err != nil {
				log.Println(err)
			}
			respond(*resp.StoppingInstances[0].InstanceId+" stopping", i, d)
			return
		}
	} else {
		respond("option does not exist", i, d)
		return
	}
}

func (h *serverCommands) startInstance(instanceID string) (*ec2.StartInstancesOutput, error) {
	startInput := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	resp, err := h.ec2Client.StartInstances(startInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return nil, err
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return nil, err
		}
	}
	return resp, nil
}

func (h *serverCommands) stopInstance(instanceID string) (*ec2.StopInstancesOutput, error) {
	stopInput := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	resp, err := h.ec2Client.StopInstances(stopInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return nil, err
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return nil, err
		}
	}
	return resp, nil
}
