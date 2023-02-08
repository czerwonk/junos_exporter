// SPDX-License-Identifier: MIT

package nat2

import (
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_nat2_statistics_"

var (
	natTotalSessionInterestDesc *prometheus.Desc

	NatPktDstInNatRouteDesc                        *prometheus.Desc
	NatFilteringSessionDesc                        *prometheus.Desc
	NatMappingSessionDesc                          *prometheus.Desc
	NatRuleLookupFailuresDesc                      *prometheus.Desc
	NatMapAllocationSuccessesDesc                  *prometheus.Desc
	NatMapAllocationFailuresDesc                   *prometheus.Desc
	NatMapFreeSuccessDesc                          *prometheus.Desc
	NatMapFreeFailuresDesc                         *prometheus.Desc
	NatEimMappingCreateFailedDesc                  *prometheus.Desc
	NatEimMappingCreatedDesc                       *prometheus.Desc
	NatEimMappingUpdatedDesc                       *prometheus.Desc
	NatEifMappingFreeDesc                          *prometheus.Desc
	NatEimMappingFreeDesc                          *prometheus.Desc
	NatTotalPktsProcessedDesc                      *prometheus.Desc
	NatTotalPktsForwardedDesc                      *prometheus.Desc
	NatTotalPktsTranslatedDesc                     *prometheus.Desc
	Nat64MtuExceedDesc                             *prometheus.Desc
	Nat64DfbitSetDesc                              *prometheus.Desc
	Nat64ErrMtuExceedBuildDesc                     *prometheus.Desc
	Nat64ErrMtuExceedSendDesc                      *prometheus.Desc
	SessionXlate464ClatPrefixNotFoundDesc          *prometheus.Desc
	SessionXlate464EmbededIpv4NotFoundDesc         *prometheus.Desc
	NatJflowLogAllocFailDesc                       *prometheus.Desc
	NatJflowLogAllocSuccessDesc                    *prometheus.Desc
	NatJflowLogFreeSuccessDesc                     *prometheus.Desc
	NatJflowLogFreeFailRecordDesc                  *prometheus.Desc
	NatJflowLogFreeFailDataDesc                    *prometheus.Desc
	NatJflowLogInvalidTransTypeDesc                *prometheus.Desc
	NatJflowLogFreeSuccessFailQueuingDesc          *prometheus.Desc
	NatJflowLogInvalidInputArgsDesc                *prometheus.Desc
	NatJflowLogInvalidAllocErrDesc                 *prometheus.Desc
	NatJflowLogRateLimitFailGetPoolDesc            *prometheus.Desc
	NatJflowLogRateLimitFailGetServiceSetDesc      *prometheus.Desc
	NatJflowLogRateLimitFailInvalidCurrentTimeDesc *prometheus.Desc

	PoolNameDesc                *prometheus.Desc
	PoolIDDesc                  *prometheus.Desc
	PortTranslationDesc         *prometheus.Desc
	PortOverloadingFactorDesc   *prometheus.Desc
	AddressAssignementDesc      *prometheus.Desc
	ClearAlarmThresholdDesc     *prometheus.Desc
	RaiseAlarmThresholdDesc     *prometheus.Desc
	TotalPoolAddressDesc        *prometheus.Desc
	AddressPoolHitsDesc         *prometheus.Desc
	BlkSizeDesc                 *prometheus.Desc
	BlkMaxPerHostDesc           *prometheus.Desc
	BlkAtvTimeoutDesc           *prometheus.Desc
	BlkInterimLogCycleDesc      *prometheus.Desc
	BlkLogDesc                  *prometheus.Desc
	BlkUsedDesc                 *prometheus.Desc
	BlkTotalDesc                *prometheus.Desc
	PortBlkEfficiencyDesc       *prometheus.Desc
	MaxBlkUsedDesc              *prometheus.Desc
	UsersDesc                   *prometheus.Desc
	EimTimeoutDesc              *prometheus.Desc
	MappingTimeoutDesc          *prometheus.Desc
	EifInboundFlowsCountDesc    *prometheus.Desc
	EifFlowLimitExceedDropsDesc *prometheus.Desc
	SinglePortSumDesc           *prometheus.Desc

	SinglePortDesc *prometheus.Desc

	OutOfPortErrorDesc          *prometheus.Desc
	OutOfAddrErrorDesc          *prometheus.Desc
	ParityPortErrorDesc         *prometheus.Desc
	PreserveRangeErrorDesc      *prometheus.Desc
	AppOutOfPortErrorDesc       *prometheus.Desc
	AppExceedPortLimitErrorDesc *prometheus.Desc
	OutOfBlkErrorDesc           *prometheus.Desc
	BlkExceedLimitErrorDesc     *prometheus.Desc
	BlkOutOfPortErrorDesc       *prometheus.Desc
	BlkMemAllocErrorDesc        *prometheus.Desc

	serviceSetCPUUtilizationDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "interface"}

	lp := []string{"target", "interface", "service_set", "pool_name", "pool_id"}
	lar := []string{"target", "interface", "service_set", "pool_name", "pool_id", "address_range_low", "address_range_high"}

	lservicesets := []string{"target", "interface", "service_set"}

	natTotalSessionInterestDesc = prometheus.NewDesc(prefix+"nat_total_session_interest", "Total Session Interest events", l, nil)

	NatPktDstInNatRouteDesc = prometheus.NewDesc(prefix+"nat_pkt_dst_in_nat_route", "Packet  Dst in NAT route", l, nil)
	NatFilteringSessionDesc = prometheus.NewDesc(prefix+"nat_filtering_session", "Session Created for EIF", l, nil)
	NatMappingSessionDesc = prometheus.NewDesc(prefix+"nat_mapping_session", "Session Created for EIM", l, nil)
	NatRuleLookupFailuresDesc = prometheus.NewDesc(prefix+"nat_rule_lookup_failures", "NAT rule lookup failures", l, nil)
	NatMapAllocationSuccessesDesc = prometheus.NewDesc(prefix+"nat_map_allocation_successes", "NAT allocation Successes", l, nil)
	NatMapAllocationFailuresDesc = prometheus.NewDesc(prefix+"nat_map_allocation_failures", "NAT allocation Failures", l, nil)
	NatMapFreeSuccessDesc = prometheus.NewDesc(prefix+"nat_map_free_success", "NAT Free Successes", l, nil)
	NatMapFreeFailuresDesc = prometheus.NewDesc(prefix+"nat_map_free_failures", "NAT Free Failures", l, nil)
	NatEimMappingCreateFailedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_create_failed", "NAT EIM mapping create failed", l, nil)
	NatEimMappingCreatedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_created", "NAT EIM mapping Created", l, nil)
	NatEimMappingUpdatedDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_updated", "NAT EIM mapping Updated", l, nil)
	NatEifMappingFreeDesc = prometheus.NewDesc(prefix+"nat_eif_mapping_free", "NAT EIF mapping Free", l, nil)
	NatEimMappingFreeDesc = prometheus.NewDesc(prefix+"nat_eim_mapping_free", "NAT EIM mapping Free", l, nil)
	NatTotalPktsProcessedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_processed", "Total Packets Processed", l, nil)
	NatTotalPktsForwardedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_forwarded", "Total Packets Forwarded", l, nil)
	NatTotalPktsTranslatedDesc = prometheus.NewDesc(prefix+"nat_total_pkts_translated", "Total Packets Translated", l, nil)
	Nat64MtuExceedDesc = prometheus.NewDesc(prefix+"nat64_mtu_exceed", "NAT64 - MTU exceeded", l, nil)
	Nat64DfbitSetDesc = prometheus.NewDesc(prefix+"nat64_dfbit_set", "NAT64 - dfbit set", l, nil)
	Nat64ErrMtuExceedBuildDesc = prometheus.NewDesc(prefix+"nat64_err_mtu_exceed_build", "NAT64 error - MTU exceed build", l, nil)
	Nat64ErrMtuExceedSendDesc = prometheus.NewDesc(prefix+"nat64_err_mtu_exceed_send", "NAT64 error - MTU exceed send", l, nil)
	SessionXlate464ClatPrefixNotFoundDesc = prometheus.NewDesc(prefix+"session_xlate464_clat_prefix_not_found", "Session xlate464 clat prefix not found", l, nil)
	SessionXlate464EmbededIpv4NotFoundDesc = prometheus.NewDesc(prefix+"session_xlate464_embeded_ipv4_not_found", "Session xlate464 embeded ipv4 not found", l, nil)
	NatJflowLogAllocFailDesc = prometheus.NewDesc(prefix+"nat_jflow_log_alloc_fail", "NAT jflow-log error - memory allocation fail", l, nil)
	NatJflowLogAllocSuccessDesc = prometheus.NewDesc(prefix+"nat_jflow_log_alloc_success", "NAT jflow-log - memory allocation success", l, nil)
	NatJflowLogFreeSuccessDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_success", "NAT jflow-log - memory free success", l, nil)
	NatJflowLogFreeFailRecordDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_fail_record", "NAT jflow-log error - memory free fail null record", l, nil)
	NatJflowLogFreeFailDataDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_fail_data", "NAT jflow-log error - memory free fail null data", l, nil)
	NatJflowLogInvalidTransTypeDesc = prometheus.NewDesc(prefix+"nat_jflow_log_invalid_trans_type", "NAT jflow-log error - invalid nat translation type", l, nil)
	NatJflowLogFreeSuccessFailQueuingDesc = prometheus.NewDesc(prefix+"nat_jflow_log_free_success_fail_queuing", "NAT jflow-log - memory free success fail queuing", l, nil)
	NatJflowLogInvalidInputArgsDesc = prometheus.NewDesc(prefix+"nat_jflow_log_invalid_input_args", "NAT jflow-log - invalid input arguments", l, nil)
	NatJflowLogInvalidAllocErrDesc = prometheus.NewDesc(prefix+"nat_jflow_log_invalid_alloc_err", "NAT jflow-log - invalid allocation error type", l, nil)
	NatJflowLogRateLimitFailGetPoolDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_get_pool", "NAT jflow-log - rate limit fail to get pool", l, nil)
	NatJflowLogRateLimitFailGetServiceSetDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_get_service_set", "NAT jflow-log - rate limit fail to get service set", l, nil)
	NatJflowLogRateLimitFailInvalidCurrentTimeDesc = prometheus.NewDesc(prefix+"nat_jflow_log_rate_limit_fail_invalid_current_time", "NAT jflow-log - rate limit fail invalid current time", l, nil)

	PoolNameDesc = prometheus.NewDesc(prefix+"pool_name", "Pool name", l, nil)
	PoolIDDesc = prometheus.NewDesc(prefix+"pool_id", "Pool id", l, nil)
	PortTranslationDesc = prometheus.NewDesc(prefix+"source_pool_port_translation", "Port", l, nil)

	PortOverloadingFactorDesc = prometheus.NewDesc(prefix+"port_overloading_factor", "Port overloading", lp, nil)
	AddressAssignementDesc = prometheus.NewDesc(prefix+"source_pool_address_assignment", "Address assignment", lp, nil)
	ClearAlarmThresholdDesc = prometheus.NewDesc(prefix+"clear_alarm_threshold", "Alarm threshold", lp, nil)
	RaiseAlarmThresholdDesc = prometheus.NewDesc(prefix+"raise_alarm_threshold", "Alarm threshold", lp, nil)
	TotalPoolAddressDesc = prometheus.NewDesc(prefix+"total_pool_address", "Total addresses", lp, nil)
	AddressPoolHitsDesc = prometheus.NewDesc(prefix+"address_pool_hits", "Translation Hits", lp, nil)
	BlkSizeDesc = prometheus.NewDesc(prefix+"source_pool_blk_size", "Port block size", lp, nil)
	BlkMaxPerHostDesc = prometheus.NewDesc(prefix+"source_pool_blk_max_per_host", "Max blocks per host", lp, nil)
	BlkAtvTimeoutDesc = prometheus.NewDesc(prefix+"source_pool_blk_atv_timeout", "Active block timeout", lp, nil)
	BlkInterimLogCycleDesc = prometheus.NewDesc(prefix+"source_pool_blk_interim_log_cycle", "Interim logging interval", lp, nil)
	BlkLogDesc = prometheus.NewDesc(prefix+"source_pool_blk_log", "PBA block log", lp, nil)
	BlkUsedDesc = prometheus.NewDesc(prefix+"source_pool_blk_used", "Used port blocks", lp, nil)
	BlkTotalDesc = prometheus.NewDesc(prefix+"source_pool_blk_total", "total port blocks", lp, nil)
	PortBlkEfficiencyDesc = prometheus.NewDesc(prefix+"source_pool_port_blk_efficiency", "Port block efficiency", lp, nil)
	MaxBlkUsedDesc = prometheus.NewDesc(prefix+"source_pool_max_blk_used", "Max number of port blocks used", lp, nil)
	UsersDesc = prometheus.NewDesc(prefix+"source_pool_users", "Unique pool users", lp, nil)
	EimTimeoutDesc = prometheus.NewDesc(prefix+"source_pool_eim_timeout", "Ei_mapping_timeout", lp, nil)
	MappingTimeoutDesc = prometheus.NewDesc(prefix+"source_pool_mapping_timeout", "Mapping_timeout", lp, nil)
	EifInboundFlowsCountDesc = prometheus.NewDesc(prefix+"source_pool_eif_inbound_flows_count", "EIF Inbound session count", lp, nil)
	EifFlowLimitExceedDropsDesc = prometheus.NewDesc(prefix+"source_pool_eif_flow_limit_exceed_drops", "EIF Inbound session limit exceeded drops", lp, nil)

	SinglePortDesc = prometheus.NewDesc(prefix+"single_port", "Ports", lar, nil)
	SinglePortSumDesc = prometheus.NewDesc(prefix+"single_port_sum", "Total used ports", lp, nil)

	OutOfPortErrorDesc = prometheus.NewDesc(prefix+"out_of_port_error", "Out of port errors", lp, nil)
	OutOfAddrErrorDesc = prometheus.NewDesc(prefix+"out_of_addr_error", "Out of address errors", lp, nil)
	ParityPortErrorDesc = prometheus.NewDesc(prefix+"parity_port_error", "Parity port errors", lp, nil)
	PreserveRangeErrorDesc = prometheus.NewDesc(prefix+"preserve_range_error", "Preserve Range errors", lp, nil)
	AppOutOfPortErrorDesc = prometheus.NewDesc(prefix+"app_out_of_port_error", "APP port allocation errors", lp, nil)
	AppExceedPortLimitErrorDesc = prometheus.NewDesc(prefix+"app_exceed_port_limit_error", "APP port limit allocation errors", lp, nil)
	OutOfBlkErrorDesc = prometheus.NewDesc(prefix+"out_of_blk_error", "Port block allocation errors", lp, nil)
	BlkExceedLimitErrorDesc = prometheus.NewDesc(prefix+"blk_exceed_limit_error", "Port blocks limit exceeded errors", lp, nil)
	BlkOutOfPortErrorDesc = prometheus.NewDesc(prefix+"blk_out_of_port_error", "Port blocks out of port errors", lp, nil)
	BlkMemAllocErrorDesc = prometheus.NewDesc(prefix+"blk_mem_alloc_error", "Port blocks memory alloc errors", lp, nil)

	serviceSetCPUUtilizationDesc = prometheus.NewDesc(prefix+"service_set_cpu_utilization", "CPU utilization for the Service Set", lservicesets, nil)

}

type natCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &natCollector{}
}

// Name returns the name of the collector
func (*natCollector) Name() string {
	return "NAT2"
}

// Describe describes the metrics
func (*natCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- natTotalSessionInterestDesc
}

// Collect collects metrics from JunOS
func (c *natCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	interfaces, err := c.natInterfaces(client)
	if err != nil {
		return err
	}
	for _, s := range interfaces {
		c.collectForInterface(s, ch, labelValues)
	}

	poolinterfaces, err := c.srcNatPools(client, ch, labelValues)
	if err != nil {
		return err
	}
	c.collectForSrcNatPool(poolinterfaces, ch, labelValues)

	servicesetscpuinterfaces, err := c.serviceSetsCPUInterfaces(client, ch, labelValues)
	for _, s := range servicesetscpuinterfaces {
		c.collectForServiceSetsCPUInterface(s, ch, labelValues)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *natCollector) natInterfaces(client *rpc.Client) ([]*iface, error) {
	var x = result{}
	err := client.RunCommandAndParse("show services nat statistics", &x)
	if err != nil {
		return nil, err
	}

	interfaces := make([]*iface, 0)
	for _, natinterface := range x.Interfaces {
		s := &iface{
			Interface:                                  natinterface.Interface,
			NatPktDstInNatRoute:                        int64(natinterface.NatPktDstInNatRoute),
			NatFilteringSession:                        int64(natinterface.NatFilteringSession),
			NatMappingSession:                          int64(natinterface.NatMappingSession),
			NatRuleLookupFailures:                      int64(natinterface.NatRuleLookupFailures),
			NatMapAllocationSuccesses:                  int64(natinterface.NatMapAllocationSuccesses),
			NatMapAllocationFailures:                   int64(natinterface.NatMapAllocationFailures),
			NatMapFreeSuccess:                          int64(natinterface.NatMapFreeSuccess),
			NatMapFreeFailures:                         int64(natinterface.NatMapFreeFailures),
			NatEimMappingCreateFailed:                  int64(natinterface.NatEimMappingCreateFailed),
			NatEimMappingCreated:                       int64(natinterface.NatEimMappingCreated),
			NatEimMappingUpdated:                       int64(natinterface.NatEimMappingUpdated),
			NatEifMappingFree:                          int64(natinterface.NatEifMappingFree),
			NatEimMappingFree:                          int64(natinterface.NatEimMappingFree),
			NatTotalPktsProcessed:                      int64(natinterface.NatTotalPktsProcessed),
			NatTotalPktsForwarded:                      int64(natinterface.NatTotalPktsForwarded),
			NatTotalPktsTranslated:                     int64(natinterface.NatTotalPktsTranslated),
			Nat64MtuExceed:                             int64(natinterface.Nat64MtuExceed),
			Nat64DfbitSet:                              int64(natinterface.Nat64DfbitSet),
			Nat64ErrMtuExceedBuild:                     int64(natinterface.Nat64ErrMtuExceedBuild),
			Nat64ErrMtuExceedSend:                      int64(natinterface.Nat64ErrMtuExceedSend),
			SessionXlate464ClatPrefixNotFound:          int64(natinterface.SessionXlate464ClatPrefixNotFound),
			SessionXlate464EmbededIpv4NotFound:         int64(natinterface.SessionXlate464EmbededIpv4NotFound),
			NatJflowLogAllocFail:                       int64(natinterface.NatJflowLogAllocFail),
			NatJflowLogAllocSuccess:                    int64(natinterface.NatJflowLogAllocSuccess),
			NatJflowLogFreeSuccess:                     int64(natinterface.NatJflowLogFreeSuccess),
			NatJflowLogFreeFailRecord:                  int64(natinterface.NatJflowLogFreeFailRecord),
			NatJflowLogFreeFailData:                    int64(natinterface.NatJflowLogFreeFailData),
			NatJflowLogInvalidTransType:                int64(natinterface.NatJflowLogInvalidTransType),
			NatJflowLogFreeSuccessFailQueuing:          int64(natinterface.NatJflowLogFreeSuccessFailQueuing),
			NatJflowLogInvalidInputArgs:                int64(natinterface.NatJflowLogInvalidInputArgs),
			NatJflowLogInvalidAllocErr:                 int64(natinterface.NatJflowLogInvalidAllocErr),
			NatJflowLogRateLimitFailGetPool:            int64(natinterface.NatJflowLogRateLimitFailGetPool),
			NatJflowLogRateLimitFailGetServiceSet:      int64(natinterface.NatJflowLogRateLimitFailGetServiceSet),
			NatJflowLogRateLimitFailInvalidCurrentTime: int64(natinterface.NatJflowLogRateLimitFailInvalidCurrentTime),
		}

		interfaces = append(interfaces, s)
	}

	return interfaces, nil
}

func (*natCollector) collectForInterface(s *iface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Interface}...)

	ch <- prometheus.MustNewConstMetric(NatPktDstInNatRouteDesc, prometheus.GaugeValue, float64(s.NatPktDstInNatRoute), l...)
	ch <- prometheus.MustNewConstMetric(NatFilteringSessionDesc, prometheus.GaugeValue, float64(s.NatFilteringSession), l...)
	ch <- prometheus.MustNewConstMetric(NatMappingSessionDesc, prometheus.GaugeValue, float64(s.NatMappingSession), l...)
	ch <- prometheus.MustNewConstMetric(NatRuleLookupFailuresDesc, prometheus.GaugeValue, float64(s.NatRuleLookupFailures), l...)
	ch <- prometheus.MustNewConstMetric(NatMapAllocationSuccessesDesc, prometheus.GaugeValue, float64(s.NatMapAllocationSuccesses), l...)
	ch <- prometheus.MustNewConstMetric(NatMapAllocationFailuresDesc, prometheus.GaugeValue, float64(s.NatMapAllocationFailures), l...)
	ch <- prometheus.MustNewConstMetric(NatMapFreeSuccessDesc, prometheus.GaugeValue, float64(s.NatMapFreeSuccess), l...)
	ch <- prometheus.MustNewConstMetric(NatMapFreeFailuresDesc, prometheus.GaugeValue, float64(s.NatMapFreeFailures), l...)
	ch <- prometheus.MustNewConstMetric(NatEimMappingCreateFailedDesc, prometheus.GaugeValue, float64(s.NatEimMappingCreateFailed), l...)
	ch <- prometheus.MustNewConstMetric(NatEimMappingCreatedDesc, prometheus.GaugeValue, float64(s.NatEimMappingCreated), l...)
	ch <- prometheus.MustNewConstMetric(NatEimMappingUpdatedDesc, prometheus.GaugeValue, float64(s.NatEimMappingUpdated), l...)
	ch <- prometheus.MustNewConstMetric(NatEifMappingFreeDesc, prometheus.GaugeValue, float64(s.NatEifMappingFree), l...)
	ch <- prometheus.MustNewConstMetric(NatEimMappingFreeDesc, prometheus.GaugeValue, float64(s.NatEimMappingFree), l...)
	ch <- prometheus.MustNewConstMetric(NatTotalPktsProcessedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsProcessed), l...)
	ch <- prometheus.MustNewConstMetric(NatTotalPktsForwardedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsForwarded), l...)
	ch <- prometheus.MustNewConstMetric(NatTotalPktsTranslatedDesc, prometheus.GaugeValue, float64(s.NatTotalPktsTranslated), l...)
	ch <- prometheus.MustNewConstMetric(Nat64MtuExceedDesc, prometheus.GaugeValue, float64(s.Nat64MtuExceed), l...)
	ch <- prometheus.MustNewConstMetric(Nat64DfbitSetDesc, prometheus.GaugeValue, float64(s.Nat64DfbitSet), l...)
	ch <- prometheus.MustNewConstMetric(Nat64ErrMtuExceedBuildDesc, prometheus.GaugeValue, float64(s.Nat64ErrMtuExceedBuild), l...)
	ch <- prometheus.MustNewConstMetric(Nat64ErrMtuExceedSendDesc, prometheus.GaugeValue, float64(s.Nat64ErrMtuExceedSend), l...)

	ch <- prometheus.MustNewConstMetric(SessionXlate464ClatPrefixNotFoundDesc, prometheus.GaugeValue, float64(s.SessionXlate464ClatPrefixNotFound), l...)
	ch <- prometheus.MustNewConstMetric(SessionXlate464EmbededIpv4NotFoundDesc, prometheus.GaugeValue, float64(s.SessionXlate464EmbededIpv4NotFound), l...)

	ch <- prometheus.MustNewConstMetric(NatJflowLogAllocFailDesc, prometheus.GaugeValue, float64(s.NatJflowLogAllocFail), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogAllocSuccessDesc, prometheus.GaugeValue, float64(s.NatJflowLogAllocSuccess), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogFreeSuccessDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeSuccess), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogFreeFailRecordDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeFailRecord), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogFreeFailDataDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeFailData), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogInvalidTransTypeDesc, prometheus.GaugeValue, float64(s.NatJflowLogInvalidTransType), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogFreeSuccessFailQueuingDesc, prometheus.GaugeValue, float64(s.NatJflowLogFreeSuccessFailQueuing), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogInvalidInputArgsDesc, prometheus.GaugeValue, float64(s.NatJflowLogInvalidInputArgs), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogInvalidAllocErrDesc, prometheus.GaugeValue, float64(s.NatJflowLogInvalidAllocErr), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogRateLimitFailGetPoolDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailGetPool), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogRateLimitFailGetServiceSetDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailGetServiceSet), l...)
	ch <- prometheus.MustNewConstMetric(NatJflowLogRateLimitFailInvalidCurrentTimeDesc, prometheus.GaugeValue, float64(s.NatJflowLogRateLimitFailInvalidCurrentTime), l...)
}

