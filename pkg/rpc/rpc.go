// SPDX-License-Identifier: MIT

package rpc

import (
	"encoding/xml"
)

type rpcReply struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Body    []byte   `xml:",innerxml"`
}
