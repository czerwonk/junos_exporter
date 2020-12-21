package nat

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_nat_statistics_"

var (
	nat64DfbitSetDesc                              *prometheus.Desc
	nat64ErrMapDstDesc                             *prometheus.Desc
	nat64ErrMapSrcDesc                             *prometheus.Desc
	nat64ErrMtuExceedBuildDesc                     *prometheus.Desc
	nat64ErrMtuExceedSendDesc                      *prometheus.Desc
	nat64ErrTtlExceedBuildDesc                     *prometheus.Desc
	nat64ErrTtlExceedSendDesc                      *prometheus.Desc
	nat64IpoptionsDropDesc                         *prometheus.Desc
	nat64MtuExceedDesc                             *prometheus.Desc
	nat64UdpCksumZeroDropDesc                      *prometheus.Desc
	nat64UnsuppHdrDropDesc                         *prometheus.Desc
	nat64UnsuppIcmpCodeDropDesc                    *prometheus.Desc
	nat64UnsuppIcmpErrorDesc                       *prometheus.Desc
	nat64UnsuppIcmpTypeDropDesc                    *prometheus.Desc
	nat64UnsuppL4DropDesc                          *prometheus.Desc
	natAlgDataSessionCreatedDesc                   *prometheus.Desc
	natAlgDataSessionInterestDesc                  *prometheus.Desc
	natCmEimLnodeCeletedDesc                       *prometheus.Desc
	natCmEimLnodeCreatedDesc                       *prometheus.Desc
	natCmSessLnodeCeletedDesc                      *prometheus.Desc
	natCmSessLnodeCreatedDesc                      *prometheus.Desc
	natCtrlSessNotXltdChldSessIgndDesc             *prometheus.Desc
	natDstIpv4RestorationsDesc                     *prometheus.Desc
	natDstIpv4TranslationsDesc                     *prometheus.Desc
	natDstIpv6RestorationsDesc                     *prometheus.Desc
	natDstIpv6TranslationsDesc                     *prometheus.Desc
	natDstPortRestorationsDesc                     *prometheus.Desc
	natDstPortTranslationsDesc                     *prometheus.Desc
	natEifMappingFreeDesc                          *prometheus.Desc
	natEimDrainInLookupDesc                        *prometheus.Desc
	natEimDuplicateMappingDesc                     *prometheus.Desc
	natEimEntryDrainedDesc                         *prometheus.Desc
	natEimLookupClearTimerDesc                     *prometheus.Desc
	natEimLookupEntryWithoutTimerDesc              *prometheus.Desc
	natEimLookupHoldSuccessDesc                    *prometheus.Desc
	natEimLookupTimeoutDesc                        *prometheus.Desc
	natEimMappingAllocFailuresDesc                 *prometheus.Desc
	natEimMappingCreateFailedDesc                  *prometheus.Desc
	natEimMappingCreatedDesc                       *prometheus.Desc
	natEimMappingCreatedWithoutEifSessLimitDesc    *prometheus.Desc
	natEimMappingEifCurrSessUpdateInvalidDesc      *prometheus.Desc
	natEimMappingFreeDesc                          *prometheus.Desc
	natEimMappingReusedDesc                        *prometheus.Desc
	natEimMappingUpdatedDesc                       *prometheus.Desc
	natEimMismatchedMappingDesc                    *prometheus.Desc
	natEimReleaseInTimeoutDesc                     *prometheus.Desc
	natEimReleaseRaceDesc                          *prometheus.Desc
	natEimReleaseSetTimeoutDesc                    *prometheus.Desc
	natEimReleaseWithoutEntryDesc                  *prometheus.Desc
	natEimTimerEntryRefreshedDesc                  *prometheus.Desc
	natEimTimerFreeMappingDesc                     *prometheus.Desc
	natEimTimerStartInvalidDesc                    *prometheus.Desc
	natEimTimerStartInvalidFailDesc                *prometheus.Desc
	natEimTimerUpdateTimeoutDesc                   *prometheus.Desc
	natEimWaitingForInitDesc                       *prometheus.Desc
	natEimWaitingForInitFailedDesc                 *prometheus.Desc
	natErrorIpVersionDesc                          *prometheus.Desc
	natErrorNoPolicyDesc                           *prometheus.Desc
	natFilteringSessionDesc                        *prometheus.Desc
	natFreeFailOnInactiveSsetDesc                  *prometheus.Desc
	natGreCallIdRestorationsDesc                   *prometheus.Desc
	natGreCallIdTranslationsDesc                   *prometheus.Desc
	natIcmpAllocationFailureDesc                   *prometheus.Desc
	natIcmpDropDesc                                *prometheus.Desc
	natIcmpErrorDstRestoredDesc                    *prometheus.Desc
	natIcmpErrorDstXlatedDesc                      *prometheus.Desc
	natIcmpErrorNewSrcXlatedDesc                   *prometheus.Desc
	natIcmpErrorOrgIpDstPortRestoredDesc           *prometheus.Desc
	natIcmpErrorOrgIpDstPortXlatedDesc             *prometheus.Desc
	natIcmpErrorOrgIpDstRestoredDesc               *prometheus.Desc
	natIcmpErrorOrgIpDstXlatedDesc                 *prometheus.Desc
	natIcmpErrorOrgIpSrcPortRestoredDesc           *prometheus.Desc
	natIcmpErrorOrgIpSrcPortXlatedDesc             *prometheus.Desc
	natIcmpErrorOrgIpSrcRestoredDesc               *prometheus.Desc
	natIcmpErrorOrgIpSrcXlatedDesc                 *prometheus.Desc
	natIcmpErrorSrcRestoredDesc                    *prometheus.Desc
	natIcmpErrorSrcXlatedDesc                      *prometheus.Desc
	natIcmpErrorTranslationsDesc                   *prometheus.Desc
	natIcmpIdTranslationsDesc                      *prometheus.Desc
	natJflowLogAllocFailDesc                       *prometheus.Desc
	natJflowLogAllocSuccessDesc                    *prometheus.Desc
	natJflowLogFreeFailDataDesc                    *prometheus.Desc
	natJflowLogFreeFailRecordDesc                  *prometheus.Desc
	natJflowLogFreeSuccessDesc                     *prometheus.Desc
	natJflowLogFreeSuccessFailQueuingDesc          *prometheus.Desc
	natJflowLogInvalidAllocErrDesc                 *prometheus.Desc
	natJflowLogInvalidInputArgsDesc                *prometheus.Desc
	natJflowLogInvalidTransTypeDesc                *prometheus.Desc
	natJflowLogNatSextNullDesc                     *prometheus.Desc
	natJflowLogRateLimitFailGetNatpoolDesc         *prometheus.Desc
	natJflowLogRateLimitFailGetNatpoolGivenIdDesc  *prometheus.Desc
	natJflowLogRateLimitFailGetServiceSetDesc      *prometheus.Desc
	natJflowLogRateLimitFailInvalidCurrentTimeDesc *prometheus.Desc
	natJflowLogRateLimitFailGetPoolNameDesc        *prometheus.Desc
	natMapAllocationFailuresDesc                   *prometheus.Desc
	natMapAllocationSuccessesDesc                  *prometheus.Desc
	natMapFreeFailuresDesc                         *prometheus.Desc
	natMapFreeSuccessDesc                          *prometheus.Desc
	natMappingSessionDesc                          *prometheus.Desc
	natNoSextInXlatePktDesc                        *prometheus.Desc
	natOcmpIdRestorationsDesc                      *prometheus.Desc
	natPktDropInBackupStateDesc                    *prometheus.Desc
	natPktDstInNatRouteDesc                        *prometheus.Desc
	natPolicyAddFailedDesc                         *prometheus.Desc
	natPolicyDeleteFailedDesc                      *prometheus.Desc
	natPoolSessionCntUpdateFailOnCloseDesc         *prometheus.Desc
	natPoolSessionCntUpdateFailOnCreateDesc        *prometheus.Desc
	natPrefixFilterAllocFailedDesc                 *prometheus.Desc
	natPrefixFilterChangedDesc                     *prometheus.Desc
	natPrefixFilterCreatedDesc                     *prometheus.Desc
	natPrefixFilterCtrlFreeDesc                    *prometheus.Desc
	natPrefixFilterMappingAddDesc                  *prometheus.Desc
	natPrefixFilterMappingFreeDesc                 *prometheus.Desc
	natPrefixFilterMappingRemoveDesc               *prometheus.Desc
	natPrefixFilterMatchDesc                       *prometheus.Desc
	natPrefixFilterNameFailedDesc                  *prometheus.Desc
	natPrefixFilterNoMatchDesc                     *prometheus.Desc
	natPrefixFilterTreeAddFailedDesc               *prometheus.Desc
	natPrefixFilterUnsuppIpVersionDesc             *prometheus.Desc
	natPrefixListCreateFailedDesc                  *prometheus.Desc
	natRuleLookupFailuresDesc                      *prometheus.Desc
	natRuleLookupForIcmpErrFailDesc                *prometheus.Desc
	natSessionExtAllocFailuresDesc                 *prometheus.Desc
	natSessionExtFreeFailedDesc                    *prometheus.Desc
	natSessionExtSetFailuresDesc                   *prometheus.Desc
	natSessionInterestPubReqDesc                   *prometheus.Desc
	natSrcIpv4RestorationsDesc                     *prometheus.Desc
	natSrcIpv4TranslationsDesc                     *prometheus.Desc
	natSrcIpv6RestorationsDesc                     *prometheus.Desc
	natSrcIpv6TranslationsDesc                     *prometheus.Desc
	natSrcPortRestorationsDesc                     *prometheus.Desc
	natSrcPortTranslationsDesc                     *prometheus.Desc
	natSubsExtAllocDesc                            *prometheus.Desc
	natSubsExtDecInvalEimCntDesc                   *prometheus.Desc
	natSubsExtDecInvalSessCntDesc                  *prometheus.Desc
	natSubsExtDelayTimerFailDesc                   *prometheus.Desc
	natSubsExtDelayTimerSuccessDesc                *prometheus.Desc
	natSubsExtErrSetStateDesc                      *prometheus.Desc
	natSubsExtInTwDuringFreeDesc                   *prometheus.Desc
	natSubsExtIncorrectStateDesc                   *prometheus.Desc
	natSubsExtInlinkSuccessDesc                    *prometheus.Desc
	natSubsExtInvalidEimrefcntDesc                 *prometheus.Desc
	natSubsExtInvalidParamDesc                     *prometheus.Desc
	natSubsExtIsInvalidDesc                        *prometheus.Desc
	natSubsExtIsInvalidSubsInTwDesc                *prometheus.Desc
	natSubsExtIsNullDesc                           *prometheus.Desc
	natSubsExtLinkExistDesc                        *prometheus.Desc
	natSubsExtLinkFailDesc                         *prometheus.Desc
	natSubsExtLinkSuccessDesc                      *prometheus.Desc
	natSubsExtLinkUnknownRetDesc                   *prometheus.Desc
	natSubsExtMissingExtDesc                       *prometheus.Desc
	natSubsExtNoMemDesc                            *prometheus.Desc
	natSubsExtPortsInUseErrDesc                    *prometheus.Desc
	natSubsExtQueueInconsistentDesc                *prometheus.Desc
	natSubsExtRefcountDecFailDesc                  *prometheus.Desc
	natSubsExtResourceInUseDesc                    *prometheus.Desc
	natSubsExtReturnToPreallocErrDesc              *prometheus.Desc
	natSubsExtReuseFromTimerDesc                   *prometheus.Desc
	natSubsExtSubsResetFailDesc                    *prometheus.Desc
	natSubsExtSubsSessionCountUpdateIgnoreDesc     *prometheus.Desc
	natSubsExtSvcSetIsNullDesc                     *prometheus.Desc
	natSubsExtSvcSetNotActiveDesc                  *prometheus.Desc
	natSubsExtTimerCbDesc                          *prometheus.Desc
	natSubsExtTimerStartFailDesc                   *prometheus.Desc
	natSubsExtTimerStartSuccessDesc                *prometheus.Desc
	natSubsExtUnlinkBusyDesc                       *prometheus.Desc
	natSubsExtUnlinkFailDesc                       *prometheus.Desc
	natSubsExtUnlinkUnkErrDesc                     *prometheus.Desc
	natSubsExtfreeDesc                             *prometheus.Desc
	natTcpPortRestorationsDesc                     *prometheus.Desc
	natTcpPortTranslationsDesc                     *prometheus.Desc
	natTotalBytesProcessedDesc                     *prometheus.Desc
	natTotalPktsDiscardedDesc                      *prometheus.Desc
	natTotalPktsForwardedDesc                      *prometheus.Desc
	natTotalPktsProcessedDesc                      *prometheus.Desc
	natTotalPktsRestoredDesc                       *prometheus.Desc
	natTotalPktsTranslatedDesc                     *prometheus.Desc
	natTotalSessionAcceptsDesc                     *prometheus.Desc
	natTotalSessionCloseDesc                       *prometheus.Desc
	natTotalSessionCreateDesc                      *prometheus.Desc
	natTotalSessionDestroyDesc                     *prometheus.Desc
	natTotalSessionDiscardsDesc                    *prometheus.Desc
	natTotalSessionIgnoresDesc                     *prometheus.Desc
	natTotalSessionInterestDesc                    *prometheus.Desc
	natTotalSessionPubReqDesc                      *prometheus.Desc
	natTotalSessionTimeEventDesc                   *prometheus.Desc
	natUdpPortRestorationsDesc                     *prometheus.Desc
	natUdpPortTranslationsDesc                     *prometheus.Desc
	natUnexpectedProtoWithPortXlationDesc          *prometheus.Desc
	natUnsupportedGreProtoDesc                     *prometheus.Desc
	natXlateFreeNullExtDesc                        *prometheus.Desc
	natunsupportedIcmpTypeNaptDesc                 *prometheus.Desc
	natunsupportedLayer4NaptDesc                   *prometheus.Desc
	portsInUseDesc                                 *prometheus.Desc
	outOfPortErrorsDesc                            *prometheus.Desc
	parityPortErrorsDesc                           *prometheus.Desc
	preserveRangeErrorsDesc                        *prometheus.Desc
	maxPortsInUseDesc                              *prometheus.Desc
	appPortErrorsDesc                              *prometheus.Desc
	appExceedPortLimitErrorsDesc                   *prometheus.Desc
	memAllocErrorsDesc                             *prometheus.Desc
	maxPortBlocksUsedDesc                          *prometheus.Desc
	blocksInUseDesc                                *prometheus.Desc
	blockAllocationErrorsDesc                      *prometheus.Desc
	blocksLimitExceededErrorsDesc                  *prometheus.Desc
	usersDesc                                      *prometheus.Desc
	eifInboundSessionCountDesc                     *prometheus.Desc
	eifInboundLimitExceedDropDesc                  *prometheus.Desc
	portBlockSizeDesc                              *prometheus.Desc
	activeBlockTimeoutDesc                         *prometheus.Desc
	maxBlocksPerAddressDesc                        *prometheus.Desc
	effectivePortBlocksDesc                        *prometheus.Desc
	effectivePortsDesc                             *prometheus.Desc
	portBlockEfficiencyDesc                        *prometheus.Desc
	serviceSetCpuUtilizationDesc                   *prometheus.Desc
)

