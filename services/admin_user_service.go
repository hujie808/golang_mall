package services

import (
	"web_iris/golang_mall/dao"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
)

type AdminUserService interface {
	Get(id int) *models.AdminUser
	GetAll(page, size int) []models.AdminUser
	CountAll() int64
	Delete(id int) error
	Update(data *models.AdminUser, columns []string) error
	LoginUser(username, password string) *models.AdminUser
	GetUserName(username string) *models.AdminUser
	Create(data *models.AdminUser) error
}

type adminUserService struct {
	dao *dao.AdminUserDao
}

func NewAdminUserService() AdminUserService {
	return &adminUserService{
		dao: dao.NewAdminUserDao(datasource.InstanceDbMaster()),
	}
}

func (s *adminUserService) Get(id int) *models.AdminUser {
	return s.dao.Get(id)
}

func (s *adminUserService) GetAll(page, size int) []models.AdminUser {
	return s.dao.GetAll(page, size)
}
func (s *adminUserService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *adminUserService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *adminUserService) Update(data *models.AdminUser, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *adminUserService) Create(data *models.AdminUser) error {
	return s.dao.Create(data)
}

func (s *adminUserService) LoginUser(username, password string) *models.AdminUser {
	return s.dao.LoginUser(username, password)
}
func (s *adminUserService) GetUserName(username string) *models.AdminUser {
	return s.dao.GetUserName(username)
}
