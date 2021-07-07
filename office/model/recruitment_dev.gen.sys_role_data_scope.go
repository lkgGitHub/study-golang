package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysRoleDataScopeMgr struct {
	*_BaseMgr
}

// SysRoleDataScopeMgr open func
func SysRoleDataScopeMgr(db *gorm.DB) *_SysRoleDataScopeMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleDataScopeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysRoleDataScopeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_role_data_scope"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleDataScopeMgr) GetTableName() string {
	return "sys_role_data_scope"
}

// Get 获取
func (obj *_SysRoleDataScopeMgr) Get() (result SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleDataScopeMgr) Gets() (results []*SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysRoleDataScopeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleID role_id获取 角色id
func (obj *_SysRoleDataScopeMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithOrgID org_id获取 机构id
func (obj *_SysRoleDataScopeMgr) WithOrgID(orgID int64) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleDataScopeMgr) GetByOption(opts ...Option) (result SysRoleDataScope, err error) {
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
func (obj *_SysRoleDataScopeMgr) GetByOptions(opts ...Option) (results []*SysRoleDataScope, err error) {
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
func (obj *_SysRoleDataScopeMgr) GetFromID(id int64) (result SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysRoleDataScopeMgr) GetBatchFromID(ids []int64) (results []*SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色id
func (obj *_SysRoleDataScopeMgr) GetFromRoleID(roleID int64) (results []*SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色id
func (obj *_SysRoleDataScopeMgr) GetBatchFromRoleID(roleIDs []int64) (results []*SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromOrgID 通过org_id获取内容 机构id
func (obj *_SysRoleDataScopeMgr) GetFromOrgID(orgID int64) (results []*SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id = ?", orgID).Find(&results).Error

	return
}

// GetBatchFromOrgID 批量唯一主键查找 机构id
func (obj *_SysRoleDataScopeMgr) GetBatchFromOrgID(orgIDs []int64) (results []*SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id IN (?)", orgIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleDataScopeMgr) FetchByPrimaryKey(id int64) (result SysRoleDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
