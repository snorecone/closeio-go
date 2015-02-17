package closeio

type Opportunity struct {
	Confidence int `json:"confidence,omitempty"`
	Status string `json:"status,omitempty"`
	Value string `json:"value,omitempty"`
	ValuePeriod string `json:"value_period,omitempty"` //Monthly, annually, one-time
	Note string `json:"note,omitempty"`
}
