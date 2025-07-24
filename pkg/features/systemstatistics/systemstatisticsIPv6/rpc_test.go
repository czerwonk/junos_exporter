package systemstatisticsIPv6

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsIPv6Unmarshaling_AllNumericalFields(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected map[string]float64
	}{
		{
			name: "all_fields_nonzero",
			xmlInput: `<rpc-reply>
	<statistics>
	  <ip6>
		<total-packets-received>1</total-packets-received>
		<ip6-packets-with-size-smaller-than-minimum>2</ip6-packets-with-size-smaller-than-minimum>
		<packets-with-datasize-less-than-data-length>3</packets-with-datasize-less-than-data-length>
		<ip6-packets-with-bad-options>4</ip6-packets-with-bad-options>
		<ip6-packets-with-incorrect-version-number>5</ip6-packets-with-incorrect-version-number>
		<ip6-fragments-received>6</ip6-fragments-received>
		<duplicate-or-out-of-space-fragments-dropped>7</duplicate-or-out-of-space-fragments-dropped>
		<ip6-fragments-dropped-after-timeout>8</ip6-fragments-dropped-after-timeout>
		<fragments-that-exceeded-limit>9</fragments-that-exceeded-limit>
		<ip6-packets-reassembled-ok>10</ip6-packets-reassembled-ok>
		<ip6-packets-for-this-host>11</ip6-packets-for-this-host>
		<ip6-packets-forwarded>12</ip6-packets-forwarded>
		<ip6-packets-not-forwardable>13</ip6-packets-not-forwardable>
		<ip6-redirects-sent>14</ip6-redirects-sent>
		<ip6-packets-sent-from-this-host>15</ip6-packets-sent-from-this-host>
		<ip6-packets-sent-with-fabricated-ip-header>16</ip6-packets-sent-with-fabricated-ip-header>
		<ip6-output-packets-dropped-due-to-no-bufs>17</ip6-output-packets-dropped-due-to-no-bufs>
		<ip6-output-packets-discarded-due-to-no-route>18</ip6-output-packets-discarded-due-to-no-route>
		<ip6-output-datagrams-fragmented>19</ip6-output-datagrams-fragmented>
		<ip6-fragments-created>20</ip6-fragments-created>
		<ip6-datagrams-that-can-not-be-fragmented>21</ip6-datagrams-that-can-not-be-fragmented>
		<packets-that-violated-scope-rules>22</packets-that-violated-scope-rules>
		<multicast-packets-which-we-do-not-join>23</multicast-packets-which-we-do-not-join>
		<ip6nh-tcp>24</ip6nh-tcp>
		<ip6nh-udp>25</ip6nh-udp>
		<ip6nh-icmp6>26</ip6nh-icmp6>
	  </ip6>
	</statistics>
  </rpc-reply>`,
			expected: map[string]float64{
				"TotalPacketsReceived":                  1,
				"Ip6PacketsWithSizeSmallerThanMinimum":  2,
				"PacketsWithDatasizeLessThanDataLength": 3,
				"Ip6PacketsWithBadOptions":              4,
				"Ip6PacketsWithIncorrectVersionNumber":  5,
				"Ip6FragmentsReceived":                  6,
				"DuplicateOrOutOfSpaceFragmentsDropped": 7,
				"Ip6FragmentsDroppedAfterTimeout":       8,
				"FragmentsThatExceededLimit":            9,
				"Ip6PacketsReassembledOk":               10,
				"Ip6PacketsForThisHost":                 11,
				"Ip6PacketsForwarded":                   12,
				"Ip6PacketsNotForwardable":              13,
				"Ip6RedirectsSent":                      14,
				"Ip6PacketsSentFromThisHost":            15,
				"Ip6PacketsSentWithFabricatedIpHeader":  16,
				"Ip6OutputPacketsDroppedDueToNoBufs":    17,
				"Ip6OutputPacketsDiscardedDueToNoRoute": 18,
				"Ip6OutputDatagramsFragmented":          19,
				"Ip6FragmentsCreated":                   20,
				"Ip6DatagramsThatCanNotBeFragmented":    21,
				"PacketsThatViolatedScopeRules":         22,
				"MulticastPacketsWhichWeDoNotJoin":      23,
				"Ip6nhTcp":                              24,
				"Ip6nhUdp":                              25,
				"Ip6nhIcmp6":                            26,
			},
		},
		{
			name: "all_fields_zero",
			xmlInput: `<rpc-reply>
	<statistics>
	  <ip6>
		<total-packets-received>0</total-packets-received>
		<ip6-packets-with-size-smaller-than-minimum>0</ip6-packets-with-size-smaller-than-minimum>
		<packets-with-datasize-less-than-data-length>0</packets-with-datasize-less-than-data-length>
		<ip6-packets-with-bad-options>0</ip6-packets-with-bad-options>
		<ip6-packets-with-incorrect-version-number>0</ip6-packets-with-incorrect-version-number>
		<ip6-fragments-received>0</ip6-fragments-received>
		<duplicate-or-out-of-space-fragments-dropped>0</duplicate-or-out-of-space-fragments-dropped>
		<ip6-fragments-dropped-after-timeout>0</ip6-fragments-dropped-after-timeout>
		<fragments-that-exceeded-limit>0</fragments-that-exceeded-limit>
		<ip6-packets-reassembled-ok>0</ip6-packets-reassembled-ok>
		<ip6-packets-for-this-host>0</ip6-packets-for-this-host>
		<ip6-packets-forwarded>0</ip6-packets-forwarded>
		<ip6-packets-not-forwardable>0</ip6-packets-not-forwardable>
		<ip6-redirects-sent>0</ip6-redirects-sent>
		<ip6-packets-sent-from-this-host>0</ip6-packets-sent-from-this-host>
		<ip6-packets-sent-with-fabricated-ip-header>0</ip6-packets-sent-with-fabricated-ip-header>
		<ip6-output-packets-dropped-due-to-no-bufs>0</ip6-output-packets-dropped-due-to-no-bufs>
		<ip6-output-packets-discarded-due-to-no-route>0</ip6-output-packets-discarded-due-to-no-route>
		<ip6-output-datagrams-fragmented>0</ip6-output-datagrams-fragmented>
		<ip6-fragments-created>0</ip6-fragments-created>
		<ip6-datagrams-that-can-not-be-fragmented>0</ip6-datagrams-that-can-not-be-fragmented>
		<packets-that-violated-scope-rules>0</packets-that-violated-scope-rules>
		<multicast-packets-which-we-do-not-join>0</multicast-packets-which-we-do-not-join>
		<ip6nh-tcp>0</ip6nh-tcp>
		<ip6nh-udp>0</ip6nh-udp>
		<ip6nh-icmp6>0</ip6nh-icmp6>
	  </ip6>
	</statistics>
  </rpc-reply>`,
			expected: map[string]float64{
				"TotalPacketsReceived":                  0,
				"Ip6PacketsWithSizeSmallerThanMinimum":  0,
				"PacketsWithDatasizeLessThanDataLength": 0,
				"Ip6PacketsWithBadOptions":              0,
				"Ip6PacketsWithIncorrectVersionNumber":  0,
				"Ip6FragmentsReceived":                  0,
				"DuplicateOrOutOfSpaceFragmentsDropped": 0,
				"Ip6FragmentsDroppedAfterTimeout":       0,
				"FragmentsThatExceededLimit":            0,
				"Ip6PacketsReassembledOk":               0,
				"Ip6PacketsForThisHost":                 0,
				"Ip6PacketsForwarded":                   0,
				"Ip6PacketsNotForwardable":              0,
				"Ip6RedirectsSent":                      0,
				"Ip6PacketsSentFromThisHost":            0,
				"Ip6PacketsSentWithFabricatedIpHeader":  0,
				"Ip6OutputPacketsDroppedDueToNoBufs":    0,
				"Ip6OutputPacketsDiscardedDueToNoRoute": 0,
				"Ip6OutputDatagramsFragmented":          0,
				"Ip6FragmentsCreated":                   0,
				"Ip6DatagramsThatCanNotBeFragmented":    0,
				"PacketsThatViolatedScopeRules":         0,
				"MulticastPacketsWhichWeDoNotJoin":      0,
				"Ip6nhTcp":                              0,
				"Ip6nhUdp":                              0,
				"Ip6nhIcmp6":                            0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual StatisticsIPv6
			err := xml.Unmarshal([]byte(tt.xmlInput), &actual)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected["TotalPacketsReceived"], actual.Statistics.Ip6.TotalPacketsReceived, "TotalPacketsReceived mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsWithSizeSmallerThanMinimum"], actual.Statistics.Ip6.Ip6PacketsWithSizeSmallerThanMinimum, "Ip6PacketsWithSizeSmallerThanMinimum mismatch")
			assert.Equal(t, tt.expected["PacketsWithDatasizeLessThanDataLength"], actual.Statistics.Ip6.PacketsWithDatasizeLessThanDataLength, "PacketsWithDatasizeLessThanDataLength mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsWithBadOptions"], actual.Statistics.Ip6.Ip6PacketsWithBadOptions, "Ip6PacketsWithBadOptions mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsWithIncorrectVersionNumber"], actual.Statistics.Ip6.Ip6PacketsWithIncorrectVersionNumber, "Ip6PacketsWithIncorrectVersionNumber mismatch")
			assert.Equal(t, tt.expected["Ip6FragmentsReceived"], actual.Statistics.Ip6.Ip6FragmentsReceived, "Ip6FragmentsReceived mismatch")
			assert.Equal(t, tt.expected["DuplicateOrOutOfSpaceFragmentsDropped"], actual.Statistics.Ip6.DuplicateOrOutOfSpaceFragmentsDropped, "DuplicateOrOutOfSpaceFragmentsDropped mismatch")
			assert.Equal(t, tt.expected["Ip6FragmentsDroppedAfterTimeout"], actual.Statistics.Ip6.Ip6FragmentsDroppedAfterTimeout, "Ip6FragmentsDroppedAfterTimeout mismatch")
			assert.Equal(t, tt.expected["FragmentsThatExceededLimit"], actual.Statistics.Ip6.FragmentsThatExceededLimit, "FragmentsThatExceededLimit mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsReassembledOk"], actual.Statistics.Ip6.Ip6PacketsReassembledOk, "Ip6PacketsReassembledOk mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsForThisHost"], actual.Statistics.Ip6.Ip6PacketsForThisHost, "Ip6PacketsForThisHost mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsForwarded"], actual.Statistics.Ip6.Ip6PacketsForwarded, "Ip6PacketsForwarded mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsNotForwardable"], actual.Statistics.Ip6.Ip6PacketsNotForwardable, "Ip6PacketsNotForwardable mismatch")
			assert.Equal(t, tt.expected["Ip6RedirectsSent"], actual.Statistics.Ip6.Ip6RedirectsSent, "Ip6RedirectsSent mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsSentFromThisHost"], actual.Statistics.Ip6.Ip6PacketsSentFromThisHost, "Ip6PacketsSentFromThisHost mismatch")
			assert.Equal(t, tt.expected["Ip6PacketsSentWithFabricatedIpHeader"], actual.Statistics.Ip6.Ip6PacketsSentWithFabricatedIpHeader, "Ip6PacketsSentWithFabricatedIpHeader mismatch")
			assert.Equal(t, tt.expected["Ip6OutputPacketsDroppedDueToNoBufs"], actual.Statistics.Ip6.Ip6OutputPacketsDroppedDueToNoBufs, "Ip6OutputPacketsDroppedDueToNoBufs mismatch")
			assert.Equal(t, tt.expected["Ip6OutputPacketsDiscardedDueToNoRoute"], actual.Statistics.Ip6.Ip6OutputPacketsDiscardedDueToNoRoute, "Ip6OutputPacketsDiscardedDueToNoRoute mismatch")
			assert.Equal(t, tt.expected["Ip6OutputDatagramsFragmented"], actual.Statistics.Ip6.Ip6OutputDatagramsFragmented, "Ip6OutputDatagramsFragmented mismatch")
			assert.Equal(t, tt.expected["Ip6FragmentsCreated"], actual.Statistics.Ip6.Ip6FragmentsCreated, "Ip6FragmentsCreated mismatch")
			assert.Equal(t, tt.expected["Ip6DatagramsThatCanNotBeFragmented"], actual.Statistics.Ip6.Ip6DatagramsThatCanNotBeFragmented, "Ip6DatagramsThatCanNotBeFragmented mismatch")
			assert.Equal(t, tt.expected["PacketsThatViolatedScopeRules"], actual.Statistics.Ip6.PacketsThatViolatedScopeRules, "PacketsThatViolatedScopeRules mismatch")
			assert.Equal(t, tt.expected["MulticastPacketsWhichWeDoNotJoin"], actual.Statistics.Ip6.MulticastPacketsWhichWeDoNotJoin, "MulticastPacketsWhichWeDoNotJoin mismatch")
			assert.Equal(t, tt.expected["Ip6nhTcp"], actual.Statistics.Ip6.Ip6nhTcp, "Ip6nhTcp mismatch")
			assert.Equal(t, tt.expected["Ip6nhUdp"], actual.Statistics.Ip6.Ip6nhUdp, "Ip6nhUdp mismatch")
			assert.Equal(t, tt.expected["Ip6nhIcmp6"], actual.Statistics.Ip6.Ip6nhIcmp6, "Ip6nhIcmp6 mismatch")
		})
	}
}