func init() {
	l := []string{"target", "interface"}
	lpool := []string{"target", "interface", "pool_name", "translation_type", "port_range", "port_block_type"}
	lservicesets := []string{"target", "interface", "service_set"}

	nat64DfbitSetDesc = prometheus.NewDesc(prefix+"nat64_dfbit_set", "NAT64 - dfbit set", l, nil)
	nat64ErrMapDstDesc = prometheus.NewDesc(prefix+"nat64_err_map_dst", "NAT64 error - mapping ipv6 destination", l, nil)
	nat64ErrMapSrcDesc = prometheus.NewDesc(prefix+"nat64_err_map_src", "NAT64 error - mapping ipv4 source", l, nil)
	nat64ErrMtuExceedBuildDesc = prometheus.NewDesc(prefix+"nat64_err_mtu_exceed_build", "NAT64 error - MTU exceed build", l, nil)
	nat64ErrMtuExceedSendDesc = prometheus.NewDesc(prefix+"nat64_err_mtu_exceed_send", "NAT64 error - MTU exceed send", l, nil)
	nat64ErrTtlExceedBuildDesc = prometheus.NewDesc(prefix+"nat64_err_ttl_exceed_build", "NAT64 error - TTL exceed build", l, nil)
	nat64ErrTtlExceedSendDesc = prometheus.NewDesc(prefix+"nat64_err_ttl_exceed_send", "NAT64 error - TTL exceed send", l, nil)
	nat64IpoptionsDropDesc = prometheus.NewDesc(prefix+"nat64_ipoptions_drop", "NAT64 - IP options drop", l, nil)
	nat64MtuExceedDesc = prometheus.NewDesc(prefix+"nat64_mtu_exceed", "NAT64 - MTU exceeded", l, nil)
	nat64UdpCksumZeroDropDesc = prometheus.NewDesc(prefix+"nat64_udp_cksum_zero_drop", "NAT64 - UDP checksum zero drop", l, nil)
	nat64UnsuppHdrDropDesc = prometheus.NewDesc(prefix+"nat64_unsupp_hdr_drop", "NAT64 - Unsupported header drop", l, nil)
	nat64UnsuppIcmpCodeDropDesc = prometheus.NewDesc(prefix+"nat64_unsupp_icmp_code_drop", "NAT64 - Unsupported ICMP code drop", l, nil)
	nat64UnsuppIcmpErrorDesc = prometheus.NewDesc(prefix+"nat64_unsupp_icmp_error", "NAT64 - Unsupported ICMP error", l, nil)
	nat64UnsuppIcmpTypeDropDesc = prometheus.NewDesc(prefix+"nat64_unsupp_icmp_type_drop", "NAT64 - Unsupported ICMP type drop", l, nil)
	nat64UnsuppL4DropDesc = prometheus.NewDesc(prefix+"nat64_unsupp_l4_drop", "NAT64 - Unsupported L4 drop", l, nil)
	natAlgDataSessionCreatedDesc = prometheus.NewDesc(prefix+"nat_alg_data_session_created", "ALG Session Create", l, nil)
	natAlgDataSessionInterestDesc = prometheus.NewDesc(prefix+"nat_alg_data_session_interest", "ALG Session interest", l, nil)
	natCmEimLnodeCeletedDesc = prometheus.NewDesc(prefix+"nat_cm_eim_lnode_deleted", "EIM List Node Deleted", l, nil)
	natCmEimLnodeCreatedDesc = prometheus.NewDesc(prefix+"nat_cm_eim_lnode_created", "EIM List Node Created", l, nil)
	natCmSessLnodeCeletedDesc = prometheus.NewDesc(prefix+"nat_cm_sess_lnode_deleted", "Session List Node Deleted", l, nil)
	natCmSessLnodeCreatedDesc = prometheus.NewDesc(prefix+"nat_cm_sess_lnode_created", "Session List Node Created", l, nil)
	natCtrlSessNotXltdChldSessIgndDesc = prometheus.NewDesc(prefix+"nat_ctrl_sess_not_xltd_chld_sess_ignd", "Control Session Not Xlated Child Sess Ignored", l, nil)
	natDstIpv4RestorationsDesc = prometheus.NewDesc(prefix+"nat_dst_ipv4_restorations", "Dst  IPv4   Restorations", l, nil)
	natDstIpv4TranslationsDesc = prometheus.NewDesc(prefix+"nat_dst_ipv4_translations", "Dst  IPv4   Translations", l, nil)
	natDstIpv6RestorationsDesc = prometheus.NewDesc(prefix+"nat_dst_ipv6_restorations", "Dst  IPv6   Restorations", l, nil)
	natDstIpv6TranslationsDesc = prometheus.NewDesc(prefix+"nat_dst_ipv6_translations", "Dst  IPv6   Translations", l, nil)
	natDstPortRestorationsDesc = prometheus.NewDesc(prefix+"nat_dst_port_restorations", "Dst  Port   Restorations", l, nil)
	natDstPortTranslationsDesc = prometheus.NewDesc(prefix+"nat_dst_port_translations", "Dst  Port   Translations", l, nil)
	natEifMappingFreeDesc = prometheus.NewDesc(prefix+"nat_eif_mapping_free", "NAT EIF mapping Free", l, nil)
	natEimDrainInLookupDesc = prometheus.NewDesc(prefix+"nat_eim_drain_in_lookup", "NAT EIM lookup timer drained", l, nil)
	natEimDuplicateMappingDesc = prometheus.NewDesc(prefix+"nat_eim_duplicate_mapping", "NAT EIM mapping duplicate entry", l, nil)
	natEimEntryDrainedDesc = prometheus.NewDesc(prefix+"nat_eim_entry_drained", "NAT EIM entry drained", l, nil)
	natEimLookupClearTimerDesc = prometheus.NewDesc(prefix+"nat_eim_lookup_clear_timer", "NAT EIM lookup timer cleared for timeout entry", l, nil)
	natEimLookupEntryWithoutTimerDesc = prometheus.NewDesc(prefix+"nat_eim_lookup_entry_without_timer", "NAT EIM lookup timeout entry without timer", l, nil)
	natEimLookupHoldSuccessDesc = prometheus.NewDesc(prefix+"nat_eim_lookup_hold_success", "NAT EIM lookup and hold success", l, nil)
	natEimLookupTimeoutDesc = prometheus.NewDesc(prefix+"nat_eim_lookup_timeout", "NAT EIM lookup entry in timeout", l, nil)
	natEimMappingAllocFailuresDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_alloc_failures", "NAT EIM mapping allocation failures", l, nil)
	natEimMappingCreateFailedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_create_failed", "NAT EIM mapping create failed", l, nil)
	natEimMappingCreatedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_created", "NAT EIM mapping Created", l, nil)
	natEimMappingCreatedWithoutEifSessLimitDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_created_without_eif_sess_limit", "NAT EIM mapping - created without eif sess limit", l, nil)
	natEimMappingEifCurrSessUpdateInvalidDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_eif_curr_sess_update_invalid", "NAT EIM mapping - eif curr session update invalid", l, nil)
	natEimMappingFreeDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_free", "NAT EIM mapping Free", l, nil)
	natEimMappingReusedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_reused", "NAT EIM mapping reused", l, nil)
	natEimMappingUpdatedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_updated", "NAT EIM mapping Updated", l, nil)
	natEimMismatchedMappingDesc = prometheus.NewDesc(prefix+"nat_eim_mismatched_mapping", "NAT EIM mapping mismatched entry", l, nil)
	natEimReleaseInTimeoutDesc = prometheus.NewDesc(prefix+"nat_eim_release_in_timeout", "NAT EIM release entry in timeout", l, nil)
	natEimReleaseRaceDesc = prometheus.NewDesc(prefix+"nat_eim_release_race", "NAT EIM release race", l, nil)
	natEimReleaseSetTimeoutDesc = prometheus.NewDesc(prefix+"nat_eim_release_set_timeout", "NAT EIM release set entry for timeout", l, nil)
	natEimReleaseWithoutEntryDesc = prometheus.NewDesc(prefix+"nat_eim_release_without_entry", "NAT EIM release without entry", l, nil)
	natEimTimerEntryRefreshedDesc = prometheus.NewDesc(prefix+"nat_eim_timer_entry_refreshed", "NAT EIM timer entry refreshed", l, nil)
	natEimTimerFreeMappingDesc = prometheus.NewDesc(prefix+"nat_eim_timer_free_mapping", "NAT EIM timer entry freed", l, nil)
	natEimTimerStartInvalidDesc = prometheus.NewDesc(prefix+"nat_eim_timer_start_invalid", "NAT EIM timer invalid timer started", l, nil)
	natEimTimerStartInvalidFailDesc = prometheus.NewDesc(prefix+"nat_eim_timer_start_invalid_fail", "NAT EIM timer invalid timer start failed", l, nil)
	natEimTimerUpdateTimeoutDesc = prometheus.NewDesc(prefix+"nat_eim_timer_update_timeout", "NAT EIM timer entry updated", l, nil)
	natEimWaitingForInitDesc = prometheus.NewDesc(prefix+"nat_eim_waiting_for_init", "NAT EIM waiting for init", l, nil)
	natEimWaitingForInitFailedDesc = prometheus.NewDesc(prefix+"nat_eim_waiting_for_init_failed", "NAT EIM waiting for init failed", l, nil)
	natErrorIpVersionDesc = prometheus.NewDesc(prefix+"nat_error_ip_version", "NAT error - IP version", l, nil)
	natErrorNoPolicyDesc = prometheus.NewDesc(prefix+"nat_error_no_policy", "NAT error - no policy", l, nil)
	natFilteringSessionDesc = prometheus.NewDesc(prefix+"nat_filtering_session", "Session Created for EIF", l, nil)
	natFreeFailOnInactiveSsetDesc = prometheus.NewDesc(prefix+"nat_free_fail_on_inactive_sset", "NAT Free failures while service set is not active", l, nil)
	natGreCallIdRestorationsDesc = prometheus.NewDesc(prefix+"nat_gre_call_id_restorations", "GRE  CallID Restorations", l, nil)
	natGreCallIdTranslationsDesc = prometheus.NewDesc(prefix+"nat_gre_call_id_translations", "GRE  CallID Translations", l, nil)
	natIcmpAllocationFailureDesc = prometheus.NewDesc(prefix+"nat_icmp_allocation_failure", "ICMP Allocation Failure", l, nil)
	natIcmpDropDesc = prometheus.NewDesc(prefix+"nat_icmp_drop", "ICMP Drops", l, nil)
	natIcmpErrorDstRestoredDesc = prometheus.NewDesc(prefix+"nat_icmp_error_dst_restored", "DST IP restored in ICMP Error", l, nil)
	natIcmpErrorDstXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_dst_xlated", "DST IP translated in ICMP Error", l, nil)
	natIcmpErrorNewSrcXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_new_src_xlated", "New SRC IP translated in ICMP Error", l, nil)
	natIcmpErrorOrgIpDstPortRestoredDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_dst_port_restored", "Inner DST port restored in ICMP Error", l, nil)
	natIcmpErrorOrgIpDstPortXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_dst_port_xlated", "Inner DST port translated in ICMP Error", l, nil)
	natIcmpErrorOrgIpDstRestoredDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_dst_restored", "Inner DST IP restored in ICMP Error", l, nil)
	natIcmpErrorOrgIpDstXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_dst_xlated", "Inner DST IP translated in ICMP Error", l, nil)
	natIcmpErrorOrgIpSrcPortRestoredDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_src_port_restored", "Inner SRC port restored in ICMP Error", l, nil)
	natIcmpErrorOrgIpSrcPortXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_src_port_xlated", "Inner SRC port translated in ICMP Error", l, nil)
	natIcmpErrorOrgIpSrcRestoredDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_src_restored", "Inner SRC IP restored in ICMP Error", l, nil)
	natIcmpErrorOrgIpSrcXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_org_ip_src_xlated", "Inner SRC IP translated in ICMP Error", l, nil)
	natIcmpErrorSrcRestoredDesc = prometheus.NewDesc(prefix+"nat_icmp_error_src_restored", "SRC IP restored in ICMP Error", l, nil)
	natIcmpErrorSrcXlatedDesc = prometheus.NewDesc(prefix+"nat_icmp_error_src_xlated", "SRC IP translated in ICMP Error", l, nil)
	natIcmpErrorTranslationsDesc = prometheus.NewDesc(prefix+"nat_icmp_error_translations", "ICMP Error  Translations", l, nil)
	natIcmpIdTranslationsDesc = prometheus.NewDesc(prefix+"nat_icmp_id_translations", "ICMP ID     Translations", l, nil)
	natJflowLogAllocFailDesc = prometheus.NewDesc(prefix+"nat_jflow_log_alloc_fail", "NAT jflow-log error - memory allocation fail", l, nil)
	natJflowLogAllocSuccessDesc = prometheus.NewDesc(prefix+"nat_jflow_log_alloc_success", "NAT jflow-log - memory allocation success", l, nil)
	natJflowLogFreeFailDataDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_fail_data", "NAT jflow-log error - memory free fail null data", l, nil)
	natJflowLogFreeFailRecordDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_fail_record", "NAT jflow-log error - memory free fail null record", l, nil)
	natJflowLogFreeSuccessDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_success", "NAT jflow-log - memory free success", l, nil)
	natJflowLogFreeSuccessFailQueuingDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_success_fail_queuing", "NAT jflow-log - memory free success fail queuing", l, nil)
	natJflowLogInvalidAllocErrDesc = prometheus.NewDesc(prefix+"nat_jflow_log_invalid_alloc_err", "NAT jflow-log - invalid allocation error type", l, nil)
	natJflowLogInvalidInputArgsDesc = prometheus.NewDesc(prefix+"nat_jflow_log_invalid_input_args", "NAT jflow-log - invalid input arguments", l, nil)
	natJflowLogInvalidTransTypeDesc = prometheus.NewDesc(prefix+"nat_jflow_log_invalid_trans_type", "NAT jflow-log error - invalid nat translation type", l, nil)
	natJflowLogNatSextNullDesc = prometheus.NewDesc(prefix+"nat_jflow_log_nat_sext_null", "NAT jflow-log error - session extension get fail", l, nil)
	natJflowLogRateLimitFailGetNatpoolDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_get_natpool", "NAT jflow-log - rate limit fail to get nat pool", l, nil)
	natJflowLogRateLimitFailGetNatpoolGivenIdDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_get_natpool_given_id", "NAT jflow-log - rate limit fail to get pool given id", l, nil)
	natJflowLogRateLimitFailGetServiceSetDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_get_service_set", "NAT jflow-log - rate limit fail to get service set", l, nil)
	natJflowLogRateLimitFailInvalidCurrentTimeDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_invalid_current_time", "NAT jflow-log - rate limit fail invalid current time", l, nil)
	natJflowLogRateLimitFailGetPoolNameDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_get_pool_name", "NAT jflow-log - rate limit fail to get pool name", l, nil)
	natMapAllocationFailuresDesc = prometheus.NewDesc(prefix+"nat_map_allocation_failures", "NAT allocation Failures", l, nil)
	natMapAllocationSuccessesDesc = prometheus.NewDesc(prefix+"nat_map_allocation_successes", "NAT allocation Successes", l, nil)
	natMapFreeFailuresDesc = prometheus.NewDesc(prefix+"nat_map_free_failures", "NAT Free Failures", l, nil)
	natMapFreeSuccessDesc = prometheus.NewDesc(prefix+"nat_map_free_success", "NAT Free Successes", l, nil)
	natMappingSessionDesc = prometheus.NewDesc(prefix+"nat_mapping_session", "Session Created for EIM", l, nil)
	natNoSextInXlatePktDesc = prometheus.NewDesc(prefix+"no_sext_in_xlate_pkt", "No NAT session ext in xlate packet", l, nil)
	natOcmpIdRestorationsDesc = prometheus.NewDesc(prefix+"nat_icmp_id_restorations", "ICMP ID     Restorations", l, nil)
	natPktDropInBackupStateDesc = prometheus.NewDesc(prefix+"nat_pkt_drop_in_backup_state", "Packet drop in backup state", l, nil)
	natPktDstInNatRouteDesc = prometheus.NewDesc(prefix+"nat_pkt_dst_in_nat_route", "Packet  Dst in NAT route", l, nil)
	natPolicyAddFailedDesc = prometheus.NewDesc(prefix+"nat_policy_add_failed", "NAT error - policy add failed", l, nil)
	natPolicyDeleteFailedDesc = prometheus.NewDesc(prefix+"nat_policy_delete_failed", "NAT error - policy delete failed", l, nil)
	natPoolSessionCntUpdateFailOnCloseDesc = prometheus.NewDesc(prefix+"pool_session_cnt_update_fail_on_close", "Pool session count update failed on close", l, nil)
	natPoolSessionCntUpdateFailOnCreateDesc = prometheus.NewDesc(prefix+"pool_session_cnt_update_fail_on_create", "Pool session count update failed on create", l, nil)
	natPrefixFilterAllocFailedDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_alloc_failed", "NAT error - prefix filter allocation failed", l, nil)
	natPrefixFilterChangedDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_changed", "NAT prefix filter changed", l, nil)
	natPrefixFilterCreatedDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_created", "NAT prefix filter created", l, nil)
	natPrefixFilterCtrlFreeDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_ctrl_free", "NAT prefix filter control free", l, nil)
	natPrefixFilterMappingAddDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_mapping_add", "NAT prefix filter mapping add", l, nil)
	natPrefixFilterMappingFreeDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_mapping_free", "NAT prefix filter mapping free", l, nil)
	natPrefixFilterMappingRemoveDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_mapping_remove", "NAT prefix filter mapping remove", l, nil)
	natPrefixFilterMatchDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_match", "NAT prefix filter match", l, nil)
	natPrefixFilterNameFailedDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_name_failed", "NAT error - prefix filter name failed", l, nil)
	natPrefixFilterNoMatchDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_no_match", "NAT prefix filter no match", l, nil)
	natPrefixFilterTreeAddFailedDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_tree_add_failed", "NAT error - prefix filter tree add failed", l, nil)
	natPrefixFilterUnsuppIpVersionDesc = prometheus.NewDesc(prefix+"nat_prefix_filter_unsupp_ip_version", "NAT prefix filter unsupported IP version", l, nil)
	natPrefixListCreateFailedDesc = prometheus.NewDesc(prefix+"nat_prefix_list_create_failed", "NAT error - prefix list create failed", l, nil)
	natRuleLookupFailuresDesc = prometheus.NewDesc(prefix+"nat_rule_lookup_failures", "NAT rule lookup failures", l, nil)
	natRuleLookupForIcmpErrFailDesc = prometheus.NewDesc(prefix+"nat_rule_lookup_for_icmp_err_fail", "ICMP Error  NAT rule lookup fail", l, nil)
	natSessionExtAllocFailuresDesc = prometheus.NewDesc(prefix+"nat_session_ext_alloc_failures", "Session Ext Alloc Failures", l, nil)
	natSessionExtFreeFailedDesc = prometheus.NewDesc(prefix+"nat_session_ext_free_failed", "NAT error - ext free failed", l, nil)
	natSessionExtSetFailuresDesc = prometheus.NewDesc(prefix+"nat_session_ext_set_failures", "Session Ext Set Failures", l, nil)
	natSessionInterestPubReqDesc = prometheus.NewDesc(prefix+"nat_session_interest_pub_req", "Session interest thru pub event", l, nil)
	natSrcIpv4RestorationsDesc = prometheus.NewDesc(prefix+"nat_src_ipv4_restorations", "Src  IPv4   Restorations", l, nil)
	natSrcIpv4TranslationsDesc = prometheus.NewDesc(prefix+"nat_src_ipv4_translations", "Src  IPv4   Translations", l, nil)
	natSrcIpv6RestorationsDesc = prometheus.NewDesc(prefix+"nat_src_ipv6_restorations", "Src  IPv6   Restorations", l, nil)
	natSrcIpv6TranslationsDesc = prometheus.NewDesc(prefix+"nat_src_ipv6_translations", "Src  IPv6   Translations", l, nil)
	natSrcPortRestorationsDesc = prometheus.NewDesc(prefix+"nat_src_port_restorations", "Src  Port   Restorations", l, nil)
	natSrcPortTranslationsDesc = prometheus.NewDesc(prefix+"nat_src_port_translations", "Src  Port   Translations", l, nil)
	natSubsExtAllocDesc = prometheus.NewDesc(prefix+"nat_subs_ext_alloc", "NAT subscriber extension allocated", l, nil)
	natSubsExtDecInvalEimCntDesc = prometheus.NewDesc(prefix+"nat_subs_ext_dec_inval_eim_cnt", "NAT subscriber extension dec invalid eim count", l, nil)
	natSubsExtDecInvalSessCntDesc = prometheus.NewDesc(prefix+"nat_subs_ext_dec_inval_sess_cnt", "NAT subscriber extension dec invalid session count", l, nil)
	natSubsExtDelayTimerFailDesc = prometheus.NewDesc(prefix+"nat_subs_ext_delay_timer_fail", "NAT subscriber extension delay timer start failed", l, nil)
	natSubsExtDelayTimerSuccessDesc = prometheus.NewDesc(prefix+"nat_subs_ext_delay_timer_success", "NAT subscriber extension delay timer start successful", l, nil)
	natSubsExtErrSetStateDesc = prometheus.NewDesc(prefix+"nat_subs_ext_err_set_state", "NAT subscriber extension error while setting state", l, nil)
	natSubsExtInTwDuringFreeDesc = prometheus.NewDesc(prefix+"nat_subs_ext_in_tw_during_free", "NAT subscriber extension is in timer wheel during free", l, nil)
	natSubsExtIncorrectStateDesc = prometheus.NewDesc(prefix+"nat_subs_ext_incorrect_state", "NAT subscriber extension incorrect state", l, nil)
	natSubsExtInlinkSuccessDesc = prometheus.NewDesc(prefix+"nat_subs_ext_unlink_success", "NAT subscriber extension unlink successful", l, nil)
	natSubsExtInvalidEimrefcntDesc = prometheus.NewDesc(prefix+"nat_subs_ext_invalid_eimrefcnt", "NAT subscriber extension unexpected eim refcount", l, nil)
	natSubsExtInvalidParamDesc = prometheus.NewDesc(prefix+"nat_subs_ext_invalid_param", "NAT subscriber extension invalid parameters", l, nil)
	natSubsExtIsInvalidDesc = prometheus.NewDesc(prefix+"nat_subs_ext_is_invalid", "NAT subscriber extension is invalid", l, nil)
	natSubsExtIsInvalidSubsInTwDesc = prometheus.NewDesc(prefix+"nat_subs_ext_is_invalid_subs_in_tw", "NAT subscriber extension is invalid and in timer wheel", l, nil)
	natSubsExtIsNullDesc = prometheus.NewDesc(prefix+"nat_subs_ext_is_null", "NAT subscriber extension is null", l, nil)
	natSubsExtLinkExistDesc = prometheus.NewDesc(prefix+"nat_subs_ext_link_exist", "NAT subscriber extension link already exists", l, nil)
	natSubsExtLinkFailDesc = prometheus.NewDesc(prefix+"nat_subs_ext_link_fail", "NAT subscriber extension link failed", l, nil)
	natSubsExtLinkSuccessDesc = prometheus.NewDesc(prefix+"nat_subs_ext_link_success", "NAT subscriber extension link successful", l, nil)
	natSubsExtLinkUnknownRetDesc = prometheus.NewDesc(prefix+"nat_subs_ext_link_unknown_ret", "NAT subscriber extension link unknown return value", l, nil)
	natSubsExtMissingExtDesc = prometheus.NewDesc(prefix+"nat_subs_ext_missing_ext", "NAT subscriber extension nat extension is missing", l, nil)
	natSubsExtNoMemDesc = prometheus.NewDesc(prefix+"nat_subs_ext_no_mem", "NAT subscriber extension no memory", l, nil)
	natSubsExtPortsInUseErrDesc = prometheus.NewDesc(prefix+"nat_subs_ext_ports_in_use_err", "NAT subscriber extension ports in use error", l, nil)
	natSubsExtQueueInconsistentDesc = prometheus.NewDesc(prefix+"nat_subs_ext_queue_inconsistent", "NAT subscriber extension queue inconsistent", l, nil)
	natSubsExtRefcountDecFailDesc = prometheus.NewDesc(prefix+"nat_subs_ext_refcount_dec_fail", "NAT subscriber extension refcount decrement failed", l, nil)
	natSubsExtResourceInUseDesc = prometheus.NewDesc(prefix+"nat_subs_ext_resource_in_use", "NAT subscriber extension resource in use", l, nil)
	natSubsExtReturnToPreallocErrDesc = prometheus.NewDesc(prefix+"nat_subs_ext_return_to_prealloc_err", "NAT subscriber extension return to prealloc queue error", l, nil)
	natSubsExtReuseFromTimerDesc = prometheus.NewDesc(prefix+"nat_subs_ext_reuse_from_timer", "NAT subscriber extension reuse from timer", l, nil)
	natSubsExtSubsResetFailDesc = prometheus.NewDesc(prefix+"nat_subs_ext_subs_reset_fail", "NAT subscriber extension subscriber reset failed", l, nil)
	natSubsExtSubsSessionCountUpdateIgnoreDesc = prometheus.NewDesc(prefix+"nat_subs_ext_subs_session_count_update_ignore", "NAT subscriber extension session count update ignored", l, nil)
	natSubsExtSvcSetIsNullDesc = prometheus.NewDesc(prefix+"nat_subs_ext_svc_set_is_null", "NAT subscriber extension svc set is null", l, nil)
	natSubsExtSvcSetNotActiveDesc = prometheus.NewDesc(prefix+"nat_subs_ext_svc_set_not_active", "NAT subscriber extension svc set is not active", l, nil)
	natSubsExtTimerCbDesc = prometheus.NewDesc(prefix+"nat_subs_ext_timer_cb", "NAT subscriber extension timer callback called", l, nil)
	natSubsExtTimerStartFailDesc = prometheus.NewDesc(prefix+"nat_subs_ext_timer_start_fail", "NAT subscriber extension timer start failed", l, nil)
	natSubsExtTimerStartSuccessDesc = prometheus.NewDesc(prefix+"nat_subs_ext_timer_start_success", "NAT subscriber extension timer start successful", l, nil)
	natSubsExtUnlinkBusyDesc = prometheus.NewDesc(prefix+"nat_subs_ext_unlink_busy", "NAT subscriber extension unlink on busy", l, nil)
	natSubsExtUnlinkFailDesc = prometheus.NewDesc(prefix+"nat_subs_ext_unlink_fail", "NAT subscriber extension unlink fail", l, nil)
	natSubsExtUnlinkUnkErrDesc = prometheus.NewDesc(prefix+"nat_subs_ext_unlink_unk_err", "NAT subscriber extension unknown error unlinking", l, nil)
	natSubsExtfreeDesc = prometheus.NewDesc(prefix+"nat_subs_ext_free", "NAT subscriber extension freed", l, nil)
	natTcpPortRestorationsDesc = prometheus.NewDesc(prefix+"nat_tcp_port_restorations", "TCP  Port   Restorations", l, nil)
	natTcpPortTranslationsDesc = prometheus.NewDesc(prefix+"nat_tcp_port_translations", "TCP  Port   Translations", l, nil)
	natTotalBytesProcessedDesc = prometheus.NewDesc(prefix+"nat_total_bytes_processed", "Total Bytes   Processed", l, nil)
	natTotalPktsDiscardedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_discarded", "Total Packets Discarded", l, nil)
	natTotalPktsForwardedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_forwarded", "Total Packets Forwarded", l, nil)
	natTotalPktsProcessedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_processed", "Total Packets Processed", l, nil)
	natTotalPktsRestoredDesc = prometheus.NewDesc(prefix+"nat_total_pkts_restored", "Total Packets Restored", l, nil)
	natTotalPktsTranslatedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_translated", "Total Packets Translated", l, nil)
	natTotalSessionAcceptsDesc = prometheus.NewDesc(prefix+"nat_total_session_accepts", "Total Session Accepts", l, nil)
	natTotalSessionCloseDesc = prometheus.NewDesc(prefix+"total_session_close", "Total Session close", l, nil)
	natTotalSessionCreateDesc = prometheus.NewDesc(prefix+"nat_total_session_create", "Total Session Create events", l, nil)
	natTotalSessionDestroyDesc = prometheus.NewDesc(prefix+"nat_total_session_destroy", "Total Session Destroy events", l, nil)
	natTotalSessionDiscardsDesc = prometheus.NewDesc(prefix+"nat_total_session_discards", "Total Session Discards", l, nil)
	natTotalSessionIgnoresDesc = prometheus.NewDesc(prefix+"nat_total_session_ignores", "Total Session Ignores", l, nil)
	natTotalSessionInterestDesc = prometheus.NewDesc(prefix+"nat_total_session_interest", "Total Session Interest events", l, nil)
	natTotalSessionPubReqDesc = prometheus.NewDesc(prefix+"nat_total_session_pub_req", "Total Session Pub Req events", l, nil)
	natTotalSessionTimeEventDesc = prometheus.NewDesc(prefix+"nat_total_session_time_event", "Total Session Time events", l, nil)
	natUdpPortRestorationsDesc = prometheus.NewDesc(prefix+"nat_udp_port_restorations", "UDP  Port   Restorations", l, nil)
	natUdpPortTranslationsDesc = prometheus.NewDesc(prefix+"nat_udp_port_translations", "UDP  Port   Translations", l, nil)
	natUnexpectedProtoWithPortXlationDesc = prometheus.NewDesc(prefix+"nat_unexpected_proto_with_port_xlation", "NAT Unexpected Protocol With Port Xlation", l, nil)
	natUnsupportedGreProtoDesc = prometheus.NewDesc(prefix+"nat_unsupported_gre_proto", "GRE  Wrong protocol value", l, nil)
	natXlateFreeNullExtDesc = prometheus.NewDesc(prefix+"nat_xlate_free_null_ext", "NAT error - xlate free called with null ext", l, nil)
	natunsupportedIcmpTypeNaptDesc = prometheus.NewDesc(prefix+"nat_unsupported_icmp_type_napt", "NAT unsupported icmp id for port translation", l, nil)
	natunsupportedLayer4NaptDesc = prometheus.NewDesc(prefix+"nat_unsupported_layer_4_napt", "NAT unsupported layer-4 header for port translation", l, nil)
	portsInUseDesc = prometheus.NewDesc(prefix+"pool_ports_in_use", "NAT ports in use", lpool, nil)
	outOfPortErrorsDesc = prometheus.NewDesc(prefix+"pool_out_of_port_errors", "NAT out of ports errors", lpool, nil)
	parityPortErrorsDesc = prometheus.NewDesc(prefix+"pool_parity_port_errors", "NAT parity port errors", lpool, nil)
	preserveRangeErrorsDesc = prometheus.NewDesc(prefix+"pool_preserve_range_errors", "NAT preserve range errors", lpool, nil)
	maxPortsInUseDesc = prometheus.NewDesc(prefix+"pool_max_ports_in_use", "NAT maximum ports in use", lpool, nil)
	appPortErrorsDesc = prometheus.NewDesc(prefix+"pool_app_port_errors", "NAT AP-P port allocation errors", lpool, nil)
	appExceedPortLimitErrorsDesc = prometheus.NewDesc(prefix+"pool_app_exceed_port_limit_errors", "NAT AP-P port limit exceeded errors", lpool, nil)
	memAllocErrorsDesc = prometheus.NewDesc(prefix+"pool_mem_alloc_errors", "NAT memory allocation errors", lpool, nil)
	maxPortBlocksUsedDesc = prometheus.NewDesc(prefix+"pool_max_port_blocks_used", "NAT max port blocks in use", lpool, nil)
	blocksInUseDesc = prometheus.NewDesc(prefix+"pool_blocks_in_use", "NAT port blocks in use", lpool, nil)
	blockAllocationErrorsDesc = prometheus.NewDesc(prefix+"pool_block_allocation_errors", "NAT port block allocation errors", lpool, nil)
	blocksLimitExceededErrorsDesc = prometheus.NewDesc(prefix+"pool_blocks_limit_exceeded_errors", "NAT port blocks limit exceeded errors", lpool, nil)
	usersDesc = prometheus.NewDesc(prefix+"pool_users", "NAT current users", lpool, nil)
	eifInboundSessionCountDesc = prometheus.NewDesc(prefix+"pool_eif_inbound_session_count", "NAT inbound EIF sessions", lpool, nil)
	eifInboundLimitExceedDropDesc = prometheus.NewDesc(prefix+"pool_eif_inbound_limit_exceed_drop", "NAT inbound EIF limit exceeded drops", lpool, nil)
	portBlockSizeDesc = prometheus.NewDesc(prefix+"pool_port_block_size", "NAT Pool port block size", lpool, nil)
	activeBlockTimeoutDesc = prometheus.NewDesc(prefix+"pool_active_block_timeout", "NAT Pool active-block-timeout", lpool, nil)
	maxBlocksPerAddressDesc = prometheus.NewDesc(prefix+"pool_max_blocks_per_address", "NAT Pool max blocks per address", lpool, nil)
	effectivePortBlocksDesc = prometheus.NewDesc(prefix+"pool_effective_port_blocks", "NAT Pool effective port blocks", lpool, nil)
	effectivePortsDesc = prometheus.NewDesc(prefix+"pool_effective_ports", "NAT Pool effective ports", lpool, nil)
	portBlockEfficiencyDesc = prometheus.NewDesc(prefix+"pool_port_block_efficiency", "NAT Pool port block efficiency", lpool, nil)
	serviceSetCpuUtilizationDesc = prometheus.NewDesc(prefix+"service_set_cpu_utlization", "CPU utilization for the Service Set", lservicesets, nil)

}

type natCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &natCollector{}
}

// Name returns the name of the collector
func (*natCollector) Name() string {
	return "NAT"
}

// Describe describes the metrics
func (*natCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- natTotalSessionInterestDesc
}

// Collect collects metrics from JunOS
func (c *natCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	interfaces, err := c.NatInterfaces(client)
	if err != nil {
		return err
	}
	for _, s := range interfaces {
		c.collectForInterface(s, ch, labelValues)
	}

	poolinterfaces, err := c.NatPoolInterfaces(client, ch, labelValues)
	if err != nil {
		return err
	}
	for _, s := range poolinterfaces {
		c.collectForPoolInterface(s, ch, labelValues)
	}

	pooldetailinterfaces, err := c.NatPoolDetailInterfaces(client, ch, labelValues)
	if err != nil {
		return err
	}
	for _, s := range pooldetailinterfaces {
		c.collectForPoolDetailInterface(s, ch, labelValues)
	}

	servicesetscpuinterfaces, err := c.ServiceSetsCpuInterfaces(client, ch, labelValues)
	for _, s := range servicesetscpuinterfaces {
		c.collectForServiceSetsCpuInterface(s, ch, labelValues)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *natCollector) NatInterfaces(client *rpc.Client) ([]*NatInterface, error) {
	var x = NatRpc{}
	err := client.RunCommandAndParse("show services nat statistics", &x)
	if err != nil {
		return nil, err
	}

	interfaces := make([]*NatInterface, 0)
	for _, natinterface := range x.Interfaces {
		s := &NatInterface{
			Interface:                                  natinterface.Interface,
			Nat64DfbitSet:                              int64(natinterface.Nat64DfbitSet),
			Nat64ErrMapDst:                             int64(natinterface.Nat64ErrMapDst),
			Nat64ErrMapSrc:                             int64(natinterface.Nat64ErrMapSrc),
			Nat64ErrMtuExceedBuild:                     int64(natinterface.Nat64ErrMtuExceedBuild),
			Nat64ErrMtuExceedSend:                      int64(natinterface.Nat64ErrMtuExceedSend),
			Nat64ErrTtlExceedBuild:                     int64(natinterface.Nat64ErrTtlExceedBuild),
			Nat64ErrTtlExceedSend:                      int64(natinterface.Nat64ErrTtlExceedSend),
			Nat64IpoptionsDrop:                         int64(natinterface.Nat64IpoptionsDrop),
			Nat64MtuExceed:                             int64(natinterface.Nat64MtuExceed),
			Nat64UdpCksumZeroDrop:                      int64(natinterface.Nat64UdpCksumZeroDrop),
			Nat64UnsuppHdrDrop:                         int64(natinterface.Nat64UnsuppHdrDrop),
			Nat64UnsuppIcmpCodeDrop:                    int64(natinterface.Nat64UnsuppIcmpCodeDrop),
			Nat64UnsuppIcmpError:                       int64(natinterface.Nat64UnsuppIcmpError),
			Nat64UnsuppIcmpTypeDrop:                    int64(natinterface.Nat64UnsuppIcmpTypeDrop),
			Nat64UnsuppL4Drop:                          int64(natinterface.Nat64UnsuppL4Drop),
			NatAlgDataSessionCreated:                   int64(natinterface.NatAlgDataSessionCreated),
			NatAlgDataSessionInterest:                  int64(natinterface.NatAlgDataSessionInterest),
			NatCmEimLnodeCeleted:                       int64(natinterface.NatCmEimLnodeCeleted),
			NatCmEimLnodeCreated:                       int64(natinterface.NatCmEimLnodeCreated),
			NatCmSessLnodeCeleted:                      int64(natinterface.NatCmSessLnodeCeleted),
			NatCmSessLnodeCreated:                      int64(natinterface.NatCmSessLnodeCreated),
			NatCtrlSessNotXltdChldSessIgnd:             int64(natinterface.NatCtrlSessNotXltdChldSessIgnd),
			NatDstIpv4Restorations:                     int64(natinterface.NatDstIpv4Restorations),
			NatDstIpv4Translations:                     int64(natinterface.NatDstIpv4Translations),
			NatDstIpv6Restorations:                     int64(natinterface.NatDstIpv6Restorations),
			NatDstIpv6Translations:                     int64(natinterface.NatDstIpv6Translations),
			NatDstPortRestorations:                     int64(natinterface.NatDstPortRestorations),
			NatDstPortTranslations:                     int64(natinterface.NatDstPortTranslations),
			NatEifMappingFree:                          int64(natinterface.NatEifMappingFree),
			NatEimDrainInLookup:                        int64(natinterface.NatEimDrainInLookup),
			NatEimDuplicateMapping:                     int64(natinterface.NatEimDuplicateMapping),
			NatEimEntryDrained:                         int64(natinterface.NatEimEntryDrained),
			NatEimLookupClearTimer:                     int64(natinterface.NatEimLookupClearTimer),
			NatEimLookupEntryWithoutTimer:              int64(natinterface.NatEimLookupEntryWithoutTimer),
			NatEimLookupHoldSuccess:                    int64(natinterface.NatEimLookupHoldSuccess),
			NatEimLookupTimeout:                        int64(natinterface.NatEimLookupTimeout),
			NatEimMappingAllocFailures:                 int64(natinterface.NatEimMappingAllocFailures),
			NatEimMappingCreateFailed:                  int64(natinterface.NatEimMappingCreateFailed),
			NatEimMappingCreated:                       int64(natinterface.NatEimMappingCreated),
			NatEimMappingCreatedWithoutEifSessLimit:    int64(natinterface.NatEimMappingCreatedWithoutEifSessLimit),
			NatEimMappingEifCurrSessUpdateInvalid:      int64(natinterface.NatEimMappingEifCurrSessUpdateInvalid),
			NatEimMappingFree:                          int64(natinterface.NatEimMappingFree),
			NatEimMappingReused:                        int64(natinterface.NatEimMappingReused),
			NatEimMappingUpdated:                       int64(natinterface.NatEimMappingUpdated),
			NatEimMismatchedMapping:                    int64(natinterface.NatEimMismatchedMapping),
			NatEimReleaseInTimeout:                     int64(natinterface.NatEimReleaseInTimeout),
			NatEimReleaseRace:                          int64(natinterface.NatEimReleaseRace),
			NatEimReleaseSetTimeout:                    int64(natinterface.NatEimReleaseSetTimeout),
			NatEimReleaseWithoutEntry:                  int64(natinterface.NatEimReleaseWithoutEntry),
			NatEimTimerEntryRefreshed:                  int64(natinterface.NatEimTimerEntryRefreshed),
			NatEimTimerFreeMapping:                     int64(natinterface.NatEimTimerFreeMapping),
			NatEimTimerStartInvalid:                    int64(natinterface.NatEimTimerStartInvalid),
			NatEimTimerStartInvalidFail:                int64(natinterface.NatEimTimerStartInvalidFail),
			NatEimTimerUpdateTimeout:                   int64(natinterface.NatEimTimerUpdateTimeout),
			NatEimWaitingForInit:                       int64(natinterface.NatEimWaitingForInit),
			NatEimWaitingForInitFailed:                 int64(natinterface.NatEimWaitingForInitFailed),
			NatErrorIpVersion:                          int64(natinterface.NatErrorIpVersion),
			NatErrorNoPolicy:                           int64(natinterface.NatErrorNoPolicy),
			NatFilteringSession:                        int64(natinterface.NatFilteringSession),
			NatFreeFailOnInactiveSset:                  int64(natinterface.NatFreeFailOnInactiveSset),
			NatGreCallIdRestorations:                   int64(natinterface.NatGreCallIdRestorations),
			NatGreCallIdTranslations:                   int64(natinterface.NatGreCallIdTranslations),
			NatIcmpAllocationFailure:                   int64(natinterface.NatIcmpAllocationFailure),
			NatIcmpDrop:                                int64(natinterface.NatIcmpDrop),
			NatIcmpErrorDstRestored:                    int64(natinterface.NatIcmpErrorDstRestored),
			NatIcmpErrorDstXlated:                      int64(natinterface.NatIcmpErrorDstXlated),
			NatIcmpErrorNewSrcXlated:                   int64(natinterface.NatIcmpErrorNewSrcXlated),
			NatIcmpErrorOrgIpDstPortRestored:           int64(natinterface.NatIcmpErrorOrgIpDstPortRestored),
			NatIcmpErrorOrgIpDstPortXlated:             int64(natinterface.NatIcmpErrorOrgIpDstPortXlated),
			NatIcmpErrorOrgIpDstRestored:               int64(natinterface.NatIcmpErrorOrgIpDstRestored),
			NatIcmpErrorOrgIpDstXlated:                 int64(natinterface.NatIcmpErrorOrgIpDstXlated),
			NatIcmpErrorOrgIpSrcPortRestored:           int64(natinterface.NatIcmpErrorOrgIpSrcPortRestored),
			NatIcmpErrorOrgIpSrcPortXlated:             int64(natinterface.NatIcmpErrorOrgIpSrcPortXlated),
			NatIcmpErrorOrgIpSrcRestored:               int64(natinterface.NatIcmpErrorOrgIpSrcRestored),
			NatIcmpErrorOrgIpSrcXlated:                 int64(natinterface.NatIcmpErrorOrgIpSrcXlated),
			NatIcmpErrorSrcRestored:                    int64(natinterface.NatIcmpErrorSrcRestored),
			NatIcmpErrorSrcXlated:                      int64(natinterface.NatIcmpErrorSrcXlated),
			NatIcmpErrorTranslations:                   int64(natinterface.NatIcmpErrorTranslations),
			NatIcmpIdTranslations:                      int64(natinterface.NatIcmpIdTranslations),
			NatJflowLogAllocFail:                       int64(natinterface.NatJflowLogAllocFail),
			NatJflowLogAllocSuccess:                    int64(natinterface.NatJflowLogAllocSuccess),
			NatJflowLogFreeFailData:                    int64(natinterface.NatJflowLogFreeFailData),
			NatJflowLogFreeFailRecord:                  int64(natinterface.NatJflowLogFreeFailRecord),
			NatJflowLogFreeSuccess:                     int64(natinterface.NatJflowLogFreeSuccess),
			NatJflowLogFreeSuccessFailQueuing:          int64(natinterface.NatJflowLogFreeSuccessFailQueuing),
			NatJflowLogInvalidAllocErr:                 int64(natinterface.NatJflowLogInvalidAllocErr),
			NatJflowLogInvalidInputArgs:                int64(natinterface.NatJflowLogInvalidInputArgs),
			NatJflowLogInvalidTransType:                int64(natinterface.NatJflowLogInvalidTransType),
			NatJflowLogNatSextNull:                     int64(natinterface.NatJflowLogNatSextNull),
			NatJflowLogRateLimitFailGetNatpool:         int64(natinterface.NatJflowLogRateLimitFailGetNatpool),
			NatJflowLogRateLimitFailGetNatpoolGivenId:  int64(natinterface.NatJflowLogRateLimitFailGetNatpoolGivenId),
			NatJflowLogRateLimitFailGetServiceSet:      int64(natinterface.NatJflowLogRateLimitFailGetServiceSet),
			NatJflowLogRateLimitFailInvalidCurrentTime: int64(natinterface.NatJflowLogRateLimitFailInvalidCurrentTime),
			NatJflowLogRateLimitFailGetPoolName:        int64(natinterface.NatJflowLogRateLimitFailGetPoolName),
			NatMapAllocationFailures:                   int64(natinterface.NatMapAllocationFailures),
			NatMapAllocationSuccesses:                  int64(natinterface.NatMapAllocationSuccesses),
			NatMapFreeFailures:                         int64(natinterface.NatMapFreeFailures),
			NatMapFreeSuccess:                          int64(natinterface.NatMapFreeSuccess),
			NatMappingSession:                          int64(natinterface.NatMappingSession),
			NatNoSextInXlatePkt:                        int64(natinterface.NatNoSextInXlatePkt),
			NatOcmpIdRestorations:                      int64(natinterface.NatOcmpIdRestorations),
			NatPktDropInBackupState:                    int64(natinterface.NatPktDropInBackupState),
			NatPktDstInNatRoute:                        int64(natinterface.NatPktDstInNatRoute),
			NatPolicyAddFailed:                         int64(natinterface.NatPolicyAddFailed),
			NatPolicyDeleteFailed:                      int64(natinterface.NatPolicyDeleteFailed),
			NatPoolSessionCntUpdateFailOnClose:         int64(natinterface.NatPoolSessionCntUpdateFailOnClose),
			NatPoolSessionCntUpdateFailOnCreate:        int64(natinterface.NatPoolSessionCntUpdateFailOnCreate),
			NatPrefixFilterAllocFailed:                 int64(natinterface.NatPrefixFilterAllocFailed),
			NatPrefixFilterChanged:                     int64(natinterface.NatPrefixFilterChanged),
			NatPrefixFilterCreated:                     int64(natinterface.NatPrefixFilterCreated),
			NatPrefixFilterCtrlFree:                    int64(natinterface.NatPrefixFilterCtrlFree),
			NatPrefixFilterMappingAdd:                  int64(natinterface.NatPrefixFilterMappingAdd),
			NatPrefixFilterMappingFree:                 int64(natinterface.NatPrefixFilterMappingFree),
			NatPrefixFilterMappingRemove:               int64(natinterface.NatPrefixFilterMappingRemove),
			NatPrefixFilterMatch:                       int64(natinterface.NatPrefixFilterMatch),
			NatPrefixFilterNameFailed:                  int64(natinterface.NatPrefixFilterNameFailed),
			NatPrefixFilterNoMatch:                     int64(natinterface.NatPrefixFilterNoMatch),
			NatPrefixFilterTreeAddFailed:               int64(natinterface.NatPrefixFilterTreeAddFailed),
			NatPrefixFilterUnsuppIpVersion:             int64(natinterface.NatPrefixFilterUnsuppIpVersion),
			NatPrefixListCreateFailed:                  int64(natinterface.NatPrefixListCreateFailed),
			NatRuleLookupFailures:                      int64(natinterface.NatRuleLookupFailures),
			NatRuleLookupForIcmpErrFail:                int64(natinterface.NatRuleLookupForIcmpErrFail),
			NatSessionExtAllocFailures:                 int64(natinterface.NatSessionExtAllocFailures),
			NatSessionExtFreeFailed:                    int64(natinterface.NatSessionExtFreeFailed),
			NatSessionExtSetFailures:                   int64(natinterface.NatSessionExtSetFailures),
			NatSessionInterestPubReq:                   int64(natinterface.NatSessionInterestPubReq),
			NatSrcIpv4Restorations:                     int64(natinterface.NatSrcIpv4Restorations),
			NatSrcIpv4Translations:                     int64(natinterface.NatSrcIpv4Translations),
			NatSrcIpv6Restorations:                     int64(natinterface.NatSrcIpv6Restorations),
			NatSrcIpv6Translations:                     int64(natinterface.NatSrcIpv6Translations),
			NatSrcPortRestorations:                     int64(natinterface.NatSrcPortRestorations),
			NatSrcPortTranslations:                     int64(natinterface.NatSrcPortTranslations),
			NatSubsExtAlloc:                            int64(natinterface.NatSubsExtAlloc),
			NatSubsExtDecInvalEimCnt:                   int64(natinterface.NatSubsExtDecInvalEimCnt),
			NatSubsExtDecInvalSessCnt:                  int64(natinterface.NatSubsExtDecInvalSessCnt),
			NatSubsExtDelayTimerFail:                   int64(natinterface.NatSubsExtDelayTimerFail),
			NatSubsExtDelayTimerSuccess:                int64(natinterface.NatSubsExtDelayTimerSuccess),
			NatSubsExtErrSetState:                      int64(natinterface.NatSubsExtErrSetState),
			NatSubsExtInTwDuringFree:                   int64(natinterface.NatSubsExtInTwDuringFree),
			NatSubsExtIncorrectState:                   int64(natinterface.NatSubsExtIncorrectState),
			NatSubsExtInlinkSuccess:                    int64(natinterface.NatSubsExtInlinkSuccess),
			NatSubsExtInvalidEimrefcnt:                 int64(natinterface.NatSubsExtInvalidEimrefcnt),
			NatSubsExtInvalidParam:                     int64(natinterface.NatSubsExtInvalidParam),
			NatSubsExtIsInvalid:                        int64(natinterface.NatSubsExtIsInvalid),
			NatSubsExtIsInvalidSubsInTw:                int64(natinterface.NatSubsExtIsInvalidSubsInTw),
			NatSubsExtIsNull:                           int64(natinterface.NatSubsExtIsNull),
			NatSubsExtLinkExist:                        int64(natinterface.NatSubsExtLinkExist),
			NatSubsExtLinkFail:                         int64(natinterface.NatSubsExtLinkFail),
			NatSubsExtLinkSuccess:                      int64(natinterface.NatSubsExtLinkSuccess),
			NatSubsExtLinkUnknownRet:                   int64(natinterface.NatSubsExtLinkUnknownRet),
			NatSubsExtMissingExt:                       int64(natinterface.NatSubsExtMissingExt),
			NatSubsExtNoMem:                            int64(natinterface.NatSubsExtNoMem),
			NatSubsExtPortsInUseErr:                    int64(natinterface.NatSubsExtPortsInUseErr),
			NatSubsExtQueueInconsistent:                int64(natinterface.NatSubsExtQueueInconsistent),
			NatSubsExtRefcountDecFail:                  int64(natinterface.NatSubsExtRefcountDecFail),
			NatSubsExtResourceInUse:                    int64(natinterface.NatSubsExtResourceInUse),
			NatSubsExtReturnToPreallocErr:              int64(natinterface.NatSubsExtReturnToPreallocErr),
			NatSubsExtReuseFromTimer:                   int64(natinterface.NatSubsExtReuseFromTimer),
			NatSubsExtSubsResetFail:                    int64(natinterface.NatSubsExtSubsResetFail),
			NatSubsExtSubsSessionCountUpdateIgnore:     int64(natinterface.NatSubsExtSubsSessionCountUpdateIgnore),
			NatSubsExtSvcSetIsNull:                     int64(natinterface.NatSubsExtSvcSetIsNull),
			NatSubsExtSvcSetNotActive:                  int64(natinterface.NatSubsExtSvcSetNotActive),
			NatSubsExtTimerCb:                          int64(natinterface.NatSubsExtTimerCb),
			NatSubsExtTimerStartFail:                   int64(natinterface.NatSubsExtTimerStartFail),
			NatSubsExtTimerStartSuccess:                int64(natinterface.NatSubsExtTimerStartSuccess),
			NatSubsExtUnlinkBusy:                       int64(natinterface.NatSubsExtUnlinkBusy),
			NatSubsExtUnlinkFail:                       int64(natinterface.NatSubsExtUnlinkFail),
			NatSubsExtUnlinkUnkErr:                     int64(natinterface.NatSubsExtUnlinkUnkErr),
			NatSubsExtfree:                             int64(natinterface.NatSubsExtfree),
			NatTcpPortRestorations:                     int64(natinterface.NatTcpPortRestorations),
			NatTcpPortTranslations:                     int64(natinterface.NatTcpPortTranslations),
			NatTotalBytesProcessed:                     int64(natinterface.NatTotalBytesProcessed),
			NatTotalPktsDiscarded:                      int64(natinterface.NatTotalPktsDiscarded),
			NatTotalPktsForwarded:                      int64(natinterface.NatTotalPktsForwarded),
			NatTotalPktsProcessed:                      int64(natinterface.NatTotalPktsProcessed),
			NatTotalPktsRestored:                       int64(natinterface.NatTotalPktsRestored),
			NatTotalPktsTranslated:                     int64(natinterface.NatTotalPktsTranslated),
			NatTotalSessionAccepts:                     int64(natinterface.NatTotalSessionAccepts),
			NatTotalSessionClose:                       int64(natinterface.NatTotalSessionClose),
			NatTotalSessionCreate:                      int64(natinterface.NatTotalSessionCreate),
			NatTotalSessionDestroy:                     int64(natinterface.NatTotalSessionDestroy),
			NatTotalSessionDiscards:                    int64(natinterface.NatTotalSessionDiscards),
			NatTotalSessionIgnores:                     int64(natinterface.NatTotalSessionIgnores),
			NatTotalSessionInterest:                    int64(natinterface.NatTotalSessionInterest),
			NatTotalSessionPubReq:                      int64(natinterface.NatTotalSessionPubReq),
			NatTotalSessionTimeEvent:                   int64(natinterface.NatTotalSessionTimeEvent),
			NatUdpPortRestorations:                     int64(natinterface.NatUdpPortRestorations),
			NatUdpPortTranslations:                     int64(natinterface.NatUdpPortTranslations),
			NatUnexpectedProtoWithPortXlation:          int64(natinterface.NatUnexpectedProtoWithPortXlation),
			NatUnsupportedGreProto:                     int64(natinterface.NatUnsupportedGreProto),
			NatXlateFreeNullExt:                        int64(natinterface.NatXlateFreeNullExt),
			NatunsupportedIcmpTypeNapt:                 int64(natinterface.NatunsupportedIcmpTypeNapt),
			NatunsupportedLayer4Napt:                   int64(natinterface.NatunsupportedLayer4Napt),
		}

		interfaces = append(interfaces, s)
	}

	return interfaces, nil
}

func (*natCollector) collectForInterface(s *NatInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Interface}...)

	ch <- prometheus.MustNewConstMetric(nat64DfbitSetDesc, prometheus.GaugeValue, float64(s.Nat64DfbitSet), l...)
	ch <- prometheus.MustNewConstMetric(nat64ErrMapDstDesc, prometheus.GaugeValue, float64(s.Nat64ErrMapDst), l...)
	ch <- prometheus.MustNewConstMetric(nat64ErrMapSrcDesc, prometheus.GaugeValue, float64(s.Nat64ErrMapSrc), l...)
	ch <- prometheus.MustNewConstMetric(nat64ErrMtuExceedBuildDesc, prometheus.GaugeValue, float64(s.Nat64ErrMtuExceedBuild), l...)
	ch <- prometheus.MustNewConstMetric(nat64ErrMtuExceedSendDesc, prometheus.GaugeValue, float64(s.Nat64ErrMtuExceedSend), l...)
	ch <- prometheus.MustNewConstMetric(nat64ErrTtlExceedBuildDesc, prometheus.GaugeValue, float64(s.Nat64ErrTtlExceedBuild), l...)
	ch <- prometheus.MustNewConstMetric(nat64ErrTtlExceedSendDesc, prometheus.GaugeValue, float64(s.Nat64ErrTtlExceedSend), l...)
	ch <- prometheus.MustNewConstMetric(nat64IpoptionsDropDesc, prometheus.GaugeValue, float64(s.Nat64IpoptionsDrop), l...)
	ch <- prometheus.MustNewConstMetric(nat64MtuExceedDesc, prometheus.GaugeValue, float64(s.Nat64MtuExceed), l...)
	ch <- prometheus.MustNewConstMetric(nat64UdpCksumZeroDropDesc, prometheus.GaugeValue, float64(s.Nat64UdpCksumZeroDrop), l...)
	ch <- prometheus.MustNewConstMetric(nat64UnsuppHdrDropDesc, prometheus.GaugeValue, float64(s.Nat64UnsuppHdrDrop), l...)
	ch <- prometheus.MustNewConstMetric(nat64UnsuppIcmpCodeDropDesc, prometheus.GaugeValue, float64(s.Nat64UnsuppIcmpCodeDrop), l...)
	ch <- prometheus.MustNewConstMetric(nat64UnsuppIcmpErrorDesc, prometheus.GaugeValue, float64(s.Nat64UnsuppIcmpError), l...)
	ch <- prometheus.MustNewConstMetric(nat64UnsuppIcmpTypeDropDesc, prometheus.GaugeValue, float64(s.Nat64UnsuppIcmpTypeDrop), l...)
	ch <- prometheus.MustNewConstMetric(nat64UnsuppL4DropDesc, prometheus.GaugeValue, float64(s.Nat64UnsuppL4Drop), l...)
	ch <- prometheus.MustNewConstMetric(natAlgDataSessionCreatedDesc, prometheus.GaugeValue, float64(s.NatAlgDataSessionCreated), l...)
	ch <- prometheus.MustNewConstMetric(natAlgDataSessionInterestDesc, prometheus.GaugeValue, float64(s.NatAlgDataSessionInterest), l...)
	ch <- prometheus.MustNewConstMetric(natCmEimLnodeCeletedDesc, prometheus.GaugeValue, float64(s.NatCmEimLnodeCeleted), l...)
	ch <- prometheus.MustNewConstMetric(natCmEimLnodeCreatedDesc, prometheus.GaugeValue, float64(s.NatCmEimLnodeCreated), l...)
	ch <- prometheus.MustNewConstMetric(natCmSessLnodeCeletedDesc, prometheus.GaugeValue, float64(s.NatCmSessLnodeCeleted), l...)
	ch <- prometheus.MustNewConstMetric(natCmSessLnodeCreatedDesc, prometheus.GaugeValue, float64(s.NatCmSessLnodeCreated), l...)
	ch <- prometheus.MustNewConstMetric(natCtrlSessNotXltdChldSessIgndDesc, prometheus.GaugeValue, float64(s.NatCtrlSessNotXltdChldSessIgnd), l...)
	ch <- prometheus.MustNewConstMetric(natDstIpv4RestorationsDesc, prometheus.GaugeValue, float64(s.NatDstIpv4Restorations), l...)
	ch <- prometheus.MustNewConstMetric(natDstIpv4TranslationsDesc, prometheus.GaugeValue, float64(s.NatDstIpv4Translations), l...)
	ch <- prometheus.MustNewConstMetric(natDstIpv6RestorationsDesc, prometheus.GaugeValue, float64(s.NatDstIpv6Restorations), l...)
	ch <- prometheus.MustNewConstMetric(natDstIpv6TranslationsDesc, prometheus.GaugeValue, float64(s.NatDstIpv6Translations), l...)
	ch <- prometheus.MustNewConstMetric(natDstPortRestorationsDesc, prometheus.GaugeValue, float64(s.NatDstPortRestorations), l...)
	ch <- prometheus.MustNewConstMetric(natDstPortTranslationsDesc, prometheus.GaugeValue, float64(s.NatDstPortTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natEifMappingFreeDesc, prometheus.GaugeValue, float64(s.NatEifMappingFree), l...)
	ch <- prometheus.MustNewConstMetric(natEimDrainInLookupDesc, prometheus.GaugeValue, float64(s.NatEimDrainInLookup), l...)
	ch <- prometheus.MustNewConstMetric(natEimDuplicateMappingDesc, prometheus.GaugeValue, float64(s.NatEimDuplicateMapping), l...)
	ch <- prometheus.MustNewConstMetric(natEimEntryDrainedDesc, prometheus.GaugeValue, float64(s.NatEimEntryDrained), l...)
	ch <- prometheus.MustNewConstMetric(natEimLookupClearTimerDesc, prometheus.GaugeValue, float64(s.NatEimLookupClearTimer), l...)
	ch <- prometheus.MustNewConstMetric(natEimLookupEntryWithoutTimerDesc, prometheus.GaugeValue, float64(s.NatEimLookupEntryWithoutTimer), l...)
	ch <- prometheus.MustNewConstMetric(natEimLookupHoldSuccessDesc, prometheus.GaugeValue, float64(s.NatEimLookupHoldSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natEimLookupTimeoutDesc, prometheus.GaugeValue, float64(s.NatEimLookupTimeout), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingAllocFailuresDesc, prometheus.GaugeValue, float64(s.NatEimMappingAllocFailures), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingCreateFailedDesc, prometheus.GaugeValue, float64(s.NatEimMappingCreateFailed), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingCreatedDesc, prometheus.GaugeValue, float64(s.NatEimMappingCreated), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingCreatedWithoutEifSessLimitDesc, prometheus.GaugeValue, float64(s.NatEimMappingCreatedWithoutEifSessLimit), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingEifCurrSessUpdateInvalidDesc, prometheus.GaugeValue, float64(s.NatEimMappingEifCurrSessUpdateInvalid), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingFreeDesc, prometheus.GaugeValue, float64(s.NatEimMappingFree), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingReusedDesc, prometheus.GaugeValue, float64(s.NatEimMappingReused), l...)
	ch <- prometheus.MustNewConstMetric(natEimMappingUpdatedDesc, prometheus.GaugeValue, float64(s.NatEimMappingUpdated), l...)
	ch <- prometheus.MustNewConstMetric(natEimMismatchedMappingDesc, prometheus.GaugeValue, float64(s.NatEimMismatchedMapping), l...)
	ch <- prometheus.MustNewConstMetric(natEimReleaseInTimeoutDesc, prometheus.GaugeValue, float64(s.NatEimReleaseInTimeout), l...)
	ch <- prometheus.MustNewConstMetric(natEimReleaseRaceDesc, prometheus.GaugeValue, float64(s.NatEimReleaseRace), l...)
	ch <- prometheus.MustNewConstMetric(natEimReleaseSetTimeoutDesc, prometheus.GaugeValue, float64(s.NatEimReleaseSetTimeout), l...)
	ch <- prometheus.MustNewConstMetric(natEimReleaseWithoutEntryDesc, prometheus.GaugeValue, float64(s.NatEimReleaseWithoutEntry), l...)
	ch <- prometheus.MustNewConstMetric(natEimTimerEntryRefreshedDesc, prometheus.GaugeValue, float64(s.NatEimTimerEntryRefreshed), l...)
	ch <- prometheus.MustNewConstMetric(natEimTimerFreeMappingDesc, prometheus.GaugeValue, float64(s.NatEimTimerFreeMapping), l...)
	ch <- prometheus.MustNewConstMetric(natEimTimerStartInvalidDesc, prometheus.GaugeValue, float64(s.NatEimTimerStartInvalid), l...)
	ch <- prometheus.MustNewConstMetric(natEimTimerStartInvalidFailDesc, prometheus.GaugeValue, float64(s.NatEimTimerStartInvalidFail), l...)
	ch <- prometheus.MustNewConstMetric(natEimTimerUpdateTimeoutDesc, prometheus.GaugeValue, float64(s.NatEimTimerUpdateTimeout), l...)
	ch <- prometheus.MustNewConstMetric(natEimWaitingForInitDesc, prometheus.GaugeValue, float64(s.NatEimWaitingForInit), l...)
	ch <- prometheus.MustNewConstMetric(natEimWaitingForInitFailedDesc, prometheus.GaugeValue, float64(s.NatEimWaitingForInitFailed), l...)
	ch <- prometheus.MustNewConstMetric(natErrorIpVersionDesc, prometheus.GaugeValue, float64(s.NatErrorIpVersion), l...)
	ch <- prometheus.MustNewConstMetric(natErrorNoPolicyDesc, prometheus.GaugeValue, float64(s.NatErrorNoPolicy), l...)
	ch <- prometheus.MustNewConstMetric(natFilteringSessionDesc, prometheus.GaugeValue, float64(s.NatFilteringSession), l...)
	ch <- prometheus.MustNewConstMetric(natFreeFailOnInactiveSsetDesc, prometheus.GaugeValue, float64(s.NatFreeFailOnInactiveSset), l...)
	ch <- prometheus.MustNewConstMetric(natGreCallIdRestorationsDesc, prometheus.GaugeValue, float64(s.NatGreCallIdRestorations), l...)
	ch <- prometheus.MustNewConstMetric(natGreCallIdTranslationsDesc, prometheus.GaugeValue, float64(s.NatGreCallIdTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpAllocationFailureDesc, prometheus.GaugeValue, float64(s.NatIcmpAllocationFailure), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpDropDesc, prometheus.GaugeValue, float64(s.NatIcmpDrop), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorDstRestoredDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorDstRestored), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorDstXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorDstXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorNewSrcXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorNewSrcXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpDstPortRestoredDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpDstPortRestored), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpDstPortXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpDstPortXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpDstRestoredDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpDstRestored), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpDstXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpDstXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpSrcPortRestoredDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpSrcPortRestored), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpSrcPortXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpSrcPortXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpSrcRestoredDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpSrcRestored), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorOrgIpSrcXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorOrgIpSrcXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorSrcRestoredDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorSrcRestored), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorSrcXlatedDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorSrcXlated), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpErrorTranslationsDesc, prometheus.GaugeValue, float64(s.NatIcmpErrorTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natIcmpIdTranslationsDesc, prometheus.GaugeValue, float64(s.NatIcmpIdTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogAllocFailDesc, prometheus.GaugeValue, float64(s.NatJflowLogAllocFail), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogAllocSuccessDesc, prometheus.GaugeValue, float64(s.NatJflowLogAllocSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogFreeFailDataDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeFailData), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogFreeFailRecordDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeFailRecord), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogFreeSuccessDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogFreeSuccessFailQueuingDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeSuccessFailQueuing), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogInvalidAllocErrDesc, prometheus.GaugeValue, float64(s.NatJflowLogInvalidAllocErr), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogInvalidInputArgsDesc, prometheus.GaugeValue, float64(s.NatJflowLogInvalidInputArgs), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogInvalidTransTypeDesc, prometheus.GaugeValue, float64(s.NatJflowLogInvalidTransType), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogNatSextNullDesc, prometheus.GaugeValue, float64(s.NatJflowLogNatSextNull), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogRateLimitFailGetNatpoolDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailGetNatpool), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogRateLimitFailGetNatpoolGivenIdDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailGetNatpoolGivenId), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogRateLimitFailGetServiceSetDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailGetServiceSet), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogRateLimitFailInvalidCurrentTimeDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailInvalidCurrentTime), l...)
	ch <- prometheus.MustNewConstMetric(natJflowLogRateLimitFailGetPoolNameDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailGetPoolName), l...)
	ch <- prometheus.MustNewConstMetric(natMapAllocationFailuresDesc, prometheus.GaugeValue, float64(s.NatMapAllocationFailures), l...)
	ch <- prometheus.MustNewConstMetric(natMapAllocationSuccessesDesc, prometheus.GaugeValue, float64(s.NatMapAllocationSuccesses), l...)
	ch <- prometheus.MustNewConstMetric(natMapFreeFailuresDesc, prometheus.GaugeValue, float64(s.NatMapFreeFailures), l...)
	ch <- prometheus.MustNewConstMetric(natMapFreeSuccessDesc, prometheus.GaugeValue, float64(s.NatMapFreeSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natMappingSessionDesc, prometheus.GaugeValue, float64(s.NatMappingSession), l...)
	ch <- prometheus.MustNewConstMetric(natNoSextInXlatePktDesc, prometheus.GaugeValue, float64(s.NatNoSextInXlatePkt), l...)
	ch <- prometheus.MustNewConstMetric(natOcmpIdRestorationsDesc, prometheus.GaugeValue, float64(s.NatOcmpIdRestorations), l...)
	ch <- prometheus.MustNewConstMetric(natPktDropInBackupStateDesc, prometheus.GaugeValue, float64(s.NatPktDropInBackupState), l...)
	ch <- prometheus.MustNewConstMetric(natPktDstInNatRouteDesc, prometheus.GaugeValue, float64(s.NatPktDstInNatRoute), l...)
	ch <- prometheus.MustNewConstMetric(natPolicyAddFailedDesc, prometheus.GaugeValue, float64(s.NatPolicyAddFailed), l...)
	ch <- prometheus.MustNewConstMetric(natPolicyDeleteFailedDesc, prometheus.GaugeValue, float64(s.NatPolicyDeleteFailed), l...)
	ch <- prometheus.MustNewConstMetric(natPoolSessionCntUpdateFailOnCloseDesc, prometheus.GaugeValue, float64(s.NatPoolSessionCntUpdateFailOnClose), l...)
	ch <- prometheus.MustNewConstMetric(natPoolSessionCntUpdateFailOnCreateDesc, prometheus.GaugeValue, float64(s.NatPoolSessionCntUpdateFailOnCreate), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterAllocFailedDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterAllocFailed), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterChangedDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterChanged), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterCreatedDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterCreated), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterCtrlFreeDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterCtrlFree), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterMappingAddDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterMappingAdd), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterMappingFreeDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterMappingFree), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterMappingRemoveDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterMappingRemove), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterMatchDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterMatch), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterNameFailedDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterNameFailed), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterNoMatchDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterNoMatch), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterTreeAddFailedDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterTreeAddFailed), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixFilterUnsuppIpVersionDesc, prometheus.GaugeValue, float64(s.NatPrefixFilterUnsuppIpVersion), l...)
	ch <- prometheus.MustNewConstMetric(natPrefixListCreateFailedDesc, prometheus.GaugeValue, float64(s.NatPrefixListCreateFailed), l...)
	ch <- prometheus.MustNewConstMetric(natRuleLookupFailuresDesc, prometheus.GaugeValue, float64(s.NatRuleLookupFailures), l...)
	ch <- prometheus.MustNewConstMetric(natRuleLookupForIcmpErrFailDesc, prometheus.GaugeValue, float64(s.NatRuleLookupForIcmpErrFail), l...)
	ch <- prometheus.MustNewConstMetric(natSessionExtAllocFailuresDesc, prometheus.GaugeValue, float64(s.NatSessionExtAllocFailures), l...)
	ch <- prometheus.MustNewConstMetric(natSessionExtFreeFailedDesc, prometheus.GaugeValue, float64(s.NatSessionExtFreeFailed), l...)
	ch <- prometheus.MustNewConstMetric(natSessionExtSetFailuresDesc, prometheus.GaugeValue, float64(s.NatSessionExtSetFailures), l...)
	ch <- prometheus.MustNewConstMetric(natSessionInterestPubReqDesc, prometheus.GaugeValue, float64(s.NatSessionInterestPubReq), l...)
	ch <- prometheus.MustNewConstMetric(natSrcIpv4RestorationsDesc, prometheus.GaugeValue, float64(s.NatSrcIpv4Restorations), l...)
	ch <- prometheus.MustNewConstMetric(natSrcIpv4TranslationsDesc, prometheus.GaugeValue, float64(s.NatSrcIpv4Translations), l...)
	ch <- prometheus.MustNewConstMetric(natSrcIpv6RestorationsDesc, prometheus.GaugeValue, float64(s.NatSrcIpv6Restorations), l...)
	ch <- prometheus.MustNewConstMetric(natSrcIpv6TranslationsDesc, prometheus.GaugeValue, float64(s.NatSrcIpv6Translations), l...)
	ch <- prometheus.MustNewConstMetric(natSrcPortRestorationsDesc, prometheus.GaugeValue, float64(s.NatSrcPortRestorations), l...)
	ch <- prometheus.MustNewConstMetric(natSrcPortTranslationsDesc, prometheus.GaugeValue, float64(s.NatSrcPortTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtAllocDesc, prometheus.GaugeValue, float64(s.NatSubsExtAlloc), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtDecInvalEimCntDesc, prometheus.GaugeValue, float64(s.NatSubsExtDecInvalEimCnt), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtDecInvalSessCntDesc, prometheus.GaugeValue, float64(s.NatSubsExtDecInvalSessCnt), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtDelayTimerFailDesc, prometheus.GaugeValue, float64(s.NatSubsExtDelayTimerFail), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtDelayTimerSuccessDesc, prometheus.GaugeValue, float64(s.NatSubsExtDelayTimerSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtErrSetStateDesc, prometheus.GaugeValue, float64(s.NatSubsExtErrSetState), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtInTwDuringFreeDesc, prometheus.GaugeValue, float64(s.NatSubsExtInTwDuringFree), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtIncorrectStateDesc, prometheus.GaugeValue, float64(s.NatSubsExtIncorrectState), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtInlinkSuccessDesc, prometheus.GaugeValue, float64(s.NatSubsExtInlinkSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtInvalidEimrefcntDesc, prometheus.GaugeValue, float64(s.NatSubsExtInvalidEimrefcnt), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtInvalidParamDesc, prometheus.GaugeValue, float64(s.NatSubsExtInvalidParam), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtIsInvalidDesc, prometheus.GaugeValue, float64(s.NatSubsExtIsInvalid), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtIsInvalidSubsInTwDesc, prometheus.GaugeValue, float64(s.NatSubsExtIsInvalidSubsInTw), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtIsNullDesc, prometheus.GaugeValue, float64(s.NatSubsExtIsNull), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtLinkExistDesc, prometheus.GaugeValue, float64(s.NatSubsExtLinkExist), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtLinkFailDesc, prometheus.GaugeValue, float64(s.NatSubsExtLinkFail), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtLinkSuccessDesc, prometheus.GaugeValue, float64(s.NatSubsExtLinkSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtLinkUnknownRetDesc, prometheus.GaugeValue, float64(s.NatSubsExtLinkUnknownRet), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtMissingExtDesc, prometheus.GaugeValue, float64(s.NatSubsExtMissingExt), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtNoMemDesc, prometheus.GaugeValue, float64(s.NatSubsExtNoMem), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtPortsInUseErrDesc, prometheus.GaugeValue, float64(s.NatSubsExtPortsInUseErr), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtQueueInconsistentDesc, prometheus.GaugeValue, float64(s.NatSubsExtQueueInconsistent), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtRefcountDecFailDesc, prometheus.GaugeValue, float64(s.NatSubsExtRefcountDecFail), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtResourceInUseDesc, prometheus.GaugeValue, float64(s.NatSubsExtResourceInUse), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtReturnToPreallocErrDesc, prometheus.GaugeValue, float64(s.NatSubsExtReturnToPreallocErr), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtReuseFromTimerDesc, prometheus.GaugeValue, float64(s.NatSubsExtReuseFromTimer), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtSubsResetFailDesc, prometheus.GaugeValue, float64(s.NatSubsExtSubsResetFail), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtSubsSessionCountUpdateIgnoreDesc, prometheus.GaugeValue, float64(s.NatSubsExtSubsSessionCountUpdateIgnore), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtSvcSetIsNullDesc, prometheus.GaugeValue, float64(s.NatSubsExtSvcSetIsNull), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtSvcSetNotActiveDesc, prometheus.GaugeValue, float64(s.NatSubsExtSvcSetNotActive), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtTimerCbDesc, prometheus.GaugeValue, float64(s.NatSubsExtTimerCb), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtTimerStartFailDesc, prometheus.GaugeValue, float64(s.NatSubsExtTimerStartFail), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtTimerStartSuccessDesc, prometheus.GaugeValue, float64(s.NatSubsExtTimerStartSuccess), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtUnlinkBusyDesc, prometheus.GaugeValue, float64(s.NatSubsExtUnlinkBusy), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtUnlinkFailDesc, prometheus.GaugeValue, float64(s.NatSubsExtUnlinkFail), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtUnlinkUnkErrDesc, prometheus.GaugeValue, float64(s.NatSubsExtUnlinkUnkErr), l...)
	ch <- prometheus.MustNewConstMetric(natSubsExtfreeDesc, prometheus.GaugeValue, float64(s.NatSubsExtfree), l...)
	ch <- prometheus.MustNewConstMetric(natTcpPortRestorationsDesc, prometheus.GaugeValue, float64(s.NatTcpPortRestorations), l...)
	ch <- prometheus.MustNewConstMetric(natTcpPortTranslationsDesc, prometheus.GaugeValue, float64(s.NatTcpPortTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natTotalBytesProcessedDesc, prometheus.GaugeValue, float64(s.NatTotalBytesProcessed), l...)
	ch <- prometheus.MustNewConstMetric(natTotalPktsDiscardedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsDiscarded), l...)
	ch <- prometheus.MustNewConstMetric(natTotalPktsForwardedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsForwarded), l...)
	ch <- prometheus.MustNewConstMetric(natTotalPktsProcessedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsProcessed), l...)
	ch <- prometheus.MustNewConstMetric(natTotalPktsRestoredDesc, prometheus.GaugeValue, float64(s.NatTotalPktsRestored), l...)
	ch <- prometheus.MustNewConstMetric(natTotalPktsTranslatedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsTranslated), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionAcceptsDesc, prometheus.GaugeValue, float64(s.NatTotalSessionAccepts), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionCloseDesc, prometheus.GaugeValue, float64(s.NatTotalSessionClose), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionCreateDesc, prometheus.GaugeValue, float64(s.NatTotalSessionCreate), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionDestroyDesc, prometheus.GaugeValue, float64(s.NatTotalSessionDestroy), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionDiscardsDesc, prometheus.GaugeValue, float64(s.NatTotalSessionDiscards), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionIgnoresDesc, prometheus.GaugeValue, float64(s.NatTotalSessionIgnores), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionInterestDesc, prometheus.GaugeValue, float64(s.NatTotalSessionInterest), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionPubReqDesc, prometheus.GaugeValue, float64(s.NatTotalSessionPubReq), l...)
	ch <- prometheus.MustNewConstMetric(natTotalSessionTimeEventDesc, prometheus.GaugeValue, float64(s.NatTotalSessionTimeEvent), l...)
	ch <- prometheus.MustNewConstMetric(natUdpPortRestorationsDesc, prometheus.GaugeValue, float64(s.NatUdpPortRestorations), l...)
	ch <- prometheus.MustNewConstMetric(natUdpPortTranslationsDesc, prometheus.GaugeValue, float64(s.NatUdpPortTranslations), l...)
	ch <- prometheus.MustNewConstMetric(natUnexpectedProtoWithPortXlationDesc, prometheus.GaugeValue, float64(s.NatUnexpectedProtoWithPortXlation), l...)
	ch <- prometheus.MustNewConstMetric(natUnsupportedGreProtoDesc, prometheus.GaugeValue, float64(s.NatUnsupportedGreProto), l...)
	ch <- prometheus.MustNewConstMetric(natXlateFreeNullExtDesc, prometheus.GaugeValue, float64(s.NatXlateFreeNullExt), l...)
	ch <- prometheus.MustNewConstMetric(natunsupportedIcmpTypeNaptDesc, prometheus.GaugeValue, float64(s.NatunsupportedIcmpTypeNapt), l...)
	ch <- prometheus.MustNewConstMetric(natunsupportedLayer4NaptDesc, prometheus.GaugeValue, float64(s.NatunsupportedLayer4Napt), l...)
}

