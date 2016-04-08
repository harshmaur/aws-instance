package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/harshmaur/aws-instance/utils"
)

func main() {

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	// fmt.Printf("%T\n %v", svc, svc)

	// // Get Running Spot Instance Ids
	// instanceIds := utils.GetRunningSpotInstanceIds(svc)
	// if instanceIds == nil {
	// 	fmt.Println("No instances running")
	// } else {
	// 	fmt.Println(instanceIds)
	// }

	utils.EvaluateSpotPriceHistory(svc)

}
