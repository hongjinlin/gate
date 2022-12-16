package models

import (
	"time"
)

// Gate [...]
type Gate struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"-"`
	Nation     string    `gorm:"column:nation;type:varchar(128);index" json:"nation"` // 国家/地区
	IP         string    `gorm:"column:ip;type:varchar(32);uniqueIndex" json:"ip"`
	Runtime    string    `gorm:"column:runtime;type:varchar(32)" json:"runtime"` // 运行时间
	Status     int       `gorm:"column:status" json:"status"`                    // 1:available 0: unavailable
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *Gate) TableName() string {
	return "gate"
}

// GateColumns get sql column name.获取数据库列名
var GateColumns = struct {
	ID         string
	Nation     string
	IP         string
	Runtime    string
	Status     string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Nation:     "nation",
	IP:         "ip",
	Runtime:    "runtime",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
