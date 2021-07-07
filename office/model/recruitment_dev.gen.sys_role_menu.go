package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysRoleMenuMgr struct {
	*_BaseMgr
}

// SysRoleMenuMgr open func
func SysRoleMenuMgr(db *gorm.DB) *_SysRoleMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysRoleMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_role_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleMenuMgr) GetTableName() string {
	return "sys_role_menu"
}

// Get 获取
func (obj *_SysRoleMenuMgr) Get() (result SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleMenuMgr) Gets() (results []*SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysRoleMenuMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleID role_id获取 角色id
func (obj *_SysRoleMenuMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithMenuID menu_id获取 菜单id
func (obj *_SysRoleMenuMgr) WithMenuID(menuID int64) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleMenuMgr) GetByOption(opts ...Option) (result SysRoleMenu, err error) {
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
func (obj *_SysRoleMenuMgr) GetByOptions(opts ...Option) (results []*SysRoleMenu, err error) {
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
func (obj *_SysRoleMenuMgr) GetFromID(id int64) (result SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysRoleMenuMgr) GetBatchFromID(ids []int64) (results []*SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色id
func (obj *_SysRoleMenuMgr) GetFromRoleID(roleID int64) (results []*SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色id
func (obj *_SysRoleMenuMgr) GetBatchFromRoleID(roleIDs []int64) (results []*SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromMenuID 通过menu_id获取内容 菜单id
func (obj *_SysRoleMenuMgr) GetFromMenuID(menuID int64) (results []*SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&results).Error

	return
}

// GetBatchFromMenuID 批量唯一主键查找 菜单id
func (obj *_SysRoleMenuMgr) GetBatchFromMenuID(menuIDs []int64) (results []*SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_id IN (?)", menuIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleMenuMgr) FetchByPrimaryKey(id int64) (result SysRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
