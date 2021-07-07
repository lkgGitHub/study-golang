package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SysEmpMgr struct {
	*_BaseMgr
}

// SysEmpMgr open func
func SysEmpMgr(db *gorm.DB) *_SysEmpMgr {
	if db == nil {
		panic(fmt.Errorf("SysEmpMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysEmpMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_emp"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysEmpMgr) GetTableName() string {
	return "sys_emp"
}

// Get 获取
func (obj *_SysEmpMgr) Get() (result SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysEmpMgr) Gets() (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysEmpMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithJobNum job_num获取 工号
func (obj *_SysEmpMgr) WithJobNum(jobNum string) Option {
	return optionFunc(func(o *options) { o.query["job_num"] = jobNum })
}

// WithOrgID org_id获取 所属机构id
func (obj *_SysEmpMgr) WithOrgID(orgID int64) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// WithOrgName org_name获取 所属机构名称
func (obj *_SysEmpMgr) WithOrgName(orgName string) Option {
	return optionFunc(func(o *options) { o.query["org_name"] = orgName })
}

// GetByOption 功能选项模式获取
func (obj *_SysEmpMgr) GetByOption(opts ...Option) (result SysEmp, err error) {
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
func (obj *_SysEmpMgr) GetByOptions(opts ...Option) (results []*SysEmp, err error) {
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
func (obj *_SysEmpMgr) GetFromID(id int64) (result SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysEmpMgr) GetBatchFromID(ids []int64) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromJobNum 通过job_num获取内容 工号
func (obj *_SysEmpMgr) GetFromJobNum(jobNum string) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_num = ?", jobNum).Find(&results).Error

	return
}

// GetBatchFromJobNum 批量唯一主键查找 工号
func (obj *_SysEmpMgr) GetBatchFromJobNum(jobNums []string) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_num IN (?)", jobNums).Find(&results).Error

	return
}

// GetFromOrgID 通过org_id获取内容 所属机构id
func (obj *_SysEmpMgr) GetFromOrgID(orgID int64) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id = ?", orgID).Find(&results).Error

	return
}

// GetBatchFromOrgID 批量唯一主键查找 所属机构id
func (obj *_SysEmpMgr) GetBatchFromOrgID(orgIDs []int64) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_id IN (?)", orgIDs).Find(&results).Error

	return
}

// GetFromOrgName 通过org_name获取内容 所属机构名称
func (obj *_SysEmpMgr) GetFromOrgName(orgName string) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_name = ?", orgName).Find(&results).Error

	return
}

// GetBatchFromOrgName 批量唯一主键查找 所属机构名称
func (obj *_SysEmpMgr) GetBatchFromOrgName(orgNames []string) (results []*SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("org_name IN (?)", orgNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysEmpMgr) FetchByPrimaryKey(id int64) (result SysEmp, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
