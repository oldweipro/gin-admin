package request

type LoginQrcode struct {
	QrcodeKey string `json:"qrcodeKey" form:"qrcodeKey"`
}
