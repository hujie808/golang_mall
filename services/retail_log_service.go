package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type RetailLogService interface {
	Get(id int) *models.RetailLog
	GetAll(page, size int) []models.RetailLog
	CountAll() int64
	Delete(id int) error
	Update(data *models.RetailLog, columns []string) error
	Create(data *models.RetailLog) error
}

type retailLogService struct {
	dao *dao.RetailLogDao
}

func NewRetailLogService() RetailLogService {
	return &retailLogService{
		dao: dao.NewRetailLogDao(datasource.InstanceDbMaster()),
	}
}

func (s *retailLogService) Get(id int) *models.RetailLog {
	return s.dao.Get(id)
}

func (s *retailLogService) GetAll(page, size int) []models.RetailLog {
	return s.dao.GetAll(page, size)
}
func (s *retailLogService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *retailLogService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *retailLogService) Update(data *models.RetailLog, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *retailLogService) Create(data *models.RetailLog) error {
	return s.dao.Create(data)
}
