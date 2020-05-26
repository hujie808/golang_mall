package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallSkuService interface {
	Get(id int) *models.MallSku
	GetAll(page, size int) []models.MallSku
	GetCommodityId(commodityId int) []models.MallSku
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallSku, columns []string) error
	Create(data *models.MallSku) error
	GetIdAll(idList []int) []models.MallSku
	GetCommodityStock(commodityId int) int
	GetIdAllCommodityTitle(idList []int) []models.JionSku
	GetTitle(id int) *models.JionSku //单独id有商品title
}

type mallSkuService struct {
	dao *dao.MallSkuDao
}

func NewMallSkuService() MallSkuService {
	return &mallSkuService{
		dao: dao.NewMallSkuDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallSkuService) Get(id int) *models.MallSku {
	return s.dao.Get(id)
}

func (s *mallSkuService) GetAll(page, size int) []models.MallSku {
	return s.dao.GetAll(page, size)
}

func (s *mallSkuService) GetCommodityId(commodityId int) []models.MallSku {
	return s.dao.GetCommodityId(commodityId)
}

//库存平均值
func (s *mallSkuService) GetCommodityStock(commodityId int) int {
	list := s.dao.GetCommodityId(commodityId)
	avg := 0
	for _, v := range list {
		avg += v.Stock
	}
	return avg
}

func (s *mallSkuService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallSkuService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallSkuService) Update(data *models.MallSku, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallSkuService) Create(data *models.MallSku) error {
	return s.dao.Create(data)
}

//连接商品表
func (s *mallSkuService) GetIdAllCommodityTitle(idList []int) []models.JionSku {
	return s.dao.GetIdAllCommodityTitle(idList)
}

func (s *mallSkuService) GetIdAll(idList []int) []models.MallSku {
	return s.dao.GetIdAll(idList)
}

func (s *mallSkuService) GetTitle(id int) *models.JionSku {
	return s.dao.GetTitle(id)
}
