package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallOrderInfoService interface {
	Get(id int) *models.MallOrderInfo
	GetAll(page, size int) []models.MallOrderInfo
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallOrderInfo, columns []string) error
	Create(data *models.MallOrderInfo) error
	GetOrder(orderId int) []models.MallOrderInfo
	GetCommodity(c_id int) []models.MallOrderInfo
	GetOrderListInfoId(orderId []int) []models.MallOrderInfo //oderIdList 找到所属的info
	ReallyDelete(orderId, wechatId int) error                //删除关于订单的详情
	GetVideoAll(userId int) []models.JoinOrderInfo
}

type mallOrderInfoService struct {
	dao *dao.MallOrderInfoDao
}

func NewMallOrderInfoService() MallOrderInfoService {
	return &mallOrderInfoService{
		dao: dao.NewMallOrderInfoDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallOrderInfoService) Get(id int) *models.MallOrderInfo {
	return s.dao.Get(id)
}

func (s *mallOrderInfoService) GetAll(page, size int) []models.MallOrderInfo {
	return s.dao.GetAll(page, size)
}

func (s *mallOrderInfoService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallOrderInfoService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallOrderInfoService) Update(data *models.MallOrderInfo, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallOrderInfoService) Create(data *models.MallOrderInfo) error {
	return s.dao.Create(data)
}

func (s *mallOrderInfoService) GetCommodity(c_id int) []models.MallOrderInfo {
	return s.dao.GetCommodity(c_id)
}

func (s *mallOrderInfoService) GetOrder(orderId int) []models.MallOrderInfo {
	return s.dao.GetOrder(orderId)
}

func (s *mallOrderInfoService) GetOrderListInfoId(listOrderId []int) []models.MallOrderInfo {
	return s.dao.GetOrderListInfoId(listOrderId)
}

func (s *mallOrderInfoService) ReallyDelete(orderId, wechatId int) error {
	return s.dao.ReallyDelete(orderId, wechatId)
}

func (s *mallOrderInfoService) GetVideoAll(userId int) []models.JoinOrderInfo {
	return s.dao.GetVideoAll(userId)
}
