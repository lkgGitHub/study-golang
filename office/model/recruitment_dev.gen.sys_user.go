package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _SysUserMgr struct {
	*_BaseMgr
}

// SysUserMgr open func
func SysUserMgr(db *gorm.DB) *_SysUserMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserMgr) GetTableName() string {
	return "sys_user"
}

// Get 获取
func (obj *_SysUserMgr) Get() (result SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysUserMgr) Gets() (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysUserMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAccount account获取 账号
func (obj *_SysUserMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithPassword password获取 密码
func (obj *_SysUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithNickName nick_name获取 昵称
func (obj *_SysUserMgr) WithNickName(nickName string) Option {
	return optionFunc(func(o *options) { o.query["nick_name"] = nickName })
}

// WithName name获取 姓名
func (obj *_SysUserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAvatar avatar获取 头像
func (obj *_SysUserMgr) WithAvatar(avatar int64) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithBirthday birthday获取 生日
func (obj *_SysUserMgr) WithBirthday(birthday datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithSex sex获取 性别(字典 1男 2女 3未知)
func (obj *_SysUserMgr) WithSex(sex int8) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithEmail email获取 邮箱
func (obj *_SysUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPhone phone获取 手机
func (obj *_SysUserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithTel tel获取 电话
func (obj *_SysUserMgr) WithTel(tel string) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithLastLoginIP last_login_ip获取 最后登陆IP
func (obj *_SysUserMgr) WithLastLoginIP(lastLoginIP string) Option {
	return optionFunc(func(o *options) { o.query["last_login_ip"] = lastLoginIP })
}

// WithLastLoginTime last_login_time获取 最后登陆时间
func (obj *_SysUserMgr) WithLastLoginTime(lastLoginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_login_time"] = lastLoginTime })
}

// WithAdminType admin_type获取 管理员类型（0超级管理员 1非管理员）
func (obj *_SysUserMgr) WithAdminType(adminType int8) Option {
	return optionFunc(func(o *options) { o.query["admin_type"] = adminType })
}

// WithStatus status获取 状态（字典 0正常 1冻结 2删除）
func (obj *_SysUserMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysUserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysUserMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysUserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysUserMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysUserMgr) GetByOption(opts ...Option) (result SysUser, err error) {
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
func (obj *_SysUserMgr) GetByOptions(opts ...Option) (results []*SysUser, err error) {
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
func (obj *_SysUserMgr) GetFromID(id int64) (result SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysUserMgr) GetBatchFromID(ids []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 账号
func (obj *_SysUserMgr) GetFromAccount(account string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量唯一主键查找 账号
func (obj *_SysUserMgr) GetBatchFromAccount(accounts []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_SysUserMgr) GetFromPassword(password string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找 密码
func (obj *_SysUserMgr) GetBatchFromPassword(passwords []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromNickName 通过nick_name获取内容 昵称
func (obj *_SysUserMgr) GetFromNickName(nickName string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nick_name = ?", nickName).Find(&results).Error

	return
}

// GetBatchFromNickName 批量唯一主键查找 昵称
func (obj *_SysUserMgr) GetBatchFromNickName(nickNames []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nick_name IN (?)", nickNames).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_SysUserMgr) GetFromName(name string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 姓名
func (obj *_SysUserMgr) GetBatchFromName(names []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_SysUserMgr) GetFromAvatar(avatar int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像
func (obj *_SysUserMgr) GetBatchFromAvatar(avatars []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容 生日
func (obj *_SysUserMgr) GetFromBirthday(birthday datatypes.Date) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量唯一主键查找 生日
func (obj *_SysUserMgr) GetBatchFromBirthday(birthdays []datatypes.Date) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别(字典 1男 2女 3未知)
func (obj *_SysUserMgr) GetFromSex(sex int8) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别(字典 1男 2女 3未知)
func (obj *_SysUserMgr) GetBatchFromSex(sexs []int8) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_SysUserMgr) GetFromEmail(email string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱
func (obj *_SysUserMgr) GetBatchFromEmail(emails []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_SysUserMgr) GetFromPhone(phone string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_SysUserMgr) GetBatchFromPhone(phones []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容 电话
func (obj *_SysUserMgr) GetFromTel(tel string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tel = ?", tel).Find(&results).Error

	return
}

// GetBatchFromTel 批量唯一主键查找 电话
func (obj *_SysUserMgr) GetBatchFromTel(tels []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tel IN (?)", tels).Find(&results).Error

	return
}

// GetFromLastLoginIP 通过last_login_ip获取内容 最后登陆IP
func (obj *_SysUserMgr) GetFromLastLoginIP(lastLoginIP string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_login_ip = ?", lastLoginIP).Find(&results).Error

	return
}

// GetBatchFromLastLoginIP 批量唯一主键查找 最后登陆IP
func (obj *_SysUserMgr) GetBatchFromLastLoginIP(lastLoginIPs []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_login_ip IN (?)", lastLoginIPs).Find(&results).Error

	return
}

// GetFromLastLoginTime 通过last_login_time获取内容 最后登陆时间
func (obj *_SysUserMgr) GetFromLastLoginTime(lastLoginTime time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_login_time = ?", lastLoginTime).Find(&results).Error

	return
}

// GetBatchFromLastLoginTime 批量唯一主键查找 最后登陆时间
func (obj *_SysUserMgr) GetBatchFromLastLoginTime(lastLoginTimes []time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_login_time IN (?)", lastLoginTimes).Find(&results).Error

	return
}

// GetFromAdminType 通过admin_type获取内容 管理员类型（0超级管理员 1非管理员）
func (obj *_SysUserMgr) GetFromAdminType(adminType int8) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("admin_type = ?", adminType).Find(&results).Error

	return
}

// GetBatchFromAdminType 批量唯一主键查找 管理员类型（0超级管理员 1非管理员）
func (obj *_SysUserMgr) GetBatchFromAdminType(adminTypes []int8) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("admin_type IN (?)", adminTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0正常 1冻结 2删除）
func (obj *_SysUserMgr) GetFromStatus(status int8) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0正常 1冻结 2删除）
func (obj *_SysUserMgr) GetBatchFromStatus(statuss []int8) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysUserMgr) GetFromCreateTime(createTime time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysUserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysUserMgr) GetFromCreateUser(createUser int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysUserMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysUserMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysUserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysUserMgr) GetFromUpdateUser(updateUser int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_SysUserMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysUserMgr) FetchByPrimaryKey(id int64) (result SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
