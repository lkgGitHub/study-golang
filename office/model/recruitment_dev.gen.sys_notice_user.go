package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysNoticeUserMgr struct {
	*_BaseMgr
}

// SysNoticeUserMgr open func
func SysNoticeUserMgr(db *gorm.DB) *_SysNoticeUserMgr {
	if db == nil {
		panic(fmt.Errorf("SysNoticeUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysNoticeUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_notice_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysNoticeUserMgr) GetTableName() string {
	return "sys_notice_user"
}

// Get 获取
func (obj *_SysNoticeUserMgr) Get() (result SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysNoticeUserMgr) Gets() (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysNoticeUserMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNoticeID notice_id获取 通知公告id
func (obj *_SysNoticeUserMgr) WithNoticeID(noticeID int64) Option {
	return optionFunc(func(o *options) { o.query["notice_id"] = noticeID })
}

// WithUserID user_id获取 用户id
func (obj *_SysNoticeUserMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatus status获取 状态（字典 0未读 1已读）
func (obj *_SysNoticeUserMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithReadTime read_time获取 阅读时间
func (obj *_SysNoticeUserMgr) WithReadTime(readTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["read_time"] = readTime })
}

// GetByOption 功能选项模式获取
func (obj *_SysNoticeUserMgr) GetByOption(opts ...Option) (result SysNoticeUser, err error) {
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
func (obj *_SysNoticeUserMgr) GetByOptions(opts ...Option) (results []*SysNoticeUser, err error) {
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
func (obj *_SysNoticeUserMgr) GetFromID(id int64) (result SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysNoticeUserMgr) GetBatchFromID(ids []int64) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromNoticeID 通过notice_id获取内容 通知公告id
func (obj *_SysNoticeUserMgr) GetFromNoticeID(noticeID int64) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_id = ?", noticeID).Find(&results).Error

	return
}

// GetBatchFromNoticeID 批量唯一主键查找 通知公告id
func (obj *_SysNoticeUserMgr) GetBatchFromNoticeID(noticeIDs []int64) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_id IN (?)", noticeIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_SysNoticeUserMgr) GetFromUserID(userID int64) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_SysNoticeUserMgr) GetBatchFromUserID(userIDs []int64) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0未读 1已读）
func (obj *_SysNoticeUserMgr) GetFromStatus(status int8) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0未读 1已读）
func (obj *_SysNoticeUserMgr) GetBatchFromStatus(statuss []int8) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromReadTime 通过read_time获取内容 阅读时间
func (obj *_SysNoticeUserMgr) GetFromReadTime(readTime time.Time) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read_time = ?", readTime).Find(&results).Error

	return
}

// GetBatchFromReadTime 批量唯一主键查找 阅读时间
func (obj *_SysNoticeUserMgr) GetBatchFromReadTime(readTimes []time.Time) (results []*SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read_time IN (?)", readTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysNoticeUserMgr) FetchByPrimaryKey(id int64) (result SysNoticeUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
