package utils

import "strconv"

// SpotPriceHistory is the response we get in Describe Spot Price History
type SpotPriceHistory struct {
	SpotPriceHistory []SpotPriceHistoryDetails `json:"SpotPriceHistory"`
}

// SpotInstanceDetails gives output for the EvaluateSpotPriceHistory function
type SpotInstanceDetails struct {
	InstanceType    string `json:"InstanceType"`
	AvailibiltyZone string `json:"AvailibilityZone"`
	SpotPrice       string `json:"SpotPrice"`
}

// SpotPriceHistoryDetails is the object inside SpotPriceHistory
type SpotPriceHistoryDetails struct {
	TimeStamp           string `json:"Timestamp"`
	ProductDescription  string `json:"ProductDescription"`
	SpotInstanceDetails        // Used from the struct defined above.
}

// EvaluateSpotPriceHistoryOutput is slice of details and also implements sort.Interface interface
type EvaluateSpotPriceHistoryOutput []SpotInstanceDetails

// Implementing methods to use sort interface
func (e EvaluateSpotPriceHistoryOutput) Len() int      { return len(e) }
func (e EvaluateSpotPriceHistoryOutput) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func (e EvaluateSpotPriceHistoryOutput) Less(i, j int) bool {
	left, _ := strconv.ParseFloat(e[i].SpotPrice, 64)
	right, _ := strconv.ParseFloat(e[j].SpotPrice, 64)
	return left < right
}
