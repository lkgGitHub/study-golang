package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CompanyInfoMgr struct {
	*_BaseMgr
}

// CompanyInfoMgr open func
func CompanyInfoMgr(db *gorm.DB) *_CompanyInfoMgr {
	if db == nil {
		panic(fmt.Errorf("CompanyInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CompanyInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("company_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CompanyInfoMgr) GetTableName() string {
	return "company_info"
}

// Get 获取
func (obj *_CompanyInfoMgr) Get() (result CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CompanyInfoMgr) Gets() (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_CompanyInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithLogo logo获取 企业logo
func (obj *_CompanyInfoMgr) WithLogo(logo string) Option {
	return optionFunc(func(o *options) { o.query["logo"] = logo })
}

// WithCompanyFullName company_full_Name获取 公司全称
func (obj *_CompanyInfoMgr) WithCompanyFullName(companyFullName string) Option {
	return optionFunc(func(o *options) { o.query["company_full_Name"] = companyFullName })
}

// WithCompanyAbbreviation company_abbreviation获取 公司简称
func (obj *_CompanyInfoMgr) WithCompanyAbbreviation(companyAbbreviation string) Option {
	return optionFunc(func(o *options) { o.query["company_abbreviation"] = companyAbbreviation })
}

// WithCompanyProfession company_profession获取 公司行业
func (obj *_CompanyInfoMgr) WithCompanyProfession(companyProfession string) Option {
	return optionFunc(func(o *options) { o.query["company_profession"] = companyProfession })
}

// WithFinanceSituation finance_situation获取 融资阶段
func (obj *_CompanyInfoMgr) WithFinanceSituation(financeSituation string) Option {
	return optionFunc(func(o *options) { o.query["finance_situation"] = financeSituation })
}

// WithCompanyIntroduction company_introduction获取 公司简介
func (obj *_CompanyInfoMgr) WithCompanyIntroduction(companyIntroduction string) Option {
	return optionFunc(func(o *options) { o.query["company_introduction"] = companyIntroduction })
}

// WithStaffSize staff_size获取 人员规模
func (obj *_CompanyInfoMgr) WithStaffSize(staffSize string) Option {
	return optionFunc(func(o *options) { o.query["staff_size"] = staffSize })
}

// WithBusinessLicense business_license获取 营业执照
func (obj *_CompanyInfoMgr) WithBusinessLicense(businessLicense string) Option {
	return optionFunc(func(o *options) { o.query["business_license"] = businessLicense })
}

// WithPublicityPictures publicity_pictures获取 宣传图片
func (obj *_CompanyInfoMgr) WithPublicityPictures(publicityPictures string) Option {
	return optionFunc(func(o *options) { o.query["publicity_pictures"] = publicityPictures })
}

// WithPublicityVideo publicity_video获取 宣传视频
func (obj *_CompanyInfoMgr) WithPublicityVideo(publicityVideo string) Option {
	return optionFunc(func(o *options) { o.query["publicity_video"] = publicityVideo })
}

// WithWorkStartTime work_start_time获取 工作开始时间
func (obj *_CompanyInfoMgr) WithWorkStartTime(workStartTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["work_start_time"] = workStartTime })
}

// WithWorkEndTime work_end_time获取 工作结束时间
func (obj *_CompanyInfoMgr) WithWorkEndTime(workEndTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["work_end_time"] = workEndTime })
}

// WithIsStress is_stress获取 是否弹性(1:弹性 0：非弹性)
func (obj *_CompanyInfoMgr) WithIsStress(isStress string) Option {
	return optionFunc(func(o *options) { o.query["is_stress"] = isStress })
}

// WithSocialInsuranceType social_insurance_type获取 社保类型（五险一金，六险一金）
func (obj *_CompanyInfoMgr) WithSocialInsuranceType(socialInsuranceType string) Option {
	return optionFunc(func(o *options) { o.query["social_insurance_type"] = socialInsuranceType })
}

// WithAccumulationFundPercent accumulation_fund_Percent获取 公积金比例默认12
func (obj *_CompanyInfoMgr) WithAccumulationFundPercent(accumulationFundPercent string) Option {
	return optionFunc(func(o *options) { o.query["accumulation_fund_Percent"] = accumulationFundPercent })
}

