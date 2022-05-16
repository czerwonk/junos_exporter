package mpls_lsp

type mpls_lspRpc struct {
	Information struct {
		Sessions []mpls_lspSession `xml:"rsvp-session"`
	} `xml:"mpls-lsp-information>rsvp-session-data"`
}

type mpls_lspSession struct {
	DstIP              string    `xml:"mpls-lsp>destination-address"`
	SrcIP              string    `xml:"mpls-lsp>source-address"`
	LSPState           string    `xml:"mpls-lsp>lsp-state"`
	Name               string    `xml:"mpls-lsp>name"`

	Path               []mpls_lspPath  `xml:"mpls-lsp>mpls-lsp-path"`
}

type mpls_lspPath struct {
	Title             string    `xml:"title"`
	Name              string    `xml:"name"`
	State             string    `xml:"path-state"`
	FlapCount         int64     `xml:"path-flap-count"`
}

