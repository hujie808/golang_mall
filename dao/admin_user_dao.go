package dao

import (
	"github.com/go-xorm/xorm"

	"log"
	"web_iris/golang_mall/models"
)

type AdminUserDao struct {
	engine *xorm.Engine
}

func NewAdminUserDao(engine *xorm.Engine) *AdminUserDao { //实例化引擎
	return &AdminUserDao{
		engine: engine,
	}
}

//get_id
func (d *AdminUserDao) Get(id int) *models.AdminUser {
	data := &models.AdminUser{Id: id}
	ok, err := d.engine.
		Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *AdminUserDao) GetAll(page, size int) []models.AdminUser {
	offset := (page - 1) * size
	datalist := make([]models.AdminUser, 0)
	err := d.engine.Desc("id").
		Limit(size, offset).
		Cols("id", "username", "last_login", "is_superuser", "sys_created").
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *AdminUserDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.AdminUser{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *AdminUserDao) Delete(id int) error {
	data := &models.AdminUser{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *AdminUserDao) Update(data *models.AdminUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *AdminUserDao) Create(data *models.AdminUser) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *AdminUserDao) LoginUser(username, password string) *models.AdminUser {
	datalist := make([]models.AdminUser, 0)
	err := d.engine.
		Where("username=?", username).
		Where("password=?", password).
		Where("sys_status=?", 0).
		Find(&datalist)
	if err != nil {
		log.Println("admin_user_dao.go LoginUser username ,password=", username, password, "err=", err)
		return nil
	} else {
		if len(datalist) > 0 {
			return &datalist[0]
		} else {
			return nil
		}

	}
}
func (d *AdminUserDao) GetUserName(username string) *models.AdminUser {
	datalist := make([]models.AdminUser, 0)
	err := d.engine.
		Where("username=?", username).
		//Where("sys_status=?", 0).
		Find(&datalist)
	if err != nil {
		log.Println("admin_user_dao.go GetUserName username ,password=", username, "err=", err)
		return nil
	} else {
		if len(datalist) > 0 {
			return &datalist[0]
		} else {
			return nil
		}

	}
}
