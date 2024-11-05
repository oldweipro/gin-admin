package system

import (
	"errors"
	"github.com/oldweipro/gin-admin/model/system"
	"github.com/oldweipro/gin-admin/pkg/app"
	"gorm.io/gorm"
)

//@author: [oldweipro](https://github.com/oldweipro)
//@function: CreateSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type DictionaryService struct{}

var DictionaryServiceApp = new(DictionaryService)

func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(app.DBClient.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = app.DBClient.Create(&sysDictionary).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: DeleteSysDictionary
//@description: 删除字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = app.DBClient.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = app.DBClient.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}

	if sysDictionary.SysDictionaryDetails != nil {
		return app.DBClient.Where("sys_dictionary_id=?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: UpdateSysDictionary
//@description: 更新字典数据
//@param: sysDictionary *model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	err = app.DBClient.Where("id = ?", sysDictionary.ID).First(&dict).Error
	if err != nil {
		app.Logger.Debug(err.Error())
		return errors.New("查询字典数据失败")
	}
	if dict.Type != sysDictionary.Type {
		if !errors.Is(app.DBClient.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = app.DBClient.Model(&dict).Updates(sysDictionaryMap).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetSysDictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDictionary model.SysDictionary

func (dictionaryService *DictionaryService) GetSysDictionary(Type string, Id uint, status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = app.DBClient.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true).Order("sort")
	}).First(&sysDictionary).Error
	return
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetSysDictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (dictionaryService *DictionaryService) GetSysDictionaryInfoList() (list interface{}, err error) {
	var sysDictionarys []system.SysDictionary
	err = app.DBClient.Find(&sysDictionarys).Error
	return sysDictionarys, err
}
