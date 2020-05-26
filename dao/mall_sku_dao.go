package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallSkuDao struct {
	engine *xorm.Engine
}

func NewMallSkuDao(engine *xorm.Engine) *MallSkuDao { //实例化引擎
	return &MallSkuDao{
		engine: engine,
	}
}

//get_id
func (d *MallSkuDao) Get(id int) *models.MallSku {
	data := &models.MallSku{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallSkuDao) GetTitle(id int) *models.JionSku {
	data := &models.JionSku{
		MallSku: models.MallSku{Id: id},
	}
	ok, err := d.engine.
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_sku.mall_commodity_id").
		Select("mall_sku.*,mall_commodity.title as c_title,mall_commodity.image as c_image,mall_commodity.com_type").
		Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallSkuDao) GetAll(page, size int) []models.MallSku {
	offset := (page - 1) * size
	datalist := make([]models.MallSku, 0)
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

func (d *MallSkuDao) GetIdAll(idList []int) []models.MallSku {
	datalist := make([]models.MallSku, 0)
	err := d.engine.
		In("id", idList).
		And("sys_status=?", 0).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallSkuDao) GetIdAllCommodityTitle(idList []int) []models.JionSku {
	datalist := make([]models.JionSku, 0)
	err := d.engine.
		In("mall_sku.id", idList).
		And("mall_sku.sys_status=?", 0).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_sku.mall_commodity_id").
		Select("mall_sku.*,mall_commodity.title as c_title").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallSkuDao) GetCommodityId(commodityId int) []models.MallSku {
	datalist := make([]models.MallSku, 0)
	err := d.engine.
		Desc("id").
		Where("mall_commodity_id=?", commodityId). // 有效的用户
		Where("sys_status=?", 0).                  // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallSkuDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallSku{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallSkuDao) Delete(id int) error {
	data := &models.MallSku{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallSkuDao) Update(data *models.MallSku, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallSkuDao) Create(data *models.MallSku) error {
	_, err := d.engine.Insert(data)
	return err
}
