package ladder

type Settings struct {
	Clients                   []Clients `json:"clients"`
	DisableInsecureEncryption bool      `json:"disableInsecureEncryption"`
}

type Clients struct {
	Id      string `json:"id"`
	AlterId int    `json:"alterId"`
}

type StreamSettings struct {
	Network     string      `json:"network"`
	Security    string      `json:"security"`
	TlsSettings TlsSettings `json:"tlsSettings"`
	TcpSettings TcpSettings `json:"tcpSettings"`
}

type TlsSettings struct {
	ServerName   string         `json:"serverName"`
	Certificates []Certificates `json:"certificates"`
}

type Certificates struct {
	CertificateFile string `json:"certificateFile"`
	KeyFile         string `json:"keyFile"`
}
type TcpSettings struct {
	Header Header `json:"header"`
}

type Header struct {
	Type string `json:"type"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
}
