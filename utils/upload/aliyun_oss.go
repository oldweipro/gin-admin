package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/oldweipro/gin-admin/global"
	"go.uber.org/zap"
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		global.Logger.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		global.Logger.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.ConfigServer.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		global.Logger.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.ConfigServer.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}
func (*AliyunOSS) UploadUrl(fileUrl, filename string) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		global.Logger.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 创建自定义的http.Client
	client := &http.Client{
		Transport: &http.Transport{
			// 设置Transport字段为自定义Transport，包含代理设置
			Proxy: func(req *http.Request) (*url.URL, error) {
				// 设置代理
				proxyURL, err := url.Parse("http://127.0.0.1:7890")
				if err != nil {
					return nil, err
				}
				return proxyURL, nil
			},
		},
	}

	// 上传文件流。
	resp, imgErr := client.Get(fileUrl)

	defer resp.Body.Close()
	if imgErr != nil {
		fmt.Println("同步图片下载失败: ", imgErr)
		return "", "", imgErr
	}

	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	if filename == "" {
		imgName := strings.Split(fileUrl, "?")[0]
		tokens := strings.Split(imgName, "/")
		filename = tokens[len(tokens)-1]
	}
	yunFileTmpPath := global.ConfigServer.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + filename

	err = bucket.PutObject(yunFileTmpPath, resp.Body)
	if err != nil {
		global.Logger.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.ConfigServer.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		global.Logger.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		global.Logger.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(global.ConfigServer.AliyunOSS.Endpoint, global.ConfigServer.AliyunOSS.AccessKeyId, global.ConfigServer.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.ConfigServer.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