// WithCompanyBenefits company_benefits获取 公司福利
func (obj *_CompanyInfoMgr) WithCompanyBenefits(companyBenefits string) Option {
	return optionFunc(func(o *options) { o.query["company_benefits"] = companyBenefits })
}

// WithOtherBenefits other_benefits获取 其他福利
func (obj *_CompanyInfoMgr) WithOtherBenefits(otherBenefits string) Option {
	return optionFunc(func(o *options) { o.query["other_benefits"] = otherBenefits })
}

// WithCommissionRule commission_rule获取 佣金规则
func (obj *_CompanyInfoMgr) WithCommissionRule(commissionRule string) Option {
	return optionFunc(func(o *options) { o.query["commission_rule"] = commissionRule })
}

// WithProvince province获取 省
func (obj *_CompanyInfoMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 市
func (obj *_CompanyInfoMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithDistrict district获取 区
func (obj *_CompanyInfoMgr) WithDistrict(district string) Option {
	return optionFunc(func(o *options) { o.query["district"] = district })
}

// WithStreet street获取 街道
func (obj *_CompanyInfoMgr) WithStreet(street string) Option {
	return optionFunc(func(o *options) { o.query["street"] = street })
}

// WithLinkMan link_man获取 联系人
func (obj *_CompanyInfoMgr) WithLinkMan(linkMan string) Option {
	return optionFunc(func(o *options) { o.query["link_man"] = linkMan })
}

// WithLinkTel link_Tel获取 联系电话
func (obj *_CompanyInfoMgr) WithLinkTel(linkTel string) Option {
	return optionFunc(func(o *options) { o.query["link_Tel"] = linkTel })
}

// WithAgreement agreement获取 合同样本
func (obj *_CompanyInfoMgr) WithAgreement(agreement string) Option {
	return optionFunc(func(o *options) { o.query["agreement"] = agreement })
}

// WithSource source获取 来源(线上，线下)
func (obj *_CompanyInfoMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithState state获取 公司状态(默认是1，无效:0)
func (obj *_CompanyInfoMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCheckTime check_time获取 审核时间
func (obj *_CompanyInfoMgr) WithCheckTime(checkTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["check_time"] = checkTime })
}

// WithApplyTime apply_time获取 申请时间
func (obj *_CompanyInfoMgr) WithApplyTime(applyTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["apply_time"] = applyTime })
}

