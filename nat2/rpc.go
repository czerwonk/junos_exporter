package nat2

type NatRpc struct {
	Interfaces []NatInterface `xml:"service-nat-statistics-information"`
}

type NatInterface struct {
	Interface                                  string `xml:"interface-name"`
	NatPktDstInNatRoute                        int64  `xml:"nat-pkt-dst-in-nat-route"`
	NatFilteringSession                        int64  `xml:"nat-filtering-session"`
	NatMappingSession                          int64  `xml:"nat-mapping-session"`
	NatRuleLookupFailures                      int64  `xml:"nat-rule-lookup-failures"`
	NatMapAllocationSuccesses                  int64  `xml:"nat-map-allocation-successes"`
	NatMapAllocationFailures                   int64  `xml:"nat-map-allocation-failures"`
	NatMapFreeSuccess                          int64  `xml:"nat-map-free-success"`
	NatMapFreeFailures                         int64  `xml:"nat-map-free-failures"`
	NatEimMappingCreateFailed                  int64  `xml:"nat-eim-mapping-create-failed"`
	NatEimMappingCreated                       int64  `xml:"nat-eim-mapping-created"`
	NatEimMappingUpdated                       int64  `xml:"nat-eim-mapping-updated"`
	NatEifMappingFree                          int64  `xml:"nat-eif-mapping-free"`
	NatEimMappingFree                          int64  `xml:"nat-eim-mapping-free"`
	NatTotalPktsProcessed                      int64  `xml:"nat-total-pkts-processed"`
	NatTotalPktsForwarded                      int64  `xml:"nat-total-pkts-forwarded"`
	NatTotalPktsTranslated                     int64  `xml:"nat-total-pkts-translated"`
	Nat64MtuExceed                             int64  `xml:"nat64-mtu-exceed"`
	Nat64DfbitSet                              int64  `xml:"nat64-dfbit-set"`
	Nat64ErrMtuExceedBuild                     int64  `xml:"nat64-err-mtu-exceed-build"`
	Nat64ErrMtuExceedSend                      int64  `xml:"nat64-err-mtu-exceed-send"`
	SessionXlate464ClatPrefixNotFound          int64  `xml:"session-xlate464-clat-prefix-not-found"`
	SessionXlate464EmbededIpv4NotFound         int64  `xml:"session-xlate464-embeded-ipv4-not-found"`
	NatJflowLogAllocFail                       int64  `xml:"nat-jflow-log-alloc-fail"`
	NatJflowLogAllocSuccess                    int64  `xml:"nat-jflow-log-alloc-success"`
	NatJflowLogFreeSuccess                     int64  `xml:"nat-jflow-log-free-success"`
	NatJflowLogFreeFailRecord                  int64  `xml:"nat-jflow-log-free-fail-record"`
	NatJflowLogFreeFailData                    int64  `xml:"nat-jflow-log-free-fail-data"`
	NatJflowLogInvalidTransType                int64  `xml:"nat-jflow_log-invalid-trans-type"` //Yes Actual bug in the junosxml having an underscore: Junos: 21.2R2.13
	NatJflowLogFreeSuccessFailQueuing          int64  `xml:"nat-jflow-log-free-success-fail-queuing"`
	NatJflowLogInvalidInputArgs                int64  `xml:"nat-jflow-log-invalid-input-args"`
	NatJflowLogInvalidAllocErr                 int64  `xml:"nat-jflow-log-invalid-alloc-err"`
	NatJflowLogRateLimitFailGetPool            int64  `xml:"nat-jflow-log-rate-limit-fail-get-pool"`
	NatJflowLogRateLimitFailGetServiceSet      int64  `xml:"nat-jflow-log-rate-limit-fail-get-service-set"`
	NatJflowLogRateLimitFailInvalidCurrentTime int64  `xml:"nat-jflow-log-rate-limit-fail-invalid-current-time"`
}

