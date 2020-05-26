package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallShoppingDao struct {
	engine *xorm.Engine
}

func NewMallShoppingDao(engine *xorm.Engine) *MallShoppingDao { //实例化引擎
	return &MallShoppingDao{
		engine: engine,
	}
}

//get_id
func (d *MallShoppingDao) Get(id int) *models.MallShopping {
	data := &models.MallShopping{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallShoppingDao) GetAll(page, size int) []models.MallShopping {
	offset := (page - 1) * size
	datalist := make([]models.MallShopping, 0)
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

func (d *MallShoppingDao) GetShopping(user_id int) []models.MallShopping {
	datalist := make([]models.MallShopping, 0)
	err := d.engine.
		Desc("id").
		Where("sys_status=?", 0). // 有效的用户
		And("type=?", 1).
		And("mall_wechat_id=?", user_id).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallShoppingDao) GetCollection(user_id int) []models.MallShopping {
	datalist := make([]models.MallShopping, 0)
	err := d.engine.
		Desc("id").
		Where("sys_status=?", 0). // 有效的用户
		And("type=?", 2).
		And("mall_wechat_id=?", user_id).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallShoppingDao) GetIsCollection(user_id, commodity_id int) []models.MallShopping {
	datalist := make([]models.MallShopping, 0)
	err := d.engine.
		Desc("id").
		And("type=?", 2).
		And("mall_wechat_id=?", user_id).
		And("mall_sku_id=?", commodity_id).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallShoppingDao) GetIsShopping(user_id, sku_id int) *models.MallShopping {
	data := &models.MallShopping{MallWechatId: user_id, MallSkuId: sku_id, Type: 1}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallShoppingDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallShopping{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallShoppingDao) Delete(id int) error {
	data := &models.MallShopping{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallShoppingDao) DeleteShopping(id int) error {
	data := &models.MallShopping{Id: id}
	_, err := d.engine.Id(id).Delete(data)
	return err
}

func (d *MallShoppingDao) DeleteShoppingAll(wechatId,skuid int) error {
	data := &models.MallShopping{MallWechatId: wechatId,MallSkuId:skuid, Type: 1}
	_, err := d.engine.Delete(data)
	return err
}

func (d *MallShoppingDao) Update(data *models.MallShopping, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallShoppingDao) Create(data *models.MallShopping) error {
	_, err := d.engine.Insert(data)
	return err
}
