package utils

import (
	"fmt"
	"sort"
	"strconv"
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

	resp, _ := svc.DescribeInstances(params) // Error handling to do

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
func TerminateSpotInstance(svc *ec2.EC2, instanceID string) {

	params := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}
	resp, _ := svc.TerminateInstances(params) // Error handling to do

	// Pretty-print the response data.
	fmt.Println(resp)
}

// RequestSpotInstance will request an instance based on current spot price in a particular region.
// Also include waiting to see if the request is fulfilled or not, otherwise make another request
func RequestSpotInstance(svc *ec2.EC2, rsi RequestSpotInput) {
	params := &ec2.RequestSpotInstancesInput{
		SpotPrice:             aws.String(rsi.SpotPrice), // Required
		AvailabilityZoneGroup: aws.String(rsi.AvailibiltyZoneGroup),
		LaunchSpecification: &ec2.RequestSpotLaunchSpecification{
			ImageId:          aws.String(rsi.ImageID),
			InstanceType:     aws.String(rsi.InstanceType),
			KeyName:          aws.String(rsi.KeyName),
			SecurityGroupIds: []*string{aws.String(rsi.SecurityGroupID)},
			UserData:         aws.String(rsi.UserData),
		},
	}
	resp, _ := svc.RequestSpotInstances(params) // Error handling to do

	// Pretty-print the response data.
	fmt.Println(resp)
}

// CheckTerminationMeta to check the status "marked-for-termination" and make necessary changes and Request a new Spot instance withing 2 minutes.
// http://169.254.169.254/latest/meta-data/spot/termination-time
func CheckTerminationMeta() {

}

// EvaluateSpotPriceHistory will evaluate prices, returns the instance type, region and the max bid to be used
// Max Budget is 0.028$ per hour for all instances.
func EvaluateSpotPriceHistory(svc *ec2.EC2, budget float64) EvaluateSpotPriceHistoryOutput {

	// PENDING::Give current Price and suggest "how likely" the price is going to exceed the budget or not.

	params := &ec2.DescribeSpotPriceHistoryInput{
		EndTime: aws.Time(time.Now()),
		Filters: []*ec2.Filter{
			{ // Availibility Zone
				Name:   aws.String("availability-zone"),
				Values: []*string{aws.String("us-east-1?")}, // Availibility Zone - us-east-1
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

	resp, _ := svc.DescribeSpotPriceHistory(params) // Error handling to do

	// Pretty-print the response data.
	// fmt.Println(resp)
	var espho SpotInstanceDetails
	var output EvaluateSpotPriceHistoryOutput
	for _, sph := range resp.SpotPriceHistory { // range over the response

		if i, _ := strconv.ParseFloat(*sph.SpotPrice, 64); i < budget { // convert to float64 and check if it is less than budget

			// add values to struct
			espho.AvailibiltyZone = *sph.AvailabilityZone
			espho.InstanceType = *sph.InstanceType
			espho.SpotPrice = *sph.SpotPrice
			output = append(output, espho) // append to the slice
		}
	}
	sort.Sort(output) // Sorting on the basis of SpotPrice
	return output
}
