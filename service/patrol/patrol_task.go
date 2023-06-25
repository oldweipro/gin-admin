package patrol

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"gorm.io/gorm"
)

type PatrolTaskService struct {
}

// CreatePatrolTask 创建PatrolTask记录
func (patrolTaskService *PatrolTaskService) CreatePatrolTask(patrolTask patrol.PatrolTask) (err error) {
	err = global.DB.Create(&patrolTask).Error
	return err
}

// DeletePatrolTask 删除PatrolTask记录
func (patrolTaskService *PatrolTaskService) DeletePatrolTask(patrolTask patrol.PatrolTask) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.PatrolTask{}).Where("id = ?", patrolTask.ID).Update("deleted_by", patrolTask.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&patrolTask).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePatrolTaskByIds 批量删除PatrolTask记录
func (patrolTaskService *PatrolTaskService) DeletePatrolTaskByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.PatrolTask{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&patrol.PatrolTask{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePatrolTask 更新PatrolTask记录
func (patrolTaskService *PatrolTaskService) UpdatePatrolTask(patrolTask patrol.PatrolTask) (err error) {
	err = global.DB.Save(&patrolTask).Error
	return err
}

// GetPatrolTask 根据id获取PatrolTask记录
func (patrolTaskService *PatrolTaskService) GetPatrolTask(id uint) (patrolTask patrol.PatrolTask, err error) {
	err = global.DB.Where("id = ?", id).First(&patrolTask).Error
	return
}

// GetPatrolTaskInfoList 分页获取PatrolTask记录
func (patrolTaskService *PatrolTaskService) GetPatrolTaskInfoList(info patrolReq.PatrolTaskSearch) (list []patrol.PatrolTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&patrol.PatrolTask{})
	var patrolTasks []patrol.PatrolTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TaskName != "" {
		db = db.Where("task_name LIKE ?", "%"+info.TaskName+"%")
	}
	if info.TaskItemIdList != "" {
		db = db.Where("task_item_id_list LIKE ?", "%"+info.TaskItemIdList+"%")
	}
	if info.PatrolTimes != nil {
		db = db.Where("patrol_times = ?", info.PatrolTimes)
	}
	if info.IntervalDuration != nil {
		db = db.Where("interval_duration = ?", info.IntervalDuration)
	}
	if info.ClockMode != nil {
		db = db.Where("clock_mode = ?", info.ClockMode)
	}
	if info.SiteIdList != "" {
		db = db.Where("site_id_list LIKE ?", "%"+info.SiteIdList+"%")
	}
	if info.TaskCycleTime != "" {
		db = db.Where("task_cycle_time LIKE ?", "%"+info.TaskCycleTime+"%")
	}
	if info.ValidTime != "" {
		db = db.Where("valid_time LIKE ?", "%"+info.ValidTime+"%")
	}
	if info.DeptId != nil {
		db = db.Where("dept_id = ?", info.DeptId)
	}
	if info.AssignTasks != "" {
		db = db.Where("assign_tasks LIKE ?", "%"+info.AssignTasks+"%")
	}
	if info.CopySupervisor != "" {
		db = db.Where("copy_supervisor LIKE ?", "%"+info.CopySupervisor+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&patrolTasks).Error
	return patrolTasks, total, err
}
