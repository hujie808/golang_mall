package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type RedLogService interface {
	Get(id int) *models.RedLog
	GetAll(page, size int) []models.RedLog
	CountAll() int64
	Delete(id int) error
	Update(data *models.RedLog, columns []string) error
	Create(data *models.RedLog) error
}

type redLogService struct {
	dao *dao.RedLogDao
}

func NewRedLogService() RedLogService {
	return &redLogService{
		dao: dao.NewRedLogDao(datasource.InstanceDbMaster()),
	}
}

func (s *redLogService) Get(id int) *models.RedLog {
	return s.dao.Get(id)
}

func (s *redLogService) GetAll(page, size int) []models.RedLog {
	return s.dao.GetAll(page, size)
}
func (s *redLogService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *redLogService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *redLogService) Update(data *models.RedLog, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *redLogService) Create(data *models.RedLog) error {
	return s.dao.Create(data)
}
