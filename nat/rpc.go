package nat

type NatRpc struct {
	Interfaces []NatInterface `xml:"service-nat-statistics-information"`
}

type NatInterface struct {
	Interface                                  string `xml:"interface-name"`
	NatTotalSessionInterest                    int64  `xml:"nat-total-session-interest"`
	NatTotalSessionCreate                      int64  `xml:"nat-total-session-create"`
	NatTotalSessionDestroy                     int64  `xml:"nat-total-session-destroy"`
	NatTotalSessionPubReq                      int64  `xml:"nat-total-session-pub-req"`
	NatTotalSessionAccepts                     int64  `xml:"nat-total-session-accepts"`
	NatTotalSessionDiscards                    int64  `xml:"nat-total-session-discards"`
	NatTotalSessionIgnores                     int64  `xml:"nat-total-session-ignores"`
	NatCtrlSessNotXltdChldSessIgnd             int64  `xml:"nat-ctrl-sess-not-xltd-chld-sess-ignd"`
	NatTotalSessionTimeEvent                   int64  `xml:"nat-total-session-time-event"`
	NatSessionInterestPubReq                   int64  `xml:"nat-session-interest-pub-req"`
	NatAlgDataSessionInterest                  int64  `xml:"nat-alg-data-session-interest"`
	NatAlgDataSessionCreated                   int64  `xml:"nat-alg-data-session-created"`
	NatTotalSessionClose                       int64  `xml:"total-session-close"`
	NatPktDstInNatRoute                        int64  `xml:"nat-pkt-dst-in-nat-route"`
	NatPktDropInBackupState                    int64  `xml:"nat-pkt-drop-in-backup-state"`
	NatSessionExtAllocFailures                 int64  `xml:"nat-session-ext-alloc-failures"`
	NatSessionExtSetFailures                   int64  `xml:"nat-session-ext-set-failures"`
	NatFilteringSession                        int64  `xml:"nat-filtering-session"`
	NatMappingSession                          int64  `xml:"nat-mapping-session"`
	NatRuleLookupFailures                      int64  `xml:"nat-rule-lookup-failures"`
	NatNoSextInXlatePkt                        int64  `xml:"no-sext-in-xlate-pkt"`
	NatPoolSessionCntUpdateFailOnCreate        int64  `xml:"pool-session-cnt-update-fail-on-create"`
	NatPoolSessionCntUpdateFailOnClose         int64  `xml:"pool-session-cnt-update-fail-on-close"`
	NatMapAllocationSuccesses                  int64  `xml:"nat-map-allocation-successes"`
	NatMapAllocationFailures                   int64  `xml:"nat-map-allocation-failures"`
	NatMapFreeSuccess                          int64  `xml:"nat-map-free-success"`
	NatMapFreeFailures                         int64  `xml:"nat-map-free-failures"`
	NatFreeFailOnInactiveSset                  int64  `xml:"nat_free_fail_on_inactive_sset"`
	NatEimMappingReused                        int64  `xml:"nat-eim-mapping-reused"`
	NatEimMappingAllocFailures                 int64  `xml:"nat-eim-mapping-alloc-failures"`
	NatEimMismatchedMapping                    int64  `xml:"nat-eim-mismatched-mapping"`
	NatEimDuplicateMapping                     int64  `xml:"nat-eim-duplicate-mapping"`
	NatEimMappingCreateFailed                  int64  `xml:"nat-eim-mapping-create-failed"`
	NatEimMappingCreated                       int64  `xml:"nat-eim-mapping-created"`
	NatEimMappingUpdated                       int64  `xml:"nat-eim-mapping-updated"`
	NatEifMappingFree                          int64  `xml:"nat-eif-mapping-free"`
	NatEimMappingFree                          int64  `xml:"nat-eim-mapping-free"`
	NatEimWaitingForInit                       int64  `xml:"nat-eim-waiting-for-init"`
	NatEimWaitingForInitFailed                 int64  `xml:"nat-eim-waiting-for-init-failed"`
	NatEimLookupHoldSuccess                    int64  `xml:"nat-eim-lookup-hold-success"`
	NatEimLookupTimeout                        int64  `xml:"nat-eim-lookup-timeout"`
	NatEimLookupClearTimer                     int64  `xml:"nat-eim-lookup-clear-timer"`
	NatEimDrainInLookup                        int64  `xml:"nat-eim-drain-in-lookup"`
	NatEimLookupEntryWithoutTimer              int64  `xml:"nat-eim-lookup-entry-without-timer"`
	NatEimReleaseWithoutEntry                  int64  `xml:"nat-eim-release-without-entry"`
	NatEimReleaseInTimeout                     int64  `xml:"nat-eim-release-in-timeout"`
	NatEimReleaseRace                          int64  `xml:"nat-eim-release-race"`
	NatEimReleaseSetTimeout                    int64  `xml:"nat-eim-release-set-timeout"`
	NatEimTimerEntryRefreshed                  int64  `xml:"nat-eim-timer-entry-refreshed"`
	NatEimTimerStartInvalid                    int64  `xml:"nat-eim-timer-start-invalid"`
	NatEimTimerStartInvalidFail                int64  `xml:"nat-eim-timer-start-invalid-fail"`
	NatEimTimerFreeMapping                     int64  `xml:"nat-eim-timer-free-mapping"`
	NatEimTimerUpdateTimeout                   int64  `xml:"nat-eim-timer-update-timeout"`
	NatEimEntryDrained                         int64  `xml:"nat-eim-entry-drained"`
	NatTotalPktsProcessed                      int64  `xml:"nat-total-pkts-processed"`
	NatTotalBytesProcessed                     int64  `xml:"nat-total-bytes-processed"`
	NatTotalPktsForwarded                      int64  `xml:"nat-total-pkts-forwarded"`
	NatTotalPktsDiscarded                      int64  `xml:"nat-total-pkts-discarded"`
	NatTotalPktsTranslated                     int64  `xml:"nat-total-pkts-translated"`
	NatTotalPktsRestored                       int64  `xml:"nat-total-pkts-restored"`
	NatCmSessLnodeCreated                      int64  `xml:"nat-cm-sess-lnode-created"`
	NatCmSessLnodeCeleted                      int64  `xml:"nat-cm-sess-lnode-deleted"`
	NatCmEimLnodeCreated                       int64  `xml:"nat-cm-eim-lnode-created"`
	NatCmEimLnodeCeleted                       int64  `xml:"nat-cm-eim-lnode-deleted"`
	NatSrcIpv4Translations                     int64  `xml:"nat-src-ipv4-translations"`
	NatSrcIpv4Restorations                     int64  `xml:"nat-src-ipv4-restorations"`
	NatDstIpv4Translations                     int64  `xml:"nat-dst-ipv4-translations"`
	NatDstIpv4Restorations                     int64  `xml:"nat-dst-ipv4-restorations"`
	NatSrcIpv6Translations                     int64  `xml:"nat-src-ipv6-translations"`
	NatSrcIpv6Restorations                     int64  `xml:"nat-src-ipv6-restorations"`
	NatDstIpv6Translations                     int64  `xml:"nat-dst-ipv6-translations"`
	NatDstIpv6Restorations                     int64  `xml:"nat-dst-ipv6-restorations"`
	NatSrcPortTranslations                     int64  `xml:"nat-src-port-translations"`
	NatSrcPortRestorations                     int64  `xml:"nat-src-port-restorations"`
	NatDstPortTranslations                     int64  `xml:"nat-dst-port-translations"`
	NatDstPortRestorations                     int64  `xml:"nat-dst-port-restorations"`
	NatIcmpIdTranslations                      int64  `xml:"nat-icmp-id-translations"`
	NatOcmpIdRestorations                      int64  `xml:"nat-icmp-id-restorations"`
	NatIcmpErrorTranslations                   int64  `xml:"nat-icmp-error-translations"`
	NatIcmpDrop                                int64  `xml:"nat-icmp-drop"`
	NatIcmpAllocationFailure                   int64  `xml:"nat-icmp-allocation-failure"`
	NatRuleLookupForIcmpErrFail                int64  `xml:"nat-rule-lookup-for-icmp-err-fail"`
	NatTcpPortTranslations                     int64  `xml:"nat-tcp-port-translations"`
	NatTcpPortRestorations                     int64  `xml:"nat-tcp-port-restorations"`
	NatUdpPortTranslations                     int64  `xml:"nat-udp-port-translations"`
	NatUdpPortRestorations                     int64  `xml:"nat-udp-port-restorations"`
	NatUnexpectedProtoWithPortXlation          int64  `xml:"nat-unexpected-proto-with-port-xlation"`
	NatGreCallIdTranslations                   int64  `xml:"nat-gre-call-id-translations"`
	NatGreCallIdRestorations                   int64  `xml:"nat-gre-call-id-restorations"`
	NatUnsupportedGreProto                     int64  `xml:"nat-unsupported-gre-proto"`
	NatIcmpErrorSrcRestored                    int64  `xml:"nat-icmp-error-src-restored"`
	NatIcmpErrorDstRestored                    int64  `xml:"nat-icmp-error-dst-restored"`
	NatIcmpErrorSrcXlated                      int64  `xml:"nat-icmp-error-src-xlated"`
	NatIcmpErrorDstXlated                      int64  `xml:"nat-icmp-error-dst-xlated"`
	NatIcmpErrorNewSrcXlated                   int64  `xml:"nat-icmp-error-new-src-xlated"`
	NatIcmpErrorOrgIpSrcRestored               int64  `xml:"nat-icmp-error-org-ip-src-restored"`
	NatIcmpErrorOrgIpSrcPortRestored           int64  `xml:"nat-icmp-error-org-ip-src-port-restored"`
	NatIcmpErrorOrgIpDstPortRestored           int64  `xml:"nat-icmp-error-org-ip-dst-port-restored"`
	NatIcmpErrorOrgIpDstRestored               int64  `xml:"nat-icmp-error-org-ip-dst-restored"`
	NatIcmpErrorOrgIpSrcXlated                 int64  `xml:"nat-icmp-error-org-ip-src-xlated"`
	NatIcmpErrorOrgIpSrcPortXlated             int64  `xml:"nat-icmp-error-org-ip-src-port-xlated"`
	NatIcmpErrorOrgIpDstPortXlated             int64  `xml:"nat-icmp-error-org-ip-dst-port-xlated"`
	NatIcmpErrorOrgIpDstXlated                 int64  `xml:"nat-icmp-error-org-ip-dst-xlated"`
	NatErrorNoPolicy                           int64  `xml:"nat-error-no-policy"`
	NatErrorIpVersion                          int64  `xml:"nat-error-ip-version"`
	NatXlateFreeNullExt                        int64  `xml:"nat-xlate-free-null-ext"`
	NatSessionExtFreeFailed                    int64  `xml:"nat-session-ext-free-failed"`
	NatPolicyAddFailed                         int64  `xml:"nat-policy-add-failed"`
	NatPolicyDeleteFailed                      int64  `xml:"nat-policy-delete-failed"`
	NatPrefixFilterAllocFailed                 int64  `xml:"nat-prefix-filter-alloc-failed"`
	NatPrefixFilterNameFailed                  int64  `xml:"nat-prefix-filter-name-failed"`
	NatPrefixListCreateFailed                  int64  `xml:"nat-prefix-list-create-failed"`
	NatPrefixFilterTreeAddFailed               int64  `xml:"nat-prefix-filter-tree-add-failed"`
	NatPrefixFilterCreated                     int64  `xml:"nat-prefix-filter-created"`
	NatPrefixFilterChanged                     int64  `xml:"nat-prefix-filter-changed"`
	NatPrefixFilterCtrlFree                    int64  `xml:"nat-prefix-filter-ctrl-free"`
	NatPrefixFilterMatch                       int64  `xml:"nat-prefix-filter-match"`
	NatPrefixFilterNoMatch                     int64  `xml:"nat-prefix-filter-no-match"`
	NatPrefixFilterMappingAdd                  int64  `xml:"nat-prefix-filter-mapping-add"`
	NatPrefixFilterMappingRemove               int64  `xml:"nat-prefix-filter-mapping-remove"`
	NatPrefixFilterMappingFree                 int64  `xml:"nat-prefix-filter-mapping-free"`
	NatPrefixFilterUnsuppIpVersion             int64  `xml:"nat-prefix-filter-unsupp-ip-version"`
	NatunsupportedLayer4Napt                   int64  `xml:"nat-unsupported-layer-4-napt"`
	NatunsupportedIcmpTypeNapt                 int64  `xml:"nat-unsupported-icmp-type-napt"`
	Nat64IpoptionsDrop                         int64  `xml:"nat64-ipoptions-drop"`
	Nat64UdpCksumZeroDrop                      int64  `xml:"nat64-udp-cksum-zero-drop"`
	Nat64UnsuppIcmpTypeDrop                    int64  `xml:"nat64-unsupp-icmp-type-drop"`
	Nat64UnsuppIcmpCodeDrop                    int64  `xml:"nat64-unsupp-icmp-code-drop"`
	Nat64UnsuppHdrDrop                         int64  `xml:"nat64-unsupp-hdr-drop"`
	Nat64UnsuppL4Drop                          int64  `xml:"nat64-unsupp-l4-drop"`
	Nat64MtuExceed                             int64  `xml:"nat64-mtu-exceed"`
	Nat64DfbitSet                              int64  `xml:"nat64-dfbit-set"`
	Nat64UnsuppIcmpError                       int64  `xml:"nat64-unsupp-icmp-error"`
	Nat64ErrMapSrc                             int64  `xml:"nat64-err-map-src"`
	Nat64ErrMapDst                             int64  `xml:"nat64-err-map-dst"`
	Nat64ErrMtuExceedBuild                     int64  `xml:"nat64-err-mtu-exceed-build"`
	Nat64ErrTtlExceedBuild                     int64  `xml:"nat64-err-ttl-exceed-build"`
	Nat64ErrMtuExceedSend                      int64  `xml:"nat64-err-mtu-exceed-send"`
	Nat64ErrTtlExceedSend                      int64  `xml:"nat64-err-ttl-exceed-send"`
	NatSubsExtAlloc                            int64  `xml:"nat-subs-ext-alloc"`
	NatSubsExtInvalidParam                     int64  `xml:"nat-subs-ext-invalid-param"`
	NatSubsExtNoMem                            int64  `xml:"nat-subs-ext-no-mem"`
	NatSubsExtfree                             int64  `xml:"nat-subs-ext-free"`
	NatSubsExtIsNull                           int64  `xml:"nat-subs-ext-is-null"`
	NatSubsExtIsInvalid                        int64  `xml:"nat-subs-ext-is-invalid"`
	NatSubsExtLinkSuccess                      int64  `xml:"nat-subs-ext-link-success"`
	NatSubsExtLinkExist                        int64  `xml:"nat-subs-ext-link-exist"`
	NatSubsExtLinkFail                         int64  `xml:"nat-subs-ext-link-fail"`
	NatSubsExtLinkUnknownRet                   int64  `xml:"nat-subs-ext-link-unknown-ret"`
	NatSubsExtInlinkSuccess                    int64  `xml:"nat-subs-ext-unlink-success"`
	NatSubsExtUnlinkFail                       int64  `xml:"nat-subs-ext-unlink-fail"`
	NatSubsExtUnlinkBusy                       int64  `xml:"nat-subs-ext-unlink-busy"`
	NatSubsExtResourceInUse                    int64  `xml:"nat-subs-ext-resource-in-use"`
	NatSubsExtSvcSetNotActive                  int64  `xml:"nat-subs-ext-svc-set-not-active"`
	NatSubsExtSvcSetIsNull                     int64  `xml:"nat-subs-ext-svc-set-is-null"`
	NatSubsExtTimerStartSuccess                int64  `xml:"nat-subs-ext-timer-start-success"`
	NatSubsExtTimerStartFail                   int64  `xml:"nat-subs-ext-timer-start-fail"`
	NatSubsExtDelayTimerSuccess                int64  `xml:"nat-subs-ext-delay-timer-success"`
	NatSubsExtDelayTimerFail                   int64  `xml:"nat-subs-ext-delay-timer-fail"`
	NatSubsExtReuseFromTimer                   int64  `xml:"nat-subs-ext-reuse-from-timer"`
	NatSubsExtTimerCb                          int64  `xml:"nat-subs-ext-timer-cb"`
	NatSubsExtRefcountDecFail                  int64  `xml:"nat-subs-ext-refcount-dec-fail"`
	NatSubsExtSubsResetFail                    int64  `xml:"nat-subs-ext-subs-reset-fail"`
	NatSubsExtSubsSessionCountUpdateIgnore     int64  `xml:"nat-subs-ext-subs-session-count-update-ignore"`
	NatSubsExtIncorrectState                   int64  `xml:"nat-subs-ext-incorrect-state"`
	NatSubsExtUnlinkUnkErr                     int64  `xml:"nat-subs-ext-unlink-unk-err"`
	NatSubsExtQueueInconsistent                int64  `xml:"nat-subs-ext-queue-inconsistent"`
	NatSubsExtReturnToPreallocErr              int64  `xml:"nat-subs-ext-return-to-prealloc-err"`
	NatSubsExtDecInvalSessCnt                  int64  `xml:"nat-subs-ext-dec-inval-sess-cnt"`
	NatSubsExtDecInvalEimCnt                   int64  `xml:"nat-subs-ext-dec-inval-eim-cnt"`
	NatSubsExtPortsInUseErr                    int64  `xml:"nat-subs-ext-ports-in-use-err"`
	NatSubsExtErrSetState                      int64  `xml:"nat-subs-ext-err-set-state"`
	NatSubsExtMissingExt                       int64  `xml:"nat-subs-ext-missing-ext"`
	NatSubsExtInvalidEimrefcnt                 int64  `xml:"nat-subs-ext-invalid-eimrefcnt"`
	NatSubsExtIsInvalidSubsInTw                int64  `xml:"nat-subs-ext-is-invalid-subs-in-tw"`
	NatSubsExtInTwDuringFree                   int64  `xml:"nat-subs-ext-in-tw-during-free"`
	NatJflowLogNatSextNull                     int64  `xml:"nat_jflow_log_nat_sext_null"`
	NatJflowLogAllocFail                       int64  `xml:"nat_jflow_log_alloc_fail"`
	NatJflowLogAllocSuccess                    int64  `xml:"nat_jflow_log_alloc_success"`
	NatJflowLogFreeSuccess                     int64  `xml:"nat_jflow_log_free_success"`
	NatJflowLogFreeFailRecord                  int64  `xml:"nat_jflow_log_free_fail_record"`
	NatJflowLogFreeFailData                    int64  `xml:"nat_jflow_log_free_fail_data"`
	NatJflowLogInvalidTransType                int64  `xml:"nat_jflow_log_invalid_trans_type"`
	NatJflowLogFreeSuccessFailQueuing          int64  `xml:"nat_jflow_log_free_success_fail_queuing"`
	NatJflowLogInvalidInputArgs                int64  `xml:"nat_jflow_log_invalid_input_args"`
	NatJflowLogInvalidAllocErr                 int64  `xml:"nat_jflow_log_invalid_alloc_err"`
	NatJflowLogRateLimitFailGetPoolName        int64  `xml:"nat_jflow_log_rate_limit_fail_get_pool_name"`
	NatJflowLogRateLimitFailGetNatpool         int64  `xml:"nat_jflow_log_rate_limit_fail_get_natpool"`
	NatJflowLogRateLimitFailGetNatpoolGivenId  int64  `xml:"nat_jflow_log_rate_limit_fail_get_natpool_given_id"`
	NatJflowLogRateLimitFailGetServiceSet      int64  `xml:"nat_jflow_log_rate_limit_fail_get_service_set"`
	NatJflowLogRateLimitFailInvalidCurrentTime int64  `xml:"nat_jflow_log_rate_limit_fail_invalid_current_time"`
	NatEimMappingEifCurrSessUpdateInvalid      int64  `xml:"nat_eim_mapping_eif_curr_sess_update_invalid"`
	NatEimMappingCreatedWithoutEifSessLimit    int64  `xml:"nat_eim_mapping_created_without_eif_sess_limit"`
}

