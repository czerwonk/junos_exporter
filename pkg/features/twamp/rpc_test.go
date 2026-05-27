// SPDX-License-Identifier: MIT

package twamp

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwamp(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S4.11-EVO/junos">
    <probe-results>
        <probe-test-results>
            <owner-name>TWAMP</owner-name>
            <test-name>RTR1_TT</test-name>
            <source-address>192.168.54.44</source-address>
            <target-address>192.168.54.99</target-address>
            <test-type>twamp</test-type>
            <test-size>30</test-size>
            <generic-sample-results>
                <sample-status>Probe response received</sample-status>
                <sample-tx-time>05/24/25 18:57:12.530240</sample-tx-time>
                <sample-rx-time>05/24/25 18:57:12.530884</sample-rx-time>
                <offload-status>Client and server offload timestamping</offload-status>
                <rtt>121</rtt>
                <rtt-jitter>34</rtt-jitter>
                <egress-jitter>20</egress-jitter>
                <ingress-jitter>14</ingress-jitter>
            </generic-sample-results>
            <generic-aggregate-results>
                <aggregate-type>current test</aggregate-type>
                <num-samples-tx>12</num-samples-tx>
                <num-samples-rx>12</num-samples-rx>
                <loss-percentage>0.00</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>12</measurement-samples>
                    <measurement-min>73</measurement-min>
                    <measurement-max>157</measurement-max>
                    <measurement-avg>121</measurement-avg>
                    <measurement-stddev>31</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>11</measurement-min>
                    <measurement-max>82</measurement-max>
                    <measurement-avg>40</measurement-avg>
                    <measurement-stddev>28</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress delay (usec)</measurement-type>
                    <measurement-samples>12</measurement-samples>
                    <measurement-min>68</measurement-min>
                    <measurement-max>134</measurement-max>
                    <measurement-avg>100</measurement-avg>
                    <measurement-stddev>22</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>14</measurement-min>
                    <measurement-max>62</measurement-max>
                    <measurement-avg>29</measurement-avg>
                    <measurement-stddev>16</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress delay (usec)</measurement-type>
                    <measurement-samples>12</measurement-samples>
                    <measurement-min>5</measurement-min>
                    <measurement-max>59</measurement-max>
                    <measurement-avg>21</measurement-avg>
                    <measurement-stddev>16</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>54</measurement-max>
                    <measurement-avg>20</measurement-avg>
                    <measurement-stddev>20</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
            <generic-aggregate-results> 
                <aggregate-type>last test</aggregate-type>
                <num-samples-tx>30</num-samples-tx>
                <num-samples-rx>30</num-samples-rx>
                <loss-percentage>0.00</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>30</measurement-samples>
                    <measurement-min>76</measurement-min>
                    <measurement-max>194</measurement-max>
                    <measurement-avg>129</measurement-avg>
                    <measurement-stddev>37</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>113</measurement-max>
                    <measurement-avg>41</measurement-avg>
                    <measurement-stddev>32</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress delay (usec)</measurement-type>
                    <measurement-samples>30</measurement-samples>
                    <measurement-min>68</measurement-min>
                    <measurement-max>157</measurement-max>
                    <measurement-avg>103</measurement-avg>
                    <measurement-stddev>30</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>89</measurement-max>
                    <measurement-avg>27</measurement-avg>
                    <measurement-stddev>23</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress delay (usec)</measurement-type>
                    <measurement-samples>30</measurement-samples>
                    <measurement-min>7</measurement-min>
                    <measurement-max>70</measurement-max>
                    <measurement-avg>26</measurement-avg>
                    <measurement-stddev>15</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>60</measurement-max>
                    <measurement-avg>17</measurement-avg>
                    <measurement-stddev>18</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
            <generic-aggregate-results>
                <aggregate-type>all tests</aggregate-type>
                <num-samples-tx>203652</num-samples-tx>
                <num-samples-rx>203502</num-samples-rx>
                <loss-percentage>0.07</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>203502</measurement-samples>
                    <measurement-min>63</measurement-min>
                    <measurement-max>336</measurement-max>
                    <measurement-avg>120</measurement-avg>
                    <measurement-stddev>35</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>203501</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>228</measurement-max>
                    <measurement-avg>31</measurement-avg>
                    <measurement-stddev>27</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress delay (usec)</measurement-type>
                    <measurement-samples>60339</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>241</measurement-max>
                    <measurement-avg>73</measurement-avg>
                    <measurement-stddev>41</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>203501</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>196</measurement-max>
                    <measurement-avg>22</measurement-avg>
                    <measurement-stddev>21</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress delay (usec)</measurement-type>
                    <measurement-samples>60339</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>263</measurement-max>
                    <measurement-avg>53</measurement-avg>
                    <measurement-stddev>39</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>203501</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>240</measurement-max>
                    <measurement-avg>15</measurement-avg>
                    <measurement-stddev>18</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
        </probe-test-results>
        <probe-test-results>
            <owner-name>TWAMP</owner-name>
            <test-name>RTR1_ZZ</test-name>
            <source-address>192.168.54.33</source-address>
            <target-address>192.168.54.44</target-address>
            <test-type>twamp</test-type>
            <test-size>30</test-size>
            <generic-sample-results>
                <sample-status>Probe response received</sample-status>
                <sample-tx-time>05/24/25 18:57:12.530394</sample-tx-time>
                <sample-rx-time>05/24/25 18:57:12.530976</sample-rx-time>
                <offload-status>Client and server offload timestamping</offload-status>
                <rtt>139</rtt>
                <rtt-jitter>17</rtt-jitter>
            </generic-sample-results>
            <generic-aggregate-results>
                <aggregate-type>current test</aggregate-type>
                <num-samples-tx>12</num-samples-tx>
                <num-samples-rx>12</num-samples-rx>
                <loss-percentage>0.00</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>12</measurement-samples>
                    <measurement-min>83</measurement-min>
                    <measurement-max>156</measurement-max>
                    <measurement-avg>119</measurement-avg>
                    <measurement-stddev>27</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>63</measurement-max>
                    <measurement-avg>27</measurement-avg>
                    <measurement-stddev>21</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>56</measurement-max>
                    <measurement-avg>22</measurement-avg>
                    <measurement-stddev>20</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>4</measurement-min>
                    <measurement-max>36</measurement-max>
                    <measurement-avg>12</measurement-avg>
                    <measurement-stddev>9</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
            <generic-aggregate-results>
                <aggregate-type>last test</aggregate-type>
                <num-samples-tx>30</num-samples-tx>
                <num-samples-rx>30</num-samples-rx>
                <loss-percentage>0.00</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>30</measurement-samples>
                    <measurement-min>72</measurement-min>
                    <measurement-max>224</measurement-max>
                    <measurement-avg>122</measurement-avg>
                    <measurement-stddev>38</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>139</measurement-max>
                    <measurement-avg>35</measurement-avg>
                    <measurement-stddev>35</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>67</measurement-max>
                    <measurement-avg>21</measurement-avg>
                    <measurement-stddev>18</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>72</measurement-max>
                    <measurement-avg>21</measurement-avg>
                    <measurement-stddev>21</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
            <generic-aggregate-results>
                <aggregate-type>all tests</aggregate-type>
                <num-samples-tx>203622</num-samples-tx>
                <num-samples-rx>203501</num-samples-rx>
                <loss-percentage>0.06</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>203501</measurement-samples>
                    <measurement-min>65</measurement-min>
                    <measurement-max>322</measurement-max>
                    <measurement-avg>121</measurement-avg>
                    <measurement-stddev>35</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>203500</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>243</measurement-max>
                    <measurement-avg>32</measurement-avg>
                    <measurement-stddev>28</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress delay (usec)</measurement-type>
                    <measurement-samples>48575</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>238</measurement-max>
                    <measurement-avg>62</measurement-avg>
                    <measurement-stddev>41</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>203500</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>198</measurement-max>
                    <measurement-avg>21</measurement-avg>
                    <measurement-stddev>21</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress delay (usec)</measurement-type>
                    <measurement-samples>48575</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>258</measurement-max>
                    <measurement-avg>67</measurement-avg>
                    <measurement-stddev>42</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>203500</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>172</measurement-max>
                    <measurement-avg>16</measurement-avg>
                    <measurement-stddev>19</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
        </probe-test-results>
        <probe-test-results>
            <owner-name>TWAMP</owner-name>
            <test-name>PE2_ZZ</test-name>
            <source-address>192.168.54.21</source-address>
            <target-address>192.168.54.12</target-address>
            <test-type>twamp</test-type>
            <test-size>30</test-size>
            <generic-sample-results>
                <sample-status>Probe response received</sample-status>
                <sample-tx-time>05/24/25 18:57:12.530297</sample-tx-time>
                <sample-rx-time>05/24/25 18:57:12.530749</sample-rx-time>
                <offload-status>Client and server offload timestamping</offload-status>
                <rtt>190</rtt>
                <rtt-jitter>91</rtt-jitter>
            </generic-sample-results>
            <generic-aggregate-results>
                <aggregate-type>current test</aggregate-type>
                <num-samples-tx>12</num-samples-tx>
                <num-samples-rx>12</num-samples-rx>
                <loss-percentage>0.00</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>12</measurement-samples>
                    <measurement-min>79</measurement-min>
                    <measurement-max>190</measurement-max>
                    <measurement-avg>134</measurement-avg>
                    <measurement-stddev>36</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>3</measurement-min>
                    <measurement-max>92</measurement-max>
                    <measurement-avg>46</measurement-avg>
                    <measurement-stddev>35</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>4</measurement-min>
                    <measurement-max>72</measurement-max>
                    <measurement-avg>33</measurement-avg>
                    <measurement-stddev>25</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>11</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>58</measurement-max>
                    <measurement-avg>15</measurement-avg>
                    <measurement-stddev>16</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
            <generic-aggregate-results>
                <aggregate-type>last test</aggregate-type>
                <num-samples-tx>30</num-samples-tx>
                <num-samples-rx>30</num-samples-rx>
                <loss-percentage>0.00</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>30</measurement-samples>
                    <measurement-min>78</measurement-min>
                    <measurement-max>216</measurement-max>
                    <measurement-avg>126</measurement-avg>
                    <measurement-stddev>37</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>138</measurement-max>
                    <measurement-avg>31</measurement-avg>
                    <measurement-stddev>31</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>81</measurement-max>
                    <measurement-avg>26</measurement-avg>
                    <measurement-stddev>23</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>29</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>57</measurement-max>
                    <measurement-avg>15</measurement-avg>
                    <measurement-stddev>15</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
            <generic-aggregate-results>
                <aggregate-type>all tests</aggregate-type>
                <num-samples-tx>203622</num-samples-tx>
                <num-samples-rx>203498</num-samples-rx>
                <loss-egress>2</loss-egress>
                <loss-percentage>0.06</loss-percentage>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip time (usec)</measurement-type>
                    <measurement-samples>203498</measurement-samples>
                    <measurement-min>65</measurement-min>
                    <measurement-max>8361</measurement-max>
                    <measurement-avg>118</measurement-avg>
                    <measurement-stddev>40</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Round trip jitter (usec)</measurement-type>
                    <measurement-samples>203497</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>8268</measurement-max>
                    <measurement-avg>31</measurement-avg>
                    <measurement-stddev>37</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress delay (usec)</measurement-type>
                    <measurement-samples>57837</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>8318</measurement-max>
                    <measurement-avg>63</measurement-avg>
                    <measurement-stddev>53</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Egress jitter (usec)</measurement-type>
                    <measurement-samples>203497</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>8252</measurement-max>
                    <measurement-avg>21</measurement-avg>
                    <measurement-stddev>32</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress delay (usec)</measurement-type>
                    <measurement-samples>57837</measurement-samples>
                    <measurement-min>1</measurement-min>
                    <measurement-max>233</measurement-max>
                    <measurement-avg>65</measurement-avg>
                    <measurement-stddev>39</measurement-stddev>
                </generic-aggregate-measurement>
                <generic-aggregate-measurement>
                    <measurement-type>Ingress jitter (usec)</measurement-type>
                    <measurement-samples>203497</measurement-samples>
                    <measurement-min>0</measurement-min>
                    <measurement-max>195</measurement-max>
                    <measurement-avg>15</measurement-avg>
                    <measurement-stddev>18</measurement-stddev>
                </generic-aggregate-measurement>
            </generic-aggregate-results>
        </probe-test-results>
    </probe-results>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	rpc := result{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	// test rtt
	assert.Equal(t, "192.168.54.99", rpc.Results.Probes[0].TargetAddress, "target-address")
	assert.Equal(t, int64(121), rpc.Results.Probes[0].GenericSampleResults.RTT, "rtt")
	//<measurement-max>194</measurement-max>
	assert.Equal(t, int64(194), rpc.Results.Probes[0].GenericAggregateResults[1].GenericAggregateMeasurement[0].MeasurementMax, "measurement-max")

}

func parseXML(b []byte, res *result) error {
	err := xml.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	return nil
}
