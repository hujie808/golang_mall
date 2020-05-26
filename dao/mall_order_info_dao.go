package dao

import (
	"github.com/go-xorm/xorm"

	"log"
	"web_iris/golang_mall/models"
)

type MallOrderInfoDao struct {
	engine *xorm.Engine
}

func NewMallOrderInfoDao(engine *xorm.Engine) *MallOrderInfoDao { //实例化引擎
	return &MallOrderInfoDao{
		engine: engine,
	}
}

//get_id
func (d *MallOrderInfoDao) Get(id int) *models.MallOrderInfo {
	data := &models.MallOrderInfo{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallOrderInfoDao) GetAll(page, size int) []models.MallOrderInfo {
	offset := (page - 1) * size
	datalist := make([]models.MallOrderInfo, 0)
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

func (d *MallOrderInfoDao) GetCommodity(c_id int) []models.MallOrderInfo {
	datalist := make([]models.MallOrderInfo, 0)
	err := d.engine.
		Desc("id").
		Where("sys_status=?", 0). // 有效的用户
		And("commodity_id=?", c_id).
		And("is_show=?", 0).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallOrderInfoDao) GetOrder(orderId int) []models.MallOrderInfo {
	datalist := make([]models.MallOrderInfo, 0)
	err := d.engine.
		Desc("id").
		Where("mall_order_id=?", orderId). // 有效的用户
		Where("sys_status=?", 0).          // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallOrderInfoDao) GetOrderListInfoId(orderId []int) []models.MallOrderInfo {
	datalist := make([]models.MallOrderInfo, 0)
	err := d.engine.
		Desc("id").
		//Where("sys_status=?", 0). // 有效的用户
		In("mall_order_id", orderId). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallOrderInfoDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallOrderInfo{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallOrderInfoDao) Delete(id int) error {
	data := &models.MallOrderInfo{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallOrderInfoDao) ReallyDelete(orderId, wechatId int) error {
	data := &models.MallOrderInfo{MallOrderId: orderId, MallWecahtId: wechatId}
	_, err := d.engine.Delete(data)
	return err
}

func (d *MallOrderInfoDao) Update(data *models.MallOrderInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallOrderInfoDao) Create(data *models.MallOrderInfo) error {
	_, err := d.engine.Insert(data)
	return err
}

//获得用户的全部已购买视频
func (d *MallOrderInfoDao) GetVideoAll(userId int) []models.JoinOrderInfo {
	datalist := make([]models.JoinOrderInfo, 0)
	err := d.engine.
		Where("mall_order_info.is_virtual=?", 1).
		And("mall_wecaht_id=?", userId).
		Join("INNER", "mall_commodity", "mall_commodity.id=mall_order_info.commodity_id").
		Join("LEFT", "mall_order", "mall_order.id=mall_order_id").
		In("mall_order.pay_status", []string{"2", "3", "4", "5"}).
		Select("mall_order_info.*,mall_commodity.video_url,mall_commodity.detail").
		Find(&datalist)
	if err != nil {
		log.Printf("mall_order_info.go GetVideoAll err=%s", err)
		return datalist
	} else {
		return datalist
	}
	return datalist
}