func (c *natCollector) NatPoolInterfaces(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) ([]*NatPoolInterface, error) {
	var x = NatPoolRpc{}
	err := client.RunCommandAndParse("show services nat pool", &x)
	if err != nil {
		return nil, err
	}

	interfaces := make([]*NatPoolInterface, 0)
	for _, natpoolinterface := range x.Information.Interfaces {
		s := &NatPoolInterface{
			Interface:       natpoolinterface.Interface,
			ServiceSetName:  natpoolinterface.ServiceSetName,
			ServiceNatPools: natpoolinterface.ServiceNatPools,
		}
		interfaces = append(interfaces, s)
	}
	return interfaces, nil
}

func (c *natCollector) collectForPoolInterface(s *NatPoolInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Interface}...)

	for _, pool := range s.ServiceNatPools {
		lp := append(l, []string{pool.Name, pool.TranslationType, pool.PortRange, pool.PortBlockType}...)

		ch <- prometheus.MustNewConstMetric(portBlockSizeDesc, prometheus.GaugeValue, float64(pool.PortBlockSize), lp...)
		ch <- prometheus.MustNewConstMetric(activeBlockTimeoutDesc, prometheus.GaugeValue, float64(pool.ActiveBlockTimeout), lp...)
		ch <- prometheus.MustNewConstMetric(maxBlocksPerAddressDesc, prometheus.GaugeValue, float64(pool.MaxBlocksPerAddress), lp...)
		ch <- prometheus.MustNewConstMetric(effectivePortBlocksDesc, prometheus.GaugeValue, float64(pool.EffectivePortBlocks), lp...)
		ch <- prometheus.MustNewConstMetric(effectivePortsDesc, prometheus.GaugeValue, float64(pool.EffectivePorts), lp...)
		ch <- prometheus.MustNewConstMetric(portBlockEfficiencyDesc, prometheus.GaugeValue, float64(pool.PortBlockEfficiency), lp...)
	}
}

