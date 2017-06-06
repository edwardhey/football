package models

import (
	"time"
)

type Base struct {
	Ctime     uint32 `gorm:"column:ctime;default:0"` //创建时间
	Mtime     uint32 `gorm:"column:mtime;default:0"` //更改时间
	isExists  bool   `gorm:"-"`
	isDeleted bool   `gorm:"-"`
}

/**
 * 是否存在对象
 */
func (m *Base) IsExists() bool {
	return m.isExists
}

/**
 * 是否新对象
 */
func (m *Base) IsNew() bool {
	return !m.isExists
}

func (m *Base) SetIsExists(b bool) {
	m.isExists = b
}

func (m *Base) SetIsDeleted(b bool) {
	m.isDeleted = b
}

/**
 * 设置成新对象
 */
func (m *Base) SetNew() {
	m.isExists = false
}

func (m *Base) SetExists() {
	m.isExists = true
}

func (m *Base) Reset() {
	panic("please implement this method!")
}

// default callbacks
func (m *Base) BeforeSave() error {
	return nil
}

func (m *Base) AfterSave() error {
	return nil
}

func (m *Base) BeforeCreate() error {
	now := uint32(time.Now().Unix())
	m.Ctime = now
	m.Mtime = now
	return nil
}

func (m *Base) AfterCreate() error {
	m.isExists = true
	m.isDeleted = false
	return nil
}

func (m *Base) BeforeUpdate() error {
	now := uint32(time.Now().Unix())
	m.Mtime = now
	return nil
}

func (m *Base) AfterUpdate() error {
	return nil
}

func (m *Base) AfterFind() error {
	m.isExists = true
	m.isDeleted = false
	return nil
}

func (m *Base) BeforeDelete() error {
	return nil
}

func (m *Base) AfterDelete() error {
	m.isExists = false
	m.isDeleted = true
	return nil
}
