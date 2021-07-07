package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CoinRecordMgr struct {
	*_BaseMgr
}

// CoinRecordMgr open func
func CoinRecordMgr(db *gorm.DB) *_CoinRecordMgr {
	if db == nil {
		panic(fmt.Errorf("CoinRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CoinRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("coin_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CoinRecordMgr) GetTableName() string {
	return "coin_record"
}

// Get 获取
func (obj *_CoinRecordMgr) Get() (result CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CoinRecordMgr) Gets() (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_CoinRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 同户ID
func (obj *_CoinRecordMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithType type获取 记录类型 1-收入 2-支出
func (obj *_CoinRecordMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCoinNum coin_num获取 币值
func (obj *_CoinRecordMgr) WithCoinNum(coinNum int) Option {
	return optionFunc(func(o *options) { o.query["coin_num"] = coinNum })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CoinRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDesc desc获取 收支描述
func (obj *_CoinRecordMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// GetByOption 功能选项模式获取
func (obj *_CoinRecordMgr) GetByOption(opts ...Option) (result CoinRecord, err error) {
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
func (obj *_CoinRecordMgr) GetByOptions(opts ...Option) (results []*CoinRecord, err error) {
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

// GetFromID 通过id获取内容 主键ID
func (obj *_CoinRecordMgr) GetFromID(id int64) (result CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键ID
func (obj *_CoinRecordMgr) GetBatchFromID(ids []int64) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 同户ID
func (obj *_CoinRecordMgr) GetFromUserID(userID int64) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 同户ID
func (obj *_CoinRecordMgr) GetBatchFromUserID(userIDs []int64) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 记录类型 1-收入 2-支出
func (obj *_CoinRecordMgr) GetFromType(_type int) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 记录类型 1-收入 2-支出
func (obj *_CoinRecordMgr) GetBatchFromType(_types []int) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromCoinNum 通过coin_num获取内容 币值
func (obj *_CoinRecordMgr) GetFromCoinNum(coinNum int) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coin_num = ?", coinNum).Find(&results).Error

	return
}

// GetBatchFromCoinNum 批量唯一主键查找 币值
func (obj *_CoinRecordMgr) GetBatchFromCoinNum(coinNums []int) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coin_num IN (?)", coinNums).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CoinRecordMgr) GetFromCreateTime(createTime time.Time) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_CoinRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 收支描述
func (obj *_CoinRecordMgr) GetFromDesc(desc string) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 收支描述
func (obj *_CoinRecordMgr) GetBatchFromDesc(descs []string) (results []*CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CoinRecordMgr) FetchByPrimaryKey(id int64) (result CoinRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
