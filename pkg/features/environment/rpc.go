// SPDX-License-Identifier: MIT

package environment

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name           `xml:"rpc-reply"`
	Results multiEngineResults `xml:"multi-routing-engine-results"`
}

type multiEngineResults struct {
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name                            string                          `xml:"re-name"`
	EnvironmentComponentInformation environmentComponentInformation `xml:"environment-component-information"`
	EnvironmentInformation          environmentInformation          `xml:"environment-information"`
}

type environmentComponentInformation struct {
	EnvironmentComponentItem []environmentComponentItem `xml:"environment-component-item"`
}

type environmentComponentItem struct {
	Name            string `xml:"name"`
	State           string `xml:"state"`
	FanSpeedReading []struct {
		FanName  string `xml:"fan-name"`
		FanSpeed string `xml:"fan-speed"`
	} `xml:"fan-speed-reading"`
	DcInformation struct {
		DcDetail struct {
			DcVoltage     float64 `xml:"dc-voltage,omitempty"`
			DcCurrent     float64 `xml:"dc-current,omitempty"`
			DcPower       float64 `xml:"dc-power,omitempty"`
			DcLoad        float64 `xml:"dc-load,omitempty"`
			Str3DcVoltage float64 `xml:"str3-dc-voltage,omitempty"`
		} `xml:"dc-detail"`
	} `xml:"dc-information"`
}

type environmentInformation struct {
	Items []environmentItem `xml:"environment-item"`
}

type environmentItem struct {
	Name        string `xml:"name"`
	Class       string `xml:"class"`
	Status      string `xml:"status"`
	Temperature *struct {
		Value float64 `xml:"celsius,attr"`
	} `xml:"temperature,omitempty"`
}

type singleEngineResult struct {
	XMLName                         xml.Name                        `xml:"rpc-reply"`
	EnvironmentComponentInformation environmentComponentInformation `xml:"environment-component-information"`
	EnvironmentInformation          environmentInformation          `xml:"environment-information"`
}

type showVersionResult struct {
	XMLName             xml.Name `xml:"rpc-reply"`
	Text                string   `xml:",chardata"`
	Junos               string   `xml:"junos,attr"`
	SoftwareInformation struct {
		Text               string `xml:",chardata"`
		HostName           string `xml:"host-name"`
		ProductModel       string `xml:"product-model"`
		ProductName        string `xml:"product-name"`
		OsName             string `xml:"os-name"`
		JunosVersion       string `xml:"junos-version"`
		PackageInformation []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			PackageName string `xml:"package-name"`
			Comment     string `xml:"comment"`
		} `xml:"package-information"`
	} `xml:"software-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type environmentResultSomeSwitches struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Text    string   `xml:",chardata"`
	Junos   string   `xml:"junos,attr"`
	EnvironmentInformation struct {
		Text  string `xml:",chardata"`
		Xmlns string `xml:"xmlns,attr"`
		EnvironmentItem []struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name"`
			Status string `xml:"status"`
			Class  string `xml:"class"`
			Temperature struct {
				Text    string `xml:",chardata"`
				Celsius string `xml:"celsius,attr"`
			} `xml:"temperature"`
			Comment string `xml:"comment"`
		} `xml:"environment-item"`
	} `xml:"environment-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type multiEngineResultSomeSwitches struct {
	XMLName                         xml.Name `xml:"rpc-reply"`
	Text                            string   `xml:",chardata"`
	Junos                           string   `xml:"junos,attr"`
	EnvironmentComponentInformation struct {
		Text                     string `xml:",chardata"`
		Xmlns                    string `xml:"xmlns,attr"`
		EnvironmentComponentItem []struct {
			Text           string `xml:",chardata"`
			Name           string `xml:"name"`
			State          string `xml:"state"`
			PsmInformation struct {
				Text               string `xml:",chardata"`
				TemperatureReading struct {
					Text            string `xml:",chardata"`
					TemperatureName string `xml:"temperature-name"`
					Temperature     struct {
						Text    string `xml:",chardata"`
						Celsius string `xml:"celsius,attr"`
					} `xml:"temperature"`
				} `xml:"temperature-reading"`
				PsmStatus struct {
					Text     string `xml:",chardata"`
					Fans     string `xml:"fans"`
					DcOutput string `xml:"dc-output"`
				} `xml:"psm-status"`
				FirmwareVersion    string `xml:"firmware-version"`
				FanSpeedReadingPsm struct {
					Text      string `xml:",chardata"`
					Fan1Name  string `xml:"fan1-name"`
					Fan1Speed string `xml:"fan1-speed"`
					Fan2Name  string `xml:"fan2-name"`
				} `xml:"fan-speed-reading-psm"`
			} `xml:"psm-information"`
			PsmHealthCheckDetail struct {
				Text                       string `xml:",chardata"`
				HealthCheckStatus          string `xml:"health-check-status"`
				HealthCheckStateStr        string `xml:"health-check-state-str"`
				HealthCheckLastResultStr   string `xml:"health-check-last-result-str"`
				HealthCheckLastExecStr     string `xml:"health-check-last-exec-str"`
				HealthCheckNextSchedRunStr string `xml:"health-check-next-sched-run-str"`
			} `xml:"psm-health-check-detail"`
		} `xml:"environment-component-item"`
	} `xml:"environment-component-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
