// SPDX-License-Identifier: MIT

package accounting

type result struct {
	Information struct {
		InlineFlow accountingInlineFlow `xml:"inline-jflow-flow-information"`
	} `xml:"services-accounting-information"`
	Error struct {
		Message string `xml:"message"`
	} `xml:"error,omitempty"`
}

type accountingInlineFlow struct {
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

	InlineIPv4FlowPackets         int64 `xml:"inline-ipv4-flow-packets"`
	InlineIPv4FlowBytes           int64 `xml:"inline-ipv4-flow-bytes"`
	InlineIPv4ActiveFlows         int64 `xml:"inline-ipv4-active-flows"`
	InlineIPv4TotalFlows          int64 `xml:"inline-ipv4-total-flows"`
	InlineIPv4FlowsExported       int64 `xml:"inline-ipv4-flows-exported"`
	InlineIPv4FlowPacketsExported int64 `xml:"inline-ipv4-flow-packets-exported"`
	InlineIPv4FlowsAged           int64 `xml:"inline-ipv4-flows-aged"`
	InlineIPv4FlowsExpired        int64 `xml:"inline-ipv4-flows-expired"`
	InlineIPv4FlowInsertCount     int64 `xml:"inline-ipv4-flow-insert-count"`

	InlineIPv6FlowPackets         int64 `xml:"inline-ipv6-flow-packets"`
	InlineIPv6FlowBytes           int64 `xml:"inline-ipv6-flow-bytes"`
	InlineIPv6ActiveFlows         int64 `xml:"inline-ipv6-active-flows"`
	InlineIPv6TotalFlows          int64 `xml:"inline-ipv6-total-flows"`
	InlineIPv6FlowsExported       int64 `xml:"inline-ipv6-flows-exported"`
	InlineIPv6FlowPacketsExported int64 `xml:"inline-ipv6-flow-packets-exported"`
	InlineIPv6FlowsAged           int64 `xml:"inline-ipv6-flows-aged"`
	InlineIPv6FlowsExpired        int64 `xml:"inline-ipv6-flows-expired"`
	InlineIPv6FlowInsertCount     int64 `xml:"inline-ipv6-flow-insert-count"`
}

type accountingFlowError struct {
	Information struct {
		InlineFlow accountingInlineFlowError `xml:"inline-jflow-error-information"`
	} `xml:"services-accounting-information"`
}

type accountingInlineFlowError struct {
	FpcSlot                        string `xml:"fpc-slot"`
	InlineFlowCreationFailures     int64  `xml:"inline-flow-creation-failures"`
	InlineRouteRecordLookupFailure int64  `xml:"inline-route-record-lookup-failure"`
	InlineAsLookupFailures         int64  `xml:"inline-as-lookup-failures"`
	InlineExportPacketFailures     int64  `xml:"inline-export-packet-failures"`

	InlineIPv4FlowCreationFailures     int64 `xml:"inline-fipv4-low-creation-failures"`
	InlineIPv4RouteRecordLookupFailure int64 `xml:"inline-ipv4-route-record-lookup-failure"`
	InlineIPv4AsLookupFailures         int64 `xml:"inline-ipv4-as-lookup-failures"`
	InlineIPv4ExportPacketFailures     int64 `xml:"inline-ipv4-export-packet-failures"`

	InlineIPv6FlowCreationFailures     int64 `xml:"inline-ipv6-flow-creation-failures"`
	InlineIPv6RouteRecordLookupFailure int64 `xml:"inline-ipv6-route-record-lookup-failure"`
	InlineIPv6AsLookupFailures         int64 `xml:"inline-ipv6-as-lookup-failures"`
	InlineIPv6ExportPacketFailures     int64 `xml:"inline-ipv6-export-packet-failures"`
}
