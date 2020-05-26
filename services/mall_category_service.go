package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallCategoryService interface {
	Get(id int) *models.MallCategory
	GetAll(page, size int) []models.MallCategory
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallCategory, columns []string) error
	Create(data *models.MallCategory) error
	GetParent(id int) []models.MallCategory
	GetHot(page, size int) []models.MallCategory
	GetParenAll(category []int) []models.MallCategory
	GetParentId(id int) []int
	GetName(name string) []models.MallCategory
}

type mallCategoryService struct {
	dao *dao.MallCategoryDao
}

func NewMallCategoryService() MallCategoryService {
	return &mallCategoryService{
		dao: dao.NewMallCategoryDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallCategoryService) Get(id int) *models.MallCategory {
	return s.dao.Get(id)
}

func (s *mallCategoryService) GetParent(id int) []models.MallCategory {
	return s.dao.GetParen(id)
}
func (s *mallCategoryService) GetParentId(id int) []int {
	list := []int{}
	l_category := s.dao.GetParen(id)
	for _, v := range l_category {
		list = append(list, v.Id)
	}
	return list
}

//func (s *mallCategoryService) GetIsParen() []models.MallCategory {
//	return s.dao.GetIsParen()
//}

func (s *mallCategoryService) GetParenAll(category []int) []models.MallCategory {
	return s.dao.GetParenAll(category)
}

func (s *mallCategoryService) GetAll(page, size int) []models.MallCategory {
	return s.dao.GetAll(page, size)
}
func (s *mallCategoryService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallCategoryService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallCategoryService) Update(data *models.MallCategory, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallCategoryService) Create(data *models.MallCategory) error {
	return s.dao.Create(data)
}
func (s *mallCategoryService) GetHot(page, size int) []models.MallCategory {
	return s.dao.GetHot(page, size)
}
func (s *mallCategoryService) GetName(name string) []models.MallCategory {
	return s.dao.GetName(name)
}
