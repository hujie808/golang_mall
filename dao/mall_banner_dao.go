package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallBannerDao struct {
	engine *xorm.Engine
}

func NewMallBannerDao(engine *xorm.Engine) *MallBannerDao { //实例化引擎
	return &MallBannerDao{
		engine: engine,
	}
}

//get_id
func (d *MallBannerDao) Get(id int) *models.MallBanner {
	data := &models.MallBanner{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallBannerDao) GetAll(page, size int) []models.MallBanner {
	offset := (page - 1) * size
	datalist := make([]models.MallBanner, 0)
	err := d.engine.Desc("desc").
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

func (d *MallBannerDao) GetTitle(page, size, title int) []models.MallBanner {
	offset := (page - 1) * size
	datalist := make([]models.MallBanner, 0)
	err := d.engine.Desc("desc").
		Desc("id").
		Limit(size, offset).
		Where("sys_status=?", 0). // 有效的用户
		And("title=?", title).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallBannerDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallBanner{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallBannerDao) Delete(id int) error {
	data := &models.MallBanner{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallBannerDao) Update(data *models.MallBanner, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallBannerDao) Create(data *models.MallBanner) error {
	_, err := d.engine.Insert(data)
	return err
}
