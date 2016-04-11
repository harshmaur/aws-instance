package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/harshmaur/aws-instance/utils"
)

// Need to read various contants from a file
const (
	budget = 0.0268
)

func main() {

	var input int // variable to store user input

	// svc to create new ec2 Session and pass to various functions
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	// fmt.Printf("%T\n %v", svc, svc)

	for {
		// Ask for user input to do various tasks
		fmt.Printf(`1) Get Running Spot Instances
	2) Evaluate Price History
	3) Request a Spot Instance
	4) Termination and Instance
	5) Exit
	Enter your input: `)

		fmt.Scan(&input) // Scan the Input

		// condition on input
		switch input {
		case 1:
			// Get Running Spot Instance Ids
			instanceIds := utils.GetRunningSpotInstanceIds(svc)
			if instanceIds == nil {
				fmt.Println("No instances running")
			} else {
				fmt.Println(instanceIds)
			}
		case 2:
			// Get Instance, Region and Current Bids
			out := utils.EvaluateSpotPriceHistory(svc, budget)
			bs, _ := json.MarshalIndent(out, "", "    ") // this will add indentation for pretty printing
			fmt.Println(string(bs))
		case 3:
			// Request an instance
			rsi := utils.RequestSpotInput{}
			utils.RequestSpotInstance(svc, rsi)
		case 4:
			// Terminate an instance
			utils.TerminateSpotInstance(svc, "someID")
		case 5:
			os.Exit(0)
		}
	}

}
