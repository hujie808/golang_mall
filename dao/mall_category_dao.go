package dao

import (
	"github.com/go-xorm/xorm"

	"web_iris/golang_mall/models"
)

type MallCategoryDao struct {
	engine *xorm.Engine
}

func NewMallCategoryDao(engine *xorm.Engine) *MallCategoryDao { //实例化引擎
	return &MallCategoryDao{
		engine: engine,
	}
}

//get_id
func (d *MallCategoryDao) Get(id int) *models.MallCategory {
	data := &models.MallCategory{Id: id, SysStatus: 0}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallCategoryDao) GetParen(id int) []models.MallCategory {
	datalist := make([]models.MallCategory, 0)
	err := d.engine.Desc("desc").
		Where("parent_id=?", id). // 有效的用户
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *MallCategoryDao) GetIsParen() int {
	count, err := d.engine.
		Where("parent_id>?", 0). // 有效的用户
		And("sys_status=?", 0).  // 有效的用户
		Count(&models.MallCategory{})
	if err != nil {
		return int(count)
	} else {
		return int(count)
	}
}

func (d *MallCategoryDao) GetParenAll(category []int) []models.MallCategory {
	datalist := make([]models.MallCategory, 0)
	err := d.engine.
		In("parent_id", category).
		And("sys_status=?", 0).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCategoryDao) GetAll(page, size int) []models.MallCategory {
	offset := (page - 1) * size
	datalist := make([]models.MallCategory, 0)
	err := d.engine.Desc("desc").
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

func (d *MallCategoryDao) GetName(name string) []models.MallCategory {
	datalist := make([]models.MallCategory, 0)
	err := d.engine.
		Desc("desc").
		Desc("id").
		Where("sys_status=?", 0). // 有效的用户
		And("name Like ?", "%"+name+"%").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCategoryDao) GetHot(page, size int) []models.MallCategory {
	offset := (page - 1) * size
	datalist := make([]models.MallCategory, 0)
	err := d.engine.
		Desc("desc").
		Desc("id").
		Limit(size, offset).
		Where("sys_status=?", 0). // 有效的用户
		And("category_hot=?", "0").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCategoryDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallCategory{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallCategoryDao) Delete(id int) error {
	data := &models.MallCategory{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallCategoryDao) Update(data *models.MallCategory, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallCategoryDao) Create(data *models.MallCategory) error {
	_, err := d.engine.Insert(data)
	return err
}