// WithCreateTime create_time获取 录入时间
func (obj *_CompanyInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 录入人
func (obj *_CompanyInfoMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CompanyInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_CompanyInfoMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_CompanyInfoMgr) GetByOption(opts ...Option) (result CompanyInfo, err error) {
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
func (obj *_CompanyInfoMgr) GetByOptions(opts ...Option) (results []*CompanyInfo, err error) {
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
func (obj *_CompanyInfoMgr) GetFromID(id int64) (result CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_CompanyInfoMgr) GetBatchFromID(ids []int64) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromLogo 通过logo获取内容 企业logo
func (obj *_CompanyInfoMgr) GetFromLogo(logo string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo = ?", logo).Find(&results).Error

	return
}

// GetBatchFromLogo 批量唯一主键查找 企业logo
func (obj *_CompanyInfoMgr) GetBatchFromLogo(logos []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo IN (?)", logos).Find(&results).Error

	return
}

// GetFromCompanyFullName 通过company_full_Name获取内容 公司全称
func (obj *_CompanyInfoMgr) GetFromCompanyFullName(companyFullName string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_full_Name = ?", companyFullName).Find(&results).Error

	return
}

// GetBatchFromCompanyFullName 批量唯一主键查找 公司全称
func (obj *_CompanyInfoMgr) GetBatchFromCompanyFullName(companyFullNames []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_full_Name IN (?)", companyFullNames).Find(&results).Error

	return
}

// GetFromCompanyAbbreviation 通过company_abbreviation获取内容 公司简称
func (obj *_CompanyInfoMgr) GetFromCompanyAbbreviation(companyAbbreviation string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_abbreviation = ?", companyAbbreviation).Find(&results).Error

	return
}

// GetBatchFromCompanyAbbreviation 批量唯一主键查找 公司简称
func (obj *_CompanyInfoMgr) GetBatchFromCompanyAbbreviation(companyAbbreviations []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_abbreviation IN (?)", companyAbbreviations).Find(&results).Error

	return
}

// GetFromCompanyProfession 通过company_profession获取内容 公司行业
func (obj *_CompanyInfoMgr) GetFromCompanyProfession(companyProfession string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_profession = ?", companyProfession).Find(&results).Error

	return
}

// GetBatchFromCompanyProfession 批量唯一主键查找 公司行业
func (obj *_CompanyInfoMgr) GetBatchFromCompanyProfession(companyProfessions []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_profession IN (?)", companyProfessions).Find(&results).Error

	return
}

// GetFromFinanceSituation 通过finance_situation获取内容 融资阶段
func (obj *_CompanyInfoMgr) GetFromFinanceSituation(financeSituation string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("finance_situation = ?", financeSituation).Find(&results).Error

	return
}

// GetBatchFromFinanceSituation 批量唯一主键查找 融资阶段
func (obj *_CompanyInfoMgr) GetBatchFromFinanceSituation(financeSituations []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("finance_situation IN (?)", financeSituations).Find(&results).Error

	return
}

// GetFromCompanyIntroduction 通过company_introduction获取内容 公司简介
func (obj *_CompanyInfoMgr) GetFromCompanyIntroduction(companyIntroduction string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_introduction = ?", companyIntroduction).Find(&results).Error

	return
}

// GetBatchFromCompanyIntroduction 批量唯一主键查找 公司简介
func (obj *_CompanyInfoMgr) GetBatchFromCompanyIntroduction(companyIntroductions []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_introduction IN (?)", companyIntroductions).Find(&results).Error

	return
}

// GetFromStaffSize 通过staff_size获取内容 人员规模
func (obj *_CompanyInfoMgr) GetFromStaffSize(staffSize string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("staff_size = ?", staffSize).Find(&results).Error

	return
}

// GetBatchFromStaffSize 批量唯一主键查找 人员规模
func (obj *_CompanyInfoMgr) GetBatchFromStaffSize(staffSizes []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("staff_size IN (?)", staffSizes).Find(&results).Error

	return
}

// GetFromBusinessLicense 通过business_license获取内容 营业执照
func (obj *_CompanyInfoMgr) GetFromBusinessLicense(businessLicense string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("business_license = ?", businessLicense).Find(&results).Error

	return
}

// GetBatchFromBusinessLicense 批量唯一主键查找 营业执照
func (obj *_CompanyInfoMgr) GetBatchFromBusinessLicense(businessLicenses []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("business_license IN (?)", businessLicenses).Find(&results).Error

	return
}

// GetFromPublicityPictures 通过publicity_pictures获取内容 宣传图片
func (obj *_CompanyInfoMgr) GetFromPublicityPictures(publicityPictures string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("publicity_pictures = ?", publicityPictures).Find(&results).Error

	return
}

// GetBatchFromPublicityPictures 批量唯一主键查找 宣传图片
func (obj *_CompanyInfoMgr) GetBatchFromPublicityPictures(publicityPicturess []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("publicity_pictures IN (?)", publicityPicturess).Find(&results).Error

	return
}

// GetFromPublicityVideo 通过publicity_video获取内容 宣传视频
func (obj *_CompanyInfoMgr) GetFromPublicityVideo(publicityVideo string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("publicity_video = ?", publicityVideo).Find(&results).Error

	return
}

// GetBatchFromPublicityVideo 批量唯一主键查找 宣传视频
func (obj *_CompanyInfoMgr) GetBatchFromPublicityVideo(publicityVideos []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("publicity_video IN (?)", publicityVideos).Find(&results).Error

	return
}

// GetFromWorkStartTime 通过work_start_time获取内容 工作开始时间
func (obj *_CompanyInfoMgr) GetFromWorkStartTime(workStartTime time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("work_start_time = ?", workStartTime).Find(&results).Error

	return
}

// GetBatchFromWorkStartTime 批量唯一主键查找 工作开始时间
func (obj *_CompanyInfoMgr) GetBatchFromWorkStartTime(workStartTimes []time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("work_start_time IN (?)", workStartTimes).Find(&results).Error

	return
}

// GetFromWorkEndTime 通过work_end_time获取内容 工作结束时间
func (obj *_CompanyInfoMgr) GetFromWorkEndTime(workEndTime time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("work_end_time = ?", workEndTime).Find(&results).Error

	return
}

// GetBatchFromWorkEndTime 批量唯一主键查找 工作结束时间
func (obj *_CompanyInfoMgr) GetBatchFromWorkEndTime(workEndTimes []time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("work_end_time IN (?)", workEndTimes).Find(&results).Error

	return
}

// GetFromIsStress 通过is_stress获取内容 是否弹性(1:弹性 0：非弹性)
func (obj *_CompanyInfoMgr) GetFromIsStress(isStress string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_stress = ?", isStress).Find(&results).Error

	return
}

// GetBatchFromIsStress 批量唯一主键查找 是否弹性(1:弹性 0：非弹性)
func (obj *_CompanyInfoMgr) GetBatchFromIsStress(isStresss []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_stress IN (?)", isStresss).Find(&results).Error

	return
}

// GetFromSocialInsuranceType 通过social_insurance_type获取内容 社保类型（五险一金，六险一金）
func (obj *_CompanyInfoMgr) GetFromSocialInsuranceType(socialInsuranceType string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("social_insurance_type = ?", socialInsuranceType).Find(&results).Error

	return
}

// GetBatchFromSocialInsuranceType 批量唯一主键查找 社保类型（五险一金，六险一金）
func (obj *_CompanyInfoMgr) GetBatchFromSocialInsuranceType(socialInsuranceTypes []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("social_insurance_type IN (?)", socialInsuranceTypes).Find(&results).Error

	return
}

// GetFromAccumulationFundPercent 通过accumulation_fund_Percent获取内容 公积金比例默认12
func (obj *_CompanyInfoMgr) GetFromAccumulationFundPercent(accumulationFundPercent string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("accumulation_fund_Percent = ?", accumulationFundPercent).Find(&results).Error

	return
}

// GetBatchFromAccumulationFundPercent 批量唯一主键查找 公积金比例默认12
func (obj *_CompanyInfoMgr) GetBatchFromAccumulationFundPercent(accumulationFundPercents []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("accumulation_fund_Percent IN (?)", accumulationFundPercents).Find(&results).Error

	return
}

// GetFromCompanyBenefits 通过company_benefits获取内容 公司福利
func (obj *_CompanyInfoMgr) GetFromCompanyBenefits(companyBenefits string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_benefits = ?", companyBenefits).Find(&results).Error

	return
}

// GetBatchFromCompanyBenefits 批量唯一主键查找 公司福利
func (obj *_CompanyInfoMgr) GetBatchFromCompanyBenefits(companyBenefitss []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company_benefits IN (?)", companyBenefitss).Find(&results).Error

	return
}

// GetFromOtherBenefits 通过other_benefits获取内容 其他福利
func (obj *_CompanyInfoMgr) GetFromOtherBenefits(otherBenefits string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other_benefits = ?", otherBenefits).Find(&results).Error

	return
}

// GetBatchFromOtherBenefits 批量唯一主键查找 其他福利
func (obj *_CompanyInfoMgr) GetBatchFromOtherBenefits(otherBenefitss []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other_benefits IN (?)", otherBenefitss).Find(&results).Error

	return
}

// GetFromCommissionRule 通过commission_rule获取内容 佣金规则
func (obj *_CompanyInfoMgr) GetFromCommissionRule(commissionRule string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("commission_rule = ?", commissionRule).Find(&results).Error

	return
}

// GetBatchFromCommissionRule 批量唯一主键查找 佣金规则
func (obj *_CompanyInfoMgr) GetBatchFromCommissionRule(commissionRules []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("commission_rule IN (?)", commissionRules).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省
func (obj *_CompanyInfoMgr) GetFromProvince(province string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("province = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量唯一主键查找 省
func (obj *_CompanyInfoMgr) GetBatchFromProvince(provinces []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("province IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 市
func (obj *_CompanyInfoMgr) GetFromCity(city string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找 市
func (obj *_CompanyInfoMgr) GetBatchFromCity(citys []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromDistrict 通过district获取内容 区
func (obj *_CompanyInfoMgr) GetFromDistrict(district string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("district = ?", district).Find(&results).Error

	return
}

// GetBatchFromDistrict 批量唯一主键查找 区
func (obj *_CompanyInfoMgr) GetBatchFromDistrict(districts []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("district IN (?)", districts).Find(&results).Error

	return
}

// GetFromStreet 通过street获取内容 街道
func (obj *_CompanyInfoMgr) GetFromStreet(street string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("street = ?", street).Find(&results).Error

	return
}

// GetBatchFromStreet 批量唯一主键查找 街道
func (obj *_CompanyInfoMgr) GetBatchFromStreet(streets []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("street IN (?)", streets).Find(&results).Error

	return
}

// GetFromLinkMan 通过link_man获取内容 联系人
func (obj *_CompanyInfoMgr) GetFromLinkMan(linkMan string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_man = ?", linkMan).Find(&results).Error

	return
}

// GetBatchFromLinkMan 批量唯一主键查找 联系人
func (obj *_CompanyInfoMgr) GetBatchFromLinkMan(linkMans []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_man IN (?)", linkMans).Find(&results).Error

	return
}

// GetFromLinkTel 通过link_Tel获取内容 联系电话
func (obj *_CompanyInfoMgr) GetFromLinkTel(linkTel string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_Tel = ?", linkTel).Find(&results).Error

	return
}

// GetBatchFromLinkTel 批量唯一主键查找 联系电话
func (obj *_CompanyInfoMgr) GetBatchFromLinkTel(linkTels []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_Tel IN (?)", linkTels).Find(&results).Error

	return
}

// GetFromAgreement 通过agreement获取内容 合同样本
func (obj *_CompanyInfoMgr) GetFromAgreement(agreement string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("agreement = ?", agreement).Find(&results).Error

	return
}

// GetBatchFromAgreement 批量唯一主键查找 合同样本
func (obj *_CompanyInfoMgr) GetBatchFromAgreement(agreements []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("agreement IN (?)", agreements).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 来源(线上，线下)
func (obj *_CompanyInfoMgr) GetFromSource(source string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("source = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量唯一主键查找 来源(线上，线下)
func (obj *_CompanyInfoMgr) GetBatchFromSource(sources []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("source IN (?)", sources).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 公司状态(默认是1，无效:0)
func (obj *_CompanyInfoMgr) GetFromState(state string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量唯一主键查找 公司状态(默认是1，无效:0)
func (obj *_CompanyInfoMgr) GetBatchFromState(states []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

// GetFromCheckTime 通过check_time获取内容 审核时间
func (obj *_CompanyInfoMgr) GetFromCheckTime(checkTime time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("check_time = ?", checkTime).Find(&results).Error

	return
}

// GetBatchFromCheckTime 批量唯一主键查找 审核时间
func (obj *_CompanyInfoMgr) GetBatchFromCheckTime(checkTimes []time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("check_time IN (?)", checkTimes).Find(&results).Error

	return
}

// GetFromApplyTime 通过apply_time获取内容 申请时间
func (obj *_CompanyInfoMgr) GetFromApplyTime(applyTime time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_time = ?", applyTime).Find(&results).Error

	return
}

// GetBatchFromApplyTime 批量唯一主键查找 申请时间
func (obj *_CompanyInfoMgr) GetBatchFromApplyTime(applyTimes []time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_time IN (?)", applyTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 录入时间
func (obj *_CompanyInfoMgr) GetFromCreateTime(createTime time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 录入时间
func (obj *_CompanyInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 录入人
func (obj *_CompanyInfoMgr) GetFromCreateUser(createUser string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 录入人
func (obj *_CompanyInfoMgr) GetBatchFromCreateUser(createUsers []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CompanyInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_CompanyInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_CompanyInfoMgr) GetFromUpdateUser(updateUser string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 更新人
func (obj *_CompanyInfoMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CompanyInfoMgr) FetchByPrimaryKey(id int64) (result CompanyInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
