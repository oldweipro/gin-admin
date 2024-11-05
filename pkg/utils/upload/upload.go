package upload

import (
	"github.com/oldweipro/gin-admin/pkg/app"
	"mime/multipart"
)

// OSS 对象存储接口
// Author [oldweipro](https://github.com/oldweipro)
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [oldweipro](https://github.com/oldweipro)
func NewOss() OSS {
	switch app.Config.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	case "cloudflare-r2":
		return &CloudflareR2{}
	default:
		return &Local{}
	}
}
