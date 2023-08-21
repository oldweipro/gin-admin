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

type XuiResponse struct {
	Success bool   `json:"success" form:"success"`
	Msg     string `json:"msg" form:"msg"`
	Obj     []Obj  `json:"obj" form:"obj"`
}

type Obj struct {
	Id             uint   `json:"id" form:"id"`
	Up             int64  `json:"up" form:"up"`
	Down           int64  `json:"down" form:"down"`
	Total          int64  `json:"total" form:"total"`
	Remark         string `json:"remark" form:"remark"`
	Enable         bool   `json:"enable" form:"enable"`
	ExpiryTime     int64  `json:"expiryTime" form:"expiryTime"`
	Listen         string `json:"listen" form:"listen"`
	Port           int64  `json:"port" form:"port"`
	Protocol       string `json:"protocol" form:"protocol"`
	Settings       string `json:"settings" form:"settings"`
	StreamSettings string `json:"streamSettings" form:"streamSettings"`
	Tag            string `json:"tag" form:"tag"`
	Sniffing       string `json:"sniffing" form:"sniffing"`
}
