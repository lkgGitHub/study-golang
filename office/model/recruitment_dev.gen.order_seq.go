package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderSeqMgr struct {
	*_BaseMgr
}

// OrderSeqMgr open func
func OrderSeqMgr(db *gorm.DB) *_OrderSeqMgr {
	if db == nil {
		panic(fmt.Errorf("OrderSeqMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderSeqMgr{_BaseMgr: &_BaseMgr{DB: db.Table("order_seq"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderSeqMgr) GetTableName() string {
	return "order_seq"
}

// Get 获取
func (obj *_OrderSeqMgr) Get() (result OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderSeqMgr) Gets() (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_OrderSeqMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderDate order_date获取 日期
func (obj *_OrderSeqMgr) WithOrderDate(orderDate string) Option {
	return optionFunc(func(o *options) { o.query["order_date"] = orderDate })
}

// WithCurrent current获取 当前值
func (obj *_OrderSeqMgr) WithCurrent(current int) Option {
	return optionFunc(func(o *options) { o.query["current"] = current })
}

// WithIncrement increment获取 自增值
func (obj *_OrderSeqMgr) WithIncrement(increment int) Option {
	return optionFunc(func(o *options) { o.query["increment"] = increment })
}

// WithCreateTime create_time获取 创建时间
func (obj *_OrderSeqMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_OrderSeqMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_OrderSeqMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_OrderSeqMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_OrderSeqMgr) GetByOption(opts ...Option) (result OrderSeq, err error) {
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
func (obj *_OrderSeqMgr) GetByOptions(opts ...Option) (results []*OrderSeq, err error) {
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
func (obj *_OrderSeqMgr) GetFromID(id int64) (result OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键ID
func (obj *_OrderSeqMgr) GetBatchFromID(ids []int64) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderDate 通过order_date获取内容 日期
func (obj *_OrderSeqMgr) GetFromOrderDate(orderDate string) (result OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_date = ?", orderDate).Find(&result).Error

	return
}

// GetBatchFromOrderDate 批量唯一主键查找 日期
func (obj *_OrderSeqMgr) GetBatchFromOrderDate(orderDates []string) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_date IN (?)", orderDates).Find(&results).Error

	return
}

// GetFromCurrent 通过current获取内容 当前值
func (obj *_OrderSeqMgr) GetFromCurrent(current int) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current = ?", current).Find(&results).Error

	return
}

// GetBatchFromCurrent 批量唯一主键查找 当前值
func (obj *_OrderSeqMgr) GetBatchFromCurrent(currents []int) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current IN (?)", currents).Find(&results).Error

	return
}

// GetFromIncrement 通过increment获取内容 自增值
func (obj *_OrderSeqMgr) GetFromIncrement(increment int) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("increment = ?", increment).Find(&results).Error

	return
}

// GetBatchFromIncrement 批量唯一主键查找 自增值
func (obj *_OrderSeqMgr) GetBatchFromIncrement(increments []int) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("increment IN (?)", increments).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_OrderSeqMgr) GetFromCreateTime(createTime time.Time) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_OrderSeqMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_OrderSeqMgr) GetFromUpdateTime(updateTime time.Time) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_OrderSeqMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_OrderSeqMgr) GetFromCreateUser(createUser int64) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_OrderSeqMgr) GetBatchFromCreateUser(createUsers []int64) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_OrderSeqMgr) GetFromUpdateUser(updateUser int64) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_OrderSeqMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderSeqMgr) FetchByPrimaryKey(id int64) (result OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOrderSeqDateIndex primay or index 获取唯一内容
func (obj *_OrderSeqMgr) FetchUniqueByOrderSeqDateIndex(orderDate string) (result OrderSeq, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_date = ?", orderDate).Find(&result).Error

	return
}
