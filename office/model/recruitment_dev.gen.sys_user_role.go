package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysUserRoleMgr struct {
	*_BaseMgr
}

// SysUserRoleMgr open func
func SysUserRoleMgr(db *gorm.DB) *_SysUserRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserRoleMgr) GetTableName() string {
	return "sys_user_role"
}

// Get 获取
func (obj *_SysUserRoleMgr) Get() (result SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysUserRoleMgr) Gets() (results []*SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysUserRoleMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_SysUserRoleMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithRoleID role_id获取 角色id
func (obj *_SysUserRoleMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// GetByOption 功能选项模式获取
func (obj *_SysUserRoleMgr) GetByOption(opts ...Option) (result SysUserRole, err error) {
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
func (obj *_SysUserRoleMgr) GetByOptions(opts ...Option) (results []*SysUserRole, err error) {
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
func (obj *_SysUserRoleMgr) GetFromID(id int64) (result SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysUserRoleMgr) GetBatchFromID(ids []int64) (results []*SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_SysUserRoleMgr) GetFromUserID(userID int64) (results []*SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_SysUserRoleMgr) GetBatchFromUserID(userIDs []int64) (results []*SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色id
func (obj *_SysUserRoleMgr) GetFromRoleID(roleID int64) (results []*SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色id
func (obj *_SysUserRoleMgr) GetBatchFromRoleID(roleIDs []int64) (results []*SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysUserRoleMgr) FetchByPrimaryKey(id int64) (result SysUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
