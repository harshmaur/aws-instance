package utils

// SpotPriceHistory is the response we get in Describe Spot Price History
type SpotPriceHistory struct {
	SpotPriceHistory []SpotPriceHistoryDetails `json:"SpotPriceHistory"`
}

// SpotPriceHistoryDetails is the object inside SpotPriceHistory
type SpotPriceHistoryDetails struct {
	TimeStamp          string `json:"Timestamp"`
	ProductDescription string `json:"ProductDescription"`
	InstanceType       string `json:"InstanceType"`
	SpotPrice          string `json:"SpotPrice"`
	AvailibilityZone   string `json:"AvailibiltyZone"`
}
