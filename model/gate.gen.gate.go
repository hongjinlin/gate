package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GateMgr struct {
	*_BaseMgr
}

// GateMgr open func
func GateMgr(db *gorm.DB) *_GateMgr {
	if db == nil {
		panic(fmt.Errorf("GateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gate"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GateMgr) GetTableName() string {
	return "gate"
}

// Reset 重置gorm会话
func (obj *_GateMgr) Reset() *_GateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GateMgr) Get() (result Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GateMgr) Gets() (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GateMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNation nation获取 国家/地区
func (obj *_GateMgr) WithNation(nation string) Option {
	return optionFunc(func(o *options) { o.query["nation"] = nation })
}

// WithIP ip获取
func (obj *_GateMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithRuntime runtime获取 运行时间
func (obj *_GateMgr) WithRuntime(runtime string) Option {
	return optionFunc(func(o *options) { o.query["runtime"] = runtime })
}

// WithCreateTime create_time获取
func (obj *_GateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_GateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_GateMgr) GetByOption(opts ...Option) (result Gate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GateMgr) GetByOptions(opts ...Option) (results []*Gate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Gate{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_GateMgr) GetFromID(id int64) (result Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GateMgr) GetBatchFromID(ids []int64) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNation 通过nation获取内容 国家/地区
func (obj *_GateMgr) GetFromNation(nation string) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`nation` = ?", nation).Find(&results).Error

	return
}

// GetBatchFromNation 批量查找 国家/地区
func (obj *_GateMgr) GetBatchFromNation(nations []string) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`nation` IN (?)", nations).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容
func (obj *_GateMgr) GetFromIP(ip string) (result Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`ip` = ?", ip).First(&result).Error

	return
}

// GetBatchFromIP 批量查找
func (obj *_GateMgr) GetBatchFromIP(ips []string) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromRuntime 通过runtime获取内容 运行时间
func (obj *_GateMgr) GetFromRuntime(runtime string) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`runtime` = ?", runtime).Find(&results).Error

	return
}

// GetBatchFromRuntime 批量查找 运行时间
func (obj *_GateMgr) GetBatchFromRuntime(runtimes []string) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`runtime` IN (?)", runtimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_GateMgr) GetFromCreateTime(createTime time.Time) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_GateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_GateMgr) GetFromUpdateTime(updateTime time.Time) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_GateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GateMgr) FetchByPrimaryKey(id int64) (result Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIPUnique primary or index 获取唯一内容
func (obj *_GateMgr) FetchUniqueByIPUnique(ip string) (result Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`ip` = ?", ip).First(&result).Error

	return
}

// FetchIndexByNationIndex  获取多个内容
func (obj *_GateMgr) FetchIndexByNationIndex(nation string) (results []*Gate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gate{}).Where("`nation` = ?", nation).Find(&results).Error

	return
}
