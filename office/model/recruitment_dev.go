package model

import (
	"time"

	"gorm.io/datatypes"
)

// CoinRecord [...]
type CoinRecord struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键ID
	UserID     int64     `gorm:"column:user_id;type:bigint"`                // 同户ID
	Type       int       `gorm:"column:type;type:int"`                      // 记录类型 1-收入 2-支出
	CoinNum    int       `gorm:"column:coin_num;type:int"`                  // 币值
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`          // 创建时间
	Desc       string    `gorm:"column:desc;type:varchar(256)"`             // 收支描述
}

// Collection [...]
type Collection struct {
	ID         int       `gorm:"primaryKey;column:id;type:int;not null"` //
	UserID     int       `gorm:"column:user_id;type:int"`                // 用户id
	JobID      int       `gorm:"column:job_id;type:int"`                 // 职位id
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`       // 收藏时间
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime"`       // 删除时间
	Status     string    `gorm:"column:status;type:varchar(255)"`        // 职位状态。收藏职位失效或者下架
	CreateUser int64     `gorm:"column:create_user;type:bigint"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"` // 修改时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`
}

// CompanyInfo 公司信息表
type CompanyInfo struct {
	ID                      int64     `gorm:"primaryKey;column:id;type:bigint;not null"`                     // 主键
	Logo                    string    `gorm:"column:logo;type:varchar(255)"`                                 // 企业logo
	CompanyFullName         string    `gorm:"column:company_full_Name;type:varchar(100);not null"`           // 公司全称
	CompanyAbbreviation     string    `gorm:"column:company_abbreviation;type:varchar(100);not null"`        // 公司简称
	CompanyProfession       string    `gorm:"column:company_profession;type:varchar(100);not null"`          // 公司行业
	FinanceSituation        string    `gorm:"column:finance_situation;type:varchar(100)"`                    // 融资阶段
	CompanyIntroduction     string    `gorm:"column:company_introduction;type:varchar(4000);not null"`       // 公司简介
	StaffSize               string    `gorm:"column:staff_size;type:varchar(20);not null"`                   // 人员规模
	BusinessLicense         string    `gorm:"column:business_license;type:varchar(100);not null"`            // 营业执照
	PublicityPictures       string    `gorm:"column:publicity_pictures;type:varchar(100)"`                   // 宣传图片
	PublicityVideo          string    `gorm:"column:publicity_video;type:varchar(100)"`                      // 宣传视频
	WorkStartTime           time.Time `gorm:"column:work_start_time;type:timestamp"`                         // 工作开始时间
	WorkEndTime             time.Time `gorm:"column:work_end_time;type:timestamp"`                           // 工作结束时间
	IsStress                string    `gorm:"column:is_stress;type:varchar(2)"`                              // 是否弹性(1:弹性 0：非弹性)
	SocialInsuranceType     string    `gorm:"column:social_insurance_type;type:varchar(2)"`                  // 社保类型（五险一金，六险一金）
	AccumulationFundPercent string    `gorm:"column:accumulation_fund_Percent;type:varchar(4);default:12"`   // 公积金比例默认12
	CompanyBenefits         string    `gorm:"column:company_benefits;type:varchar(40);default:12"`           // 公司福利
	OtherBenefits           string    `gorm:"column:other_benefits;type:varchar(500);default:12"`            // 其他福利
	CommissionRule          string    `gorm:"column:commission_rule;type:varchar(1000);not null;default:12"` // 佣金规则
	Province                string    `gorm:"column:province;type:varchar(100);not null;default:12"`         // 省
	City                    string    `gorm:"column:city;type:varchar(100);not null;default:12"`             // 市
	District                string    `gorm:"column:district;type:varchar(100);not null;default:12"`         // 区
	Street                  string    `gorm:"column:street;type:varchar(200);not null;default:12"`           // 街道
	LinkMan                 string    `gorm:"column:link_man;type:varchar(100);not null;default:12"`         // 联系人
	LinkTel                 string    `gorm:"column:link_Tel;type:varchar(100);not null;default:12"`         // 联系电话
	Agreement               string    `gorm:"column:agreement;type:varchar(100)"`                            // 合同样本
	Source                  string    `gorm:"column:source;type:varchar(100);default:12"`                    // 来源(线上，线下)
	State                   string    `gorm:"column:state;type:varchar(100);not null;default:12"`            // 公司状态(默认是1，无效:0)
	CheckTime               time.Time `gorm:"column:check_time;type:timestamp"`                              // 审核时间
	ApplyTime               time.Time `gorm:"column:apply_time;type:timestamp"`                              // 申请时间
	CreateTime              time.Time `gorm:"column:create_time;type:timestamp;not null"`                    // 录入时间
	CreateUser              string    `gorm:"column:create_user;type:varchar(100);not null"`                 // 录入人
	UpdateTime              time.Time `gorm:"column:update_time;type:timestamp"`                             // 更新时间
	UpdateUser              string    `gorm:"column:update_user;type:varchar(100)"`                          // 更新人
}

