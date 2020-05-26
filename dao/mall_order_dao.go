package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallOrderDao struct {
	engine *xorm.Engine
}

func NewMallOrderDao(engine *xorm.Engine) *MallOrderDao { //实例化引擎
	return &MallOrderDao{
		engine: engine,
	}
}

//get_id
func (d *MallOrderDao) Get(id int) *models.MallOrder {
	data := &models.MallOrder{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallOrderDao) GetAll(page, size int) []models.MallOrder {
	offset := (page - 1) * size
	datalist := make([]models.MallOrder, 0)
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

func (d *MallOrderDao) SearchPageIs(page, size int, pay_status, tel string) ([]models.JionOrder, int) {
	offset := (page - 1) * size
	datalist := make([]models.JionOrder, 0)
	err := d.engine.
		Desc("mall_order.click_time").
		Where("mall_order.pay_status Like ?", "%"+pay_status+"%").
		And("mall_order.order_tel Like ?", "%"+tel+"%").
		Join("INNER", "mall_wechat", "mall_wechat.id=mall_order.mall_wechat_id").
		//Where("mall_order.sys_status=?", 0). // 有效的用户
		Limit(size, offset).
		Find(&datalist)

	total, t_err := d.engine.Where("mall_order.pay_status Like ?", "%"+pay_status+"%").
		And("mall_order.order_tel Like ?", "%"+tel+"%").
		Join("INNER", "mall_wechat", "mall_wechat.id=mall_order.mall_wechat_id").Count(&models.JionOrder{})
	if t_err != nil {
		return datalist, 0
	}
	if err != nil {
		return datalist, int(total)
	} else {
		return datalist, int(total)
	}
}

func (d *MallOrderDao) SearchUserOrder(isType, userId int) []models.MallOrder {
	datalist := make([]models.MallOrder, 0)
	var err error
	if isType == 0 {
		err = d.engine.
			Desc("id").
			Where("mall_wechat_id=?", userId).
			Find(&datalist)
	} else {
		err = d.engine.
			Desc("id").
			Where("mall_wechat_id=?", userId).
			And("pay_status=?", isType).
			Find(&datalist)
	}

	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallOrderDao) SearchUserOrderId(isType, userId, id int) []models.MallOrder {
	datalist := make([]models.MallOrder, 0)
	var err error
	if isType == 0 {
		err = d.engine.
			Desc("id").
			Where("mall_wechat_id=?", userId).
			And("id=?", id).
			Find(&datalist)
	} else {
		err = d.engine.
			Desc("id").
			Where("mall_wechat_id=?", userId).
			And("id=?", id).
			And("pay_status=?", isType).
			Find(&datalist)
	}

	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallOrderDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallOrder{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallOrderDao) Delete(id int) error {
	data := &models.MallOrder{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallOrderDao) ReallyDelete(id, wecahtId int) error {
	data := &models.MallOrder{Id: id, MallWechatId: wecahtId}
	_, err := d.engine.Id(id).Delete(data)
	return err
}

func (d *MallOrderDao) CancelDelete(id int) error { //取消订单
	data := &models.MallOrder{Id: id, SysStatus: 1, PayStatus: "6"}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallOrderDao) Update(data *models.MallOrder, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallOrderDao) Create(data *models.MallOrder) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *MallOrderDao) CreateId(data *models.MallOrder) (int, error) {
	_, err := d.engine.Insert(data)
	return data.Id, err
}

func (d *MallOrderDao) GetOutTradeNo(out_trade_no string) *models.MallOrder {
	data := &models.MallOrder{OrderNumber: out_trade_no}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}
