package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallLevelDao struct {
	engine *xorm.Engine
}

func NewMallLevelDao(engine *xorm.Engine) *MallLevelDao { //实例化引擎
	return &MallLevelDao{
		engine: engine,
	}
}

//get_id
func (d *MallLevelDao) Get(id int) *models.MallLevel {
	data := &models.MallLevel{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallLevelDao) GetAll(page, size int) []models.MallLevel {
	offset := (page - 1) * size
	datalist := make([]models.MallLevel, 0)
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
func (d *MallLevelDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallLevel{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallLevelDao) Delete(id int) error {
	data := &models.MallLevel{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallLevelDao) Update(data *models.MallLevel, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallLevelDao) Create(data *models.MallLevel) error {
	_, err := d.engine.Insert(data)
	return err
}
