package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _FriendMgr struct {
	*_BaseMgr
}

// FriendMgr open func
func FriendMgr(db *gorm.DB) *_FriendMgr {
	if db == nil {
		panic(fmt.Errorf("FriendMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FriendMgr{_BaseMgr: &_BaseMgr{DB: db.Table("friend"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FriendMgr) GetTableName() string {
	return "friend"
}

// Get 获取
func (obj *_FriendMgr) Get() (result Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FriendMgr) Gets() (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FriendMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_FriendMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithFriendID friend_id获取
func (obj *_FriendMgr) WithFriendID(friendID int64) Option {
	return optionFunc(func(o *options) { o.query["friend_id"] = friendID })
}

// WithCreateTime create_time获取
func (obj *_FriendMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取
func (obj *_FriendMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// GetByOption 功能选项模式获取
func (obj *_FriendMgr) GetByOption(opts ...Option) (result Friend, err error) {
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
func (obj *_FriendMgr) GetByOptions(opts ...Option) (results []*Friend, err error) {
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
func (obj *_FriendMgr) GetFromID(id uint64) (result Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_FriendMgr) GetBatchFromID(ids []uint64) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_FriendMgr) GetFromUserID(userID int64) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_FriendMgr) GetBatchFromUserID(userIDs []int64) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromFriendID 通过friend_id获取内容
func (obj *_FriendMgr) GetFromFriendID(friendID int64) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("friend_id = ?", friendID).Find(&results).Error

	return
}

// GetBatchFromFriendID 批量唯一主键查找
func (obj *_FriendMgr) GetBatchFromFriendID(friendIDs []int64) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("friend_id IN (?)", friendIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_FriendMgr) GetFromCreateTime(createTime time.Time) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_FriendMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容
func (obj *_FriendMgr) GetFromDeleteTime(deleteTime time.Time) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找
func (obj *_FriendMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_FriendMgr) FetchByPrimaryKey(id uint64) (result Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUserID primay or index 获取唯一内容
func (obj *_FriendMgr) FetchUniqueIndexByUserID(userID int64, friendID int64) (result Friend, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ? AND friend_id = ?", userID, friendID).Find(&result).Error

	return
}
