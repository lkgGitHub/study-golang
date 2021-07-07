package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _JobDetailsMgr struct {
	*_BaseMgr
}

// JobDetailsMgr open func
func JobDetailsMgr(db *gorm.DB) *_JobDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("JobDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_JobDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("job_details"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_JobDetailsMgr) GetTableName() string {
	return "job_details"
}

// Get 获取
func (obj *_JobDetailsMgr) Get() (result JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_JobDetailsMgr) Gets() (results []*JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_JobDetailsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithJobResponsibilities job_responsibilities获取
func (obj *_JobDetailsMgr) WithJobResponsibilities(jobResponsibilities string) Option {
	return optionFunc(func(o *options) { o.query["job_responsibilities"] = jobResponsibilities })
}

// WithJobRequirements job_requirements获取
func (obj *_JobDetailsMgr) WithJobRequirements(jobRequirements string) Option {
	return optionFunc(func(o *options) { o.query["job_requirements"] = jobRequirements })
}

// GetByOption 功能选项模式获取
func (obj *_JobDetailsMgr) GetByOption(opts ...Option) (result JobDetails, err error) {
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
func (obj *_JobDetailsMgr) GetByOptions(opts ...Option) (results []*JobDetails, err error) {
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
func (obj *_JobDetailsMgr) GetFromID(id int64) (result JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_JobDetailsMgr) GetBatchFromID(ids []int64) (results []*JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromJobResponsibilities 通过job_responsibilities获取内容
func (obj *_JobDetailsMgr) GetFromJobResponsibilities(jobResponsibilities string) (results []*JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_responsibilities = ?", jobResponsibilities).Find(&results).Error

	return
}

// GetBatchFromJobResponsibilities 批量唯一主键查找
func (obj *_JobDetailsMgr) GetBatchFromJobResponsibilities(jobResponsibilitiess []string) (results []*JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_responsibilities IN (?)", jobResponsibilitiess).Find(&results).Error

	return
}

// GetFromJobRequirements 通过job_requirements获取内容
func (obj *_JobDetailsMgr) GetFromJobRequirements(jobRequirements string) (results []*JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_requirements = ?", jobRequirements).Find(&results).Error

	return
}

// GetBatchFromJobRequirements 批量唯一主键查找
func (obj *_JobDetailsMgr) GetBatchFromJobRequirements(jobRequirementss []string) (results []*JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_requirements IN (?)", jobRequirementss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_JobDetailsMgr) FetchByPrimaryKey(id int64) (result JobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
