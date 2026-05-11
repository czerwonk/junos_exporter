// SPDX-License-Identifier: MIT

package virtualchassis

import "encoding/xml"

type VirtualChassisReply struct {
    XMLName             xml.Name               `xml:"rpc-reply"`
    VirtChassInfo       VirtualChassisInfo     `xml:"virtual-chassis-information"`
}

type VirtualChassisInfo struct {
    // Junos uses different tags depending on the provisioning mode
    PreProvisVCInfo     VirtChassIDInfo           `xml:"preprovisioned-virtual-chassis-information"`
    NonProvisVCInfo     VirtChassIDInfo           `xml:"virtual-chassis-id-information"`
    MemberList          []VirtualChassisMember    `xml:"member-list>member"`
}

type VirtChassIDInfo struct {
    VirtChassID         string      `xml:"virtual-chassis-id"`
    VirtChassMode       string      `xml:"virtual-chassis-mode"`
}

type VirtualChassisMember struct {
    //XMLName              xml.Name   `xml:"member"`
    MemberID             int        `xml:"member-id"`
    MemberRole           string     `xml:"member-role"`
    MemberStatus         string     `xml:"member-status"`
    MemberSerial         string     `xml:"member-serial-number"`
    MemberModel          string     `xml:"member-model"`
    MemberPriority       int        `xml:"member-priority"`
    MastershipPriority	 int        `xml:"mastership-priority"`
    MemberMixedMode	 string     `xml:"member-mixed-mode"`
}


type VirtualChassisPortReply struct {
    XMLName    xml.Name                   `xml:"rpc-reply"`
    Results    MultiRoutingEngineResults  `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
    Items      []MultiRoutingEngineItem   `xml:"multi-routing-engine-item"`
}

type MultiRoutingEngineItem struct {
    Name          string                   `xml:"re-name"` // e.g., fpc0, fpc1
    VCPPortInfo   VirtualChassisPortInfo   `xml:"virtual-chassis-port-information"`
}

type VirtualChassisPortInfo struct {
    PortList   []VCPPort     `xml:"port-list>port-information"`
}

type VCPPort struct {
    PortName     string   `xml:"port-name"`   // e.g., vcp-0 or 0/0
    PortStatus   string   `xml:"port-status"`
    PortSpeed    string   `xml:"port-speed"`
    Neighbor     string   `xml:"neighbor-id"` // Member ID of the switch on the other end
}

