package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type MallWechatService interface {
	Get(id int) *models.MallWechat
	GetAll(page, size int) []models.MallWechat
	CountAll(nickName, tel string) int64
	Delete(id int) error
	Update(data *models.MallWechat, columns []string) error
	Create(data *models.MallWechat) error
	SearchNickname(page, size int, nickName, tel string) []models.MallWechat
	GetTel(tel string) *models.MallWechat
	GetOpenid(openid string) *models.MallWechat
}

type mallWechatService struct {
	dao *dao.MallWechatDao
}

func NewMallWechatService() MallWechatService {
	return &mallWechatService{
		dao: dao.NewMallWechatDao(datasource.InstanceDbMaster()),
	}
}

func (s *mallWechatService) Get(id int) *models.MallWechat {
	return s.dao.Get(id)
}

func (s *mallWechatService) GetTel(tel string) *models.MallWechat {
	return s.dao.GetTel(tel)
}

func (s *mallWechatService) GetAll(page, size int) []models.MallWechat {
	return s.dao.GetAll(page, size)
}
func (s *mallWechatService) CountAll(nickName, tel string) int64 {
	return s.dao.CountAll(nickName, tel)
}
func (s *mallWechatService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *mallWechatService) Update(data *models.MallWechat, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *mallWechatService) Create(data *models.MallWechat) error {
	return s.dao.Create(data)
}

func (s *mallWechatService) SearchNickname(page, size int, nickName, tel string) []models.MallWechat {
	return s.dao.SearchNickname(page, size, nickName, tel)
}
func (s *mallWechatService) GetOpenid(openid string) *models.MallWechat {
	return s.dao.GetOpenid(openid)
}
