package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/oldweipro/gin-admin/global"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type AwsS3 struct{}

func (s *AwsS3) UploadUrl(fileUrl, filename string) (string, string, error) {
	//TODO implement me
	panic("implement me")
}
func (s *AwsS3) UploadBase64(base64Str, filename string) (string, string, error) {
	//TODO implement me
	panic("implement me")
}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: UploadFile
//@description: Upload file to Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-upload-file-to-bucket
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	s := newSession()
	uploader := s3manager.NewUploader(s)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := global.ConfigServer.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		global.Logger.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			global.Logger.Error("创建文件关闭流失败", zap.Error(err))
		}
	}(f) // 创建文件 defer 关闭

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(global.ConfigServer.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		global.Logger.Error("function uploader.Upload() Filed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return global.ConfigServer.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: DeleteFile
//@description: Delete file from Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-delete-bucket-item
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) DeleteFile(key string) error {
	s := newSession()
	svc := s3.New(s)
	filename := global.ConfigServer.AwsS3.PathPrefix + "/" + key
	bucket := global.ConfigServer.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		global.Logger.Error("function svc.DeleteObject() Filed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() Filed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

// newSession Create S3 session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(global.ConfigServer.AwsS3.Region),
		Endpoint:         aws.String(global.ConfigServer.AwsS3.Endpoint), //minio在这里设置地址,可以兼容
		S3ForcePathStyle: aws.Bool(global.ConfigServer.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(global.ConfigServer.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			global.ConfigServer.AwsS3.SecretID,
			global.ConfigServer.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
