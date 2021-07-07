package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CollectionMgr struct {
	*_BaseMgr
}

// CollectionMgr open func
func CollectionMgr(db *gorm.DB) *_CollectionMgr {
	if db == nil {
		panic(fmt.Errorf("CollectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CollectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("collection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CollectionMgr) GetTableName() string {
	return "collection"
}

// Get 获取
func (obj *_CollectionMgr) Get() (result Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CollectionMgr) Gets() (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CollectionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_CollectionMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithJobID job_id获取 职位id
func (obj *_CollectionMgr) WithJobID(jobID int) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithCreateTime create_time获取 收藏时间
func (obj *_CollectionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_CollectionMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// WithStatus status获取 职位状态。收藏职位失效或者下架
func (obj *_CollectionMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateUser create_user获取
func (obj *_CollectionMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_CollectionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取
func (obj *_CollectionMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_CollectionMgr) GetByOption(opts ...Option) (result Collection, err error) {
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
func (obj *_CollectionMgr) GetByOptions(opts ...Option) (results []*Collection, err error) {
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
func (obj *_CollectionMgr) GetFromID(id int) (result Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CollectionMgr) GetBatchFromID(ids []int) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_CollectionMgr) GetFromUserID(userID int) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_CollectionMgr) GetBatchFromUserID(userIDs []int) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容 职位id
func (obj *_CollectionMgr) GetFromJobID(jobID int) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量唯一主键查找 职位id
func (obj *_CollectionMgr) GetBatchFromJobID(jobIDs []int) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 收藏时间
func (obj *_CollectionMgr) GetFromCreateTime(createTime time.Time) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 收藏时间
func (obj *_CollectionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_CollectionMgr) GetFromDeleteTime(deleteTime time.Time) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找 删除时间
func (obj *_CollectionMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 职位状态。收藏职位失效或者下架
func (obj *_CollectionMgr) GetFromStatus(status string) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 职位状态。收藏职位失效或者下架
func (obj *_CollectionMgr) GetBatchFromStatus(statuss []string) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容
func (obj *_CollectionMgr) GetFromCreateUser(createUser int64) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找
func (obj *_CollectionMgr) GetBatchFromCreateUser(createUsers []int64) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_CollectionMgr) GetFromUpdateTime(updateTime time.Time) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_CollectionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容
func (obj *_CollectionMgr) GetFromUpdateUser(updateUser int64) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找
func (obj *_CollectionMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CollectionMgr) FetchByPrimaryKey(id int) (result Collection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
