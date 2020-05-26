package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallCouponsService interface {
	Get(id int) *models.MallCoupons
	GetAll(page, size int) []models.MallCoupons
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallCoupons, columns []string) error
	Create(data *models.MallCoupons) error
	GetStart(page, size, nowTime int) []models.JionCoupons
	GetEnd(page, size, nowTime int) []models.JionCoupons
	GetOnline(page, size, nowTime int) []models.JionCoupons
	GetTest(page, size int) []models.JionCoupons
}

type mallCouponsService struct {
	dao *dao.MallCouponsDao
}

func NewMallCouponsService() MallCouponsService {
	return &mallCouponsService{
		dao: dao.NewMallCouponsDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallCouponsService) Get(id int) *models.MallCoupons {
	return s.dao.Get(id)
}

func (s *mallCouponsService) GetAll(page, size int) []models.MallCoupons {
	return s.dao.GetAll(page, size)
}
func (s *mallCouponsService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallCouponsService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallCouponsService) Update(data *models.MallCoupons, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallCouponsService) Create(data *models.MallCoupons) error {
	return s.dao.Create(data)
}

func (s *mallCouponsService) GetStart(page, size, nowTime int) []models.JionCoupons {
	return s.dao.GetStart(page, size, nowTime)
}
func (s *mallCouponsService) GetEnd(page, size, nowTime int) []models.JionCoupons {
	return s.dao.GetEnd(page, size, nowTime)
}
func (s *mallCouponsService) GetOnline(page, size, nowTime int) []models.JionCoupons {
	return s.dao.GetOnline(page, size, nowTime)
}

func (s *mallCouponsService) GetTest(page, size int) []models.JionCoupons {
	return s.dao.GetTest(page, size)
}
