package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallShoppingService interface {
	Get(id int) *models.MallShopping
	GetAll(page, size int) []models.MallShopping
	CountAll() int64
	Delete(id int) error
	Update(data *models.MallShopping, columns []string) error
	Create(data *models.MallShopping) error
	GetShopping(user_id int) []int
	GetCollection(user_id int) []int
	GetIsCollection(user_id, commodity_id int) []models.MallShopping
	DeleteShopping(id int) error          //彻底删除购物车
	DeleteShoppingAll(wechatId,skuid int) error //彻底删除购物车
	GetIsShopping(user_id, sku_id int) *models.MallShopping //用户的购物车 sku
	GetShoppingList(user_id int) []models.MallShopping
}

type mallShoppingService struct {
	dao *dao.MallShoppingDao
}

func NewMallShoppingService() MallShoppingService {
	return &mallShoppingService{
		dao: dao.NewMallShoppingDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallShoppingService) Get(id int) *models.MallShopping {
	return s.dao.Get(id)
}

func (s *mallShoppingService) GetAll(page, size int) []models.MallShopping {
	return s.dao.GetAll(page, size)
}
func (s *mallShoppingService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *mallShoppingService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallShoppingService) Update(data *models.MallShopping, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallShoppingService) Create(data *models.MallShopping) error {
	return s.dao.Create(data)
}

func (s *mallShoppingService) GetShopping(user_id int) []int {
	sku_list := []int{}
	shopping_dao := s.dao.GetShopping(user_id)
	for _, v := range shopping_dao {
		if v.MallSkuId > 0 {
			sku_list = append(sku_list, v.MallSkuId)
		}
	}
	return sku_list
}

func (s *mallShoppingService) GetShoppingList(user_id int) []models.MallShopping {

	return s.dao.GetShopping(user_id)

}

func (s *mallShoppingService) GetCollection(user_id int) []int {
	sku_list := []int{}
	shopping_dao := s.dao.GetCollection(user_id)
	for _, v := range shopping_dao {
		if v.MallSkuId > 0 {
			sku_list = append(sku_list, v.MallSkuId)
		}
	}
	return sku_list
}

func (s *mallShoppingService) GetIsShopping(user_id, sku_id int) *models.MallShopping {
	return s.dao.GetIsShopping(user_id, sku_id)
}

//用户id 商品id
func (s *mallShoppingService) GetIsCollection(user_id, commodity_id int) []models.MallShopping {
	return s.dao.GetIsCollection(user_id, commodity_id)
}

//删除购物车
func (s *mallShoppingService) DeleteShopping(id int) error {
	return s.dao.DeleteShopping(id)
}
func (s *mallShoppingService) DeleteShoppingAll(wechatId,skuid int) error {
	return s.dao.DeleteShoppingAll(wechatId,skuid)
}