// Friend [...]
type Friend struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`
	UserID     int64     `gorm:"uniqueIndex:user_id;column:user_id;type:bigint"`
	FriendID   int64     `gorm:"uniqueIndex:user_id;column:friend_id;type:bigint"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime"`
}

// Job [...]
type Job struct {
	ID                      int64     `gorm:"primaryKey;column:id;type:bigint;not null"`            // 主键id
	Name                    string    `gorm:"column:name;type:varchar(255);not null"`               // 职位名称
	Status                  string    `gorm:"column:status;type:varchar(255)"`                      // 职位状态
	RecruitingNumbers       int       `gorm:"column:recruiting_numbers;type:int;not null"`          // 招聘人数
	SuccessfulNumbers       int       `gorm:"column:successful_numbers;type:int"`                   // 通过人数
	JobDepartment           string    `gorm:"column:job_department;type:varchar(255)"`              // 所在部门
	UnderlingNumber         int       `gorm:"column:underling_number;type:int;default:0"`           // 下属人数
	SalaryStart             int       `gorm:"column:salary_start;type:int;not null"`                // 薪水起始值
	SalaryEnd               int       `gorm:"column:salary_end;type:int;not null"`                  // 薪水结束值
	City                    string    `gorm:"column:city;type:varchar(255);not null"`               // 公司城市
	Address                 string    `gorm:"column:address;type:varchar(255);not null"`            // 公司地址（具体到区）
	DetailedLocation        string    `gorm:"column:detailed_location;type:varchar(255)"`           // 详细地点
	InterviewAddr           string    `gorm:"column:interview_addr;type:varchar(255)"`              // 面试地址
	Education               string    `gorm:"column:education;type:varchar(255)"`                   // 学历
	WorkingYearsType        string    `gorm:"column:working_years_type;type:varchar(255);not null"` // 工作年限类型
	WorkingYearsStart       int       `gorm:"column:working_years_start;type:int"`                  // 工作年限起始值
	WorkingYearsEnd         int       `gorm:"column:working_years_end;type:int"`                    // 工作年限结束值
	Industry                string    `gorm:"column:industry;type:varchar(255)"`                    // 行业
	JobBackdrop             string    `gorm:"column:job_backdrop;type:varchar(255)"`                // 背景经验
	JobHopFre               float64   `gorm:"column:job_hop_fre;type:double"`                       // 跳槽频率
	AgeRequirement          string    `gorm:"column:age_requirement;type:varchar(255)"`             // 年龄要求
	InterviewRounds         int       `gorm:"column:interview_rounds;type:int"`                     // 面试轮数
	JobCommissionCashStart  float64   `gorm:"column:job_commission_cash_start;type:decimal(10,2)"`  // 佣金兑现标准起始值
	JobCommissionCashEnd    float64   `gorm:"column:job_commission_cash_end;type:decimal(10,2)"`    // 佣金兑现标准结束值
	JobCommissionCashPeriod string    `gorm:"column:job_commission_cash_period;type:varchar(255)"`  // 佣金兑现期限
	JobKeyword              string    `gorm:"column:job_keyword;type:varchar(255)"`                 // 岗位关键字
	CompanyID               int64     `gorm:"column:company_id;type:bigint"`                        // 关联企业主键ID
	IsUrgent                []uint8   `gorm:"column:is_urgent;type:bit(1);default:b'0'"`            // 是否为紧急招聘, 默认为否
	JobDetailsID            int64     `gorm:"column:job_details_id;type:bigint"`                    // 工作详情
	Deleted                 []uint8   `gorm:"column:deleted;type:bit(1)"`                           // 逻辑删除字段
	CreateTime              time.Time `gorm:"column:create_time;type:datetime"`                     // 创建时间
	CreateUser              int64     `gorm:"column:create_user;type:bigint"`
	UpdateTime              time.Time `gorm:"column:update_time;type:datetime"` // 修改时间
	UpdateUser              int64     `gorm:"column:update_user;type:bigint"`
	JobSalaryStructure      int       `gorm:"column:job_salary_structure;type:int"` // 薪资结构
}