type NatPoolRpc struct {
	Information struct {
		Interfaces []NatPoolInterface `xml:"sfw-per-service-set-nat-pool"`
	} `xml:"service-nat-pool-information"`
}

type NatPoolInterface struct {
	Interface       string           `xml:"interface-name"`
	ServiceSetName  string           `xml:"service-set-name"`
	ServiceNatPools []ServiceNatPool `xml:"service-nat-pool"`
}

type ServiceNatPool struct {
	Name                string  `xml:"pool-name"`
	TranslationType     string  `xml:"description"`
	PortRange           string  `xml:"pool-port-range"`
	PortBlockType       string  `xml:"port-block-type"`
	PortBlockSize       int64   `xml:"port-block-size"`
	ActiveBlockTimeout  int64   `xml:"active-block-timeout"`
	MaxBlocksPerAddress int64   `xml:"max-blocks-per-address"`
	EffectivePortBlocks int64   `xml:"effective-port-blocks"`
	EffectivePorts      int64   `xml:"effective-ports"`
	PortBlockEfficiency float64 `xml:"port-block-efficiency"`
}

type NatPoolDetailRpc struct {
	Information struct {
		Interfaces []NatPoolDetailInterface `xml:"sfw-per-service-set-nat-pool"`
	} `xml:"service-nat-pool-information"`
}

