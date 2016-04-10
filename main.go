package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/harshmaur/aws-instance/utils"
)

func main() {

	// Need to read various contants from a file
	// Need to ask for user input to do various tasks

	budget := 0.0268
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	// fmt.Printf("%T\n %v", svc, svc)

	// Get Running Spot Instance Ids
	instanceIds := utils.GetRunningSpotInstanceIds(svc)
	if instanceIds == nil {
		fmt.Println("No instances running")
	} else {
		fmt.Println(instanceIds)
	}

	// Get Instance, Region and Current Bids
	out := utils.EvaluateSpotPriceHistory(svc, budget)
	bs, _ := json.MarshalIndent(out, "", "    ") // this will add indentation for pretty printing
	fmt.Println(string(bs))

	// Terminate an instance
	// utils.TerminateSpotInstance(svc, "someID")

	// Request an instance
	rsi := utils.RequestSpotInput{}
	utils.RequestSpotInstance(svc, rsi)
}
