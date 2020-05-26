package dao

import (
	"github.com/go-xorm/xorm"
	"log"

	"web_iris/golang_mall/models"
)

type MallUserCouponsDao struct {
	engine *xorm.Engine
}

func NewMallUserCouponsDao(engine *xorm.Engine) *MallUserCouponsDao { //实例化引擎
	return &MallUserCouponsDao{
		engine: engine,
	}
}

//get_id
func (d *MallUserCouponsDao) Get(id int) *models.MallUserCoupons {
	data := &models.MallUserCoupons{Id: id, IsEmploy: 0}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

//拿取微信用户指定的优惠券
func (d *MallUserCouponsDao) GetCouponsId(userID, couponsID int) *models.MallUserCoupons {
	data := &models.MallUserCoupons{MallWechatId: userID, MallCouponsId: couponsID}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallUserCouponsDao) GetUserCoupons(wechatId, nowTime int) []models.JionUserCoupons {
	datalist := make([]models.JionUserCoupons, 0)
	err := d.engine.
		Desc("mall_user_coupons.id").
		Where("mall_coupons.end_time>=?", nowTime).
		And("mall_coupons.start_time<=?", nowTime).
		And("mall_user_coupons.mall_wechat_id=?", wechatId).
		And("mall_coupons.sys_status=?", 0). // 有效优惠券
		And("mall_user_coupons.is_employ=?", 1). // 有效优惠券
		Select("mall_coupons.*,mall_user_coupons.*,mall_coupons.id as mall_coupons_id,mall_user_coupons.id as mall_user_coupons_id").
		Join("INNER", "mall_coupons", "mall_coupons.id=mall_user_coupons.mall_coupons_id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

//微信用户id
func (d *MallUserCouponsDao) GetWechatId(id int) []models.MallUserCoupons {
	datalist := make([]models.MallUserCoupons, 0)
	err := d.engine.
		Where("mall_wechat_id=?", id). // 有效的用户
		//Where("is_employ=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

//用户列表未使用优惠券
func (d *MallUserCouponsDao) GetNoEmployCoupons(wechat_id,nowTime int) []models.MallUserCoupons {
	datalist := make([]models.MallUserCoupons, 0)
	log.Println(nowTime)
	err := d.engine.
		Where("mall_wechat_id=?", wechat_id).
		//And("end_time>=?", nowTime).
		//And("start_time<=?", nowTime).
		And("is_employ=?", 1). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

//已使用优惠券,做判断
func (d *MallUserCouponsDao) GetUserCouponsId(wechat_id, coupons_id int) []models.MallUserCoupons {
	datalist := make([]models.MallUserCoupons, 0)
	err := d.engine.
		Where("mall_coupons_id=?", coupons_id). // 有效的用户
		And("mall_wechat_id=?", wechat_id).
		And("is_employ=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallUserCouponsDao) GetAll(page, size int) []models.MallUserCoupons {
	offset := (page - 1) * size
	datalist := make([]models.MallUserCoupons, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Where("is_employ=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *MallUserCouponsDao) CountAll() int64 {
	num, err := d.engine.
		Where("is_employ=?", 0).
		Count(&models.MallUserCoupons{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallUserCouponsDao) Delete(id int) error {
	data := &models.MallUserCoupons{Id: id, IsEmploy: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallUserCouponsDao) Update(data *models.MallUserCoupons, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallUserCouponsDao) Create(data *models.MallUserCoupons) error {
	_, err := d.engine.Insert(data)
	return err
}
