package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysMenuMgr struct {
	*_BaseMgr
}

// SysMenuMgr open func
func SysMenuMgr(db *gorm.DB) *_SysMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SysMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysMenuMgr) GetTableName() string {
	return "sys_menu"
}

// Get 获取
func (obj *_SysMenuMgr) Get() (result SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysMenuMgr) Gets() (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysMenuMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 父id
func (obj *_SysMenuMgr) WithPid(pid int64) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithPids pids获取 父ids
func (obj *_SysMenuMgr) WithPids(pids string) Option {
	return optionFunc(func(o *options) { o.query["pids"] = pids })
}

// WithName name获取 名称
func (obj *_SysMenuMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 编码
func (obj *_SysMenuMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithType type获取 菜单类型（字典 0目录 1菜单 2按钮）
func (obj *_SysMenuMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIcon icon获取 图标
func (obj *_SysMenuMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithRouter router获取 路由地址
func (obj *_SysMenuMgr) WithRouter(router string) Option {
	return optionFunc(func(o *options) { o.query["router"] = router })
}

// WithComponent component获取 组件地址
func (obj *_SysMenuMgr) WithComponent(component string) Option {
	return optionFunc(func(o *options) { o.query["component"] = component })
}

// WithPermission permission获取 权限标识
func (obj *_SysMenuMgr) WithPermission(permission string) Option {
	return optionFunc(func(o *options) { o.query["permission"] = permission })
}

// WithApplication application获取 应用分类（应用编码）
func (obj *_SysMenuMgr) WithApplication(application string) Option {
	return optionFunc(func(o *options) { o.query["application"] = application })
}

// WithOpenType open_type获取 打开方式（字典 0无 1组件 2内链 3外链）
func (obj *_SysMenuMgr) WithOpenType(openType int8) Option {
	return optionFunc(func(o *options) { o.query["open_type"] = openType })
}

// WithVisible visible获取 是否可见（Y-是，N-否）
func (obj *_SysMenuMgr) WithVisible(visible string) Option {
	return optionFunc(func(o *options) { o.query["visible"] = visible })
}

// WithLink link获取 链接地址
func (obj *_SysMenuMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// WithRedirect redirect获取 重定向地址
func (obj *_SysMenuMgr) WithRedirect(redirect string) Option {
	return optionFunc(func(o *options) { o.query["redirect"] = redirect })
}

// WithWeight weight获取 权重（字典 1系统权重 2业务权重）
func (obj *_SysMenuMgr) WithWeight(weight int8) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithSort sort获取 排序
func (obj *_SysMenuMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SysMenuMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 状态（字典 0正常 1停用 2删除）
func (obj *_SysMenuMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysMenuMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysMenuMgr) WithCreateUser(createUser int64) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SysMenuMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 修改人
func (obj *_SysMenuMgr) WithUpdateUser(updateUser int64) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// GetByOption 功能选项模式获取
func (obj *_SysMenuMgr) GetByOption(opts ...Option) (result SysMenu, err error) {
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
func (obj *_SysMenuMgr) GetByOptions(opts ...Option) (results []*SysMenu, err error) {
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
func (obj *_SysMenuMgr) GetFromID(id int64) (result SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_SysMenuMgr) GetBatchFromID(ids []int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 父id
func (obj *_SysMenuMgr) GetFromPid(pid int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pid = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量唯一主键查找 父id
func (obj *_SysMenuMgr) GetBatchFromPid(pids []int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pid IN (?)", pids).Find(&results).Error

	return
}

// GetFromPids 通过pids获取内容 父ids
func (obj *_SysMenuMgr) GetFromPids(pids string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pids = ?", pids).Find(&results).Error

	return
}

// GetBatchFromPids 批量唯一主键查找 父ids
func (obj *_SysMenuMgr) GetBatchFromPids(pidss []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pids IN (?)", pidss).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_SysMenuMgr) GetFromName(name string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_SysMenuMgr) GetBatchFromName(names []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 编码
func (obj *_SysMenuMgr) GetFromCode(code string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量唯一主键查找 编码
func (obj *_SysMenuMgr) GetBatchFromCode(codes []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 菜单类型（字典 0目录 1菜单 2按钮）
func (obj *_SysMenuMgr) GetFromType(_type int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 菜单类型（字典 0目录 1菜单 2按钮）
func (obj *_SysMenuMgr) GetBatchFromType(_types []int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标
func (obj *_SysMenuMgr) GetFromIcon(icon string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标
func (obj *_SysMenuMgr) GetBatchFromIcon(icons []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromRouter 通过router获取内容 路由地址
func (obj *_SysMenuMgr) GetFromRouter(router string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("router = ?", router).Find(&results).Error

	return
}

// GetBatchFromRouter 批量唯一主键查找 路由地址
func (obj *_SysMenuMgr) GetBatchFromRouter(routers []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("router IN (?)", routers).Find(&results).Error

	return
}

// GetFromComponent 通过component获取内容 组件地址
func (obj *_SysMenuMgr) GetFromComponent(component string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("component = ?", component).Find(&results).Error

	return
}

// GetBatchFromComponent 批量唯一主键查找 组件地址
func (obj *_SysMenuMgr) GetBatchFromComponent(components []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("component IN (?)", components).Find(&results).Error

	return
}

// GetFromPermission 通过permission获取内容 权限标识
func (obj *_SysMenuMgr) GetFromPermission(permission string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("permission = ?", permission).Find(&results).Error

	return
}

// GetBatchFromPermission 批量唯一主键查找 权限标识
func (obj *_SysMenuMgr) GetBatchFromPermission(permissions []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("permission IN (?)", permissions).Find(&results).Error

	return
}

// GetFromApplication 通过application获取内容 应用分类（应用编码）
func (obj *_SysMenuMgr) GetFromApplication(application string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("application = ?", application).Find(&results).Error

	return
}

// GetBatchFromApplication 批量唯一主键查找 应用分类（应用编码）
func (obj *_SysMenuMgr) GetBatchFromApplication(applications []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("application IN (?)", applications).Find(&results).Error

	return
}

// GetFromOpenType 通过open_type获取内容 打开方式（字典 0无 1组件 2内链 3外链）
func (obj *_SysMenuMgr) GetFromOpenType(openType int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("open_type = ?", openType).Find(&results).Error

	return
}

// GetBatchFromOpenType 批量唯一主键查找 打开方式（字典 0无 1组件 2内链 3外链）
func (obj *_SysMenuMgr) GetBatchFromOpenType(openTypes []int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("open_type IN (?)", openTypes).Find(&results).Error

	return
}

// GetFromVisible 通过visible获取内容 是否可见（Y-是，N-否）
func (obj *_SysMenuMgr) GetFromVisible(visible string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visible = ?", visible).Find(&results).Error

	return
}

// GetBatchFromVisible 批量唯一主键查找 是否可见（Y-是，N-否）
func (obj *_SysMenuMgr) GetBatchFromVisible(visibles []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visible IN (?)", visibles).Find(&results).Error

	return
}

// GetFromLink 通过link获取内容 链接地址
func (obj *_SysMenuMgr) GetFromLink(link string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

// GetBatchFromLink 批量唯一主键查找 链接地址
func (obj *_SysMenuMgr) GetBatchFromLink(links []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

// GetFromRedirect 通过redirect获取内容 重定向地址
func (obj *_SysMenuMgr) GetFromRedirect(redirect string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect = ?", redirect).Find(&results).Error

	return
}

// GetBatchFromRedirect 批量唯一主键查找 重定向地址
func (obj *_SysMenuMgr) GetBatchFromRedirect(redirects []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect IN (?)", redirects).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 权重（字典 1系统权重 2业务权重）
func (obj *_SysMenuMgr) GetFromWeight(weight int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找 权重（字典 1系统权重 2业务权重）
func (obj *_SysMenuMgr) GetBatchFromWeight(weights []int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SysMenuMgr) GetFromSort(sort int) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序
func (obj *_SysMenuMgr) GetBatchFromSort(sorts []int) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysMenuMgr) GetFromRemark(remark string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysMenuMgr) GetBatchFromRemark(remarks []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（字典 0正常 1停用 2删除）
func (obj *_SysMenuMgr) GetFromStatus(status int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（字典 0正常 1停用 2删除）
func (obj *_SysMenuMgr) GetBatchFromStatus(statuss []int8) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysMenuMgr) GetFromCreateTime(createTime time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysMenuMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysMenuMgr) GetFromCreateUser(createUser int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysMenuMgr) GetBatchFromCreateUser(createUsers []int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_SysMenuMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_SysMenuMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改人
func (obj *_SysMenuMgr) GetFromUpdateUser(updateUser int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_SysMenuMgr) GetBatchFromUpdateUser(updateUsers []int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysMenuMgr) FetchByPrimaryKey(id int64) (result SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
