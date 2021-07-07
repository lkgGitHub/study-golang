package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SendResumeRecordMgr struct {
	*_BaseMgr
}

// SendResumeRecordMgr open func
func SendResumeRecordMgr(db *gorm.DB) *_SendResumeRecordMgr {
	if db == nil {
		panic(fmt.Errorf("SendResumeRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SendResumeRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("send_resume_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SendResumeRecordMgr) GetTableName() string {
	return "send_resume_record"
}

// Get 获取
func (obj *_SendResumeRecordMgr) Get() (result SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SendResumeRecordMgr) Gets() (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_SendResumeRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPresenter presenter获取 推荐人id
func (obj *_SendResumeRecordMgr) WithPresenter(presenter int64) Option {
	return optionFunc(func(o *options) { o.query["presenter"] = presenter })
}

// WithUserID user_id获取 投递人用户id
func (obj *_SendResumeRecordMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithJobID job_id获取 职位id
func (obj *_SendResumeRecordMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithResumeID resume_id获取 简历id
func (obj *_SendResumeRecordMgr) WithResumeID(resumeID int64) Option {
	return optionFunc(func(o *options) { o.query["resume_id"] = resumeID })
}

// WithJobRecommendID job_recommend_id获取 职位推荐记录id
func (obj *_SendResumeRecordMgr) WithJobRecommendID(jobRecommendID int64) Option {
	return optionFunc(func(o *options) { o.query["job_recommend_id"] = jobRecommendID })
}

// WithStatus status获取 简历投递状态 1-已投递 2-通过初筛 3-面试中 4-通过面试 5-已入职 6-已转正 7-不合适 8-待确认中
func (obj *_SendResumeRecordMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 收藏时间
func (obj *_SendResumeRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_SendResumeRecordMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// GetByOption 功能选项模式获取
func (obj *_SendResumeRecordMgr) GetByOption(opts ...Option) (result SendResumeRecord, err error) {
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
func (obj *_SendResumeRecordMgr) GetByOptions(opts ...Option) (results []*SendResumeRecord, err error) {
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

// GetFromID 通过id获取内容 主键id
func (obj *_SendResumeRecordMgr) GetFromID(id int64) (result SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键id
func (obj *_SendResumeRecordMgr) GetBatchFromID(ids []int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPresenter 通过presenter获取内容 推荐人id
func (obj *_SendResumeRecordMgr) GetFromPresenter(presenter int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("presenter = ?", presenter).Find(&results).Error

	return
}

// GetBatchFromPresenter 批量唯一主键查找 推荐人id
func (obj *_SendResumeRecordMgr) GetBatchFromPresenter(presenters []int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("presenter IN (?)", presenters).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 投递人用户id
func (obj *_SendResumeRecordMgr) GetFromUserID(userID int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 投递人用户id
func (obj *_SendResumeRecordMgr) GetBatchFromUserID(userIDs []int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容 职位id
func (obj *_SendResumeRecordMgr) GetFromJobID(jobID int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量唯一主键查找 职位id
func (obj *_SendResumeRecordMgr) GetBatchFromJobID(jobIDs []int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromResumeID 通过resume_id获取内容 简历id
func (obj *_SendResumeRecordMgr) GetFromResumeID(resumeID int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resume_id = ?", resumeID).Find(&results).Error

	return
}

// GetBatchFromResumeID 批量唯一主键查找 简历id
func (obj *_SendResumeRecordMgr) GetBatchFromResumeID(resumeIDs []int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resume_id IN (?)", resumeIDs).Find(&results).Error

	return
}

// GetFromJobRecommendID 通过job_recommend_id获取内容 职位推荐记录id
func (obj *_SendResumeRecordMgr) GetFromJobRecommendID(jobRecommendID int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_recommend_id = ?", jobRecommendID).Find(&results).Error

	return
}

// GetBatchFromJobRecommendID 批量唯一主键查找 职位推荐记录id
func (obj *_SendResumeRecordMgr) GetBatchFromJobRecommendID(jobRecommendIDs []int64) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_recommend_id IN (?)", jobRecommendIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 简历投递状态 1-已投递 2-通过初筛 3-面试中 4-通过面试 5-已入职 6-已转正 7-不合适 8-待确认中
func (obj *_SendResumeRecordMgr) GetFromStatus(status int) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 简历投递状态 1-已投递 2-通过初筛 3-面试中 4-通过面试 5-已入职 6-已转正 7-不合适 8-待确认中
func (obj *_SendResumeRecordMgr) GetBatchFromStatus(statuss []int) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 收藏时间
func (obj *_SendResumeRecordMgr) GetFromCreateTime(createTime time.Time) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 收藏时间
func (obj *_SendResumeRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_SendResumeRecordMgr) GetFromDeleteTime(deleteTime time.Time) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找 删除时间
func (obj *_SendResumeRecordMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SendResumeRecordMgr) FetchByPrimaryKey(id int64) (result SendResumeRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
