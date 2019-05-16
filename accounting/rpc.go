package accounting

type AccountingFlowRpc struct {
	Information struct {
		InlineFlow AccountingInlineFlow `xml:"inline-jflow-flow-information"`
	} `xml:"services-accounting-information"`
	Error struct {
		Message string `xml:"message"`
	} `xml:"error,omitempty"`
}

type AccountingInlineFlow struct {
	FpcSlot                    string `xml:"fpc-slot"`
	InlineFlowPackets          int64  `xml:"inline-flow-packets"`
	InlieFlowBytes             int64  `xml:"inline-flow-bytes"`
	InlineActiveFlows          int64  `xml:"inline-active-flowst"`
	InlineFlows                int64  `xml:"inline-flows"`
	InlineFlowsExported        int64  `xml:"inline-flows-exported"`
	InlineFlowsPacketsExported int64  `xml:"inline-flow-packets-exported"`
	InlineFlowInsertAged       int64  `xml:"inline-flow-insert-aged"`
	InlineFlowInsertExpired    int64  `xml:"inline-flow-insert-expired"`
	InlineFlowInsertCount      int64  `xml:"inline-flow-insert-count"`

	InlineIpv4FlowPackets         int64 `xml:"inline-ipv4-flow-packets"`
	InlineIpv4FlowBytes           int64 `xml:"inline-ipv4-flow-bytes"`
	InlineIpv4ActiveFlows         int64 `xml:"inline-ipv4-active-flows"`
	InlineIpv4TotalFlows          int64 `xml:"inline-ipv4-total-flows"`
	InlineIpv4FlowsExported       int64 `xml:"inline-ipv4-flows-exported"`
	InlineIpv4FlowPacketsExported int64 `xml:"inline-ipv4-flow-packets-exported"`
	InlineIpv4FlowsAged           int64 `xml:"inline-ipv4-flows-aged"`
	InlineIpv4FlowsExpired        int64 `xml:"inline-ipv4-flows-expired"`
	InlineIpv4FlowInsertCount     int64 `xml:"inline-ipv4-flow-insert-count"`

	InlineIpv6FlowPackets         int64 `xml:"inline-ipv6-flow-packets"`
	InlineIpv6FlowBytes           int64 `xml:"inline-ipv6-flow-bytes"`
	InlineIpv6ActiveFlows         int64 `xml:"inline-ipv6-active-flows"`
	InlineIpv6TotalFlows          int64 `xml:"inline-ipv6-total-flows"`
	InlineIpv6FlowsExported       int64 `xml:"inline-ipv6-flows-exported"`
	InlineIpv6FlowPacketsExported int64 `xml:"inline-ipv6-flow-packets-exported"`
	InlineIpv6FlowsAged           int64 `xml:"inline-ipv6-flows-aged"`
	InlineIpv6FlowsExpired        int64 `xml:"inline-ipv6-flows-expired"`
	InlineIpv6FlowInsertCount     int64 `xml:"inline-ipv6-flow-insert-count"`
}

type AccountingFlowErrorRpc struct {
	Information struct {
		InlineFlow AccountingInlineFlowError `xml:"inline-jflow-error-information"`
	} `xml:"services-accounting-information"`
}

type AccountingInlineFlowError struct {
	FpcSlot                        string `xml:"fpc-slot"`
	InlineFlowCreationFailures     int64  `xml:"inline-flow-creation-failures"`
	InlineRouteRecordLookupFailure int64  `xml:"inline-route-record-lookup-failure"`
	InlineAsLookupFailures         int64  `xml:"inline-as-lookup-failures"`
	InlineExportPacketFailures     int64  `xml:"inline-export-packet-failures"`

	InlineIpv4FlowCreationFailures     int64 `xml:"inline-fipv4-low-creation-failures"`
	InlineIpv4RouteRecordLookupFailure int64 `xml:"inline-ipv4-route-record-lookup-failure"`
	InlineIpv4AsLookupFailures         int64 `xml:"inline-ipv4-as-lookup-failures"`
	InlineIpv4ExportPacketFailures     int64 `xml:"inline-ipv4-export-packet-failures"`

	InlineIpv6FlowCreationFailures     int64 `xml:"inline-ipv6-flow-creation-failures"`
	InlineIpv6RouteRecordLookupFailure int64 `xml:"inline-ipv6-route-record-lookup-failure"`
	InlineIpv6AsLookupFailures         int64 `xml:"inline-ipv6-as-lookup-failures"`
	InlineIpv6ExportPacketFailures     int64 `xml:"inline-ipv6-export-packet-failures"`
}