// JobDetails [...]
type JobDetails struct {
	ID                  int64  `gorm:"primaryKey;column:id;type:bigint;not null"`
	JobResponsibilities string `gorm:"column:job_responsibilities;type:varchar(255)"`
	JobRequirements     string `gorm:"column:job_requirements;type:varchar(255)"`
}

// JobRecommendRecord [...]
type JobRecommendRecord struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` //
	Presenter  int64     `gorm:"column:presenter;type:bigint;not null"`     // 推荐者
	UserID     int64     `gorm:"column:user_id;type:bigint;not null"`       // 用户id
	JobID      int64     `gorm:"column:job_id;type:bigint;not null"`        // 职位id
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`          // 收藏时间
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime"`          // 删除时间
}

// Member 面币表
type Member struct {
	ID         int       `gorm:"primaryKey;column:id;type:int;not null"` // 主键ID
	UserID     int       `gorm:"unique;column:user_id;type:int"`         // 用户ID
	IsVip      int       `gorm:"column:is_vip;type:int;default:0"`       // 是否vip 0-普通用户 1-推荐师
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
}

// Order [...]
type Order struct {
	ID              int64     `gorm:"primaryKey;column:id;type:bigint;not null"`               // 主键ID
	CompanyID       int64     `gorm:"column:company_id;type:bigint"`                           // 企业ID
	JobID           int64     `gorm:"column:job_id;type:bigint"`                               // 职位ID
	UserID          int64     `gorm:"column:user_id;type:bigint"`                              // 用户ID
	OrderNo         string    `gorm:"unique;column:order_no;type:varchar(20);not null"`        // 订单编号
	OrderStatus     int       `gorm:"column:order_status;type:int;not null"`                   // 订单状态 1-初始化 2-待结算 3-结算中 4-已完成 5-已关闭
	BaseCommission  float64   `gorm:"column:base_commission;type:decimal(10,2);not null"`      // 佣金计算基数
	OrderAmount     float64   `gorm:"column:order_amount;type:decimal(10,2);default:0.00"`     // 订单金额
	SettledAmount   float64   `gorm:"column:settled_amount;type:decimal(10,2);default:0.00"`   // 已结算金额
	UnsettledAmount float64   `gorm:"column:unsettled_amount;type:decimal(10,2);default:0.00"` // 未结算金额
	CompleteTime    time.Time `gorm:"column:complete_time;type:datetime"`                      // 订单完成时间
	CreateTime      time.Time `gorm:"column:create_time;type:datetime"`                        // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime"`                        // 更新时间
	CreateUser      int64     `gorm:"column:create_user;type:bigint"`
	UpdateUser      int64     `gorm:"column:update_user;type:bigint"`
}

// OrderSeq [...]
type OrderSeq struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`          // 主键ID
	OrderDate  string    `gorm:"unique;column:order_date;type:varchar(10);not null"` // 日期
	Current    int       `gorm:"column:current;type:int;not null;default:1"`         // 当前值
	Increment  int       `gorm:"column:increment;type:int;default:1"`                // 自增值
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`                   // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`                   // 更新时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`                     // 创建人
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`                     // 更新人
}

// RegisterRecommendRecord [...]
type RegisterRecommendRecord struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` //
	ReferrerID int64     `gorm:"column:referrer_id;type:bigint;not null"`   // 推荐者
	UserID     int64     `gorm:"column:user_id;type:bigint;not null"`       // 用户id
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`          // 推荐注册时间
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime"`          // 删除时间
}

// Resume [...]
type Resume struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`
	ResumePath string    `gorm:"column:resume_path;type:varchar(255)"` // 简历地址
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime"`     // 删除时间
	State      string    `gorm:"column:state;type:varchar(2)"`         // 简历状态(0:有效，1:删除)
	ResumeName string    `gorm:"column:resume_name;type:varchar(255)"` // 简历文件名称
	UserPhone  string    `gorm:"column:user_phone;type:varchar(20)"`   // 用户手机号
	Format     string    `gorm:"column:format;type:varchar(10)"`       // 文件格式
	Size       string    `gorm:"column:size;type:varchar(10)"`         // 简历大小
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`     // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`     // 修改时间
	CreateUser string    `gorm:"column:create_user;type:varchar(255)"` // 创建人
	UpdateUser string    `gorm:"column:update_user;type:varchar(255)"` // 修改人
}

