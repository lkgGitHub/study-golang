package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ResumeMgr struct {
	*_BaseMgr
}

// ResumeMgr open func
func ResumeMgr(db *gorm.DB) *_ResumeMgr {
	if db == nil {
		panic(fmt.Errorf("ResumeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ResumeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("resume"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ResumeMgr) GetTableName() string {
	return "resume"
}

// Get 获取
func (obj *_ResumeMgr) Get() (result Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ResumeMgr) Gets() (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ResumeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithResumePath resume_path获取 简历地址
func (obj *_ResumeMgr) WithResumePath(resumePath string) Option {
	return optionFunc(func(o *options) { o.query["resume_path"] = resumePath })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_ResumeMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// WithState state获取 简历状态(0:有效，1:删除)
func (obj *_ResumeMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithResumeName resume_name获取 简历文件名称
func (obj *_ResumeMgr) WithResumeName(resumeName string) Option {
	return optionFunc(func(o *options) { o.query["resume_name"] = resumeName })
}

// WithUserPhone user_phone获取 用户手机号
func (obj *_ResumeMgr) WithUserPhone(userPhone string) Option {
	return optionFunc(func(o *options) { o.query["user_phone"] = userPhone })
}

// WithFormat format获取 文件格式
func (obj *_ResumeMgr) WithFormat(format string) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// WithSize size获取 简历大小
func (obj *_ResumeMgr) WithSize(size string) Option {
	return optionFunc(func(o *options) { o.query["size"] = size })
}

// WithCreateTime create_time获取 创建时间
func (obj *_ResumeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_ResumeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_ResumeMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_ResumeMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_ResumeMgr) GetByOption(opts ...Option) (result Resume, err error) {
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
func (obj *_ResumeMgr) GetByOptions(opts ...Option) (results []*Resume, err error) {
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
func (obj *_ResumeMgr) GetFromID(id int64) (result Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ResumeMgr) GetBatchFromID(ids []int64) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromResumePath 通过resume_path获取内容 简历地址
func (obj *_ResumeMgr) GetFromResumePath(resumePath string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resume_path = ?", resumePath).Find(&results).Error

	return
}

// GetBatchFromResumePath 批量唯一主键查找 简历地址
func (obj *_ResumeMgr) GetBatchFromResumePath(resumePaths []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resume_path IN (?)", resumePaths).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_ResumeMgr) GetFromDeleteTime(deleteTime time.Time) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找 删除时间
func (obj *_ResumeMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 简历状态(0:有效，1:删除)
func (obj *_ResumeMgr) GetFromState(state string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量唯一主键查找 简历状态(0:有效，1:删除)
func (obj *_ResumeMgr) GetBatchFromState(states []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

// GetFromResumeName 通过resume_name获取内容 简历文件名称
func (obj *_ResumeMgr) GetFromResumeName(resumeName string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resume_name = ?", resumeName).Find(&results).Error

	return
}

// GetBatchFromResumeName 批量唯一主键查找 简历文件名称
func (obj *_ResumeMgr) GetBatchFromResumeName(resumeNames []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resume_name IN (?)", resumeNames).Find(&results).Error

	return
}

// GetFromUserPhone 通过user_phone获取内容 用户手机号
func (obj *_ResumeMgr) GetFromUserPhone(userPhone string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_phone = ?", userPhone).Find(&results).Error

	return
}

// GetBatchFromUserPhone 批量唯一主键查找 用户手机号
func (obj *_ResumeMgr) GetBatchFromUserPhone(userPhones []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_phone IN (?)", userPhones).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容 文件格式
func (obj *_ResumeMgr) GetFromFormat(format string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("format = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量唯一主键查找 文件格式
func (obj *_ResumeMgr) GetBatchFromFormat(formats []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("format IN (?)", formats).Find(&results).Error

	return
}

// GetFromSize 通过size获取内容 简历大小
func (obj *_ResumeMgr) GetFromSize(size string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("size = ?", size).Find(&results).Error

	return
}

// GetBatchFromSize 批量唯一主键查找 简历大小
func (obj *_ResumeMgr) GetBatchFromSize(sizes []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("size IN (?)", sizes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_ResumeMgr) GetFromCreateTime(createTime time.Time) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_ResumeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_ResumeMgr) GetFromUpdateTime(updateTime time.Time) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_ResumeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_ResumeMgr) GetFromCreateUser(createUser string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_ResumeMgr) GetBatchFromCreateUser(createUsers []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改人
func (obj *_ResumeMgr) GetFromUpdateUser(updateUser string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_ResumeMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ResumeMgr) FetchByPrimaryKey(id int64) (result Resume, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
