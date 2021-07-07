package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysTimersMgr struct {
	*_BaseMgr
}

// SysTimersMgr open func
func SysTimersMgr(db *gorm.DB) *_SysTimersMgr {
	if db == nil {
		panic(fmt.Errorf("SysTimersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysTimersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_timers"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysTimersMgr) GetTableName() string {
	return "sys_timers"
}

// Get 获取
func (obj *_SysTimersMgr) Get() (result SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysTimersMgr) Gets() (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 定时器id
func (obj *_SysTimersMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTimerName timer_name获取 任务名称
func (obj *_SysTimersMgr) WithTimerName(timerName string) Option {
	return optionFunc(func(o *options) { o.query["timer_name"] = timerName })
}

// WithActionClass action_class获取 执行任务的class的类名（实现了TimerTaskRunner接口的类的全称）
func (obj *_SysTimersMgr) WithActionClass(actionClass string) Option {
	return optionFunc(func(o *options) { o.query["action_class"] = actionClass })
}

// WithCron cron获取 定时任务表达式
func (obj *_SysTimersMgr) WithCron(cron string) Option {
	return optionFunc(func(o *options) { o.query["cron"] = cron })
}

// WithJobStatus job_status获取 状态（字典 1运行  2停止）
func (obj *_SysTimersMgr) WithJobStatus(jobStatus int8) Option {
	return optionFunc(func(o *options) { o.query["job_status"] = jobStatus })
}

// WithRemark remark获取 备注信息
func (obj *_SysTimersMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysTimersMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysTimersMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysTimersMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysTimersMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysTimersMgr) GetByOption(opts ...Option) (result SysTimers, err error) {
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
func (obj *_SysTimersMgr) GetByOptions(opts ...Option) (results []*SysTimers, err error) {
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

// GetFromID 通过id获取内容 定时器id
func (obj *_SysTimersMgr) GetFromID(id int64) (result SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 定时器id
func (obj *_SysTimersMgr) GetBatchFromID(ids []int64) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTimerName 通过timer_name获取内容 任务名称
func (obj *_SysTimersMgr) GetFromTimerName(timerName string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("timer_name = ?", timerName).Find(&results).Error

	return
}

// GetBatchFromTimerName 批量唯一主键查找 任务名称
func (obj *_SysTimersMgr) GetBatchFromTimerName(timerNames []string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("timer_name IN (?)", timerNames).Find(&results).Error

	return
}

// GetFromActionClass 通过action_class获取内容 执行任务的class的类名（实现了TimerTaskRunner接口的类的全称）
func (obj *_SysTimersMgr) GetFromActionClass(actionClass string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("action_class = ?", actionClass).Find(&results).Error

	return
}

// GetBatchFromActionClass 批量唯一主键查找 执行任务的class的类名（实现了TimerTaskRunner接口的类的全称）
func (obj *_SysTimersMgr) GetBatchFromActionClass(actionClasss []string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("action_class IN (?)", actionClasss).Find(&results).Error

	return
}

// GetFromCron 通过cron获取内容 定时任务表达式
func (obj *_SysTimersMgr) GetFromCron(cron string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cron = ?", cron).Find(&results).Error

	return
}

// GetBatchFromCron 批量唯一主键查找 定时任务表达式
func (obj *_SysTimersMgr) GetBatchFromCron(crons []string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cron IN (?)", crons).Find(&results).Error

	return
}

// GetFromJobStatus 通过job_status获取内容 状态（字典 1运行  2停止）
func (obj *_SysTimersMgr) GetFromJobStatus(jobStatus int8) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_status = ?", jobStatus).Find(&results).Error

	return
}

// GetBatchFromJobStatus 批量唯一主键查找 状态（字典 1运行  2停止）
func (obj *_SysTimersMgr) GetBatchFromJobStatus(jobStatuss []int8) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_status IN (?)", jobStatuss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注信息
func (obj *_SysTimersMgr) GetFromRemark(remark string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注信息
func (obj *_SysTimersMgr) GetBatchFromRemark(remarks []string) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysTimersMgr) GetFromCreateTime(createTime time.Time) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysTimersMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysTimersMgr) GetFromCreateUser(createUser int64) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysTimersMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysTimersMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysTimersMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysTimersMgr) GetFromUpdateUser(updateUser int64) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_SysTimersMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysTimersMgr) FetchByPrimaryKey(id int64) (result SysTimers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