// SendResumeRecord 投递记录表
type SendResumeRecord struct {
	ID             int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键id
	Presenter      int64     `gorm:"column:presenter;type:bigint"`              // 推荐人id
	UserID         int64     `gorm:"column:user_id;type:bigint;not null"`       // 投递人用户id
	JobID          int64     `gorm:"column:job_id;type:bigint"`                 // 职位id
	ResumeID       int64     `gorm:"column:resume_id;type:bigint"`              // 简历id
	JobRecommendID int64     `gorm:"column:job_recommend_id;type:bigint"`       // 职位推荐记录id
	Status         int       `gorm:"column:status;type:int"`                    // 简历投递状态 1-已投递 2-通过初筛 3-面试中 4-通过面试 5-已入职 6-已转正 7-不合适 8-待确认中
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`          // 收藏时间
	DeleteTime     time.Time `gorm:"column:delete_time;type:datetime"`          // 删除时间
}

// SysApp 系统应用表
type SysApp struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键id
	Name       string    `gorm:"column:name;type:varchar(100);not null"`    // 应用名称
	Code       string    `gorm:"column:code;type:varchar(50);not null"`     // 编码
	Active     string    `gorm:"column:active;type:varchar(1)"`             // 是否默认激活（Y-是，N-否）
	Status     int8      `gorm:"column:status;type:tinyint;not null"`       // 状态（字典 0正常 1停用 2删除）
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`          // 创建时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`            // 创建人
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`          // 修改时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`            // 修改人
}

// SysConfig 系统参数配置表
type SysConfig struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`                    // 主键
	Name       string    `gorm:"column:name;type:varchar(100);not null"`                       // 名称
	Code       string    `gorm:"column:code;type:varchar(50);not null"`                        // 编码
	Value      string    `gorm:"column:value;type:varchar(255);not null"`                      // 值
	SysFlag    string    `gorm:"column:sys_flag;type:char(1);not null"`                        // 是否是系统参数（Y-是，N-否）
	Remark     string    `gorm:"column:remark;type:varchar(255)"`                              // 备注
	Status     int8      `gorm:"column:status;type:tinyint;not null"`                          // 状态（字典 0正常 1停用 2删除）
	GroupCode  string    `gorm:"column:group_code;type:varchar(255);not null;default:DEFAULT"` // 常量所属分类的编码，来自于“常量的分类”字典
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`                             // 创建时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`                               // 创建人
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`                             // 更新时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`                               // 更新人
}

// SysDictData 系统字典值表
type SysDictData struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	TypeID     int64     `gorm:"column:type_id;type:bigint;not null"`       // 字典类型id
	Value      string    `gorm:"column:value;type:text;not null"`           // 值
	Code       string    `gorm:"column:code;type:varchar(50);not null"`     // 编码
	Sort       int       `gorm:"column:sort;type:int;not null"`             // 排序
	Remark     string    `gorm:"column:remark;type:varchar(255)"`           // 备注
	Status     int8      `gorm:"column:status;type:tinyint;not null"`       // 状态（字典 0正常 1停用 2删除）
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`          // 创建时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`            // 创建人
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`          // 更新时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`            // 更新人
}

// SysDictType 系统字典类型表
type SysDictType struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	Name       string    `gorm:"column:name;type:varchar(100);not null"`    // 名称
	Code       string    `gorm:"column:code;type:varchar(50);not null"`     // 编码
	Sort       int       `gorm:"column:sort;type:int;not null"`             // 排序
	Remark     string    `gorm:"column:remark;type:varchar(255)"`           // 备注
	Status     int8      `gorm:"column:status;type:tinyint;not null"`       // 状态（字典 0正常 1停用 2删除）
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`          // 创建时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`            // 创建人
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`          // 更新时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`            // 更新人
}

// SysEmp 员工表
type SysEmp struct {
	ID      int64  `gorm:"primaryKey;column:id;type:bigint;not null"`  // 主键
	JobNum  string `gorm:"column:job_num;type:varchar(100)"`           // 工号
	OrgID   int64  `gorm:"column:org_id;type:bigint;not null"`         // 所属机构id
	OrgName string `gorm:"column:org_name;type:varchar(100);not null"` // 所属机构名称
}

// SysEmpExtOrgPos 员工附属机构岗位表
type SysEmpExtOrgPos struct {
	ID    int64 `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	EmpID int64 `gorm:"column:emp_id;type:bigint;not null"`        // 员工id
	OrgID int64 `gorm:"column:org_id;type:bigint;not null"`        // 机构id
	PosID int64 `gorm:"column:pos_id;type:bigint;not null"`        // 岗位id
}

// SysEmpPos 员工职位关联表
type SysEmpPos struct {
	ID    int64 `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	EmpID int64 `gorm:"column:emp_id;type:bigint;not null"`        // 员工id
	PosID int64 `gorm:"column:pos_id;type:bigint;not null"`        // 职位id
}

