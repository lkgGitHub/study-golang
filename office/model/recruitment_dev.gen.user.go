package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_UserMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 用户名
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithUserID user_id获取 用户ID
func (obj *_UserMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithWxOpenID wx_open_id获取 微信openID
func (obj *_UserMgr) WithWxOpenID(wxOpenID string) Option {
	return optionFunc(func(o *options) { o.query["wx_open_id"] = wxOpenID })
}

// WithProgress progress获取 完成简历进度
func (obj *_UserMgr) WithProgress(progress int8) Option {
	return optionFunc(func(o *options) { o.query["progress"] = progress })
}

// WithJobLevel job_level获取 岗位级别
func (obj *_UserMgr) WithJobLevel(jobLevel string) Option {
	return optionFunc(func(o *options) { o.query["job_level"] = jobLevel })
}

// WithTotalAccount total_account获取 钱包余额
func (obj *_UserMgr) WithTotalAccount(totalAccount int) Option {
	return optionFunc(func(o *options) { o.query["total_account"] = totalAccount })
}

// WithCoin coin获取 面币余额
func (obj *_UserMgr) WithCoin(coin int) Option {
	return optionFunc(func(o *options) { o.query["coin"] = coin })
}

// WithRecLevel rec_level获取 推荐事级别
func (obj *_UserMgr) WithRecLevel(recLevel int8) Option {
	return optionFunc(func(o *options) { o.query["rec_level"] = recLevel })
}

// WithCreateTime create_time获取
func (obj *_UserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取
func (obj *_UserMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取
func (obj *_UserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取
func (obj *_UserMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
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
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
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
func (obj *_UserMgr) GetFromID(id int64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_UserMgr) GetBatchFromID(ids []int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 用户名
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 用户名
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_UserMgr) GetFromUserID(userID int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *_UserMgr) GetBatchFromUserID(userIDs []int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromWxOpenID 通过wx_open_id获取内容 微信openID
func (obj *_UserMgr) GetFromWxOpenID(wxOpenID string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wx_open_id = ?", wxOpenID).Find(&results).Error

	return
}

// GetBatchFromWxOpenID 批量唯一主键查找 微信openID
func (obj *_UserMgr) GetBatchFromWxOpenID(wxOpenIDs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wx_open_id IN (?)", wxOpenIDs).Find(&results).Error

	return
}

// GetFromProgress 通过progress获取内容 完成简历进度
func (obj *_UserMgr) GetFromProgress(progress int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("progress = ?", progress).Find(&results).Error

	return
}

// GetBatchFromProgress 批量唯一主键查找 完成简历进度
func (obj *_UserMgr) GetBatchFromProgress(progresss []int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("progress IN (?)", progresss).Find(&results).Error

	return
}

// GetFromJobLevel 通过job_level获取内容 岗位级别
func (obj *_UserMgr) GetFromJobLevel(jobLevel string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_level = ?", jobLevel).Find(&results).Error

	return
}

// GetBatchFromJobLevel 批量唯一主键查找 岗位级别
func (obj *_UserMgr) GetBatchFromJobLevel(jobLevels []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_level IN (?)", jobLevels).Find(&results).Error

	return
}

// GetFromTotalAccount 通过total_account获取内容 钱包余额
func (obj *_UserMgr) GetFromTotalAccount(totalAccount int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_account = ?", totalAccount).Find(&results).Error

	return
}

// GetBatchFromTotalAccount 批量唯一主键查找 钱包余额
func (obj *_UserMgr) GetBatchFromTotalAccount(totalAccounts []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_account IN (?)", totalAccounts).Find(&results).Error

	return
}

// GetFromCoin 通过coin获取内容 面币余额
func (obj *_UserMgr) GetFromCoin(coin int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coin = ?", coin).Find(&results).Error

	return
}

// GetBatchFromCoin 批量唯一主键查找 面币余额
func (obj *_UserMgr) GetBatchFromCoin(coins []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coin IN (?)", coins).Find(&results).Error

	return
}

// GetFromRecLevel 通过rec_level获取内容 推荐事级别
func (obj *_UserMgr) GetFromRecLevel(recLevel int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rec_level = ?", recLevel).Find(&results).Error

	return
}

// GetBatchFromRecLevel 批量唯一主键查找 推荐事级别
func (obj *_UserMgr) GetBatchFromRecLevel(recLevels []int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rec_level IN (?)", recLevels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_UserMgr) GetFromCreateTime(createTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容
func (obj *_UserMgr) GetFromCreateUser(createUser int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromCreateUser(createUsers []int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_UserMgr) GetFromUpdateTime(updateTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容
func (obj *_UserMgr) GetFromUpdateUser(updateUser int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id int64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
