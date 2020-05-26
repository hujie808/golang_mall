package models

type ObjUserHistory struct {
	CommodityId       int     //商品id
	CommodityTitle    string  //商品标题
	CommodityNowPrice float32 //商品价格
	CommodityPrice    float32 //商品价格
	CommodityImage    string  //图片
	HTime             string  //时间:7月30日
	Sort              int     //排序
}
