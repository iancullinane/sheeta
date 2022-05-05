package server

import "github.com/aws/aws-sdk-go/service/ec2"

type EC2InstanceClient interface {
	DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
	StartInstances(input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)
	StopInstances(input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)
}
