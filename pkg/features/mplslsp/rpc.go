package mplslsp

type result struct {
	Information struct {
		Sessions []lspSession `xml:"rsvp-session"`
	} `xml:"mpls-lsp-information>rsvp-session-data"`
}

type lspSession struct {
	DstIP    string `xml:"mpls-lsp>destination-address"`
	SrcIP    string `xml:"mpls-lsp>source-address"`
	LSPState string `xml:"mpls-lsp>lsp-state"`
	Name     string `xml:"mpls-lsp>name"`

	Path []lspPath `xml:"mpls-lsp>mpls-lsp-path"`
}

type lspPath struct {
	Title     string `xml:"title"`
	Name      string `xml:"name"`
	State     string `xml:"path-state"`
	FlapCount int64  `xml:"path-flap-count"`
}
