package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallWechatDao struct {
	engine *xorm.Engine
}

func NewMallWechatDao(engine *xorm.Engine) *MallWechatDao { //实例化引擎
	return &MallWechatDao{
		engine: engine,
	}
}

//get_id
func (d *MallWechatDao) Get(id int) *models.MallWechat {
	data := &models.MallWechat{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallWechatDao) GetOpenid(openid string) *models.MallWechat {
	data := &models.MallWechat{Openid: openid}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

//get_id
func (d *MallWechatDao) GetTel(tel string) *models.MallWechat {
	data := &models.MallWechat{Tel: tel}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallWechatDao) GetAll(page, size int) []models.MallWechat {
	offset := (page - 1) * size
	datalist := make([]models.MallWechat, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallWechatDao) SearchNickname(page, size int, nickName, tel string) []models.MallWechat {
	offset := (page - 1) * size
	datalist := make([]models.MallWechat, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Where("nickname Like ?", "%"+nickName+"%").
		And("tel Like ?", "%"+tel+"%").
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallWechatDao) SearchTel(tel string) []models.MallWechat {
	datalist := make([]models.MallWechat, 0)
	err := d.engine.
		Desc("id").
		Where("tel Like ?", "%"+tel+"%").
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallWechatDao) CountAll(nickName, tel string) int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		And("nickname Like ?", "%"+nickName+"%").
		And("tel Like ?", "%"+tel+"%").
		Count(&models.MallWechat{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallWechatDao) Delete(id int) error {
	data := &models.MallWechat{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallWechatDao) Update(data *models.MallWechat, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallWechatDao) Create(data *models.MallWechat) error {
	_, err := d.engine.Insert(data)
	return err
}
