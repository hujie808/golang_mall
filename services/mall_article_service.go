package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallArticleService interface {
	Get(id int) *models.MallArticle
	GetAll(page, size int) []models.MallArticle
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallArticle, columns []string) error
	Create(data *models.MallArticle) error
	GetAllWz(page, size int) []models.MallArticle
	GetAllYfs(page, size int) []models.MallArticle
}

type mallArticleService struct {
	dao *dao.MallArticleDao
}

func NewMallArticleService() MallArticleService {
	return &mallArticleService{
		dao: dao.NewMallArticleDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallArticleService) Get(id int) *models.MallArticle {
	return s.dao.Get(id)
}

func (s *mallArticleService) GetAll(page, size int) []models.MallArticle {
	return s.dao.GetAll(page, size)
}

func (s *mallArticleService) GetAllWz(page, size int) []models.MallArticle {
	return s.dao.GetAllWz(page, size)
}

func (s *mallArticleService) GetAllYfs(page, size int) []models.MallArticle {
	return s.dao.GetAllYfs(page, size)
}

func (s *mallArticleService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallArticleService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallArticleService) Update(data *models.MallArticle, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallArticleService) Create(data *models.MallArticle) error {
	return s.dao.Create(data)
}
