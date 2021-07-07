package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysFileInfoMgr struct {
	*_BaseMgr
}

// SysFileInfoMgr open func
func SysFileInfoMgr(db *gorm.DB) *_SysFileInfoMgr {
	if db == nil {
		panic(fmt.Errorf("SysFileInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysFileInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_file_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysFileInfoMgr) GetTableName() string {
	return "sys_file_info"
}

// Get 获取
func (obj *_SysFileInfoMgr) Get() (result SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysFileInfoMgr) Gets() (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_SysFileInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFileLocation file_location获取 文件存储位置（1:阿里云，2:腾讯云，3:minio，4:本地）
func (obj *_SysFileInfoMgr) WithFileLocation(fileLocation int8) Option {
	return optionFunc(func(o *options) { o.query["file_location"] = fileLocation })
}

// WithFileBucket file_bucket获取 文件仓库
func (obj *_SysFileInfoMgr) WithFileBucket(fileBucket string) Option {
	return optionFunc(func(o *options) { o.query["file_bucket"] = fileBucket })
}

// WithFileOriginName file_origin_name获取 文件名称（上传时候的文件名）
func (obj *_SysFileInfoMgr) WithFileOriginName(fileOriginName string) Option {
	return optionFunc(func(o *options) { o.query["file_origin_name"] = fileOriginName })
}

// WithFileSuffix file_suffix获取 文件后缀
func (obj *_SysFileInfoMgr) WithFileSuffix(fileSuffix string) Option {
	return optionFunc(func(o *options) { o.query["file_suffix"] = fileSuffix })
}

// WithFileSizeKb file_size_kb获取 文件大小kb
func (obj *_SysFileInfoMgr) WithFileSizeKb(fileSizeKb int64) Option {
	return optionFunc(func(o *options) { o.query["file_size_kb"] = fileSizeKb })
}

// WithFileSizeInfo file_size_info获取 文件大小信息，计算后的
func (obj *_SysFileInfoMgr) WithFileSizeInfo(fileSizeInfo string) Option {
	return optionFunc(func(o *options) { o.query["file_size_info"] = fileSizeInfo })
}

// WithFileObjectName file_object_name获取 存储到bucket的名称（文件唯一标识id）
func (obj *_SysFileInfoMgr) WithFileObjectName(fileObjectName string) Option {
	return optionFunc(func(o *options) { o.query["file_object_name"] = fileObjectName })
}

// WithFilePath file_path获取 存储路径
func (obj *_SysFileInfoMgr) WithFilePath(filePath string) Option {
	return optionFunc(func(o *options) { o.query["file_path"] = filePath })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysFileInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_SysFileInfoMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SysFileInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 修改用户
func (obj *_SysFileInfoMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysFileInfoMgr) GetByOption(opts ...Option) (result SysFileInfo, err error) {
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
func (obj *_SysFileInfoMgr) GetByOptions(opts ...Option) (results []*SysFileInfo, err error) {
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
func (obj *_SysFileInfoMgr) GetFromID(id int64) (result SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键id
func (obj *_SysFileInfoMgr) GetBatchFromID(ids []int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromFileLocation 通过file_location获取内容 文件存储位置（1:阿里云，2:腾讯云，3:minio，4:本地）
func (obj *_SysFileInfoMgr) GetFromFileLocation(fileLocation int8) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_location = ?", fileLocation).Find(&results).Error

	return
}

// GetBatchFromFileLocation 批量唯一主键查找 文件存储位置（1:阿里云，2:腾讯云，3:minio，4:本地）
func (obj *_SysFileInfoMgr) GetBatchFromFileLocation(fileLocations []int8) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_location IN (?)", fileLocations).Find(&results).Error

	return
}

// GetFromFileBucket 通过file_bucket获取内容 文件仓库
func (obj *_SysFileInfoMgr) GetFromFileBucket(fileBucket string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_bucket = ?", fileBucket).Find(&results).Error

	return
}

// GetBatchFromFileBucket 批量唯一主键查找 文件仓库
func (obj *_SysFileInfoMgr) GetBatchFromFileBucket(fileBuckets []string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_bucket IN (?)", fileBuckets).Find(&results).Error

	return
}

// GetFromFileOriginName 通过file_origin_name获取内容 文件名称（上传时候的文件名）
func (obj *_SysFileInfoMgr) GetFromFileOriginName(fileOriginName string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_origin_name = ?", fileOriginName).Find(&results).Error

	return
}

// GetBatchFromFileOriginName 批量唯一主键查找 文件名称（上传时候的文件名）
func (obj *_SysFileInfoMgr) GetBatchFromFileOriginName(fileOriginNames []string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_origin_name IN (?)", fileOriginNames).Find(&results).Error

	return
}

// GetFromFileSuffix 通过file_suffix获取内容 文件后缀
func (obj *_SysFileInfoMgr) GetFromFileSuffix(fileSuffix string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_suffix = ?", fileSuffix).Find(&results).Error

	return
}

// GetBatchFromFileSuffix 批量唯一主键查找 文件后缀
func (obj *_SysFileInfoMgr) GetBatchFromFileSuffix(fileSuffixs []string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_suffix IN (?)", fileSuffixs).Find(&results).Error

	return
}

// GetFromFileSizeKb 通过file_size_kb获取内容 文件大小kb
func (obj *_SysFileInfoMgr) GetFromFileSizeKb(fileSizeKb int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size_kb = ?", fileSizeKb).Find(&results).Error

	return
}

// GetBatchFromFileSizeKb 批量唯一主键查找 文件大小kb
func (obj *_SysFileInfoMgr) GetBatchFromFileSizeKb(fileSizeKbs []int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size_kb IN (?)", fileSizeKbs).Find(&results).Error

	return
}

// GetFromFileSizeInfo 通过file_size_info获取内容 文件大小信息，计算后的
func (obj *_SysFileInfoMgr) GetFromFileSizeInfo(fileSizeInfo string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size_info = ?", fileSizeInfo).Find(&results).Error

	return
}

// GetBatchFromFileSizeInfo 批量唯一主键查找 文件大小信息，计算后的
func (obj *_SysFileInfoMgr) GetBatchFromFileSizeInfo(fileSizeInfos []string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size_info IN (?)", fileSizeInfos).Find(&results).Error

	return
}

// GetFromFileObjectName 通过file_object_name获取内容 存储到bucket的名称（文件唯一标识id）
func (obj *_SysFileInfoMgr) GetFromFileObjectName(fileObjectName string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_object_name = ?", fileObjectName).Find(&results).Error

	return
}

// GetBatchFromFileObjectName 批量唯一主键查找 存储到bucket的名称（文件唯一标识id）
func (obj *_SysFileInfoMgr) GetBatchFromFileObjectName(fileObjectNames []string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_object_name IN (?)", fileObjectNames).Find(&results).Error

	return
}

// GetFromFilePath 通过file_path获取内容 存储路径
func (obj *_SysFileInfoMgr) GetFromFilePath(filePath string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_path = ?", filePath).Find(&results).Error

	return
}

// GetBatchFromFilePath 批量唯一主键查找 存储路径
func (obj *_SysFileInfoMgr) GetBatchFromFilePath(filePaths []string) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_path IN (?)", filePaths).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysFileInfoMgr) GetFromCreateTime(createTime time.Time) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysFileInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_SysFileInfoMgr) GetFromCreateUser(createUser int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建用户
func (obj *_SysFileInfoMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_SysFileInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_SysFileInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改用户
func (obj *_SysFileInfoMgr) GetFromUpdateUser(updateUser int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改用户
func (obj *_SysFileInfoMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysFileInfoMgr) FetchByPrimaryKey(id int64) (result SysFileInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
