package upload

import (
	"go.uber.org/zap"
	"mime/multipart"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/oldweipro/gin-admin/global"
	"github.com/pkg/errors"
)

var HuaWeiObs = new(Obs)

type Obs struct{}

func (o *Obs) UploadUrl(fileUrl, filename string) (string, string, error) {
	//TODO implement me
	panic("implement me")
}
func (o *Obs) UploadBase64(base64Str, filename string) (string, string, error) {
	//TODO implement me
	panic("implement me")
}

func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.ConfigServer.HuaWeiObs.AccessKey, global.ConfigServer.HuaWeiObs.SecretKey, global.ConfigServer.HuaWeiObs.Endpoint)
}

func (o *Obs) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// var open multipart.File
	open, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			global.Logger.Error("上传文件关闭流失败", zap.Error(err))
		}
	}(open)
	filename := file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.ConfigServer.HuaWeiObs.Bucket,
				Key:    filename,
			},
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = NewHuaWeiObsClient()
	if err != nil {
		return "", "", errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return "", "", errors.Wrap(err, "文件上传失败!")
	}
	filepath := global.ConfigServer.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}

func (o *Obs) DeleteFile(key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.ConfigServer.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}
