package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallLevelService interface {
	Get(id int) *models.MallLevel
	GetAll(page, size int) []models.MallLevel
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallLevel, columns []string) error
	Create(data *models.MallLevel) error
}

type mallLevelService struct {
	dao *dao.MallLevelDao
}

func NewMallLevelService() MallLevelService {
	return &mallLevelService{
		dao: dao.NewMallLevelDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallLevelService) Get(id int) *models.MallLevel {
	return s.dao.Get(id)
}

func (s *mallLevelService) GetAll(page, size int) []models.MallLevel {
	return s.dao.GetAll(page, size)
}
func (s *mallLevelService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallLevelService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallLevelService) Update(data *models.MallLevel, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallLevelService) Create(data *models.MallLevel) error {
	return s.dao.Create(data)
}
