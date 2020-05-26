package dao

import (
	"github.com/go-xorm/xorm"
	"web_iris/golang_mall/models"
)

type RedLogDao struct {
	engine *xorm.Engine
}

func NewRedLogDao(engine *xorm.Engine) *RedLogDao { //实例化引擎
	return &RedLogDao{
		engine: engine,
	}
}

//get_id
func (d *RedLogDao) Get(id int) *models.RedLog {
	data := &models.RedLog{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *RedLogDao) GetAll(page, size int) []models.RedLog {
	offset := (page - 1) * size
	datalist := make([]models.RedLog, 0)
	err := d.engine.
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *RedLogDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.RedLog{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *RedLogDao) Delete(id int) error {
	data := &models.RedLog{Id: id}
	_, err := d.engine.Id(id).Delete(data)
	return err
}

func (d *RedLogDao) Update(data *models.RedLog, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *RedLogDao) Create(data *models.RedLog) error {
	_, err := d.engine.Insert(data)
	return err
}