func (c *natCollector) srcNatPools(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) ([]srcNatPool, error) {
	var x = srcNatPoolResult{}
	err := client.RunCommandAndParse("show services nat source pool all", &x)
	if err != nil {
		return nil, err
	}

	return x.Information.Pools[:], nil
}

// func (c *natCollector) collectForSrcNatPool(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) {
func (c *natCollector) collectForSrcNatPool(s []srcNatPool, ch chan<- prometheus.Metric, labelValues []string) {

	for _, pool := range s {
		lp := append(labelValues, []string{pool.Interface, pool.ServiceSetName, pool.PoolName, pool.PoolID}...)

		ch <- prometheus.MustNewConstMetric(PortOverloadingFactorDesc, prometheus.GaugeValue, float64(pool.PortOverloadingFactor), lp...)
		ch <- prometheus.MustNewConstMetric(TotalPoolAddressDesc, prometheus.GaugeValue, float64(pool.TotalPoolAddress), lp...)
		ch <- prometheus.MustNewConstMetric(AddressPoolHitsDesc, prometheus.GaugeValue, float64(pool.AddressPoolHits), lp...)
		ch <- prometheus.MustNewConstMetric(BlkSizeDesc, prometheus.GaugeValue, float64(pool.BlkSize), lp...)
		ch <- prometheus.MustNewConstMetric(BlkMaxPerHostDesc, prometheus.GaugeValue, float64(pool.BlkMaxPerHost), lp...)
		ch <- prometheus.MustNewConstMetric(BlkAtvTimeoutDesc, prometheus.GaugeValue, float64(pool.BlkAtvTimeout), lp...)
		ch <- prometheus.MustNewConstMetric(BlkInterimLogCycleDesc, prometheus.GaugeValue, float64(pool.BlkInterimLogCycle), lp...)

		ch <- prometheus.MustNewConstMetric(BlkUsedDesc, prometheus.GaugeValue, float64(pool.BlkUsed), lp...)
		ch <- prometheus.MustNewConstMetric(BlkTotalDesc, prometheus.GaugeValue, float64(pool.BlkTotal), lp...)

		fpercentage, err := strconv.ParseFloat(strings.Trim(pool.PortBlkEfficiency, "%"), 64)
		if err != nil {
			fpercentage = float64(-1)
		}

		ch <- prometheus.MustNewConstMetric(PortBlkEfficiencyDesc, prometheus.GaugeValue, float64(fpercentage), lp...)

		ch <- prometheus.MustNewConstMetric(MaxBlkUsedDesc, prometheus.GaugeValue, float64(pool.MaxBlkUsed), lp...)
		ch <- prometheus.MustNewConstMetric(UsersDesc, prometheus.GaugeValue, float64(pool.Users), lp...)
		ch <- prometheus.MustNewConstMetric(EimTimeoutDesc, prometheus.GaugeValue, float64(pool.EimTimeout), lp...)
		ch <- prometheus.MustNewConstMetric(MappingTimeoutDesc, prometheus.GaugeValue, float64(pool.MappingTimeout), lp...)
		ch <- prometheus.MustNewConstMetric(EifInboundFlowsCountDesc, prometheus.GaugeValue, float64(pool.EifInboundFlowsCount), lp...)
		ch <- prometheus.MustNewConstMetric(EifFlowLimitExceedDropsDesc, prometheus.GaugeValue, float64(pool.EifFlowLimitExceedDrops), lp...)

		//  not working because the xml parsing is wrong.
		//		for _, par := range pool.SrcPoolAddressRanges {
		//			lar := append(lp, []string{par.AddressRangeLow, par.AddressRangeHigh}...)
		//			ch <- prometheus.MustNewConstMetric(SinglePortDesc, prometheus.GaugeValue, float64(par.SinglePort), lar...)
		//		}

		for i, arl := range pool.SrcPoolAddressRanges.AddressRangeLow {
			arh := pool.SrcPoolAddressRanges.AddressRangeHigh[i]
			sp := pool.SrcPoolAddressRanges.SinglePort[i]
			lar := append(lp, []string{arl, arh}...)
			ch <- prometheus.MustNewConstMetric(SinglePortDesc, prometheus.GaugeValue, float64(sp), lar...)
		}

		ch <- prometheus.MustNewConstMetric(SinglePortSumDesc, prometheus.GaugeValue, float64(pool.SrcPoolAddressRangeSum.SinglePortSum), lp...)

		ch <- prometheus.MustNewConstMetric(OutOfPortErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.OutOfPortError), lp...)
		ch <- prometheus.MustNewConstMetric(OutOfAddrErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.OutOfAddrError), lp...)
		ch <- prometheus.MustNewConstMetric(ParityPortErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.ParityPortError), lp...)
		ch <- prometheus.MustNewConstMetric(PreserveRangeErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.PreserveRangeError), lp...)
		ch <- prometheus.MustNewConstMetric(AppOutOfPortErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.AppOutOfPortError), lp...)
		ch <- prometheus.MustNewConstMetric(AppExceedPortLimitErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.AppExceedPortLimitError), lp...)
		ch <- prometheus.MustNewConstMetric(OutOfBlkErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.OutOfBlkError), lp...)
		ch <- prometheus.MustNewConstMetric(BlkExceedLimitErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.BlkExceedLimitError), lp...)
		ch <- prometheus.MustNewConstMetric(BlkOutOfPortErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.BlkOutOfPortError), lp...)
		ch <- prometheus.MustNewConstMetric(BlkMemAllocErrorDesc, prometheus.GaugeValue, float64(pool.SrcPoolErrorCounters.BlkMemAllocError), lp...)
	}
}

func (c *natCollector) serviceSetsCPUInterfaces(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) ([]*serviceSetsCPUInterface, error) {
	var x = serviceSetsCPUResult{}
	err := client.RunCommandAndParse("show services service-sets cpu-usage", &x)
	if err != nil {
		return nil, err
	}

	interfacesdetail := make([]*serviceSetsCPUInterface, 0)
	for _, servicesetscpuinterface := range x.Information.Interfaces {
		s := &serviceSetsCPUInterface{
			Interface:             servicesetscpuinterface.Interface,
			ServiceSetName:        servicesetscpuinterface.ServiceSetName,
			CPUUtilizationPercent: servicesetscpuinterface.CPUUtilizationPercent,
		}
		interfacesdetail = append(interfacesdetail, s)
	}
	return interfacesdetail, nil
}

func (c *natCollector) collectForServiceSetsCPUInterface(s *serviceSetsCPUInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Interface, s.ServiceSetName}...)

	ch <- prometheus.MustNewConstMetric(serviceSetCPUUtilizationDesc, prometheus.GaugeValue, float64(s.CPUUtilizationPercent), l...)
}
