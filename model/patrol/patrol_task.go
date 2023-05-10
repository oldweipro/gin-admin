// 自动生成模板PatrolTask
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// PatrolTask 结构体
type PatrolTask struct {
	global.GVA_MODEL
	TaskName         string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;"`
	TaskItemIdList   string `json:"taskItemIdList" form:"taskItemIdList" gorm:"column:task_item_id_list;comment:任务内容项列表;"`
	PatrolTimes      *int   `json:"patrolTimes" form:"patrolTimes" gorm:"column:patrol_times;comment:巡检次数;"`
	IntervalDuration *int   `json:"intervalDuration" form:"intervalDuration" gorm:"column:interval_duration;comment:多次巡更间隔时长;"`
	ClockMode        *int   `json:"clockMode" form:"clockMode" gorm:"column:clock_mode;comment:巡检打卡方式;"`
	SiteIdList       string `json:"siteIdList" form:"siteIdList" gorm:"column:site_id_list;comment:巡检地点ID集合;"`
	TaskCycleTime    string `json:"taskCycleTime" form:"taskCycleTime" gorm:"column:task_cycle_time;comment:任务时间周期;"`
	ValidTime        string `json:"validTime" form:"validTime" gorm:"column:valid_time;comment:当天巡检有效时间;"`
	DeptId           *int   `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:所属项目;"`
	AssignTasks      string `json:"assignTasks" form:"assignTasks" gorm:"column:assign_tasks;comment:分配任务;"`
	CopySupervisor   string `json:"copySupervisor" form:"copySupervisor" gorm:"column:copy_supervisor;comment:抄送至监管员;"`
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName PatrolTask 表名
func (PatrolTask) TableName() string {
	return "patrol_task"
}
