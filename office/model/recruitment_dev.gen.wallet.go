package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WalletMgr struct {
	*_BaseMgr
}

// WalletMgr open func
func WalletMgr(db *gorm.DB) *_WalletMgr {
	if db == nil {
		panic(fmt.Errorf("WalletMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WalletMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wallet"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WalletMgr) GetTableName() string {
	return "wallet"
}

// Get 获取
func (obj *_WalletMgr) Get() (result Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WalletMgr) Gets() (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WalletMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_WalletMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithWallet wallet获取
func (obj *_WalletMgr) WithWallet(wallet float64) Option {
	return optionFunc(func(o *options) { o.query["wallet"] = wallet })
}

// WithCreateTime create_time获取
func (obj *_WalletMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取
func (obj *_WalletMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// GetByOption 功能选项模式获取
func (obj *_WalletMgr) GetByOption(opts ...Option) (result Wallet, err error) {
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
func (obj *_WalletMgr) GetByOptions(opts ...Option) (results []*Wallet, err error) {
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
func (obj *_WalletMgr) GetFromID(id int) (result Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WalletMgr) GetBatchFromID(ids []int) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_WalletMgr) GetFromUserID(userID int) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_WalletMgr) GetBatchFromUserID(userIDs []int) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromWallet 通过wallet获取内容
func (obj *_WalletMgr) GetFromWallet(wallet float64) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wallet = ?", wallet).Find(&results).Error

	return
}

// GetBatchFromWallet 批量唯一主键查找
func (obj *_WalletMgr) GetBatchFromWallet(wallets []float64) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wallet IN (?)", wallets).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_WalletMgr) GetFromCreateTime(createTime time.Time) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_WalletMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容
func (obj *_WalletMgr) GetFromDeleteTime(deleteTime time.Time) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量唯一主键查找
func (obj *_WalletMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_time IN (?)", deleteTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WalletMgr) FetchByPrimaryKey(id int) (result Wallet, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
