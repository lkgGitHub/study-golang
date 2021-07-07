package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysDictDataMgr struct {
	*_BaseMgr
}

// SysDictDataMgr open func
func SysDictDataMgr(db *gorm.DB) *_SysDictDataMgr {
	if db == nil {
		panic(fmt.Errorf("SysDictDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysDictDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_dict_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysDictDataMgr) GetTableName() string {
	return "sys_dict_data"
}

// Get 获取
func (obj *_SysDictDataMgr) Get() (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysDictDataMgr) Gets() (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysDictDataMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTypeID type_id获取 字典类型id
func (obj *_SysDictDataMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithValue value获取 值
func (obj *_SysDictDataMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithCode code获取 编码
func (obj *_SysDictDataMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithSort sort获取 排序
func (obj *_SysDictDataMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SysDictDataMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 状态（字典 0正常 1停用 2删除）
func (obj *_SysDictDataMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysDictDataMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysDictDataMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysDictDataMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysDictDataMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysDictDataMgr) GetByOption(opts ...Option) (result SysDictData, err error) {
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
func (obj *_SysDictDataMgr) GetByOptions(opts ...Option) (results []*SysDictData, err error) {
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
func (obj *_SysDictDataMgr) GetFromID(id int64) (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysDictDataMgr) GetBatchFromID(ids []int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容 字典类型id
func (obj *_SysDictDataMgr) GetFromTypeID(typeID int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量唯一主键查找 字典类型id
func (obj *_SysDictDataMgr) GetBatchFromTypeID(typeIDs []int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_id IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 值
func (obj *_SysDictDataMgr) GetFromValue(value string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 值
func (obj *_SysDictDataMgr) GetBatchFromValue(values []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 编码
func (obj *_SysDictDataMgr) GetFromCode(code string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量唯一主键查找 编码
func (obj *_SysDictDataMgr) GetBatchFromCode(codes []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SysDictDataMgr) GetFromSort(sort int) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序
func (obj *_SysDictDataMgr) GetBatchFromSort(sorts []int) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysDictDataMgr) GetFromRemark(remark string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysDictDataMgr) GetBatchFromRemark(remarks []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0正常 1停用 2删除）
func (obj *_SysDictDataMgr) GetFromStatus(status int8) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0正常 1停用 2删除）
func (obj *_SysDictDataMgr) GetBatchFromStatus(statuss []int8) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysDictDataMgr) GetFromCreateTime(createTime time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysDictDataMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysDictDataMgr) GetFromCreateUser(createUser int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysDictDataMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysDictDataMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysDictDataMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysDictDataMgr) GetFromUpdateUser(updateUser int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_SysDictDataMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysDictDataMgr) FetchByPrimaryKey(id int64) (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