// SysFileInfo 文件信息表
type SysFileInfo struct {
	ID             int64     `gorm:"primaryKey;column:id;type:bigint;not null"`          // 主键id
	FileLocation   int8      `gorm:"column:file_location;type:tinyint;not null"`         // 文件存储位置（1:阿里云，2:腾讯云，3:minio，4:本地）
	FileBucket     string    `gorm:"column:file_bucket;type:varchar(1000)"`              // 文件仓库
	FileOriginName string    `gorm:"column:file_origin_name;type:varchar(100);not null"` // 文件名称（上传时候的文件名）
	FileSuffix     string    `gorm:"column:file_suffix;type:varchar(50)"`                // 文件后缀
	FileSizeKb     int64     `gorm:"column:file_size_kb;type:bigint"`                    // 文件大小kb
	FileSizeInfo   string    `gorm:"column:file_size_info;type:varchar(100)"`            // 文件大小信息，计算后的
	FileObjectName string    `gorm:"column:file_object_name;type:varchar(100);not null"` // 存储到bucket的名称（文件唯一标识id）
	FilePath       string    `gorm:"column:file_path;type:varchar(1000)"`                // 存储路径
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`                   // 创建时间
	CreateUser     int64     `gorm:"column:create_user;type:bigint"`                     // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime"`                   // 修改时间
	UpdateUser     int64     `gorm:"column:update_user;type:bigint"`                     // 修改用户
}

// SysMenu 系统菜单表
type SysMenu struct {
	ID          int64     `gorm:"primaryKey;column:id;type:bigint;not null"`     // 主键
	Pid         int64     `gorm:"column:pid;type:bigint;not null"`               // 父id
	Pids        string    `gorm:"column:pids;type:text;not null"`                // 父ids
	Name        string    `gorm:"column:name;type:varchar(100);not null"`        // 名称
	Code        string    `gorm:"column:code;type:varchar(50);not null"`         // 编码
	Type        int8      `gorm:"column:type;type:tinyint;not null;default:1"`   // 菜单类型（字典 0目录 1菜单 2按钮）
	Icon        string    `gorm:"column:icon;type:varchar(255)"`                 // 图标
	Router      string    `gorm:"column:router;type:varchar(255)"`               // 路由地址
	Component   string    `gorm:"column:component;type:varchar(255)"`            // 组件地址
	Permission  string    `gorm:"column:permission;type:varchar(255)"`           // 权限标识
	Application string    `gorm:"column:application;type:varchar(50);not null"`  // 应用分类（应用编码）
	OpenType    int8      `gorm:"column:open_type;type:tinyint;not null"`        // 打开方式（字典 0无 1组件 2内链 3外链）
	Visible     string    `gorm:"column:visible;type:char(1);not null"`          // 是否可见（Y-是，N-否）
	Link        string    `gorm:"column:link;type:varchar(255)"`                 // 链接地址
	Redirect    string    `gorm:"column:redirect;type:varchar(255)"`             // 重定向地址
	Weight      int8      `gorm:"column:weight;type:tinyint"`                    // 权重（字典 1系统权重 2业务权重）
	Sort        int       `gorm:"column:sort;type:int;not null"`                 // 排序
	Remark      string    `gorm:"column:remark;type:varchar(255)"`               // 备注
	Status      int8      `gorm:"column:status;type:tinyint;not null;default:0"` // 状态（字典 0正常 1停用 2删除）
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`              // 创建时间
	CreateUser  int64     `gorm:"column:create_user;type:bigint"`                // 创建人
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`              // 修改时间
	UpdateUser  int64     `gorm:"column:update_user;type:bigint"`                // 修改人
}