type SrcNatPoolRpc struct {
	Information struct {
		TotalSourcePools string       `xml:"total-source-pools"`
		Pools            []SrcNatPool `xml:"source-nat-pool-info-entry"`
	} `xml:"source-nat-pool-detail-information"`
}

type SrcNatPool struct {
	Interface               string `xml:"interface-name"`
	ServiceSetName          string `xml:"service-set-name"`
	PoolName                string `xml:"pool-name"`
	PoolId                  string `xml:"pool-id"`
	PortTranslation         string `xml:"source-pool-port-translation"`
	PortOverloadingFactor   int64  `xml:"port-overloading-factor"`
	AddressAssignement      string `xml:"source-pool-address-assignment"`
	ClearAlarmThreshold     string `xml:"clear-alarm-threshold"`
	RaiseAlarmThreshold     string `xml:"raise-alarm-threshold"`
	TotalPoolAddress        int64  `xml:"total-pool-address"`
	AddressPoolHits         int64  `xml:"address-pool-hits"`
	BlkSize                 int64  `xml:"source-pool-blk-size"`
	BlkMaxPerHost           int64  `xml:"source-pool-blk-max-per-host"`
	BlkAtvTimeout           int64  `xml:"source-pool-blk-atv-timeout"`
	BlkInterimLogCycle      int64  `xml:"source-pool-blk-interim-log-cycle"`
	BlkLog                  string `xml:"source-pool-blk-log"`
	BlkUsed                 int64  `xml:"source-pool-blk-used"`
	BlkTotal                int64  `xml:"source-pool-blk-total"`
	PortBlkEfficiency       string `xml:"source-pool-port-blk-efficiency"`
	MaxBlkUsed              int64  `xml:"source-pool-max-blk-used"`
	Users                   int64  `xml:"source-pool-users"`
	EimTimeout              int64  `xml:"source-pool-eim-timeout"`
	MappingTimeout          int64  `xml:"source-pool-mapping-timeout"`
	EifInboundFlowsCount    int64  `xml:"source-pool-eif-inbound-flows-count"`
	EifFlowLimitExceedDrops int64  `xml:"source-pool-eif-flow-limit-exceed-drops"`

	//     []SrcPoolAddressRange `xml:"source-pool-address-range"` //Annoying Junos implementation: same xml 'categories' multiple times
	SrcPoolAddressRanges struct {
		AddressRangeLow  []string `xml:"address-range-low"`
		AddressRangeHigh []string `xml:"address-range-high"`
		SinglePort       []int64  `xml:"single-port"`
	} `xml:"source-pool-address-range"`

	SrcPoolAddressRangeSum struct {
		SinglePortSum int64 `xml:"single-port-sum"`
	} `xml:"source-pool-address-range-sum"`
	SrcPoolErrorCounters struct {
		OutOfPortError          int64 `xml:"out-of-port-error"`
		OutOfAddrError          int64 `xml:"out-of-addr-error"`
		ParityPortError         int64 `xml:"parity-port-error"`
		PreserveRangeError      int64 `xml:"preserve-range-error"`
		AppOutOfPortError       int64 `xml:"app-out-of-port-error"`
		AppExceedPortLimitError int64 `xml:"app-exceed-port-limit-error"`
		OutOfBlkError           int64 `xml:"out-of-blk-error"`
		BlkExceedLimitError     int64 `xml:"blk-exceed-limit-error"`
		BlkOutOfPortError       int64 `xml:"blk-out-of-port-error"`
		BlkMemAllocError        int64 `xml:"blk-mem-alloc-error"`
	} `xml:"source-pool-error-counters"`
}

type ServiceSetsCpuRpc struct {
	Information struct {
		Interfaces []ServiceSetsCpuInterface `xml:"service-set-cpu-statistics"`
	} `xml:"service-set-cpu-statistics-information"`
}

type ServiceSetsCpuInterface struct {
	Interface             string  `xml:"interface-name"`
	ServiceSetName        string  `xml:"service-set-name"`
	CpuUtilizationPercent float64 `xml:"cpu-utilization-percent"`
}
