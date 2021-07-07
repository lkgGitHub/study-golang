package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysSmsMgr struct {
	*_BaseMgr
}

// SysSmsMgr open func
func SysSmsMgr(db *gorm.DB) *_SysSmsMgr {
	if db == nil {
		panic(fmt.Errorf("SysSmsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysSmsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_sms"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysSmsMgr) GetTableName() string {
	return "sys_sms"
}

// Get 获取
func (obj *_SysSmsMgr) Get() (result SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysSmsMgr) Gets() (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysSmsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPhoneNumbers phone_numbers获取 手机号
func (obj *_SysSmsMgr) WithPhoneNumbers(phoneNumbers string) Option {
	return optionFunc(func(o *options) { o.query["phone_numbers"] = phoneNumbers })
}

// WithValidateCode validate_code获取 短信验证码
func (obj *_SysSmsMgr) WithValidateCode(validateCode string) Option {
	return optionFunc(func(o *options) { o.query["validate_code"] = validateCode })
}

// WithTemplateCode template_code获取 短信模板ID
func (obj *_SysSmsMgr) WithTemplateCode(templateCode string) Option {
	return optionFunc(func(o *options) { o.query["template_code"] = templateCode })
}

// WithBizID biz_id获取 回执id，可根据该id查询具体的发送状态
func (obj *_SysSmsMgr) WithBizID(bizID string) Option {
	return optionFunc(func(o *options) { o.query["biz_id"] = bizID })
}

// WithStatus status获取 发送状态（字典 0 未发送，1 发送成功，2 发送失败，3 失效）
func (obj *_SysSmsMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSource source获取 来源（字典 1 app， 2 pc， 3 其他）
func (obj *_SysSmsMgr) WithSource(source int8) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithInvalidTime invalid_time获取 失效时间
func (obj *_SysSmsMgr) WithInvalidTime(invalidTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["invalid_time"] = invalidTime })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysSmsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysSmsMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysSmsMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysSmsMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysSmsMgr) GetByOption(opts ...Option) (result SysSms, err error) {
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
func (obj *_SysSmsMgr) GetByOptions(opts ...Option) (results []*SysSms, err error) {
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
func (obj *_SysSmsMgr) GetFromID(id int64) (result SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysSmsMgr) GetBatchFromID(ids []int64) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPhoneNumbers 通过phone_numbers获取内容 手机号
func (obj *_SysSmsMgr) GetFromPhoneNumbers(phoneNumbers string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_numbers = ?", phoneNumbers).Find(&results).Error

	return
}

// GetBatchFromPhoneNumbers 批量唯一主键查找 手机号
func (obj *_SysSmsMgr) GetBatchFromPhoneNumbers(phoneNumberss []string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_numbers IN (?)", phoneNumberss).Find(&results).Error

	return
}

// GetFromValidateCode 通过validate_code获取内容 短信验证码
func (obj *_SysSmsMgr) GetFromValidateCode(validateCode string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validate_code = ?", validateCode).Find(&results).Error

	return
}

// GetBatchFromValidateCode 批量唯一主键查找 短信验证码
func (obj *_SysSmsMgr) GetBatchFromValidateCode(validateCodes []string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validate_code IN (?)", validateCodes).Find(&results).Error

	return
}

// GetFromTemplateCode 通过template_code获取内容 短信模板ID
func (obj *_SysSmsMgr) GetFromTemplateCode(templateCode string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_code = ?", templateCode).Find(&results).Error

	return
}

// GetBatchFromTemplateCode 批量唯一主键查找 短信模板ID
func (obj *_SysSmsMgr) GetBatchFromTemplateCode(templateCodes []string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_code IN (?)", templateCodes).Find(&results).Error

	return
}

// GetFromBizID 通过biz_id获取内容 回执id，可根据该id查询具体的发送状态
func (obj *_SysSmsMgr) GetFromBizID(bizID string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("biz_id = ?", bizID).Find(&results).Error

	return
}

// GetBatchFromBizID 批量唯一主键查找 回执id，可根据该id查询具体的发送状态
func (obj *_SysSmsMgr) GetBatchFromBizID(bizIDs []string) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("biz_id IN (?)", bizIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 发送状态（字典 0 未发送，1 发送成功，2 发送失败，3 失效）
func (obj *_SysSmsMgr) GetFromStatus(status int8) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 发送状态（字典 0 未发送，1 发送成功，2 发送失败，3 失效）
func (obj *_SysSmsMgr) GetBatchFromStatus(statuss []int8) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 来源（字典 1 app， 2 pc， 3 其他）
func (obj *_SysSmsMgr) GetFromSource(source int8) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("source = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量唯一主键查找 来源（字典 1 app， 2 pc， 3 其他）
func (obj *_SysSmsMgr) GetBatchFromSource(sources []int8) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("source IN (?)", sources).Find(&results).Error

	return
}

// GetFromInvalidTime 通过invalid_time获取内容 失效时间
func (obj *_SysSmsMgr) GetFromInvalidTime(invalidTime time.Time) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invalid_time = ?", invalidTime).Find(&results).Error

	return
}

// GetBatchFromInvalidTime 批量唯一主键查找 失效时间
func (obj *_SysSmsMgr) GetBatchFromInvalidTime(invalidTimes []time.Time) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invalid_time IN (?)", invalidTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysSmsMgr) GetFromCreateTime(createTime time.Time) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysSmsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysSmsMgr) GetFromCreateUser(createUser int64) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysSmsMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysSmsMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysSmsMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysSmsMgr) GetFromUpdateUser(updateUser int64) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_SysSmsMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysSmsMgr) FetchByPrimaryKey(id int64) (result SysSms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
