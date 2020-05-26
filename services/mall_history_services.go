package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallHistoryService interface {
	Get(id int) *models.MallHistory
	GetAll(page, size int) []models.MallHistory
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallHistory, columns []string) error
	Create(data *models.MallHistory) error
	GetWechatId(id int) *models.MallHistory
}

type mallHistoryService struct {
	dao *dao.MallHistoryDao
}

func NewMallHistoryService() MallHistoryService {
	return &mallHistoryService{
		dao: dao.NewMallHistoryDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallHistoryService) Get(id int) *models.MallHistory {
	return s.dao.Get(id)
}

func (s *mallHistoryService) GetAll(page, size int) []models.MallHistory {
	return s.dao.GetAll(page, size)
}
func (s *mallHistoryService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallHistoryService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallHistoryService) Update(data *models.MallHistory, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallHistoryService) Create(data *models.MallHistory) error {
	return s.dao.Create(data)
}

func (s *mallHistoryService) GetWechatId(id int) *models.MallHistory {
	return s.dao.GetWechatId(id)
}
