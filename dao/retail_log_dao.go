package dao

import (
	"github.com/go-xorm/xorm"
	"web_iris/golang_mall/models"
)

type RetailLogDao struct {
	engine *xorm.Engine
}

func NewRetailLogDao(engine *xorm.Engine) *RetailLogDao { //实例化引擎
	return &RetailLogDao{
		engine: engine,
	}
}

//get_id
func (d *RetailLogDao) Get(id int) *models.RetailLog {
	data := &models.RetailLog{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *RetailLogDao) GetAll(page, size int) []models.RetailLog {
	offset := (page - 1) * size
	datalist := make([]models.RetailLog, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *RetailLogDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.RetailLog{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *RetailLogDao) Delete(id int) error {
	data := &models.RetailLog{Id: id}
	_, err := d.engine.Id(id).Delete(data)
	return err
}

func (d *RetailLogDao) Update(data *models.RetailLog, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *RetailLogDao) Create(data *models.RetailLog) error {
	_, err := d.engine.Insert(data)
	return err
}