type NatPoolDetailInterface struct {
	Interface             string                 `xml:"interface-name"`
	ServiceSetName        string                 `xml:"service-set-name"`
	ServiceNatPoolsDetail []ServiceNatPoolDetail `xml:"service-nat-pool"`
}

type ServiceNatPoolDetail struct {
	Name                      string `xml:"pool-name"`
	TranslationType           string `xml:"description"`
	PortRange                 string `xml:"pool-port-range"`
	PortsInUse                int64  `xml:"pool-ports-in-use"`
	OutOfPortErrors           int64  `xml:"pool-out-of-port-errors"`
	ParityPortErrors          int64  `xml:"pool-parity-port-errors"`
	PreserveRangeErrors       int64  `xml:"pool-preserve-range-errors"`
	MaxPortsInUse             int64  `xml:"pool-max-ports-in-use"`
	AppPortErrors             int64  `xml:"pool-app-port-errors"`
	AppExceedPortLimitErrors  int64  `xml:"pool-app-exceed-port-limit-errors"`
	MemAllocErrors            int64  `xml:"pool-mem-alloc-errors"`
	PortBlockType             string `xml:"port-block-type"`
	MaxPortBlocksUsed         int64  `xml:"max-port-blocks-used"`
	BlocksInUse               int64  `xml:"port-blocks-in-use"`
	BlockAllocationErrors     int64  `xml:"port-block-allocation-errors"`
	BlocksLimitExceededErrors int64  `xml:"port-blocks-limit-exceeded-errors"`
	Users                     int64  `xml:"pool-users"`
	EifInboundSessionCount    int64  `xml:"eif-inbound-session-count"`
	EifInboundLimitExceedDrop int64  `xml:"eif-inbound-session-limit-exceed-drop"`
}


type ServiceSetsCpuRpc struct {
	Information struct {
		Interfaces []ServiceSetsCpuInterface `xml:"service-set-cpu-statistics"`
	} `xml:"service-set-cpu-statistics-information"`
}

type ServiceSetsCpuInterface struct {
	Interface             string                 `xml:"interface-name"`
	ServiceSetName        string                 `xml:"service-set-name"`
	CpuUtilizationPercent float64 				 `xml:"cpu-utilization-percent"`
}
