package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysConfigMgr struct {
	*_BaseMgr
}

// SysConfigMgr open func
func SysConfigMgr(db *gorm.DB) *_SysConfigMgr {
	if db == nil {
		panic(fmt.Errorf("SysConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysConfigMgr) GetTableName() string {
	return "sys_config"
}

// Get 获取
func (obj *_SysConfigMgr) Get() (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysConfigMgr) Gets() (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysConfigMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_SysConfigMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 编码
func (obj *_SysConfigMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithValue value获取 值
func (obj *_SysConfigMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithSysFlag sys_flag获取 是否是系统参数（Y-是，N-否）
func (obj *_SysConfigMgr) WithSysFlag(sysFlag string) Option {
	return optionFunc(func(o *options) { o.query["sys_flag"] = sysFlag })
}

// WithRemark remark获取 备注
func (obj *_SysConfigMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 状态（字典 0正常 1停用 2删除）
func (obj *_SysConfigMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithGroupCode group_code获取 常量所属分类的编码，来自于“常量的分类”字典
func (obj *_SysConfigMgr) WithGroupCode(groupCode string) Option {
	return optionFunc(func(o *options) { o.query["group_code"] = groupCode })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysConfigMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysConfigMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysConfigMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysConfigMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysConfigMgr) GetByOption(opts ...Option) (result SysConfig, err error) {
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
func (obj *_SysConfigMgr) GetByOptions(opts ...Option) (results []*SysConfig, err error) {
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
func (obj *_SysConfigMgr) GetFromID(id int64) (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysConfigMgr) GetBatchFromID(ids []int64) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_SysConfigMgr) GetFromName(name string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_SysConfigMgr) GetBatchFromName(names []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 编码
func (obj *_SysConfigMgr) GetFromCode(code string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量唯一主键查找 编码
func (obj *_SysConfigMgr) GetBatchFromCode(codes []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 值
func (obj *_SysConfigMgr) GetFromValue(value string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 值
func (obj *_SysConfigMgr) GetBatchFromValue(values []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

// GetFromSysFlag 通过sys_flag获取内容 是否是系统参数（Y-是，N-否）
func (obj *_SysConfigMgr) GetFromSysFlag(sysFlag string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sys_flag = ?", sysFlag).Find(&results).Error

	return
}

// GetBatchFromSysFlag 批量唯一主键查找 是否是系统参数（Y-是，N-否）
func (obj *_SysConfigMgr) GetBatchFromSysFlag(sysFlags []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sys_flag IN (?)", sysFlags).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysConfigMgr) GetFromRemark(remark string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysConfigMgr) GetBatchFromRemark(remarks []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0正常 1停用 2删除）
func (obj *_SysConfigMgr) GetFromStatus(status int8) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0正常 1停用 2删除）
func (obj *_SysConfigMgr) GetBatchFromStatus(statuss []int8) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromGroupCode 通过group_code获取内容 常量所属分类的编码，来自于“常量的分类”字典
func (obj *_SysConfigMgr) GetFromGroupCode(groupCode string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_code = ?", groupCode).Find(&results).Error

	return
}

// GetBatchFromGroupCode 批量唯一主键查找 常量所属分类的编码，来自于“常量的分类”字典
func (obj *_SysConfigMgr) GetBatchFromGroupCode(groupCodes []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_code IN (?)", groupCodes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysConfigMgr) GetFromCreateTime(createTime time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysConfigMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysConfigMgr) GetFromCreateUser(createUser int64) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysConfigMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysConfigMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysConfigMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysConfigMgr) GetFromUpdateUser(updateUser int64) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_SysConfigMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysConfigMgr) FetchByPrimaryKey(id int64) (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
