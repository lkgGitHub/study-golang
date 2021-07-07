package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _RegisterRecommendRecordMgr struct {
	*_BaseMgr
}

// RegisterRecommendRecordMgr open func
func RegisterRecommendRecordMgr(db *gorm.DB) *_RegisterRecommendRecordMgr {
	if db == nil {
		panic(fmt.Errorf("RegisterRecommendRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RegisterRecommendRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("register_recommend_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RegisterRecommendRecordMgr) GetTableName() string {
	return "register_recommend_record"
}

// Get 获取
func (obj *_RegisterRecommendRecordMgr) Get() (result RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RegisterRecommendRecordMgr) Gets() (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RegisterRecommendRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithReferrerID referrer_id获取 推荐者
func (obj *_RegisterRecommendRecordMgr) WithReferrerID(referrerID int64) Option {
	return optionFunc(func(o *options) { o.query["referrer_id"] = referrerID })
}

// WithUserID user_id获取 用户id
func (obj *_RegisterRecommendRecordMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCreateTime create_time获取 推荐注册时间
func (obj *_RegisterRecommendRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_RegisterRecommendRecordMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// GetByOption 功能选项模式获取
func (obj *_RegisterRecommendRecordMgr) GetByOption(opts ...Option) (result RegisterRecommendRecord, err error) {
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
func (obj *_RegisterRecommendRecordMgr) GetByOptions(opts ...Option) (results []*RegisterRecommendRecord, err error) {
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

// GetFromID 通过id获取内容
func (obj *_RegisterRecommendRecordMgr) GetFromID(id int64) (result RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_RegisterRecommendRecordMgr) GetBatchFromID(ids []int64) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromReferrerID 通过referrer_id获取内容 推荐者
func (obj *_RegisterRecommendRecordMgr) GetFromReferrerID(referrerID int64) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("referrer_id = ?", referrerID).Find(&results).Error

	return
}

// GetBatchFromReferrerID 批量唯一主键查找 推荐者
func (obj *_RegisterRecommendRecordMgr) GetBatchFromReferrerID(referrerIDs []int64) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("referrer_id IN (?)", referrerIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_RegisterRecommendRecordMgr) GetFromUserID(userID int64) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_RegisterRecommendRecordMgr) GetBatchFromUserID(userIDs []int64) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 推荐注册时间
func (obj *_RegisterRecommendRecordMgr) GetFromCreateTime(createTime time.Time) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 推荐注册时间
func (obj *_RegisterRecommendRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_RegisterRecommendRecordMgr) GetFromDeleteTime(deleteTime time.Time) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找 删除时间
func (obj *_RegisterRecommendRecordMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_RegisterRecommendRecordMgr) FetchByPrimaryKey(id int64) (result RegisterRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
