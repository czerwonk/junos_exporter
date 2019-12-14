package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		resp string
	}{
		{
			name: "no nulti routing engine",
			resp: `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.2R1/junos">
			<multi-routing-engine-results>
				<multi-routing-engine-item>
					<re-name>fpc0</re-name>
					<system-storage-information junos:style="brief">
						<filesystem>
							<filesystem-name>/dev/gpt/junos</filesystem-name>
							<total-blocks junos:format="1.3G">2796512</total-blocks>
							<used-blocks junos:format="814M">1667792</used-blocks>
							<available-blocks junos:format="442M">905000</available-blocks>
							<used-percent> 65</used-percent>
							<mounted-on>/.mount</mounted-on>
						</filesystem>
					</system-storage-information>
				</multi-routing-engine-item>
			</multi-routing-engine-results>
			<cli>
				<banner>{master:0}</banner>
			</cli>
		</rpc-reply>`,
		},
		{
			name: "multi routing engine",
			resp: `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.2R1/junos">
			<system-storage-information junos:style="brief">
				<filesystem>
					<filesystem-name>/dev/gpt/junos</filesystem-name>
					<total-blocks junos:format="1.3G">2796512</total-blocks>
					<used-blocks junos:format="814M">1667792</used-blocks>
					<available-blocks junos:format="442M">905000</available-blocks>
					<used-percent> 65</used-percent>
					<mounted-on>/.mount</mounted-on>
				</filesystem>#
		    </system-storage-information>
			<cli>
				<banner>{master:0}</banner>
			</cli>
		</rpc-reply>`,
		},
	}

	t.Parallel()

	for _, test := range tests {
		t.Run(test.name, func(te *testing.T) {
			rpc := MultiRoutingEngineResults{}
			err := parseXML([]byte(test.resp), &rpc)
			if err != nil {
				te.Fatal(err)
			}

			if assert.Equal(te, len(rpc.Results), 1, "Engines") {
				assert.Equal(te, len(rpc.Results[0].Storage.Information.Filesystems), 1, "Filesystems")
			}
		})
	}
}
