package accounting

type AccountingFlow struct {
	FpcSlot               string
	InlineActiveFlows     float64
	InlineIpv4ActiveFlows float64
	InlineIpv6ActiveFlows float64

	InlineFlows          float64
	InlineIpv4TotalFlows float64
	InlineIpv6TotalFlows float64
}

type AccountingError struct {
	FpcSlot                        string
	InlineFlowCreationFailures     float64
	InlineIpv4FlowCreationFailures float64
	InlineIpv6FlowCreationFailures float64
}
