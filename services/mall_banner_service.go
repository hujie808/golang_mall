package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallBannerService interface {
	Get(id int) *models.MallBanner
	GetAll(page, size int) []models.MallBanner
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallBanner, columns []string) error
	Create(data *models.MallBanner) error
	GetTitle(page, size, title int) []models.MallBanner
}

type mallBannerService struct {
	dao *dao.MallBannerDao
}

func NewMallBannerService() MallBannerService {
	return &mallBannerService{
		dao: dao.NewMallBannerDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallBannerService) Get(id int) *models.MallBanner {
	return s.dao.Get(id)
}

func (s *mallBannerService) GetAll(page, size int) []models.MallBanner {
	return s.dao.GetAll(page, size)
}
func (s *mallBannerService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallBannerService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallBannerService) Update(data *models.MallBanner, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallBannerService) Create(data *models.MallBanner) error {
	return s.dao.Create(data)
}

func (s *mallBannerService) GetTitle(page, size, title int) []models.MallBanner {
	return s.dao.GetTitle(page, size, title)
}
