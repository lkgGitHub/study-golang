package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MemberMgr struct {
	*_BaseMgr
}

// MemberMgr open func
func MemberMgr(db *gorm.DB) *_MemberMgr {
	if db == nil {
		panic(fmt.Errorf("MemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MemberMgr) GetTableName() string {
	return "member"
}

// Get 获取
func (obj *_MemberMgr) Get() (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MemberMgr) Gets() (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_MemberMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户ID
func (obj *_MemberMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithIsVip is_vip获取 是否vip 0-普通用户 1-推荐师
func (obj *_MemberMgr) WithIsVip(isVip int) Option {
	return optionFunc(func(o *options) { o.query["is_vip"] = isVip })
}

// WithCreateTime create_time获取
func (obj *_MemberMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_MemberMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_MemberMgr) GetByOption(opts ...Option) (result Member, err error) {
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
func (obj *_MemberMgr) GetByOptions(opts ...Option) (results []*Member, err error) {
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

// GetFromID 通过id获取内容 主键ID
func (obj *_MemberMgr) GetFromID(id int) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键ID
func (obj *_MemberMgr) GetBatchFromID(ids []int) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MemberMgr) GetFromUserID(userID int) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *_MemberMgr) GetBatchFromUserID(userIDs []int) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromIsVip 通过is_vip获取内容 是否vip 0-普通用户 1-推荐师
func (obj *_MemberMgr) GetFromIsVip(isVip int) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_vip = ?", isVip).Find(&results).Error

	return
}

// GetBatchFromIsVip 批量唯一主键查找 是否vip 0-普通用户 1-推荐师
func (obj *_MemberMgr) GetBatchFromIsVip(isVips []int) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_vip IN (?)", isVips).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_MemberMgr) GetFromCreateTime(createTime time.Time) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_MemberMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_MemberMgr) GetFromUpdateTime(updateTime time.Time) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_MemberMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MemberMgr) FetchByPrimaryKey(id int) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByMemberUserIDIndex primay or index 获取唯一内容
func (obj *_MemberMgr) FetchUniqueByMemberUserIDIndex(userID int) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}
