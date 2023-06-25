package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/transaction"
	openfishReq "github.com/oldweipro/gin-admin/model/transaction/request"
)

type WalletsService struct {
}

// CreateWallets 创建Wallets记录
func (walletsService *WalletsService) CreateWallets(wallets *transaction.Wallets) (err error) {
	err = global.DB.Create(wallets).Error
	return err
}

// DeleteWallets 删除Wallets记录
func (walletsService *WalletsService) DeleteWallets(wallets transaction.Wallets) (err error) {
	err = global.DB.Delete(&wallets).Error
	return err
}

// DeleteWalletsByIds 批量删除Wallets记录
func (walletsService *WalletsService) DeleteWalletsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]transaction.Wallets{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWallets 更新Wallets记录
func (walletsService *WalletsService) UpdateWallets(wallets transaction.Wallets) (err error) {
	err = global.DB.Save(&wallets).Error
	return err
}

// GetWallets 根据id获取Wallets记录
func (walletsService *WalletsService) GetWallets(id uint) (wallets transaction.Wallets, err error) {
	err = global.DB.Where("id = ?", id).First(&wallets).Error
	return
}

// GetWalletsInfoList 分页获取Wallets记录
func (walletsService *WalletsService) GetWalletsInfoList(info openfishReq.WalletsSearch) (list []transaction.Wallets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&transaction.Wallets{})
	var walletss []transaction.Wallets
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

// GetCurrentUserWallets 获取当前用户钱包
func (walletsService *WalletsService) GetCurrentUserWallets(userId uint) (wallets transaction.Wallets, err error) {
	err = global.DB.Where("user_id = ?", userId).First(&wallets).Error
	return
}
