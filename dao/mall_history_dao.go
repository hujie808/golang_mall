package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallHistoryDao struct {
	engine *xorm.Engine
}

func NewMallHistoryDao(engine *xorm.Engine) *MallHistoryDao { //实例化引擎
	return &MallHistoryDao{
		engine: engine,
	}
}

//get_id
func (d *MallHistoryDao) Get(id int) *models.MallHistory {
	data := &models.MallHistory{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

//get   mall_wechat_id
func (d *MallHistoryDao) GetWechatId(id int) *models.MallHistory {
	data := &models.MallHistory{MallWechatId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallHistoryDao) GetAll(page, size int) []models.MallHistory {
	offset := (page - 1) * size
	datalist := make([]models.MallHistory, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *MallHistoryDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallHistory{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallHistoryDao) Delete(id int) error {
	data := &models.MallHistory{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallHistoryDao) Update(data *models.MallHistory, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallHistoryDao) Create(data *models.MallHistory) error {
	_, err := d.engine.Insert(data)
	return err
}
