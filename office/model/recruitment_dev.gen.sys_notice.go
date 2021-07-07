package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysNoticeMgr struct {
	*_BaseMgr
}

// SysNoticeMgr open func
func SysNoticeMgr(db *gorm.DB) *_SysNoticeMgr {
	if db == nil {
		panic(fmt.Errorf("SysNoticeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysNoticeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_notice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysNoticeMgr) GetTableName() string {
	return "sys_notice"
}

// Get 获取
func (obj *_SysNoticeMgr) Get() (result SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysNoticeMgr) Gets() (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysNoticeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 标题
func (obj *_SysNoticeMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_SysNoticeMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithType type获取 类型（字典 1通知 2公告）
func (obj *_SysNoticeMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithPublicUserID public_user_id获取 发布人id
func (obj *_SysNoticeMgr) WithPublicUserID(publicUserID int64) Option {
	return optionFunc(func(o *options) { o.query["public_user_id"] = publicUserID })
}

// WithPublicUserName public_user_name获取 发布人姓名
func (obj *_SysNoticeMgr) WithPublicUserName(publicUserName string) Option {
	return optionFunc(func(o *options) { o.query["public_user_name"] = publicUserName })
}

// WithPublicOrgID public_org_id获取 发布机构id
func (obj *_SysNoticeMgr) WithPublicOrgID(publicOrgID int64) Option {
	return optionFunc(func(o *options) { o.query["public_org_id"] = publicOrgID })
}

// WithPublicOrgName public_org_name获取 发布机构名称
func (obj *_SysNoticeMgr) WithPublicOrgName(publicOrgName string) Option {
	return optionFunc(func(o *options) { o.query["public_org_name"] = publicOrgName })
}

// WithPublicTime public_time获取 发布时间
func (obj *_SysNoticeMgr) WithPublicTime(publicTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["public_time"] = publicTime })
}

// WithCancelTime cancel_time获取 撤回时间
func (obj *_SysNoticeMgr) WithCancelTime(cancelTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["cancel_time"] = cancelTime })
}

// WithStatus status获取 状态（字典 0草稿 1发布 2撤回 3删除）
func (obj *_SysNoticeMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysNoticeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysNoticeMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SysNoticeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 修改人
func (obj *_SysNoticeMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysNoticeMgr) GetByOption(opts ...Option) (result SysNotice, err error) {
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
func (obj *_SysNoticeMgr) GetByOptions(opts ...Option) (results []*SysNotice, err error) {
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
func (obj *_SysNoticeMgr) GetFromID(id int64) (result SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysNoticeMgr) GetBatchFromID(ids []int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_SysNoticeMgr) GetFromTitle(title string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 标题
func (obj *_SysNoticeMgr) GetBatchFromTitle(titles []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_SysNoticeMgr) GetFromContent(content string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 内容
func (obj *_SysNoticeMgr) GetBatchFromContent(contents []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型（字典 1通知 2公告）
func (obj *_SysNoticeMgr) GetFromType(_type int8) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型（字典 1通知 2公告）
func (obj *_SysNoticeMgr) GetBatchFromType(_types []int8) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromPublicUserID 通过public_user_id获取内容 发布人id
func (obj *_SysNoticeMgr) GetFromPublicUserID(publicUserID int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_user_id = ?", publicUserID).Find(&results).Error

	return
}

// GetBatchFromPublicUserID 批量唯一主键查找 发布人id
func (obj *_SysNoticeMgr) GetBatchFromPublicUserID(publicUserIDs []int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_user_id IN (?)", publicUserIDs).Find(&results).Error

	return
}

// GetFromPublicUserName 通过public_user_name获取内容 发布人姓名
func (obj *_SysNoticeMgr) GetFromPublicUserName(publicUserName string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_user_name = ?", publicUserName).Find(&results).Error

	return
}

// GetBatchFromPublicUserName 批量唯一主键查找 发布人姓名
func (obj *_SysNoticeMgr) GetBatchFromPublicUserName(publicUserNames []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_user_name IN (?)", publicUserNames).Find(&results).Error

	return
}

// GetFromPublicOrgID 通过public_org_id获取内容 发布机构id
func (obj *_SysNoticeMgr) GetFromPublicOrgID(publicOrgID int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_org_id = ?", publicOrgID).Find(&results).Error

	return
}

// GetBatchFromPublicOrgID 批量唯一主键查找 发布机构id
func (obj *_SysNoticeMgr) GetBatchFromPublicOrgID(publicOrgIDs []int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_org_id IN (?)", publicOrgIDs).Find(&results).Error

	return
}

// GetFromPublicOrgName 通过public_org_name获取内容 发布机构名称
func (obj *_SysNoticeMgr) GetFromPublicOrgName(publicOrgName string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_org_name = ?", publicOrgName).Find(&results).Error

	return
}

// GetBatchFromPublicOrgName 批量唯一主键查找 发布机构名称
func (obj *_SysNoticeMgr) GetBatchFromPublicOrgName(publicOrgNames []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_org_name IN (?)", publicOrgNames).Find(&results).Error

	return
}

// GetFromPublicTime 通过public_time获取内容 发布时间
func (obj *_SysNoticeMgr) GetFromPublicTime(publicTime time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_time = ?", publicTime).Find(&results).Error

	return
}

// GetBatchFromPublicTime 批量唯一主键查找 发布时间
func (obj *_SysNoticeMgr) GetBatchFromPublicTime(publicTimes []time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_time IN (?)", publicTimes).Find(&results).Error

	return
}

// GetFromCancelTime 通过cancel_time获取内容 撤回时间
func (obj *_SysNoticeMgr) GetFromCancelTime(cancelTime time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cancel_time = ?", cancelTime).Find(&results).Error

	return
}

// GetBatchFromCancelTime 批量唯一主键查找 撤回时间
func (obj *_SysNoticeMgr) GetBatchFromCancelTime(cancelTimes []time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cancel_time IN (?)", cancelTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0草稿 1发布 2撤回 3删除）
func (obj *_SysNoticeMgr) GetFromStatus(status int8) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0草稿 1发布 2撤回 3删除）
func (obj *_SysNoticeMgr) GetBatchFromStatus(statuss []int8) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysNoticeMgr) GetFromCreateTime(createTime time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysNoticeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysNoticeMgr) GetFromCreateUser(createUser int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysNoticeMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_SysNoticeMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_SysNoticeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改人
func (obj *_SysNoticeMgr) GetFromUpdateUser(updateUser int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_SysNoticeMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysNoticeMgr) FetchByPrimaryKey(id int64) (result SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
