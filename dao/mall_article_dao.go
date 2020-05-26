package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallArticleDao struct {
	engine *xorm.Engine
}

func NewMallArticleDao(engine *xorm.Engine) *MallArticleDao { //实例化引擎
	return &MallArticleDao{
		engine: engine,
	}
}

//get_id
func (d *MallArticleDao) Get(id int) *models.MallArticle {
	data := &models.MallArticle{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallArticleDao) GetAll(page, size int) []models.MallArticle {
	offset := (page - 1) * size
	datalist := make([]models.MallArticle, 0)
	err := d.engine.
		Desc("order","id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallArticleDao) GetAllWz(page, size int) []models.MallArticle {
	offset := (page - 1) * size
	datalist := make([]models.MallArticle, 0)
	err := d.engine.
		Desc("order","id").
		Where("a_type=?", "wz"). // 有效的用户
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallArticleDao) GetAllYfs(page, size int) []models.MallArticle {
	offset := (page - 1) * size
	datalist := make([]models.MallArticle, 0)
	err := d.engine.
		Desc("order","id").
		Where("a_type=?", "yfs"). // 有效的用户
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}


func (d *MallArticleDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.MallArticle{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallArticleDao) Delete(id int) error {
	data := &models.MallArticle{Id: id}
	_, err := d.engine.Id(id).Delete(data)
	return err
}

func (d *MallArticleDao) Update(data *models.MallArticle, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallArticleDao) Create(data *models.MallArticle) error {
	_, err := d.engine.Insert(data)
	return err
}