// SysNotice 通知表
type SysNotice struct {
	ID             int64     `gorm:"primaryKey;column:id;type:bigint;not null"`          // 主键
	Title          string    `gorm:"column:title;type:varchar(1000)"`                    // 标题
	Content        string    `gorm:"column:content;type:text"`                           // 内容
	Type           int8      `gorm:"column:type;type:tinyint;not null"`                  // 类型（字典 1通知 2公告）
	PublicUserID   int64     `gorm:"column:public_user_id;type:bigint;not null"`         // 发布人id
	PublicUserName string    `gorm:"column:public_user_name;type:varchar(100);not null"` // 发布人姓名
	PublicOrgID    int64     `gorm:"column:public_org_id;type:bigint"`                   // 发布机构id
	PublicOrgName  string    `gorm:"column:public_org_name;type:varchar(50)"`            // 发布机构名称
	PublicTime     time.Time `gorm:"column:public_time;type:datetime"`                   // 发布时间
	CancelTime     time.Time `gorm:"column:cancel_time;type:datetime"`                   // 撤回时间
	Status         int8      `gorm:"column:status;type:tinyint;not null"`                // 状态（字典 0草稿 1发布 2撤回 3删除）
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`                   // 创建时间
	CreateUser     int64     `gorm:"column:create_user;type:bigint"`                     // 创建人
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime"`                   // 修改时间
	UpdateUser     int64     `gorm:"column:update_user;type:bigint"`                     // 修改人
}

// SysNoticeUser 系统用户数据范围表
type SysNoticeUser struct {
	ID       int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	NoticeID int64     `gorm:"column:notice_id;type:bigint;not null"`     // 通知公告id
	UserID   int64     `gorm:"column:user_id;type:bigint;not null"`       // 用户id
	Status   int8      `gorm:"column:status;type:tinyint;not null"`       // 状态（字典 0未读 1已读）
	ReadTime time.Time `gorm:"column:read_time;type:datetime"`            // 阅读时间
}

// SysOpLog 系统操作日志表
type SysOpLog struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	Name       string    `gorm:"column:name;type:varchar(50)"`              // 名称
	OpType     int8      `gorm:"column:op_type;type:tinyint"`               // 操作类型
	Success    string    `gorm:"column:success;type:char(1)"`               // 是否执行成功（Y-是，N-否）
	Message    string    `gorm:"column:message;type:text"`                  // 具体消息
	IP         string    `gorm:"column:ip;type:varchar(255)"`               // ip
	Location   string    `gorm:"column:location;type:varchar(255)"`         // 地址
	Browser    string    `gorm:"column:browser;type:varchar(255)"`          // 浏览器
	Os         string    `gorm:"column:os;type:varchar(255)"`               // 操作系统
	URL        string    `gorm:"column:url;type:varchar(500)"`              // 请求地址
	ClassName  string    `gorm:"column:class_name;type:varchar(500)"`       // 类名称
	MethodName string    `gorm:"column:method_name;type:varchar(500)"`      // 方法名称
	ReqMethod  string    `gorm:"column:req_method;type:varchar(255)"`       // 请求方式（GET POST PUT DELETE)
	Param      string    `gorm:"column:param;type:text"`                    // 请求参数
	Result     string    `gorm:"column:result;type:longtext"`               // 返回结果
	OpTime     time.Time `gorm:"column:op_time;type:datetime"`              // 操作时间
	Account    string    `gorm:"column:account;type:varchar(50)"`           // 操作账号
}

// SysOrg 系统组织机构表
type SysOrg struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`     // 主键
	Pid        int64     `gorm:"column:pid;type:bigint;not null"`               // 父id
	Pids       string    `gorm:"column:pids;type:text;not null"`                // 父ids
	Name       string    `gorm:"column:name;type:varchar(100);not null"`        // 名称
	Code       string    `gorm:"column:code;type:varchar(50);not null"`         // 编码
	Sort       int       `gorm:"column:sort;type:int;not null"`                 // 排序
	Remark     string    `gorm:"column:remark;type:varchar(255)"`               // 描述
	Status     int8      `gorm:"column:status;type:tinyint;not null;default:0"` // 状态（字典 0正常 1停用 2删除）
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`              // 创建时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`                // 创建人
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`              // 更新时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`                // 更新人
}

// SysPos 系统职位表
type SysPos struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`     // 主键
	Name       string    `gorm:"column:name;type:varchar(100);not null"`        // 名称
	Code       string    `gorm:"unique;column:code;type:varchar(50);not null"`  // 编码
	Sort       int       `gorm:"column:sort;type:int;not null"`                 // 排序
	Remark     string    `gorm:"column:remark;type:varchar(255)"`               // 备注
	Status     int8      `gorm:"column:status;type:tinyint;not null;default:0"` // 状态（字典 0正常 1停用 2删除）
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`              // 创建时间
	CreateUser int64     `gorm:"column:create_user;type:bigint"`                // 创建人
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`              // 更新时间
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`                // 更新人
}

