// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID          int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // 自增ID
	Phone       string                `gorm:"column:phone" json:"phone"`                            // 手机号码
	Password    string                `gorm:"column:password" json:"password"`                      // 密码
	UserName    string                `gorm:"column:user_name" json:"user_name"`                    // 用户姓名
	UserHeadURL string                `gorm:"column:user_head_url" json:"user_head_url"`            // 用户头像
	Status      string                `gorm:"column:status;default:1" json:"status"`                // 用户状态(0:禁用,1:正常)
	UserType    string                `gorm:"column:user_type;default:1" json:"user_type"`          // 用户类型(0:管理员,1:巡视员)
	LoginTime   time.Time             `gorm:"column:login_time" json:"login_time"`                  // 登录时间
	CreateTime  int64                 `gorm:"column:create_time;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime  int64                 `gorm:"column:update_time;autoUpdateTime" json:"update_time"` // 更新时间
	ModifyTime  time.Time             `gorm:"column:modify_time;autoUpdateTime" json:"modify_time"` // 维护字段-更新时间
	DeleteTime  soft_delete.DeletedAt `gorm:"column:delete_time;not null" json:"delete_time"`       // 删除时间
	OrgID       int64                 `gorm:"column:org_id;not null" json:"org_id"`                 // 组织ID
	RoomIds     string                `gorm:"column:room_ids" json:"room_ids"`                      // 警务室ID
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}

// IsEmpty 判断空
func (m *SysUser) IsEmpty() bool {
	if m == nil {
		return true
	}

	return m.ID == 0
}
