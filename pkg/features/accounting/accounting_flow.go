package accounting

type accountingFlow struct {
	FpcSlot               string
	InlineActiveFlows     float64
	InlineIPv4ActiveFlows float64
	InlineIPv6ActiveFlows float64

	InlineFlows          float64
	InlineIPv4TotalFlows float64
	InlineIPv6TotalFlows float64
}

type accountingError struct {
	FpcSlot                        string
	InlineFlowCreationFailures     float64
	InlineIPv4FlowCreationFailures float64
	InlineIPv6FlowCreationFailures float64
}
