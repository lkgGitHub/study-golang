package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _JobMgr struct {
	*_BaseMgr
}

// JobMgr open func
func JobMgr(db *gorm.DB) *_JobMgr {
	if db == nil {
		panic(fmt.Errorf("JobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_JobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("job"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_JobMgr) GetTableName() string {
	return "job"
}

// Get 获取
func (obj *_JobMgr) Get() (result Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_JobMgr) Gets() (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_JobMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 职位名称
func (obj *_JobMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithStatus status获取 职位状态
func (obj *_JobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRecruitingNumbers recruiting_numbers获取 招聘人数
func (obj *_JobMgr) WithRecruitingNumbers(recruitingNumbers int) Option {
	return optionFunc(func(o *options) { o.query["recruiting_numbers"] = recruitingNumbers })
}

// WithSuccessfulNumbers successful_numbers获取 通过人数
func (obj *_JobMgr) WithSuccessfulNumbers(successfulNumbers int) Option {
	return optionFunc(func(o *options) { o.query["successful_numbers"] = successfulNumbers })
}

// WithJobDepartment job_department获取 所在部门
func (obj *_JobMgr) WithJobDepartment(jobDepartment string) Option {
	return optionFunc(func(o *options) { o.query["job_department"] = jobDepartment })
}

// WithUnderlingNumber underling_number获取 下属人数
func (obj *_JobMgr) WithUnderlingNumber(underlingNumber int) Option {
	return optionFunc(func(o *options) { o.query["underling_number"] = underlingNumber })
}

// WithSalaryStart salary_start获取 薪水起始值
func (obj *_JobMgr) WithSalaryStart(salaryStart int) Option {
	return optionFunc(func(o *options) { o.query["salary_start"] = salaryStart })
}

// WithSalaryEnd salary_end获取 薪水结束值
func (obj *_JobMgr) WithSalaryEnd(salaryEnd int) Option {
	return optionFunc(func(o *options) { o.query["salary_end"] = salaryEnd })
}

// WithCity city获取 公司城市
func (obj *_JobMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithAddress address获取 公司地址（具体到区）
func (obj *_JobMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithDetailedLocation detailed_location获取 详细地点
func (obj *_JobMgr) WithDetailedLocation(detailedLocation string) Option {
	return optionFunc(func(o *options) { o.query["detailed_location"] = detailedLocation })
}

// WithInterviewAddr interview_addr获取 面试地址
func (obj *_JobMgr) WithInterviewAddr(interviewAddr string) Option {
	return optionFunc(func(o *options) { o.query["interview_addr"] = interviewAddr })
}

// WithEducation education获取 学历
func (obj *_JobMgr) WithEducation(education string) Option {
	return optionFunc(func(o *options) { o.query["education"] = education })
}

// WithWorkingYearsType working_years_type获取 工作年限类型
func (obj *_JobMgr) WithWorkingYearsType(workingYearsType string) Option {
	return optionFunc(func(o *options) { o.query["working_years_type"] = workingYearsType })
}

// WithWorkingYearsStart working_years_start获取 工作年限起始值
func (obj *_JobMgr) WithWorkingYearsStart(workingYearsStart int) Option {
	return optionFunc(func(o *options) { o.query["working_years_start"] = workingYearsStart })
}

// WithWorkingYearsEnd working_years_end获取 工作年限结束值
func (obj *_JobMgr) WithWorkingYearsEnd(workingYearsEnd int) Option {
	return optionFunc(func(o *options) { o.query["working_years_end"] = workingYearsEnd })
}

// WithIndustry industry获取 行业
func (obj *_JobMgr) WithIndustry(industry string) Option {
	return optionFunc(func(o *options) { o.query["industry"] = industry })
}

// WithJobBackdrop job_backdrop获取 背景经验
func (obj *_JobMgr) WithJobBackdrop(jobBackdrop string) Option {
	return optionFunc(func(o *options) { o.query["job_backdrop"] = jobBackdrop })
}

// WithJobHopFre job_hop_fre获取 跳槽频率
func (obj *_JobMgr) WithJobHopFre(jobHopFre float64) Option {
	return optionFunc(func(o *options) { o.query["job_hop_fre"] = jobHopFre })
}

// WithAgeRequirement age_requirement获取 年龄要求
func (obj *_JobMgr) WithAgeRequirement(ageRequirement string) Option {
	return optionFunc(func(o *options) { o.query["age_requirement"] = ageRequirement })
}

// WithInterviewRounds interview_rounds获取 面试轮数
func (obj *_JobMgr) WithInterviewRounds(interviewRounds int) Option {
	return optionFunc(func(o *options) { o.query["interview_rounds"] = interviewRounds })
}

// WithJobCommissionCashStart job_commission_cash_start获取 佣金兑现标准起始值
func (obj *_JobMgr) WithJobCommissionCashStart(jobCommissionCashStart float64) Option {
	return optionFunc(func(o *options) { o.query["job_commission_cash_start"] = jobCommissionCashStart })
}

// WithJobCommissionCashEnd job_commission_cash_end获取 佣金兑现标准结束值
func (obj *_JobMgr) WithJobCommissionCashEnd(jobCommissionCashEnd float64) Option {
	return optionFunc(func(o *options) { o.query["job_commission_cash_end"] = jobCommissionCashEnd })
}

// WithJobCommissionCashPeriod job_commission_cash_period获取 佣金兑现期限
func (obj *_JobMgr) WithJobCommissionCashPeriod(jobCommissionCashPeriod string) Option {
	return optionFunc(func(o *options) { o.query["job_commission_cash_period"] = jobCommissionCashPeriod })
}

// WithJobKeyword job_keyword获取 岗位关键字
func (obj *_JobMgr) WithJobKeyword(jobKeyword string) Option {
	return optionFunc(func(o *options) { o.query["job_keyword"] = jobKeyword })
}

// WithCompanyID company_id获取 关联企业主键ID
func (obj *_JobMgr) WithCompanyID(companyID int64) Option {
	return optionFunc(func(o *options) { o.query["company_id"] = companyID })
}

// WithIsUrgent is_urgent获取 是否为紧急招聘, 默认为否
func (obj *_JobMgr) WithIsUrgent(isUrgent []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_urgent"] = isUrgent })
}

// WithJobDetailsID job_details_id获取 工作详情
func (obj *_JobMgr) WithJobDetailsID(jobDetailsID int64) Option {
	return optionFunc(func(o *options) { o.query["job_details_id"] = jobDetailsID })
}

// WithDeleted deleted获取 逻辑删除字段
func (obj *_JobMgr) WithDeleted(deleted []uint8) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithCreateTime create_time获取 创建时间
func (obj *_JobMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取
func (obj *_JobMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_JobMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取
func (obj *_JobMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithJobSalaryStructure job_salary_structure获取 薪资结构
func (obj *_JobMgr) WithJobSalaryStructure(jobSalaryStructure int) Option {
	return optionFunc(func(o *options) { o.query["job_salary_structure"] = jobSalaryStructure })
}

// GetByOption 功能选项模式获取
func (obj *_JobMgr) GetByOption(opts ...Option) (result Job, err error) {
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
func (obj *_JobMgr) GetByOptions(opts ...Option) (results []*Job, err error) {
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
func (obj *_JobMgr) GetFromID(id int64) (result Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键id
func (obj *_JobMgr) GetBatchFromID(ids []int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 职位名称
func (obj *_JobMgr) GetFromName(name string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 职位名称
func (obj *_JobMgr) GetBatchFromName(names []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 职位状态
func (obj *_JobMgr) GetFromStatus(status string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 职位状态
func (obj *_JobMgr) GetBatchFromStatus(statuss []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRecruitingNumbers 通过recruiting_numbers获取内容 招聘人数
func (obj *_JobMgr) GetFromRecruitingNumbers(recruitingNumbers int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recruiting_numbers = ?", recruitingNumbers).Find(&results).Error

	return
}

// GetBatchFromRecruitingNumbers 批量唯一主键查找 招聘人数
func (obj *_JobMgr) GetBatchFromRecruitingNumbers(recruitingNumberss []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recruiting_numbers IN (?)", recruitingNumberss).Find(&results).Error

	return
}

// GetFromSuccessfulNumbers 通过successful_numbers获取内容 通过人数
func (obj *_JobMgr) GetFromSuccessfulNumbers(successfulNumbers int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("successful_numbers = ?", successfulNumbers).Find(&results).Error

	return
}

// GetBatchFromSuccessfulNumbers 批量唯一主键查找 通过人数
func (obj *_JobMgr) GetBatchFromSuccessfulNumbers(successfulNumberss []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("successful_numbers IN (?)", successfulNumberss).Find(&results).Error

	return
}

// GetFromJobDepartment 通过job_department获取内容 所在部门
func (obj *_JobMgr) GetFromJobDepartment(jobDepartment string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_department = ?", jobDepartment).Find(&results).Error

	return
}

// GetBatchFromJobDepartment 批量唯一主键查找 所在部门
func (obj *_JobMgr) GetBatchFromJobDepartment(jobDepartments []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_department IN (?)", jobDepartments).Find(&results).Error

	return
}

// GetFromUnderlingNumber 通过underling_number获取内容 下属人数
func (obj *_JobMgr) GetFromUnderlingNumber(underlingNumber int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("underling_number = ?", underlingNumber).Find(&results).Error

	return
}

// GetBatchFromUnderlingNumber 批量唯一主键查找 下属人数
func (obj *_JobMgr) GetBatchFromUnderlingNumber(underlingNumbers []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("underling_number IN (?)", underlingNumbers).Find(&results).Error

	return
}

// GetFromSalaryStart 通过salary_start获取内容 薪水起始值
func (obj *_JobMgr) GetFromSalaryStart(salaryStart int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("salary_start = ?", salaryStart).Find(&results).Error

	return
}

// GetBatchFromSalaryStart 批量唯一主键查找 薪水起始值
func (obj *_JobMgr) GetBatchFromSalaryStart(salaryStarts []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("salary_start IN (?)", salaryStarts).Find(&results).Error

	return
}

// GetFromSalaryEnd 通过salary_end获取内容 薪水结束值
func (obj *_JobMgr) GetFromSalaryEnd(salaryEnd int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("salary_end = ?", salaryEnd).Find(&results).Error

	return
}

// GetBatchFromSalaryEnd 批量唯一主键查找 薪水结束值
func (obj *_JobMgr) GetBatchFromSalaryEnd(salaryEnds []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("salary_end IN (?)", salaryEnds).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 公司城市
func (obj *_JobMgr) GetFromCity(city string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找 公司城市
func (obj *_JobMgr) GetBatchFromCity(citys []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 公司地址（具体到区）
func (obj *_JobMgr) GetFromAddress(address string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量唯一主键查找 公司地址（具体到区）
func (obj *_JobMgr) GetBatchFromAddress(addresss []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address IN (?)", addresss).Find(&results).Error

	return
}

// GetFromDetailedLocation 通过detailed_location获取内容 详细地点
func (obj *_JobMgr) GetFromDetailedLocation(detailedLocation string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("detailed_location = ?", detailedLocation).Find(&results).Error

	return
}

// GetBatchFromDetailedLocation 批量唯一主键查找 详细地点
func (obj *_JobMgr) GetBatchFromDetailedLocation(detailedLocations []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("detailed_location IN (?)", detailedLocations).Find(&results).Error

	return
}

// GetFromInterviewAddr 通过interview_addr获取内容 面试地址
func (obj *_JobMgr) GetFromInterviewAddr(interviewAddr string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interview_addr = ?", interviewAddr).Find(&results).Error

	return
}

// GetBatchFromInterviewAddr 批量唯一主键查找 面试地址
func (obj *_JobMgr) GetBatchFromInterviewAddr(interviewAddrs []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interview_addr IN (?)", interviewAddrs).Find(&results).Error

	return
}

// GetFromEducation 通过education获取内容 学历
func (obj *_JobMgr) GetFromEducation(education string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("education = ?", education).Find(&results).Error

	return
}

// GetBatchFromEducation 批量唯一主键查找 学历
func (obj *_JobMgr) GetBatchFromEducation(educations []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("education IN (?)", educations).Find(&results).Error

	return
}

// GetFromWorkingYearsType 通过working_years_type获取内容 工作年限类型
func (obj *_JobMgr) GetFromWorkingYearsType(workingYearsType string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("working_years_type = ?", workingYearsType).Find(&results).Error

	return
}

// GetBatchFromWorkingYearsType 批量唯一主键查找 工作年限类型
func (obj *_JobMgr) GetBatchFromWorkingYearsType(workingYearsTypes []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("working_years_type IN (?)", workingYearsTypes).Find(&results).Error

	return
}

// GetFromWorkingYearsStart 通过working_years_start获取内容 工作年限起始值
func (obj *_JobMgr) GetFromWorkingYearsStart(workingYearsStart int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("working_years_start = ?", workingYearsStart).Find(&results).Error

	return
}

// GetBatchFromWorkingYearsStart 批量唯一主键查找 工作年限起始值
func (obj *_JobMgr) GetBatchFromWorkingYearsStart(workingYearsStarts []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("working_years_start IN (?)", workingYearsStarts).Find(&results).Error

	return
}

// GetFromWorkingYearsEnd 通过working_years_end获取内容 工作年限结束值
func (obj *_JobMgr) GetFromWorkingYearsEnd(workingYearsEnd int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("working_years_end = ?", workingYearsEnd).Find(&results).Error

	return
}

// GetBatchFromWorkingYearsEnd 批量唯一主键查找 工作年限结束值
func (obj *_JobMgr) GetBatchFromWorkingYearsEnd(workingYearsEnds []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("working_years_end IN (?)", workingYearsEnds).Find(&results).Error

	return
}

// GetFromIndustry 通过industry获取内容 行业
func (obj *_JobMgr) GetFromIndustry(industry string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("industry = ?", industry).Find(&results).Error

	return
}

// GetBatchFromIndustry 批量唯一主键查找 行业
func (obj *_JobMgr) GetBatchFromIndustry(industrys []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("industry IN (?)", industrys).Find(&results).Error

	return
}

// GetFromJobBackdrop 通过job_backdrop获取内容 背景经验
func (obj *_JobMgr) GetFromJobBackdrop(jobBackdrop string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_backdrop = ?", jobBackdrop).Find(&results).Error

	return
}

// GetBatchFromJobBackdrop 批量唯一主键查找 背景经验
func (obj *_JobMgr) GetBatchFromJobBackdrop(jobBackdrops []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_backdrop IN (?)", jobBackdrops).Find(&results).Error

	return
}

// GetFromJobHopFre 通过job_hop_fre获取内容 跳槽频率
func (obj *_JobMgr) GetFromJobHopFre(jobHopFre float64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_hop_fre = ?", jobHopFre).Find(&results).Error

	return
}

// GetBatchFromJobHopFre 批量唯一主键查找 跳槽频率
func (obj *_JobMgr) GetBatchFromJobHopFre(jobHopFres []float64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_hop_fre IN (?)", jobHopFres).Find(&results).Error

	return
}

// GetFromAgeRequirement 通过age_requirement获取内容 年龄要求
func (obj *_JobMgr) GetFromAgeRequirement(ageRequirement string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("age_requirement = ?", ageRequirement).Find(&results).Error

	return
}

// GetBatchFromAgeRequirement 批量唯一主键查找 年龄要求
func (obj *_JobMgr) GetBatchFromAgeRequirement(ageRequirements []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("age_requirement IN (?)", ageRequirements).Find(&results).Error

	return
}

// GetFromInterviewRounds 通过interview_rounds获取内容 面试轮数
func (obj *_JobMgr) GetFromInterviewRounds(interviewRounds int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interview_rounds = ?", interviewRounds).Find(&results).Error

	return
}

// GetBatchFromInterviewRounds 批量唯一主键查找 面试轮数
func (obj *_JobMgr) GetBatchFromInterviewRounds(interviewRoundss []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interview_rounds IN (?)", interviewRoundss).Find(&results).Error

	return
}

// GetFromJobCommissionCashStart 通过job_commission_cash_start获取内容 佣金兑现标准起始值
func (obj *_JobMgr) GetFromJobCommissionCashStart(jobCommissionCashStart float64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_commission_cash_start = ?", jobCommissionCashStart).Find(&results).Error

	return
}

// GetBatchFromJobCommissionCashStart 批量唯一主键查找 佣金兑现标准起始值
func (obj *_JobMgr) GetBatchFromJobCommissionCashStart(jobCommissionCashStarts []float64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_commission_cash_start IN (?)", jobCommissionCashStarts).Find(&results).Error

	return
}

// GetFromJobCommissionCashEnd 通过job_commission_cash_end获取内容 佣金兑现标准结束值
func (obj *_JobMgr) GetFromJobCommissionCashEnd(jobCommissionCashEnd float64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_commission_cash_end = ?", jobCommissionCashEnd).Find(&results).Error

	return
}

// GetBatchFromJobCommissionCashEnd 批量唯一主键查找 佣金兑现标准结束值
func (obj *_JobMgr) GetBatchFromJobCommissionCashEnd(jobCommissionCashEnds []float64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_commission_cash_end IN (?)", jobCommissionCashEnds).Find(&results).Error

	return
}

// GetFromJobCommissionCashPeriod 通过job_commission_cash_period获取内容 佣金兑现期限
func (obj *_JobMgr) GetFromJobCommissionCashPeriod(jobCommissionCashPeriod string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_commission_cash_period = ?", jobCommissionCashPeriod).Find(&results).Error

	return
}

// GetBatchFromJobCommissionCashPeriod 批量唯一主键查找 佣金兑现期限
func (obj *_JobMgr) GetBatchFromJobCommissionCashPeriod(jobCommissionCashPeriods []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_commission_cash_period IN (?)", jobCommissionCashPeriods).Find(&results).Error

	return
}

// GetFromJobKeyword 通过job_keyword获取内容 岗位关键字
func (obj *_JobMgr) GetFromJobKeyword(jobKeyword string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_keyword = ?", jobKeyword).Find(&results).Error

	return
}

// GetBatchFromJobKeyword 批量唯一主键查找 岗位关键字
func (obj *_JobMgr) GetBatchFromJobKeyword(jobKeywords []string) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_keyword IN (?)", jobKeywords).Find(&results).Error

	return
}

// GetFromCompanyID 通过company_id获取内容 关联企业主键ID
func (obj *_JobMgr) GetFromCompanyID(companyID int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_id = ?", companyID).Find(&results).Error

	return
}

// GetBatchFromCompanyID 批量唯一主键查找 关联企业主键ID
func (obj *_JobMgr) GetBatchFromCompanyID(companyIDs []int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_id IN (?)", companyIDs).Find(&results).Error

	return
}

// GetFromIsUrgent 通过is_urgent获取内容 是否为紧急招聘, 默认为否
func (obj *_JobMgr) GetFromIsUrgent(isUrgent []uint8) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_urgent = ?", isUrgent).Find(&results).Error

	return
}

// GetBatchFromIsUrgent 批量唯一主键查找 是否为紧急招聘, 默认为否
func (obj *_JobMgr) GetBatchFromIsUrgent(isUrgents [][]uint8) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_urgent IN (?)", isUrgents).Find(&results).Error

	return
}

// GetFromJobDetailsID 通过job_details_id获取内容 工作详情
func (obj *_JobMgr) GetFromJobDetailsID(jobDetailsID int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_details_id = ?", jobDetailsID).Find(&results).Error

	return
}

// GetBatchFromJobDetailsID 批量唯一主键查找 工作详情
func (obj *_JobMgr) GetBatchFromJobDetailsID(jobDetailsIDs []int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_details_id IN (?)", jobDetailsIDs).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除字段
func (obj *_JobMgr) GetFromDeleted(deleted []uint8) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找 逻辑删除字段
func (obj *_JobMgr) GetBatchFromDeleted(deleteds [][]uint8) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_JobMgr) GetFromCreateTime(createTime time.Time) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_JobMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容
func (obj *_JobMgr) GetFromCreateUser(createUser int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找
func (obj *_JobMgr) GetBatchFromCreateUser(createUsers []int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_JobMgr) GetFromUpdateTime(updateTime time.Time) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_JobMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容
func (obj *_JobMgr) GetFromUpdateUser(updateUser int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找
func (obj *_JobMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromJobSalaryStructure 通过job_salary_structure获取内容 薪资结构
func (obj *_JobMgr) GetFromJobSalaryStructure(jobSalaryStructure int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_salary_structure = ?", jobSalaryStructure).Find(&results).Error

	return
}

// GetBatchFromJobSalaryStructure 批量唯一主键查找 薪资结构
func (obj *_JobMgr) GetBatchFromJobSalaryStructure(jobSalaryStructures []int) (results []*Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_salary_structure IN (?)", jobSalaryStructures).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_JobMgr) FetchByPrimaryKey(id int64) (result Job, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
