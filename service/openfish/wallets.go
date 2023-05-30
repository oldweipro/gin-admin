package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
)

type WalletsService struct {
}

// CreateWallets 创建Wallets记录
// Author [piexlmax](https://github.com/piexlmax)
func (walletsService *WalletsService) CreateWallets(wallets *openfish.Wallets) (err error) {
	err = global.GVA_DB.Create(wallets).Error
	return err
}

// DeleteWallets 删除Wallets记录
// Author [piexlmax](https://github.com/piexlmax)
func (walletsService *WalletsService) DeleteWallets(wallets openfish.Wallets) (err error) {
	err = global.GVA_DB.Delete(&wallets).Error
	return err
}

// DeleteWalletsByIds 批量删除Wallets记录
// Author [piexlmax](https://github.com/piexlmax)
func (walletsService *WalletsService) DeleteWalletsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]openfish.Wallets{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWallets 更新Wallets记录
// Author [piexlmax](https://github.com/piexlmax)
func (walletsService *WalletsService) UpdateWallets(wallets openfish.Wallets) (err error) {
	err = global.GVA_DB.Save(&wallets).Error
	return err
}

// GetWallets 根据id获取Wallets记录
// Author [piexlmax](https://github.com/piexlmax)
func (walletsService *WalletsService) GetWallets(id uint) (wallets openfish.Wallets, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wallets).Error
	return
}

// GetWalletsInfoList 分页获取Wallets记录
// Author [piexlmax](https://github.com/piexlmax)
func (walletsService *WalletsService) GetWalletsInfoList(info openfishReq.WalletsSearch) (list []openfish.Wallets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openfish.Wallets{})
	var walletss []openfish.Wallets
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.WalletName != "" {
		db = db.Where("wallet_name LIKE ?", "%"+info.WalletName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&walletss).Error
	return walletss, total, err
}
