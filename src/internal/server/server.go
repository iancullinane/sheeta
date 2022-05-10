package server

import (
	"fmt"
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

func (h *serverCommands) Handler(data discordgo.ApplicationCommandInteractionData) string {

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(data.Options))
	for _, opt := range data.Options {
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
				return fmt.Sprintf("%s: %s", "aws error", aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return err.Error()
		}
	}

	if v, ok := optionMap["start-server"]; ok {
		if v.Value == true {
			resp, err := h.startInstance(*describeResult.Reservations[0].Instances[0].InstanceId)
			if err != nil {
				log.Println(err)
			}
			return *resp.StartingInstances[0].InstanceId + " starting"
		} else {
			resp, err := h.stopInstance(*describeResult.Reservations[0].Instances[0].InstanceId)
			if err != nil {
				log.Println(err)
			}

			return *resp.StoppingInstances[0].InstanceId + " stopping"
		}
	} else {
		return "option does not exist"
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
