package utils

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

type fileNode struct {
	Name      string      `json:"title"`
	Id        int         `json:"key"`
	Image     string      `json:"image"`
	FileNodes []*fileNode `json:"children"`
}

type allcategory struct {
	All []models.MallCategory
}

var lastId = []int{}

func (a allcategory) GetParent(id int) allcategory {
	datalist := make([]models.MallCategory, 0)
	for _, v := range a.All {
		if v.ParentId == id {
			datalist = append(datalist, v)
		}
	}
	return allcategory{datalist}
}

func walks(id int, node *fileNode, allCategory allcategory) {

	parentId := allCategory.GetParent(id)
	for _, yiji := range parentId.All {
		child := fileNode{yiji.Name, yiji.Id, conf.Host+yiji.Image,[]*fileNode{}}
		node.FileNodes = append(node.FileNodes, &child)
		parentId := allCategory.GetParent(id)
		if parentId.All != nil {
			lastId = append(lastId, yiji.Id)
			walks(yiji.Id, &child, allCategory)
			continue
		}
	}
	return
}

func GetLastCategory(s services.MallCategoryService) []int {
	count := s.CountAll()
	allCategory := s.GetAll(1, int(count))
	root := fileNode{"frist", 0,"", []*fileNode{}}
	walks(0, &root, allcategory{All: allCategory})
	//, _ := json.Marshal(root)
	return lastId
}
