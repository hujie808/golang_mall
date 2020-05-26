package services

import (
	"errors"
	"math/rand"
	"time"
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallCommodityService interface {
	Get(id int) *models.MallCommodity
	GetAll(page, size int) []models.MallCommodity
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallCommodity, columns []string) error
	Create(data *models.MallCommodity) error
	GetCategory(category int) []models.MallCommodity
	GetSearchPage(page, size int, title string) (int, []models.MallCommodity)
	CreateId(data *models.MallCommodity) (int, error)
	GetRandomHot() []models.MallCommodity
	GetCategoryAll(category []int) []models.MallCommodity
	GetCommodityAll(IdList []int) []models.MallCommodity //commodityId_list
	GetAllIndex(page, size int) []models.MallCommodity
}

type mallCommodityService struct {
	dao *dao.MallCommodityDao
}

func NewMallCommodityService() MallCommodityService {
	return &mallCommodityService{
		dao: dao.NewMallCommodityDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallCommodityService) Get(id int) *models.MallCommodity {
	return s.dao.Get(id)
}

func (s *mallCommodityService) GetAll(page, size int) []models.MallCommodity {
	return s.dao.GetAll(page, size)
}
func (s *mallCommodityService) GetAllIndex(page, size int) []models.MallCommodity{
	return s.dao.GetAllIndex(page, size)
}
func (s *mallCommodityService) GetSearchPage(page, size int, title string) (int, []models.MallCommodity){
	return s.dao.GetSearchPage(page, size, title)
}

func (s *mallCommodityService) GetCategory(category int) []models.MallCommodity {
	return s.dao.GetCategory(category)
}

func (s *mallCommodityService) GetCategoryAll(category []int) []models.MallCommodity {
	return s.dao.GetCategoryAll(category)
}
func (s *mallCommodityService) GetCommodityAll(IdList []int) []models.MallCommodity {
	return s.dao.GetCommodityAll(IdList)
}

func (s *mallCommodityService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallCommodityService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallCommodityService) Update(data *models.MallCommodity, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallCommodityService) Create(data *models.MallCommodity) error {
	return s.dao.Create(data)
}

func (s *mallCommodityService) CreateId(data *models.MallCommodity) (int, error) {
	return s.dao.CreateId(data)
}

func (s *mallCommodityService) GetRandomHot() []models.MallCommodity {
	commodity := s.dao.GetRandomHot()
	if len(commodity) > 6 {
		commodity, _ = random(commodity, 6)
	}
	return commodity
}

func random(strings []models.MallCommodity, length int) ([]models.MallCommodity, error) {
	rand.Seed(time.Now().Unix())
	if len(strings) <= 0 {
		return nil, errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(strings) <= length {
		return nil, errors.New("the size of the parameter length illegal")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := make([]models.MallCommodity, 0)
	for i := 0; i < length; i++ {
		str = append(str, strings[i])
	}
	return str, nil
}
