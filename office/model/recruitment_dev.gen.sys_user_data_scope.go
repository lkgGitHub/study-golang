package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysUserDataScopeMgr struct {
	*_BaseMgr
}

// SysUserDataScopeMgr open func
func SysUserDataScopeMgr(db *gorm.DB) *_SysUserDataScopeMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserDataScopeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserDataScopeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user_data_scope"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserDataScopeMgr) GetTableName() string {
	return "sys_user_data_scope"
}

// Get 获取
func (obj *_SysUserDataScopeMgr) Get() (result SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysUserDataScopeMgr) Gets() (results []*SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysUserDataScopeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_SysUserDataScopeMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithOrgID org_id获取 机构id
func (obj *_SysUserDataScopeMgr) WithOrgID(orgID int64) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// GetByOption 功能选项模式获取
func (obj *_SysUserDataScopeMgr) GetByOption(opts ...Option) (result SysUserDataScope, err error) {
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
func (obj *_SysUserDataScopeMgr) GetByOptions(opts ...Option) (results []*SysUserDataScope, err error) {
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
func (obj *_SysUserDataScopeMgr) GetFromID(id int64) (result SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysUserDataScopeMgr) GetBatchFromID(ids []int64) (results []*SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_SysUserDataScopeMgr) GetFromUserID(userID int64) (results []*SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_SysUserDataScopeMgr) GetBatchFromUserID(userIDs []int64) (results []*SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromOrgID 通过org_id获取内容 机构id
func (obj *_SysUserDataScopeMgr) GetFromOrgID(orgID int64) (results []*SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id = ?", orgID).Find(&results).Error

	return
}

// GetBatchFromOrgID 批量唯一主键查找 机构id
func (obj *_SysUserDataScopeMgr) GetBatchFromOrgID(orgIDs []int64) (results []*SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id IN (?)", orgIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysUserDataScopeMgr) FetchByPrimaryKey(id int64) (result SysUserDataScope, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
