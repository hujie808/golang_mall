package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallOrderService interface {
	Get(id int) *models.MallOrder
	GetAll(page, size int) []models.MallOrder
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallOrder, columns []string) error
	Create(data *models.MallOrder) error
	SearchPageIs(page, size int, pay_status, tel string) ([]models.JionOrder, int)
	CreateId(data *models.MallOrder) (int, error)
	SearchUserOrder(isType, userId int) []models.MallOrder //找到关于user的order
	ReallyDelete(id, wecahtId int) error                   //删除订单
	CancelDelete(id int) error                             //取消订单
	SearchUserOrderId(isType, userId, id int) []models.MallOrder
	GetOutTradeNo(out_trade_no string) *models.MallOrder
}

type mallOrderService struct {
	dao *dao.MallOrderDao
}

func NewMallOrderService() MallOrderService {
	return &mallOrderService{
		dao: dao.NewMallOrderDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallOrderService) Get(id int) *models.MallOrder {
	return s.dao.Get(id)
}

func (s *mallOrderService) GetAll(page, size int) []models.MallOrder {
	return s.dao.GetAll(page, size)
}

func (s *mallOrderService) SearchPageIs(page, size int, pay_status, tel string) ([]models.JionOrder, int) {
	return s.dao.SearchPageIs(page, size, pay_status, tel)
}

func (s *mallOrderService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallOrderService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallOrderService) Update(data *models.MallOrder, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallOrderService) Create(data *models.MallOrder) error {
	return s.dao.Create(data)
}

func (s *mallOrderService) SearchUserOrder(isType, userId int) []models.MallOrder {
	return s.dao.SearchUserOrder(isType, userId)
}

func (s *mallOrderService) SearchUserOrderId(isType, userId, id int) []models.MallOrder {
	return s.dao.SearchUserOrderId(isType, userId, id)
}

func (s *mallOrderService) CreateId(data *models.MallOrder) (int, error) {
	return s.dao.CreateId(data)
}

func (s *mallOrderService) ReallyDelete(id, wecahtId int) error {
	return s.dao.ReallyDelete(id, wecahtId)
}

func (s *mallOrderService) CancelDelete(id int) error {
	return s.dao.CancelDelete(id)
}

func (s *mallOrderService) GetOutTradeNo(out_trade_no string) *models.MallOrder {
	return s.dao.GetOutTradeNo(out_trade_no)
}
