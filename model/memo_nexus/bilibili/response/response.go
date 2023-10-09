package response

type CommonResponse[T QrcodePoll | LoginQrcodeGenerate | ProfileData] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    T      `json:"data"`
}

type QrcodePoll struct {
	Url          string `json:"url" form:"url"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	Timestamp    uint64 `json:"timestamp" form:"timestamp"`
	Code         int    `json:"code" form:"code"`
	Message      string `json:"message" form:"message"`
}

type LoginQrcodeGenerate struct {
	Url       string `json:"url" form:"url"`
	QrcodeKey string `json:"qrcode_key" form:"qrcode_key"`
}
