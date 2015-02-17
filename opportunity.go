package closeio

type Opportunity struct {
	Confidence int
	Status string
	Value string
	ValuePeriod string //Monthly, annually, one-time
	Note string
}
