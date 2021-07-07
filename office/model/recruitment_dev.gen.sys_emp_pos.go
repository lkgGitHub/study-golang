package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysEmpPosMgr struct {
	*_BaseMgr
}

// SysEmpPosMgr open func
func SysEmpPosMgr(db *gorm.DB) *_SysEmpPosMgr {
	if db == nil {
		panic(fmt.Errorf("SysEmpPosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysEmpPosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_emp_pos"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysEmpPosMgr) GetTableName() string {
	return "sys_emp_pos"
}

// Get 获取
func (obj *_SysEmpPosMgr) Get() (result SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysEmpPosMgr) Gets() (results []*SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysEmpPosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEmpID emp_id获取 员工id
func (obj *_SysEmpPosMgr) WithEmpID(empID int64) Option {
	return optionFunc(func(o *options) { o.query["emp_id"] = empID })
}

// WithPosID pos_id获取 职位id
func (obj *_SysEmpPosMgr) WithPosID(posID int64) Option {
	return optionFunc(func(o *options) { o.query["pos_id"] = posID })
}

// GetByOption 功能选项模式获取
func (obj *_SysEmpPosMgr) GetByOption(opts ...Option) (result SysEmpPos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysEmpPosMgr) GetByOptions(opts ...Option) (results []*SysEmpPos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SysEmpPosMgr) GetFromID(id int64) (result SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysEmpPosMgr) GetBatchFromID(ids []int64) (results []*SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromEmpID 通过emp_id获取内容 员工id
func (obj *_SysEmpPosMgr) GetFromEmpID(empID int64) (results []*SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("emp_id = ?", empID).Find(&results).Error

	return
}

// GetBatchFromEmpID 批量唯一主键查找 员工id
func (obj *_SysEmpPosMgr) GetBatchFromEmpID(empIDs []int64) (results []*SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("emp_id IN (?)", empIDs).Find(&results).Error

	return
}

// GetFromPosID 通过pos_id获取内容 职位id
func (obj *_SysEmpPosMgr) GetFromPosID(posID int64) (results []*SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pos_id = ?", posID).Find(&results).Error

	return
}

// GetBatchFromPosID 批量唯一主键查找 职位id
func (obj *_SysEmpPosMgr) GetBatchFromPosID(posIDs []int64) (results []*SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pos_id IN (?)", posIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysEmpPosMgr) FetchByPrimaryKey(id int64) (result SysEmpPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
