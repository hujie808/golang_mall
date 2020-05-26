package utils

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

type FileNode struct {
	Name      string      `json:"title"`
	Id        int         `json:"key"`
	Image     string      `json:"image"`
	FileNodes []*FileNode `json:"children"`
}

type Allcategory struct {
	All []models.MallCategory
}

func (a Allcategory) GetParent(id int) Allcategory {
	datalist := make([]models.MallCategory, 0)
	for _, v := range a.All {
		if v.ParentId == id {
			datalist = append(datalist, v)
		}
	}
	return Allcategory{datalist}
}

func walk(id int, node *FileNode, allCategory Allcategory) {
	parentId := allCategory.GetParent(id)
	for _, yiji := range parentId.All {
		child := FileNode{yiji.Name, yiji.Id,conf.Host+yiji.Image, []*FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)
		parentId := allCategory.GetParent(id)
		if parentId.All != nil {
			walk(yiji.Id, &child, allCategory)
		}
	}
	return
}

func GetAllCategory(s services.MallCategoryService) string {
	count := s.CountAll()
	allCategory := s.GetAll(1, int(count))
	root := FileNode{"frist", 0,"", []*FileNode{}}
	walk(0, &root, Allcategory{All: allCategory})
	data, _ := json.Marshal(root)
	return gjson.Get(string(data), "children").String()
}