// SysRole 系统角色表
type SysRole struct {
	ID            int64     `gorm:"primaryKey;column:id;type:bigint;not null"`              // 主键id
	Name          string    `gorm:"column:name;type:varchar(100);not null"`                 // 名称
	Code          string    `gorm:"column:code;type:varchar(50);not null"`                  // 编码
	Sort          int       `gorm:"column:sort;type:int;not null"`                          // 序号
	DataScopeType int8      `gorm:"column:data_scope_type;type:tinyint;not null;default:1"` // 数据范围类型（字典 1全部数据 2本部门及以下数据 3本部门数据 4仅本人数据 5自定义数据）
	Remark        string    `gorm:"column:remark;type:varchar(255)"`                        // 备注
	Status        int8      `gorm:"column:status;type:tinyint;not null;default:0"`          // 状态（字典 0正常 1停用 2删除）
	CreateTime    time.Time `gorm:"column:create_time;type:datetime"`                       // 创建时间
	CreateUser    int64     `gorm:"column:create_user;type:bigint"`                         // 创建人
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime"`                       // 更新时间
	UpdateUser    int64     `gorm:"column:update_user;type:bigint"`                         // 更新人
}

// SysRoleDataScope 系统角色数据范围表
type SysRoleDataScope struct {
	ID     int64 `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	RoleID int64 `gorm:"column:role_id;type:bigint;not null"`       // 角色id
	OrgID  int64 `gorm:"column:org_id;type:bigint;not null"`        // 机构id
}

// SysRoleMenu 系统角色菜单表
type SysRoleMenu struct {
	ID     int64 `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	RoleID int64 `gorm:"column:role_id;type:bigint;not null"`       // 角色id
	MenuID int64 `gorm:"column:menu_id;type:bigint;not null"`       // 菜单id
}

// SysSms 短信信息发送表
type SysSms struct {
	ID           int64     `gorm:"primaryKey;column:id;type:bigint;not null"`                  // 主键
	PhoneNumbers string    `gorm:"column:phone_numbers;type:varchar(200);not null"`            // 手机号
	ValidateCode string    `gorm:"column:validate_code;type:varchar(255)"`                     // 短信验证码
	TemplateCode string    `gorm:"column:template_code;type:varchar(255)"`                     // 短信模板ID
	BizID        string    `gorm:"column:biz_id;type:varchar(255)"`                            // 回执id，可根据该id查询具体的发送状态
	Status       int8      `gorm:"column:status;type:tinyint;not null"`                        // 发送状态（字典 0 未发送，1 发送成功，2 发送失败，3 失效）
	Source       int8      `gorm:"column:source;type:tinyint;not null"`                        // 来源（字典 1 app， 2 pc， 3 其他）
	InvalidTime  time.Time `gorm:"column:invalid_time;type:datetime"`                          // 失效时间
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP"` // 创建时间
	CreateUser   int64     `gorm:"column:create_user;type:bigint"`                             // 创建人
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime"`                           // 更新时间
	UpdateUser   int64     `gorm:"column:update_user;type:bigint"`                             // 更新人
}

// SysTimers 定时任务
type SysTimers struct {
	ID          int64     `gorm:"primaryKey;column:id;type:bigint;not null"`      // 定时器id
	TimerName   string    `gorm:"column:timer_name;type:varchar(255);default:''"` // 任务名称
	ActionClass string    `gorm:"column:action_class;type:varchar(255)"`          // 执行任务的class的类名（实现了TimerTaskRunner接口的类的全称）
	Cron        string    `gorm:"column:cron;type:varchar(255);default:''"`       // 定时任务表达式
	JobStatus   int8      `gorm:"column:job_status;type:tinyint;default:0"`       // 状态（字典 1运行  2停止）
	Remark      string    `gorm:"column:remark;type:varchar(1000);default:''"`    // 备注信息
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`               // 创建时间
	CreateUser  int64     `gorm:"column:create_user;type:bigint"`                 // 创建人
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`               // 更新时间
	UpdateUser  int64     `gorm:"column:update_user;type:bigint"`                 // 更新人
}

