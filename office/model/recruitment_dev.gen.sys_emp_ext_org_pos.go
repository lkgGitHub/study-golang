package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysEmpExtOrgPosMgr struct {
	*_BaseMgr
}

// SysEmpExtOrgPosMgr open func
func SysEmpExtOrgPosMgr(db *gorm.DB) *_SysEmpExtOrgPosMgr {
	if db == nil {
		panic(fmt.Errorf("SysEmpExtOrgPosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysEmpExtOrgPosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_emp_ext_org_pos"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysEmpExtOrgPosMgr) GetTableName() string {
	return "sys_emp_ext_org_pos"
}

// Get 获取
func (obj *_SysEmpExtOrgPosMgr) Get() (result SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysEmpExtOrgPosMgr) Gets() (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysEmpExtOrgPosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEmpID emp_id获取 员工id
func (obj *_SysEmpExtOrgPosMgr) WithEmpID(empID int64) Option {
	return optionFunc(func(o *options) { o.query["emp_id"] = empID })
}

// WithOrgID org_id获取 机构id
func (obj *_SysEmpExtOrgPosMgr) WithOrgID(orgID int64) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// WithPosID pos_id获取 岗位id
func (obj *_SysEmpExtOrgPosMgr) WithPosID(posID int64) Option {
	return optionFunc(func(o *options) { o.query["pos_id"] = posID })
}

// GetByOption 功能选项模式获取
func (obj *_SysEmpExtOrgPosMgr) GetByOption(opts ...Option) (result SysEmpExtOrgPos, err error) {
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
func (obj *_SysEmpExtOrgPosMgr) GetByOptions(opts ...Option) (results []*SysEmpExtOrgPos, err error) {
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
func (obj *_SysEmpExtOrgPosMgr) GetFromID(id int64) (result SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysEmpExtOrgPosMgr) GetBatchFromID(ids []int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromEmpID 通过emp_id获取内容 员工id
func (obj *_SysEmpExtOrgPosMgr) GetFromEmpID(empID int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("emp_id = ?", empID).Find(&results).Error

	return
}

// GetBatchFromEmpID 批量唯一主键查找 员工id
func (obj *_SysEmpExtOrgPosMgr) GetBatchFromEmpID(empIDs []int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("emp_id IN (?)", empIDs).Find(&results).Error

	return
}

// GetFromOrgID 通过org_id获取内容 机构id
func (obj *_SysEmpExtOrgPosMgr) GetFromOrgID(orgID int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id = ?", orgID).Find(&results).Error

	return
}

// GetBatchFromOrgID 批量唯一主键查找 机构id
func (obj *_SysEmpExtOrgPosMgr) GetBatchFromOrgID(orgIDs []int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id IN (?)", orgIDs).Find(&results).Error

	return
}

// GetFromPosID 通过pos_id获取内容 岗位id
func (obj *_SysEmpExtOrgPosMgr) GetFromPosID(posID int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pos_id = ?", posID).Find(&results).Error

	return
}

// GetBatchFromPosID 批量唯一主键查找 岗位id
func (obj *_SysEmpExtOrgPosMgr) GetBatchFromPosID(posIDs []int64) (results []*SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pos_id IN (?)", posIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysEmpExtOrgPosMgr) FetchByPrimaryKey(id int64) (result SysEmpExtOrgPos, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
