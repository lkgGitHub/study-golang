package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysVisLogMgr struct {
	*_BaseMgr
}

// SysVisLogMgr open func
func SysVisLogMgr(db *gorm.DB) *_SysVisLogMgr {
	if db == nil {
		panic(fmt.Errorf("SysVisLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysVisLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_vis_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysVisLogMgr) GetTableName() string {
	return "sys_vis_log"
}

// Get 获取
func (obj *_SysVisLogMgr) Get() (result SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysVisLogMgr) Gets() (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysVisLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_SysVisLogMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSuccess success获取 是否执行成功（Y-是，N-否）
func (obj *_SysVisLogMgr) WithSuccess(success string) Option {
	return optionFunc(func(o *options) { o.query["success"] = success })
}

// WithMessage message获取 具体消息
func (obj *_SysVisLogMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithIP ip获取 ip
func (obj *_SysVisLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithLocation location获取 地址
func (obj *_SysVisLogMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithBrowser browser获取 浏览器
func (obj *_SysVisLogMgr) WithBrowser(browser string) Option {
	return optionFunc(func(o *options) { o.query["browser"] = browser })
}

// WithOs os获取 操作系统
func (obj *_SysVisLogMgr) WithOs(os string) Option {
	return optionFunc(func(o *options) { o.query["os"] = os })
}

// WithVisType vis_type获取 操作类型（字典 1登入 2登出）
func (obj *_SysVisLogMgr) WithVisType(visType int8) Option {
	return optionFunc(func(o *options) { o.query["vis_type"] = visType })
}

// WithVisTime vis_time获取 访问时间
func (obj *_SysVisLogMgr) WithVisTime(visTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["vis_time"] = visTime })
}

// WithAccount account获取 访问账号
func (obj *_SysVisLogMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// GetByOption 功能选项模式获取
func (obj *_SysVisLogMgr) GetByOption(opts ...Option) (result SysVisLog, err error) {
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
func (obj *_SysVisLogMgr) GetByOptions(opts ...Option) (results []*SysVisLog, err error) {
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
func (obj *_SysVisLogMgr) GetFromID(id int64) (result SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysVisLogMgr) GetBatchFromID(ids []int64) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_SysVisLogMgr) GetFromName(name string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_SysVisLogMgr) GetBatchFromName(names []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromSuccess 通过success获取内容 是否执行成功（Y-是，N-否）
func (obj *_SysVisLogMgr) GetFromSuccess(success string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("success = ?", success).Find(&results).Error

	return
}

// GetBatchFromSuccess 批量唯一主键查找 是否执行成功（Y-是，N-否）
func (obj *_SysVisLogMgr) GetBatchFromSuccess(successs []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("success IN (?)", successs).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容 具体消息
func (obj *_SysVisLogMgr) GetFromMessage(message string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量唯一主键查找 具体消息
func (obj *_SysVisLogMgr) GetBatchFromMessage(messages []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip
func (obj *_SysVisLogMgr) GetFromIP(ip string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量唯一主键查找 ip
func (obj *_SysVisLogMgr) GetBatchFromIP(ips []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容 地址
func (obj *_SysVisLogMgr) GetFromLocation(location string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找 地址
func (obj *_SysVisLogMgr) GetBatchFromLocation(locations []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

// GetFromBrowser 通过browser获取内容 浏览器
func (obj *_SysVisLogMgr) GetFromBrowser(browser string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser = ?", browser).Find(&results).Error

	return
}

// GetBatchFromBrowser 批量唯一主键查找 浏览器
func (obj *_SysVisLogMgr) GetBatchFromBrowser(browsers []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser IN (?)", browsers).Find(&results).Error

	return
}

// GetFromOs 通过os获取内容 操作系统
func (obj *_SysVisLogMgr) GetFromOs(os string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os = ?", os).Find(&results).Error

	return
}

// GetBatchFromOs 批量唯一主键查找 操作系统
func (obj *_SysVisLogMgr) GetBatchFromOs(oss []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os IN (?)", oss).Find(&results).Error

	return
}

// GetFromVisType 通过vis_type获取内容 操作类型（字典 1登入 2登出）
func (obj *_SysVisLogMgr) GetFromVisType(visType int8) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vis_type = ?", visType).Find(&results).Error

	return
}

// GetBatchFromVisType 批量唯一主键查找 操作类型（字典 1登入 2登出）
func (obj *_SysVisLogMgr) GetBatchFromVisType(visTypes []int8) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vis_type IN (?)", visTypes).Find(&results).Error

	return
}

// GetFromVisTime 通过vis_time获取内容 访问时间
func (obj *_SysVisLogMgr) GetFromVisTime(visTime time.Time) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vis_time = ?", visTime).Find(&results).Error

	return
}

// GetBatchFromVisTime 批量唯一主键查找 访问时间
func (obj *_SysVisLogMgr) GetBatchFromVisTime(visTimes []time.Time) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vis_time IN (?)", visTimes).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 访问账号
func (obj *_SysVisLogMgr) GetFromAccount(account string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量唯一主键查找 访问账号
func (obj *_SysVisLogMgr) GetBatchFromAccount(accounts []string) (results []*SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account IN (?)", accounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysVisLogMgr) FetchByPrimaryKey(id int64) (result SysVisLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