// SysUser 系统用户表
type SysUser struct {
	ID            int64          `gorm:"primaryKey;column:id;type:bigint;not null"`         // 主键
	Account       string         `gorm:"column:account;type:varchar(50);not null"`          // 账号
	Password      string         `gorm:"column:password;type:varchar(100);not null"`        // 密码
	NickName      string         `gorm:"column:nick_name;type:varchar(50)"`                 // 昵称
	Name          string         `gorm:"column:name;type:varchar(100);not null"`            // 姓名
	Avatar        int64          `gorm:"column:avatar;type:bigint"`                         // 头像
	Birthday      datatypes.Date `gorm:"column:birthday;type:date"`                         // 生日
	Sex           int8           `gorm:"column:sex;type:tinyint;not null"`                  // 性别(字典 1男 2女 3未知)
	Email         string         `gorm:"column:email;type:varchar(50)"`                     // 邮箱
	Phone         string         `gorm:"column:phone;type:varchar(50)"`                     // 手机
	Tel           string         `gorm:"column:tel;type:varchar(50)"`                       // 电话
	LastLoginIP   string         `gorm:"column:last_login_ip;type:varchar(100)"`            // 最后登陆IP
	LastLoginTime time.Time      `gorm:"column:last_login_time;type:datetime"`              // 最后登陆时间
	AdminType     int8           `gorm:"column:admin_type;type:tinyint;not null;default:0"` // 管理员类型（0超级管理员 1非管理员）
	Status        int8           `gorm:"column:status;type:tinyint;not null;default:0"`     // 状态（字典 0正常 1冻结 2删除）
	CreateTime    time.Time      `gorm:"column:create_time;type:datetime"`                  // 创建时间
	CreateUser    int64          `gorm:"column:create_user;type:bigint"`                    // 创建人
	UpdateTime    time.Time      `gorm:"column:update_time;type:datetime"`                  // 更新时间
	UpdateUser    int64          `gorm:"column:update_user;type:bigint"`                    // 更新人
}

// SysUserDataScope 系统用户数据范围表
type SysUserDataScope struct {
	ID     int64 `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	UserID int64 `gorm:"column:user_id;type:bigint;not null"`       // 用户id
	OrgID  int64 `gorm:"column:org_id;type:bigint;not null"`        // 机构id
}

// SysUserRole 系统用户角色表
type SysUserRole struct {
	ID     int64 `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	UserID int64 `gorm:"column:user_id;type:bigint;not null"`       // 用户id
	RoleID int64 `gorm:"column:role_id;type:bigint;not null"`       // 角色id
}

// SysVisLog 系统访问日志表
type SysVisLog struct {
	ID       int64     `gorm:"primaryKey;column:id;type:bigint;not null"` // 主键
	Name     string    `gorm:"column:name;type:varchar(50)"`              // 名称
	Success  string    `gorm:"column:success;type:char(1)"`               // 是否执行成功（Y-是，N-否）
	Message  string    `gorm:"column:message;type:text"`                  // 具体消息
	IP       string    `gorm:"column:ip;type:varchar(255)"`               // ip
	Location string    `gorm:"column:location;type:varchar(255)"`         // 地址
	Browser  string    `gorm:"column:browser;type:varchar(255)"`          // 浏览器
	Os       string    `gorm:"column:os;type:varchar(255)"`               // 操作系统
	VisType  int8      `gorm:"column:vis_type;type:tinyint;not null"`     // 操作类型（字典 1登入 2登出）
	VisTime  time.Time `gorm:"column:vis_time;type:datetime"`             // 访问时间
	Account  string    `gorm:"column:account;type:varchar(50)"`           // 访问账号
}

// User [...]
type User struct {
	ID           int64     `gorm:"primaryKey;column:id;type:bigint;not null"`    // 主键
	Name         string    `gorm:"column:name;type:varchar(255)"`                // 用户名
	UserID       int64     `gorm:"column:user_id;type:bigint;not null"`          // 用户ID
	WxOpenID     string    `gorm:"column:wx_open_id;type:varchar(255);not null"` // 微信openID
	Progress     int8      `gorm:"column:progress;type:tinyint"`                 // 完成简历进度
	JobLevel     string    `gorm:"column:job_level;type:varchar(255)"`           // 岗位级别
	TotalAccount int       `gorm:"column:total_account;type:int"`                // 钱包余额
	Coin         int       `gorm:"column:coin;type:int"`                         // 面币余额
	RecLevel     int8      `gorm:"column:rec_level;type:tinyint"`                // 推荐事级别
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	CreateUser   int64     `gorm:"column:create_user;type:bigint"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime"`
	UpdateUser   int64     `gorm:"column:update_user;type:bigint"`
}

// Wallet 钱包表
type Wallet struct {
	ID         int       `gorm:"primaryKey;column:id;type:int;not null"`
	UserID     int       `gorm:"column:user_id;type:int"`
	Wallet     float64   `gorm:"column:wallet;type:double(11,0) unsigned zerofill"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime"`
}
