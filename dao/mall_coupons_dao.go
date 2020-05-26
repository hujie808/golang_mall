package dao

import (
	"github.com/go-xorm/xorm"

	"log"
	"web_iris/golang_mall/models"
)

type MallCouponsDao struct {
	engine *xorm.Engine
}

func NewMallCouponsDao(engine *xorm.Engine) *MallCouponsDao { //实例化引擎
	return &MallCouponsDao{
		engine: engine,
	}
}

//get_id
func (d *MallCouponsDao) Get(id int) *models.MallCoupons {
	data := &models.MallCoupons{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallCouponsDao) GetAll(page, size int) []models.MallCoupons {
	offset := (page - 1) * size
	datalist := make([]models.MallCoupons, 0)

	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_coupons.mall_commodity_id").
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCouponsDao) GetTest(page, size int) []models.JionCoupons {
	offset := (page - 1) * size
	datalist := make([]models.JionCoupons, 0)
	err := d.engine.
		Desc("mall_coupons.id").
		Limit(size, offset).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_coupons.mall_commodity_id").
		Select("mall_coupons.*,mall_commodity.title as c_title").
		Where("mall_coupons.sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		log.Println(datalist)
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCouponsDao) GetStart(page, size, nowTime int) []models.JionCoupons {
	offset := (page - 1) * size
	datalist := make([]models.JionCoupons, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_coupons.mall_commodity_id").
		Select("mall_coupons.*,mall_commodity.title as c_title").
		Where("mall_coupons.start_time>?", nowTime).
		Where("mall_coupons.sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCouponsDao) GetEnd(page, size, nowTime int) []models.JionCoupons {
	offset := (page - 1) * size
	datalist := make([]models.JionCoupons, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Where("mall_coupons.end_time<=?", nowTime).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_coupons.mall_commodity_id").
		Select("mall_coupons.*,mall_commodity.title as c_title").
		Where("mall_coupons.sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCouponsDao) GetOnline(page, size, nowTime int) []models.JionCoupons {
	offset := (page - 1) * size
	datalist := make([]models.JionCoupons, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Where("mall_coupons.end_time>=?", nowTime).
		And("mall_coupons.start_time<=?", nowTime).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_coupons.mall_commodity_id").
		Select("mall_coupons.*,mall_commodity.title as c_title").
		Where("mall_coupons.sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCouponsDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallCoupons{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallCouponsDao) Delete(id int) error {
	data := &models.MallCoupons{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallCouponsDao) Update(data *models.MallCoupons, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallCouponsDao) Create(data *models.MallCoupons) error {
	_, err := d.engine.Insert(data)
	return err
}
