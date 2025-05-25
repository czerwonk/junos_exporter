// SPDX-License-Identifier: MIT

package twamp

type result struct {
	Results struct {
		Probes []probe `xml:"probe-test-results"`
	} `xml:"probe-results"`
}

// ProbeTestResults holds the details of a single probe test
type probe struct {
	Owner                   string                    `xml:"owner-name"`
	Test                    string                    `xml:"test-name"`
	SourceAddress           string                    `xml:"source-address"`
	TargetAddress           string                    `xml:"target-address"`
	Type                    string                    `xml:"test-type"`
	Size                    int64                     `xml:"test-size"`
	GenericSampleResults    GenericSampleResults      `xml:"generic-sample-results"`    // Assuming one per test result
	GenericAggregateResults []GenericAggregateResults `xml:"generic-aggregate-results"` // Slice for multiple aggregate results
}

// GenericSampleResults holds individual sample data
type GenericSampleResults struct {
	SampleStatus  string `xml:"sample-status"`
	SampleTxTime  string `xml:"sample-tx-time"` // Consider using time.Time with custom unmarshalling if needed
	SampleRxTime  string `xml:"sample-rx-time"` // Consider using time.Time with custom unmarshalling if needed
	OffloadStatus string `xml:"offload-status"`
	RTT           int64  `xml:"rtt"`
	RTTJitter     int64  `xml:"rtt-jitter"`
}

// GenericAggregateResults holds aggregated data for different periods
type GenericAggregateResults struct {
	AggregateType               string                        `xml:"aggregate-type"`
	NumSamplesTx                int64                         `xml:"num-samples-tx"`
	NumSamplesRx                int64                         `xml:"num-samples-rx"`
	LossPercentage              float64                       `xml:"loss-percentage"`
	GenericAggregateMeasurement []GenericAggregateMeasurement `xml:"generic-aggregate-measurement"` // Slice for multiple measurements
}

// GenericAggregateMeasurement holds specific measurement statistics
type GenericAggregateMeasurement struct {
	MeasurementType    string `xml:"measurement-type"`
	MeasurementSamples int64  `xml:"measurement-samples"`
	MeasurementMin     int64  `xml:"measurement-min"`
	MeasurementMax     int64  `xml:"measurement-max"`
	MeasurementAvg     int64  `xml:"measurement-avg"`
	MeasurementStddev  int64  `xml:"measurement-stddev"`
}
