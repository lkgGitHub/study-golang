package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _JobRecommendRecordMgr struct {
	*_BaseMgr
}

// JobRecommendRecordMgr open func
func JobRecommendRecordMgr(db *gorm.DB) *_JobRecommendRecordMgr {
	if db == nil {
		panic(fmt.Errorf("JobRecommendRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_JobRecommendRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("job_recommend_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_JobRecommendRecordMgr) GetTableName() string {
	return "job_recommend_record"
}

// Get 获取
func (obj *_JobRecommendRecordMgr) Get() (result JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_JobRecommendRecordMgr) Gets() (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_JobRecommendRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPresenter presenter获取 推荐者
func (obj *_JobRecommendRecordMgr) WithPresenter(presenter int64) Option {
	return optionFunc(func(o *options) { o.query["presenter"] = presenter })
}

// WithUserID user_id获取 用户id
func (obj *_JobRecommendRecordMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithJobID job_id获取 职位id
func (obj *_JobRecommendRecordMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithCreateTime create_time获取 收藏时间
func (obj *_JobRecommendRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_JobRecommendRecordMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// GetByOption 功能选项模式获取
func (obj *_JobRecommendRecordMgr) GetByOption(opts ...Option) (result JobRecommendRecord, err error) {
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
func (obj *_JobRecommendRecordMgr) GetByOptions(opts ...Option) (results []*JobRecommendRecord, err error) {
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
func (obj *_JobRecommendRecordMgr) GetFromID(id int64) (result JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_JobRecommendRecordMgr) GetBatchFromID(ids []int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPresenter 通过presenter获取内容 推荐者
func (obj *_JobRecommendRecordMgr) GetFromPresenter(presenter int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("presenter = ?", presenter).Find(&results).Error

	return
}

// GetBatchFromPresenter 批量唯一主键查找 推荐者
func (obj *_JobRecommendRecordMgr) GetBatchFromPresenter(presenters []int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("presenter IN (?)", presenters).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_JobRecommendRecordMgr) GetFromUserID(userID int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_JobRecommendRecordMgr) GetBatchFromUserID(userIDs []int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容 职位id
func (obj *_JobRecommendRecordMgr) GetFromJobID(jobID int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量唯一主键查找 职位id
func (obj *_JobRecommendRecordMgr) GetBatchFromJobID(jobIDs []int64) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 收藏时间
func (obj *_JobRecommendRecordMgr) GetFromCreateTime(createTime time.Time) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 收藏时间
func (obj *_JobRecommendRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_JobRecommendRecordMgr) GetFromDeleteTime(deleteTime time.Time) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找 删除时间
func (obj *_JobRecommendRecordMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_JobRecommendRecordMgr) FetchByPrimaryKey(id int64) (result JobRecommendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
