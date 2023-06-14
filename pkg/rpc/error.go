// SPDX-License-Identifier: MIT

package rpc

type RpcReplyXmlParserError struct {
	RawResponse []byte
	Err         error
}

func (e *RpcReplyXmlParserError) Error() string {
	return "Failed to decode RPC reply XML: " + e.Err.Error()
}

func (e *RpcReplyXmlParserError) Unwrap() error { return e.Err }
