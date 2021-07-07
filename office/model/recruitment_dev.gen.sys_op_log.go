package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysOpLogMgr struct {
	*_BaseMgr
}

// SysOpLogMgr open func
func SysOpLogMgr(db *gorm.DB) *_SysOpLogMgr {
	if db == nil {
		panic(fmt.Errorf("SysOpLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysOpLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_op_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysOpLogMgr) GetTableName() string {
	return "sys_op_log"
}

// Get 获取
func (obj *_SysOpLogMgr) Get() (result SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysOpLogMgr) Gets() (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysOpLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_SysOpLogMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOpType op_type获取 操作类型
func (obj *_SysOpLogMgr) WithOpType(opType int8) Option {
	return optionFunc(func(o *options) { o.query["op_type"] = opType })
}

// WithSuccess success获取 是否执行成功（Y-是，N-否）
func (obj *_SysOpLogMgr) WithSuccess(success string) Option {
	return optionFunc(func(o *options) { o.query["success"] = success })
}

// WithMessage message获取 具体消息
func (obj *_SysOpLogMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithIP ip获取 ip
func (obj *_SysOpLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithLocation location获取 地址
func (obj *_SysOpLogMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithBrowser browser获取 浏览器
func (obj *_SysOpLogMgr) WithBrowser(browser string) Option {
	return optionFunc(func(o *options) { o.query["browser"] = browser })
}

// WithOs os获取 操作系统
func (obj *_SysOpLogMgr) WithOs(os string) Option {
	return optionFunc(func(o *options) { o.query["os"] = os })
}

// WithURL url获取 请求地址
func (obj *_SysOpLogMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithClassName class_name获取 类名称
func (obj *_SysOpLogMgr) WithClassName(className string) Option {
	return optionFunc(func(o *options) { o.query["class_name"] = className })
}

// WithMethodName method_name获取 方法名称
func (obj *_SysOpLogMgr) WithMethodName(methodName string) Option {
	return optionFunc(func(o *options) { o.query["method_name"] = methodName })
}

// WithReqMethod req_method获取 请求方式（GET POST PUT DELETE)
func (obj *_SysOpLogMgr) WithReqMethod(reqMethod string) Option {
	return optionFunc(func(o *options) { o.query["req_method"] = reqMethod })
}

// WithParam param获取 请求参数
func (obj *_SysOpLogMgr) WithParam(param string) Option {
	return optionFunc(func(o *options) { o.query["param"] = param })
}

// WithResult result获取 返回结果
func (obj *_SysOpLogMgr) WithResult(result string) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithOpTime op_time获取 操作时间
func (obj *_SysOpLogMgr) WithOpTime(opTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["op_time"] = opTime })
}

// WithAccount account获取 操作账号
func (obj *_SysOpLogMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// GetByOption 功能选项模式获取
func (obj *_SysOpLogMgr) GetByOption(opts ...Option) (result SysOpLog, err error) {
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
func (obj *_SysOpLogMgr) GetByOptions(opts ...Option) (results []*SysOpLog, err error) {
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
func (obj *_SysOpLogMgr) GetFromID(id int64) (result SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysOpLogMgr) GetBatchFromID(ids []int64) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_SysOpLogMgr) GetFromName(name string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_SysOpLogMgr) GetBatchFromName(names []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromOpType 通过op_type获取内容 操作类型
func (obj *_SysOpLogMgr) GetFromOpType(opType int8) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("op_type = ?", opType).Find(&results).Error

	return
}

// GetBatchFromOpType 批量唯一主键查找 操作类型
func (obj *_SysOpLogMgr) GetBatchFromOpType(opTypes []int8) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("op_type IN (?)", opTypes).Find(&results).Error

	return
}

// GetFromSuccess 通过success获取内容 是否执行成功（Y-是，N-否）
func (obj *_SysOpLogMgr) GetFromSuccess(success string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("success = ?", success).Find(&results).Error

	return
}

// GetBatchFromSuccess 批量唯一主键查找 是否执行成功（Y-是，N-否）
func (obj *_SysOpLogMgr) GetBatchFromSuccess(successs []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("success IN (?)", successs).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容 具体消息
func (obj *_SysOpLogMgr) GetFromMessage(message string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量唯一主键查找 具体消息
func (obj *_SysOpLogMgr) GetBatchFromMessage(messages []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip
func (obj *_SysOpLogMgr) GetFromIP(ip string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量唯一主键查找 ip
func (obj *_SysOpLogMgr) GetBatchFromIP(ips []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容 地址
func (obj *_SysOpLogMgr) GetFromLocation(location string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找 地址
func (obj *_SysOpLogMgr) GetBatchFromLocation(locations []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

// GetFromBrowser 通过browser获取内容 浏览器
func (obj *_SysOpLogMgr) GetFromBrowser(browser string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser = ?", browser).Find(&results).Error

	return
}

// GetBatchFromBrowser 批量唯一主键查找 浏览器
func (obj *_SysOpLogMgr) GetBatchFromBrowser(browsers []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser IN (?)", browsers).Find(&results).Error

	return
}

// GetFromOs 通过os获取内容 操作系统
func (obj *_SysOpLogMgr) GetFromOs(os string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os = ?", os).Find(&results).Error

	return
}

// GetBatchFromOs 批量唯一主键查找 操作系统
func (obj *_SysOpLogMgr) GetBatchFromOs(oss []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os IN (?)", oss).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 请求地址
func (obj *_SysOpLogMgr) GetFromURL(url string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找 请求地址
func (obj *_SysOpLogMgr) GetBatchFromURL(urls []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromClassName 通过class_name获取内容 类名称
func (obj *_SysOpLogMgr) GetFromClassName(className string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name = ?", className).Find(&results).Error

	return
}

// GetBatchFromClassName 批量唯一主键查找 类名称
func (obj *_SysOpLogMgr) GetBatchFromClassName(classNames []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name IN (?)", classNames).Find(&results).Error

	return
}

// GetFromMethodName 通过method_name获取内容 方法名称
func (obj *_SysOpLogMgr) GetFromMethodName(methodName string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_name = ?", methodName).Find(&results).Error

	return
}

// GetBatchFromMethodName 批量唯一主键查找 方法名称
func (obj *_SysOpLogMgr) GetBatchFromMethodName(methodNames []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_name IN (?)", methodNames).Find(&results).Error

	return
}

// GetFromReqMethod 通过req_method获取内容 请求方式（GET POST PUT DELETE)
func (obj *_SysOpLogMgr) GetFromReqMethod(reqMethod string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("req_method = ?", reqMethod).Find(&results).Error

	return
}

// GetBatchFromReqMethod 批量唯一主键查找 请求方式（GET POST PUT DELETE)
func (obj *_SysOpLogMgr) GetBatchFromReqMethod(reqMethods []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("req_method IN (?)", reqMethods).Find(&results).Error

	return
}

// GetFromParam 通过param获取内容 请求参数
func (obj *_SysOpLogMgr) GetFromParam(param string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("param = ?", param).Find(&results).Error

	return
}

// GetBatchFromParam 批量唯一主键查找 请求参数
func (obj *_SysOpLogMgr) GetBatchFromParam(params []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("param IN (?)", params).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容 返回结果
func (obj *_SysOpLogMgr) GetFromResult(result string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("result = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量唯一主键查找 返回结果
func (obj *_SysOpLogMgr) GetBatchFromResult(r []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("result IN (?)", r).Find(&results).Error

	return
}

// GetFromOpTime 通过op_time获取内容 操作时间
func (obj *_SysOpLogMgr) GetFromOpTime(opTime time.Time) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("op_time = ?", opTime).Find(&results).Error

	return
}

// GetBatchFromOpTime 批量唯一主键查找 操作时间
func (obj *_SysOpLogMgr) GetBatchFromOpTime(opTimes []time.Time) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("op_time IN (?)", opTimes).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 操作账号
func (obj *_SysOpLogMgr) GetFromAccount(account string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量唯一主键查找 操作账号
func (obj *_SysOpLogMgr) GetBatchFromAccount(accounts []string) (results []*SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account IN (?)", accounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysOpLogMgr) FetchByPrimaryKey(id int64) (result SysOpLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
