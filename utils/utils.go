package utils

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetRunningSpotInstanceIds uses DescribeSpotInstanceRequests to get the requests and checks the status and their instance id.
func GetRunningSpotInstanceIds(svc *ec2.EC2) []string {

	params := &ec2.DescribeInstancesInput{

		// Filter spot instances
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-lifecycle"), // "spot" instance lifecycle
				Values: []*string{aws.String("spot")},
			},
			{
				Name:   aws.String("instance-state-name"), // "running" instance state
				Values: []*string{aws.String("running")},
			},
		},
	}

	resp, _ := svc.DescribeInstances(params)

	var instanceIds []string
	for i := range resp.Reservations {
		for _, inst := range resp.Reservations[i].Instances {
			instanceIds = append(instanceIds, *inst.InstanceId)
			// fmt.Println(" - Instance id ", *inst.InstanceId)
		}
	}
	return instanceIds
}

// TerminateSpotInstance will terminate the spot instance

// RequestSpotInstance will request an instance based on current spot price in a particular region.
// Also include waiting to see if the request is fulfilled or not, otherwise make another request
func RequestSpotInstance() {

}

// Need a function to check the status "marked-for-termination" and make necessary changes and Request a new Spot instance withing 2 minutes.
// http://169.254.169.254/latest/meta-data/spot/termination-time

// EvaluateSpotPriceHistory will evaluate prices, returns the instance type and the max bid to be used
// Availibility Zone - us-east-1
// Max Budget is 0.028$ per hour for all instances.
func EvaluateSpotPriceHistory(svc *ec2.EC2) {

	// Budget
	// Give current Price and suggest "how likely" the price is going to exceed the budget or not.

	params := &ec2.DescribeSpotPriceHistoryInput{
		EndTime: aws.Time(time.Now()),
		Filters: []*ec2.Filter{
			{ // Availibility Zone
				Name:   aws.String("availability-zone"),
				Values: []*string{aws.String("us-east-1?")},
			},
			{ // Instance Types
				Name: aws.String("instance-type"),
				Values: []*string{
					aws.String("m4.large"),
					aws.String("m3.medium"),
					aws.String("m3.large"),
					aws.String("m3.xlarge"),
					aws.String("c4.large"),
					aws.String("c3.large"),
					aws.String("c3.xlarge"),
					aws.String("r3.large"),
				},
			},
			{ // Product Description
				Name:   aws.String("product-description"),
				Values: []*string{aws.String("Linux/UNIX")},
			},
		},
		StartTime: aws.Time(time.Now()), // setting now will give current spot prices
	}
	resp, err := svc.DescribeSpotPriceHistory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
