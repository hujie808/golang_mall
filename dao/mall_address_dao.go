package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallAddressDao struct {
	engine *xorm.Engine
}

func NewMallAddressDao(engine *xorm.Engine) *MallAddressDao { //实例化引擎
	return &MallAddressDao{
		engine: engine,
	}
}

//get_id
func (d *MallAddressDao) Get(id int) *models.MallAddress {
	data := &models.MallAddress{Id: id, SysStatus: 0}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

//get_id
func (d *MallAddressDao) GetUserDefault(userId int) *models.MallAddress {
	data := &models.MallAddress{MallWechatId: userId, IdDefautl: 1}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallAddressDao) GetWechatId(id int) []models.MallAddress {
	datalist := make([]models.MallAddress, 0)
	err := d.engine.
		Where("mall_wechat_id=?", id). // 有效的用户
		Where("sys_status=?", 0).      // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallAddressDao) GetAll(page, size int) []models.MallAddress {
	offset := (page - 1) * size
	datalist := make([]models.MallAddress, 0)
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
func (d *MallAddressDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallAddress{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallAddressDao) Delete(id int) error {
	data := &models.MallAddress{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallAddressDao) Update(data *models.MallAddress, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallAddressDao) Create(data *models.MallAddress) error {
	_, err := d.engine.Insert(data)
	return err
}
