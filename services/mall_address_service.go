package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallAddressService interface {
	Get(id int) *models.MallAddress
	GetAll(page, size int) []models.MallAddress
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallAddress, columns []string) error
	Create(data *models.MallAddress) error
	GetWechatId(id int) []models.MallAddress
	GetUserDefault(userId int) *models.MallAddress
}

type mallAddressService struct {
	dao *dao.MallAddressDao
}

func NewMallAddressService() MallAddressService {
	return &mallAddressService{
		dao: dao.NewMallAddressDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallAddressService) Get(id int) *models.MallAddress {
	return s.dao.Get(id)
}

func (s *mallAddressService) GetAll(page, size int) []models.MallAddress {
	return s.dao.GetAll(page, size)
}
func (s *mallAddressService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallAddressService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallAddressService) Update(data *models.MallAddress, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallAddressService) Create(data *models.MallAddress) error {
	return s.dao.Create(data)
}

func (s *mallAddressService) GetWechatId(id int) []models.MallAddress {
	return s.dao.GetWechatId(id)
}

func (s *mallAddressService) GetUserDefault(userId int) *models.MallAddress {
	return s.dao.GetUserDefault(userId)
}
