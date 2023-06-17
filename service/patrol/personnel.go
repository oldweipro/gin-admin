package patrol

import (
	"encoding/json"
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"gorm.io/gorm/clause"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type PersonnelService struct {
}

// SyncPersonnel 同步人员信息
// Author [oldweipro](https://github.com/oldweipro)
func (personnelService *PersonnelService) SyncPersonnel() (err error) {
	// 创建db
	db := global.DB.Model(&patrol.Personnel{})
	// 查询当前数据库的记录数
	var count int64
	db.Count(&count)
	// 第几页
	var pageNum int64 = 1
	// 每页数据量
	var pageSize int64 = 1000
	// 获取远端数据
	data := getPersonnel(pageNum, pageSize)
	// 如果远端的总数和数据库的总数是相等的，说明数据已经是同步的
	if count == int64(data.Data.Total) {
		return err
	}
	// 全量同步
	var pages = int64(data.Data.Pages)
	fmt.Println("一共多少页: ", pages)
	// 循环增加数据
	go func() {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).CreateInBatches(&data.Data.List, len(data.Data.List))
		for i := pageNum + 1; i <= pages; i++ {
			personnel := getPersonnel(i, pageSize)
			// 获取当前数据总量
			fmt.Println("同步第", i, "页")
			// Update columns to default value on `id` conflict
			db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "id"}},
				UpdateAll: true,
			}).CreateInBatches(&personnel.Data.List, len(personnel.Data.List))
			time.Sleep(time.Second)
		}
		fmt.Println("同步完成")
	}()
	return err
}

func getPersonnel(pageNum, pageSize int64) patrolReq.PersonnelQueryResult {
	var personnelQueryResult patrolReq.PersonnelQueryResult
	var urlBuild strings.Builder
	urlBuild.WriteString("")
	// strconv.FormatInt的第二个参数base表示进制，10表示十进制
	urlBuild.WriteString(strconv.FormatInt(pageNum, 10))
	urlBuild.WriteString("&pageSize=")
	urlBuild.WriteString(strconv.FormatInt(pageSize, 10))
	url := urlBuild.String()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("远端获取人员数据接口错误: ", err)
		return personnelQueryResult
	}
	if resp.StatusCode == http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(data, &personnelQueryResult)
		if err != nil {
			fmt.Println("反序列化错误: ", err)
		}
	}
	return personnelQueryResult
}

// SyncPersonnelImg 同步人员信息图片
// Author [oldweipro](https://github.com/oldweipro)
func (personnelService *PersonnelService) SyncPersonnelImg() (err error) {
	// 创建db
	db := global.DB.Model(&patrol.Personnel{})
	var personnels []patrol.Personnel
	err = db.Select("person_url").Find(&personnels).Error
	// 循环数据
	go func() {
		// 当前路径
		pwd, _ := os.Getwd()
		for i, personnel := range personnels {
			// 文件绝对路径
			imgPath := pwd + personnel.PersonUrl
			_, err := os.Stat(imgPath)
			if err != nil {
				fmt.Println("开始处理第", i, "张图片")
				// 分割出来文件名
				personUrlSplit := strings.Split(personnel.PersonUrl, "/")
				personImgName := personUrlSplit[len(personUrlSplit)-1]
				// 分割出来路径
				imgPathTrimImgName := strings.Trim(personnel.PersonUrl, personImgName)
				// 文件夹绝对路径
				imgDir := pwd + imgPathTrimImgName
				_, err := os.Stat(imgPath)
				// 检测文件夹是否存在
				if err != nil {
					// 创建文件夹
					err = os.MkdirAll(imgDir, os.ModePerm)
					if err != nil {
						fmt.Println("创建文件夹失败:", err)
						continue
					}
				}
				// 创建文件
				_, err = os.Create(imgPath)
				resp, err := http.Get("" + personnel.PersonUrl)
				bytes, _ := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("同步图片下载失败: ", err)
					continue
				}
				// 存储图片到本地
				err = os.WriteFile(imgPath, bytes, os.ModePerm)
				if err != nil {
					fmt.Println("同步图片存储失败: ", err)
					continue
				}
			}
		}
		fmt.Println("同步完成")
	}()
	return err
}

// CreatePersonnel 创建Personnel记录
// Author [piexlmax](https://github.com/piexlmax)
func (personnelService *PersonnelService) CreatePersonnel(personnel patrol.Personnel) (err error) {
	err = global.DB.Create(&personnel).Error
	return err
}

// DeletePersonnel 删除Personnel记录
// Author [piexlmax](https://github.com/piexlmax)
func (personnelService *PersonnelService) DeletePersonnel(personnel patrol.Personnel) (err error) {
	err = global.DB.Delete(&personnel).Error
	return err
}

// DeletePersonnelByIds 批量删除Personnel记录
// Author [piexlmax](https://github.com/piexlmax)
func (personnelService *PersonnelService) DeletePersonnelByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]patrol.Personnel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePersonnel 更新Personnel记录
// Author [piexlmax](https://github.com/piexlmax)
func (personnelService *PersonnelService) UpdatePersonnel(personnel patrol.Personnel) (err error) {
	err = global.DB.Save(&personnel).Error
	return err
}

// GetPersonnel 根据id获取Personnel记录
// Author [piexlmax](https://github.com/piexlmax)
func (personnelService *PersonnelService) GetPersonnel(id uint) (personnel patrol.Personnel, err error) {
	err = global.DB.Where("id = ?", id).First(&personnel).Error
	return
}

// GetPersonnelInfoList 分页获取Personnel记录
// Author [piexlmax](https://github.com/piexlmax)
func (personnelService *PersonnelService) GetPersonnelInfoList(info patrolReq.PersonnelSearch) (list []patrol.Personnel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&patrol.Personnel{})
	var personnels []patrol.Personnel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Sex != nil {
		db = db.Where("sex = ?", info.Sex)
	}
	if info.PersonName != "" {
		db = db.Where("person_name LIKE ?", "%"+info.PersonName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&personnels).Error
	return personnels, total, err
}
