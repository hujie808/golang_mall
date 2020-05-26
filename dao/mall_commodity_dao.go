package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"web_iris/golang_mall/models"
)

type MallCommodityDao struct {
	engine *xorm.Engine
}

func NewMallCommodityDao(engine *xorm.Engine) *MallCommodityDao { //实例化引擎
	return &MallCommodityDao{
		engine: engine,
	}
}

//get_id
func (d *MallCommodityDao) Get(id int) *models.MallCommodity {
	data := &models.MallCommodity{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MallCommodityDao) GetAll(page, size int) []models.MallCommodity {
	offset := (page - 1) * size
	datalist := make([]models.MallCommodity, 0)
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

func (d *MallCommodityDao) GetAllIndex(page, size int) []models.MallCommodity {
	offset := (page - 1) * size
	datalist := make([]models.MallCommodity, 0)
	err := d.engine.Desc("desc").
		Desc("id").
		Limit(size, offset).
		Where("sys_status=?", 0). // 有效的用户
		And("com_type=?","sw").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCommodityDao) GetRandomHot() []models.MallCommodity {
	datalist := make([]models.MallCommodity, 0)
	err := d.engine.Desc("desc").
		Desc("id").
		Where("sys_status=?", 0). // 有效的用户
		And("com_type=?","sw").
		And("commodity_hot=?", "0").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCommodityDao) GetSearchPage(page, size int, title string) (int, []models.MallCommodity) {
	offset := (page - 1) * size
	datalist := make([]models.MallCommodity, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		And("title Like ?", "%"+title+"%").
		//And("tel Like ?", "%"+tel+"%").
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)

	number, errcount := d.engine.
		Desc("id").
		Where("sys_status=?", 0). // 有效的用户
		And("title Like ?", "%"+title+"%").
		Count(&models.MallCommodity{})
	if errcount != nil {
		log.Println("mall_commodity_dao.go GetSearchPage errcount=", errcount)
	}
	if err != nil {
		return int(number), datalist
	} else {
		return int(number), datalist
	}
}

func (d *MallCommodityDao) GetCategory(category int) []models.MallCommodity {
	datalist := make([]models.MallCommodity, 0)
	err := d.engine.Desc("desc").
		Desc("id").
		Where("category_id=?", category). // 分类查询
		Where("sys_status=?", 0). // 有效的用户
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCommodityDao) GetCategoryAll(category []int) []models.MallCommodity {
	datalist := make([]models.MallCommodity, 0)
	err := d.engine.
		In("category_id", category, "sys_status=?", 0).
		And("sys_status=?", 0).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCommodityDao) GetCommodityAll(IdList []int) []models.MallCommodity {
	datalist := make([]models.MallCommodity, 0)
	err := d.engine.
		In("id", IdList, "sys_status=?", 0).
		And("sys_status=?", 0).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MallCommodityDao) CountAll() int64 {
	num, err := d.engine.
		Where("sys_status=?", 0).
		Count(&models.MallCommodity{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *MallCommodityDao) Delete(id int) error {
	data := &models.MallCommodity{Id: id, SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *MallCommodityDao) Update(data *models.MallCommodity, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *MallCommodityDao) Create(data *models.MallCommodity) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *MallCommodityDao) CreateId(data *models.MallCommodity) (int, error) {
	_, err := d.engine.Insert(data)
	return data.Id, err
}
