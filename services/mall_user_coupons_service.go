package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallUserCouponsService interface {
	Get(id int) *models.MallUserCoupons
	GetAll(page, size int) []models.MallUserCoupons
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallUserCoupons, columns []string) error
	Create(data *models.MallUserCoupons) error
	GetNoEmployCoupons(wechat_id,nowTime int) []models.MallUserCoupons     //未使用优惠券
	GetUserCoupons(wechatId, nowTime int) []models.JionUserCoupons //join 同时拿取coupon的信息
	GetCouponsId(userID, couponsID int) *models.MallUserCoupons    //指定用户 指定优惠券查询
	GetWechatId(id int) []models.MallUserCoupons                   //用户所有优惠券
}

type mallUserCouponsService struct {
	dao *dao.MallUserCouponsDao
}

func NewMallUserCouponsService() MallUserCouponsService {
	return &mallUserCouponsService{
		dao: dao.NewMallUserCouponsDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallUserCouponsService) Get(id int) *models.MallUserCoupons {
	return s.dao.Get(id)
}

func (s *mallUserCouponsService) GetAll(page, size int) []models.MallUserCoupons {
	return s.dao.GetAll(page, size)
}
func (s *mallUserCouponsService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallUserCouponsService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallUserCouponsService) Update(data *models.MallUserCoupons, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallUserCouponsService) Create(data *models.MallUserCoupons) error {
	return s.dao.Create(data)
}
func (s *mallUserCouponsService) GetNoEmployCoupons(wechat_id,nowTime int) []models.MallUserCoupons {
	return s.dao.GetNoEmployCoupons(wechat_id,nowTime)
}
func (s *mallUserCouponsService) GetUserCoupons(wechatId, nowTime int) []models.JionUserCoupons {
	return s.dao.GetUserCoupons(wechatId, nowTime)
}

//指定微信 指定优惠券Id
func (s *mallUserCouponsService) GetCouponsId(userID, couponsID int) *models.MallUserCoupons {
	return s.dao.GetCouponsId(userID, couponsID)
}

func (s *mallUserCouponsService) GetWechatId(id int) []models.MallUserCoupons {
	return s.dao.GetWechatId(id)
}
