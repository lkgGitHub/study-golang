package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderMgr struct {
	*_BaseMgr
}

// OrderMgr open func
func OrderMgr(db *gorm.DB) *_OrderMgr {
	if db == nil {
		panic(fmt.Errorf("OrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderMgr) GetTableName() string {
	return "order"
}

// Get 获取
func (obj *_OrderMgr) Get() (result Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderMgr) Gets() (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_OrderMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCompanyID company_id获取 企业ID
func (obj *_OrderMgr) WithCompanyID(companyID int64) Option {
	return optionFunc(func(o *options) { o.query["company_id"] = companyID })
}

// WithJobID job_id获取 职位ID
func (obj *_OrderMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithUserID user_id获取 用户ID
func (obj *_OrderMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithOrderNo order_no获取 订单编号
func (obj *_OrderMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithOrderStatus order_status获取 订单状态 1-初始化 2-待结算 3-结算中 4-已完成 5-已关闭
func (obj *_OrderMgr) WithOrderStatus(orderStatus int) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithBaseCommission base_commission获取 佣金计算基数
func (obj *_OrderMgr) WithBaseCommission(baseCommission float64) Option {
	return optionFunc(func(o *options) { o.query["base_commission"] = baseCommission })
}

// WithOrderAmount order_amount获取 订单金额
func (obj *_OrderMgr) WithOrderAmount(orderAmount float64) Option {
	return optionFunc(func(o *options) { o.query["order_amount"] = orderAmount })
}

// WithSettledAmount settled_amount获取 已结算金额
func (obj *_OrderMgr) WithSettledAmount(settledAmount float64) Option {
	return optionFunc(func(o *options) { o.query["settled_amount"] = settledAmount })
}

// WithUnsettledAmount unsettled_amount获取 未结算金额
func (obj *_OrderMgr) WithUnsettledAmount(unsettledAmount float64) Option {
	return optionFunc(func(o *options) { o.query["unsettled_amount"] = unsettledAmount })
}

// WithCompleteTime complete_time获取 订单完成时间
func (obj *_OrderMgr) WithCompleteTime(completeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["complete_time"] = completeTime })
}

// WithCreateTime create_time获取 创建时间
func (obj *_OrderMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_OrderMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取
func (obj *_OrderMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取
func (obj *_OrderMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_OrderMgr) GetByOption(opts ...Option) (result Order, err error) {
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
func (obj *_OrderMgr) GetByOptions(opts ...Option) (results []*Order, err error) {
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
func (obj *_OrderMgr) GetFromID(id int64) (result Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键ID
func (obj *_OrderMgr) GetBatchFromID(ids []int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCompanyID 通过company_id获取内容 企业ID
func (obj *_OrderMgr) GetFromCompanyID(companyID int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_id = ?", companyID).Find(&results).Error

	return
}

// GetBatchFromCompanyID 批量唯一主键查找 企业ID
func (obj *_OrderMgr) GetBatchFromCompanyID(companyIDs []int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_id IN (?)", companyIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容 职位ID
func (obj *_OrderMgr) GetFromJobID(jobID int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量唯一主键查找 职位ID
func (obj *_OrderMgr) GetBatchFromJobID(jobIDs []int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_OrderMgr) GetFromUserID(userID int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *_OrderMgr) GetBatchFromUserID(userIDs []int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_OrderMgr) GetFromOrderNo(orderNo string) (result Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_no = ?", orderNo).Find(&result).Error

	return
}

// GetBatchFromOrderNo 批量唯一主键查找 订单编号
func (obj *_OrderMgr) GetBatchFromOrderNo(orderNos []string) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_no IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态 1-初始化 2-待结算 3-结算中 4-已完成 5-已关闭
func (obj *_OrderMgr) GetFromOrderStatus(orderStatus int) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_status = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量唯一主键查找 订单状态 1-初始化 2-待结算 3-结算中 4-已完成 5-已关闭
func (obj *_OrderMgr) GetBatchFromOrderStatus(orderStatuss []int) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_status IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromBaseCommission 通过base_commission获取内容 佣金计算基数
func (obj *_OrderMgr) GetFromBaseCommission(baseCommission float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("base_commission = ?", baseCommission).Find(&results).Error

	return
}

// GetBatchFromBaseCommission 批量唯一主键查找 佣金计算基数
func (obj *_OrderMgr) GetBatchFromBaseCommission(baseCommissions []float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("base_commission IN (?)", baseCommissions).Find(&results).Error

	return
}

// GetFromOrderAmount 通过order_amount获取内容 订单金额
func (obj *_OrderMgr) GetFromOrderAmount(orderAmount float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_amount = ?", orderAmount).Find(&results).Error

	return
}

// GetBatchFromOrderAmount 批量唯一主键查找 订单金额
func (obj *_OrderMgr) GetBatchFromOrderAmount(orderAmounts []float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_amount IN (?)", orderAmounts).Find(&results).Error

	return
}

// GetFromSettledAmount 通过settled_amount获取内容 已结算金额
func (obj *_OrderMgr) GetFromSettledAmount(settledAmount float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("settled_amount = ?", settledAmount).Find(&results).Error

	return
}

// GetBatchFromSettledAmount 批量唯一主键查找 已结算金额
func (obj *_OrderMgr) GetBatchFromSettledAmount(settledAmounts []float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("settled_amount IN (?)", settledAmounts).Find(&results).Error

	return
}

// GetFromUnsettledAmount 通过unsettled_amount获取内容 未结算金额
func (obj *_OrderMgr) GetFromUnsettledAmount(unsettledAmount float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unsettled_amount = ?", unsettledAmount).Find(&results).Error

	return
}

// GetBatchFromUnsettledAmount 批量唯一主键查找 未结算金额
func (obj *_OrderMgr) GetBatchFromUnsettledAmount(unsettledAmounts []float64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unsettled_amount IN (?)", unsettledAmounts).Find(&results).Error

	return
}

// GetFromCompleteTime 通过complete_time获取内容 订单完成时间
func (obj *_OrderMgr) GetFromCompleteTime(completeTime time.Time) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("complete_time = ?", completeTime).Find(&results).Error

	return
}

// GetBatchFromCompleteTime 批量唯一主键查找 订单完成时间
func (obj *_OrderMgr) GetBatchFromCompleteTime(completeTimes []time.Time) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("complete_time IN (?)", completeTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_OrderMgr) GetFromCreateTime(createTime time.Time) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_OrderMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_OrderMgr) GetFromUpdateTime(updateTime time.Time) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_OrderMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容
func (obj *_OrderMgr) GetFromCreateUser(createUser int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找
func (obj *_OrderMgr) GetBatchFromCreateUser(createUsers []int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容
func (obj *_OrderMgr) GetFromUpdateUser(updateUser int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找
func (obj *_OrderMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderMgr) FetchByPrimaryKey(id int64) (result Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOrderNoIndex primay or index 获取唯一内容
func (obj *_OrderMgr) FetchUniqueByOrderNoIndex(orderNo string) (result Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_no = ?", orderNo).Find(&result).Error

	return
}
