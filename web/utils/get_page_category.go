package utils

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

type fileNodePage struct {
	Name          string          `json:"title"`
	Id            int             `json:"key"`
	FileNodePages []*fileNodePage `json:"children"`
	Superior      string          `json:"Image"`
}

type AllPagecategory struct {
	All []models.MallCategory
}

func (a AllPagecategory) GetParent(id int) AllPagecategory {
	datalist := make([]models.MallCategory, 0)
	for _, v := range a.All {
		if v.ParentId == id {
			datalist = append(datalist, v)
		}
	}
	return AllPagecategory{datalist}
}

func walk_page(id int, node *fileNodePage, allCategory AllPagecategory) {
	parentId := allCategory.GetParent(id)
	for _, yiji := range parentId.All {
		var child fileNodePage
		if yiji.ParentId == 0 {
			child = fileNodePage{yiji.Name, yiji.Id, []*fileNodePage{}, conf.Host + yiji.Image}
		} else {
			child = fileNodePage{yiji.Name, yiji.Id, []*fileNodePage{}, conf.Host + yiji.Image}
		}

		node.FileNodePages = append(node.FileNodePages, &child)
		parentId := allCategory.GetParent(id)
		if parentId.All != nil {
			walk_page(yiji.Id, &child, allCategory)
		}
	}
	return
}

func GetAllPageCategory(s services.MallCategoryService) string {
	count := s.CountAll()
	allCategory := s.GetAll(1, int(count))
	root := fileNodePage{"frist", 0, []*fileNodePage{}, conf.Host + ""}
	walk_page(0, &root, AllPagecategory{All: allCategory})
	data, _ := json.Marshal(root)
	return gjson.Get(string(data), "children").String()
}