func (c *natCollector) NatPoolDetailInterfaces(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) ([]*NatPoolDetailInterface, error) {
	var x = NatPoolDetailRpc{}
	err := client.RunCommandAndParse("show services nat pool detail", &x)
	if err != nil {
		return nil, err
	}

	interfacesdetail := make([]*NatPoolDetailInterface, 0)
	for _, natpooldetailinterface := range x.Information.Interfaces {
		s := &NatPoolDetailInterface{
			Interface:             natpooldetailinterface.Interface,
			ServiceSetName:        natpooldetailinterface.ServiceSetName,
			ServiceNatPoolsDetail: natpooldetailinterface.ServiceNatPoolsDetail,
		}
		interfacesdetail = append(interfacesdetail, s)
	}
	return interfacesdetail, nil
}

func (c *natCollector) collectForPoolDetailInterface(s *NatPoolDetailInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Interface}...)

	for _, pool := range s.ServiceNatPoolsDetail {
		lp := append(l, []string{pool.Name, pool.TranslationType, pool.PortRange, pool.PortBlockType}...)
		ch <- prometheus.MustNewConstMetric(portsInUseDesc, prometheus.GaugeValue, float64(pool.PortsInUse), lp...)
		ch <- prometheus.MustNewConstMetric(outOfPortErrorsDesc, prometheus.GaugeValue, float64(pool.OutOfPortErrors), lp...)
		ch <- prometheus.MustNewConstMetric(parityPortErrorsDesc, prometheus.GaugeValue, float64(pool.ParityPortErrors), lp...)
		ch <- prometheus.MustNewConstMetric(preserveRangeErrorsDesc, prometheus.GaugeValue, float64(pool.PreserveRangeErrors), lp...)
		ch <- prometheus.MustNewConstMetric(maxPortsInUseDesc, prometheus.GaugeValue, float64(pool.MaxPortsInUse), lp...)
		ch <- prometheus.MustNewConstMetric(appPortErrorsDesc, prometheus.GaugeValue, float64(pool.AppPortErrors), lp...)
		ch <- prometheus.MustNewConstMetric(appExceedPortLimitErrorsDesc, prometheus.GaugeValue, float64(pool.AppExceedPortLimitErrors), lp...)
		ch <- prometheus.MustNewConstMetric(memAllocErrorsDesc, prometheus.GaugeValue, float64(pool.MemAllocErrors), lp...)
		ch <- prometheus.MustNewConstMetric(maxPortBlocksUsedDesc, prometheus.GaugeValue, float64(pool.MaxPortBlocksUsed), lp...)
		ch <- prometheus.MustNewConstMetric(blocksInUseDesc, prometheus.GaugeValue, float64(pool.BlocksInUse), lp...)
		ch <- prometheus.MustNewConstMetric(blockAllocationErrorsDesc, prometheus.GaugeValue, float64(pool.BlockAllocationErrors), lp...)
		ch <- prometheus.MustNewConstMetric(blocksLimitExceededErrorsDesc, prometheus.GaugeValue, float64(pool.BlocksLimitExceededErrors), lp...)
		ch <- prometheus.MustNewConstMetric(usersDesc, prometheus.GaugeValue, float64(pool.Users), lp...)
		ch <- prometheus.MustNewConstMetric(eifInboundSessionCountDesc, prometheus.GaugeValue, float64(pool.EifInboundSessionCount), lp...)
		ch <- prometheus.MustNewConstMetric(eifInboundLimitExceedDropDesc, prometheus.GaugeValue, float64(pool.EifInboundLimitExceedDrop), lp...)
	}
}


func (c *natCollector) ServiceSetsCpuInterfaces(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) ([]*ServiceSetsCpuInterface, error) {
	var x = ServiceSetsCpuRpc{}
	err := client.RunCommandAndParse("show services service-sets cpu-usage", &x)
	if err != nil {
		return nil, err
	}

	interfacesdetail := make([]*ServiceSetsCpuInterface, 0)
	for _, servicesetscpuinterface := range x.Information.Interfaces {
		s := &ServiceSetsCpuInterface{
			Interface:             servicesetscpuinterface.Interface,
			ServiceSetName:        servicesetscpuinterface.ServiceSetName,
			CpuUtilizationPercent: servicesetscpuinterface.CpuUtilizationPercent,
		}
		interfacesdetail = append(interfacesdetail, s)
	}
	return interfacesdetail, nil
}

func (c *natCollector) collectForServiceSetsCpuInterface(s *ServiceSetsCpuInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Interface, s.ServiceSetName}...)

	ch <- prometheus.MustNewConstMetric(serviceSetCpuUtilizationDesc, prometheus.GaugeValue, float64(s.CpuUtilizationPercent), l...)
}
