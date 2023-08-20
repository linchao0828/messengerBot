// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameConfig = "config"

// Config mapped from table <config>
type Config struct {
	ID          int64     `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Key         string    `gorm:"column:key;type:varchar(32);not null" json:"key"`
	Value       string    `gorm:"column:value;type:varchar(128);not null" json:"value"`
	CreatedTime time.Time `gorm:"column:created_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime time.Time `gorm:"column:deleted_time;type:timestamp;not null;default:2038-01-19 03:14:07" json:"deleted_time"`
}

// TableName Config's table name
func (*Config) TableName() string {
	return TableNameConfig
}
