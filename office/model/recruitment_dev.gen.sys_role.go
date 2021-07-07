package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysRoleMgr struct {
	*_BaseMgr
}

// SysRoleMgr open func
func SysRoleMgr(db *gorm.DB) *_SysRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleMgr) GetTableName() string {
	return "sys_role"
}

// Get 获取
func (obj *_SysRoleMgr) Get() (result SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleMgr) Gets() (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_SysRoleMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_SysRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 编码
func (obj *_SysRoleMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithSort sort获取 序号
func (obj *_SysRoleMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithDataScopeType data_scope_type获取 数据范围类型（字典 1全部数据 2本部门及以下数据 3本部门数据 4仅本人数据 5自定义数据）
func (obj *_SysRoleMgr) WithDataScopeType(dataScopeType int8) Option {
	return optionFunc(func(o *options) { o.query["data_scope_type"] = dataScopeType })
}

// WithRemark remark获取 备注
func (obj *_SysRoleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 状态（字典 0正常 1停用 2删除）
func (obj *_SysRoleMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysRoleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysRoleMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysRoleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysRoleMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleMgr) GetByOption(opts ...Option) (result SysRole, err error) {
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
func (obj *_SysRoleMgr) GetByOptions(opts ...Option) (results []*SysRole, err error) {
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
func (obj *_SysRoleMgr) GetFromID(id int64) (result SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键id
func (obj *_SysRoleMgr) GetBatchFromID(ids []int64) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_SysRoleMgr) GetFromName(name string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_SysRoleMgr) GetBatchFromName(names []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 编码
func (obj *_SysRoleMgr) GetFromCode(code string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量唯一主键查找 编码
func (obj *_SysRoleMgr) GetBatchFromCode(codes []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 序号
func (obj *_SysRoleMgr) GetFromSort(sort int) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 序号
func (obj *_SysRoleMgr) GetBatchFromSort(sorts []int) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromDataScopeType 通过data_scope_type获取内容 数据范围类型（字典 1全部数据 2本部门及以下数据 3本部门数据 4仅本人数据 5自定义数据）
func (obj *_SysRoleMgr) GetFromDataScopeType(dataScopeType int8) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data_scope_type = ?", dataScopeType).Find(&results).Error

	return
}

// GetBatchFromDataScopeType 批量唯一主键查找 数据范围类型（字典 1全部数据 2本部门及以下数据 3本部门数据 4仅本人数据 5自定义数据）
func (obj *_SysRoleMgr) GetBatchFromDataScopeType(dataScopeTypes []int8) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data_scope_type IN (?)", dataScopeTypes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysRoleMgr) GetFromRemark(remark string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysRoleMgr) GetBatchFromRemark(remarks []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0正常 1停用 2删除）
func (obj *_SysRoleMgr) GetFromStatus(status int8) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0正常 1停用 2删除）
func (obj *_SysRoleMgr) GetBatchFromStatus(statuss []int8) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysRoleMgr) GetFromCreateTime(createTime time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysRoleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysRoleMgr) GetFromCreateUser(createUser int64) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysRoleMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysRoleMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysRoleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysRoleMgr) GetFromUpdateUser(updateUser int64) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_SysRoleMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleMgr) FetchByPrimaryKey(id int64) (result SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
