package controllers

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
	"sort"
	"strconv"
	"strings"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

func historySave(s services.MallHistoryService, userId, commodityId int) {
	history := s.GetWechatId(userId)
	if history.Id > 0 {
		viewHistory := historyJson(history.ViewHistory, commodityId)
		history.ViewHistory = viewHistory
		err:=s.Update(history, []string{"view_history"})
		if err!=nil{
			log.Println("历史更新失败")
		}
	} else {
		create_h := historyJson("", commodityId)
		newHistory := &models.MallHistory{
			MallWechatId: userId,
			ViewHistory:  create_h,
			SysStatus:    0,
		}
		err := s.Create(newHistory)
		if err != nil {
			log.Println("history_api.go historySave err=", err)
		}
	}
}

func historyJson(his string, commodityId int) string {

	list := gjson.Parse(his).Map()
	result := make(map[string]int)

	result[strconv.Itoa(commodityId)] = int(comm.NowUnix())
	for k, v := range list {
		result[k] = int(v.Int())
		if k==strconv.Itoa(commodityId){
			result[k] =int(comm.NowUnix())
		}
	}
	s, _ := json.Marshal(result)
	mString := string(s)
	return mString
}

//拿到筛选完的数据
func historyGetAll(s services.MallHistoryService, c services.MallCommodityService, userId int) []map[string][]models.ObjUserHistory {
	history := s.GetWechatId(userId)
	var result ResultH
	results := make(map[string][]models.ObjUserHistory, 0)
	commodityList := commodityAllMap(c.GetAll(0, 0)) //商品列表
	if history.Id > 0 {
		r_historys := HistoryJsonsDate(history.ViewHistory)
		for _, v := range r_historys {
			v_int, _ := strconv.Atoi(v.Key)
			if commodityList[v_int].Id == 0 {
				continue
			}
			r_date := comm.FormatFromUnixTimeShort(int64(v.Value))
			userHistory := models.ObjUserHistory{
				CommodityId:       commodityList[v_int].Id,
				CommodityTitle:    commodityList[v_int].Title,
				CommodityNowPrice: commodityList[v_int].PriceNow,
				CommodityPrice:    commodityList[v_int].Price,
				CommodityImage:    commodityList[v_int].Image,
				HTime:             r_date,
				Sort:              v.Value,
			}
			results[r_date] = append(results[r_date], userHistory)
		}
		for k, v := range results {
			resultInList := make(map[string][]models.ObjUserHistory, 0)
			resultInList[k] = v
			result = append(result, resultInList)
		}
	}
	sort.Sort(result)
	return result
}

//筛选5个数据
//拿到筛选完的数据
func historyGetFive(s services.MallHistoryService, c services.MallCommodityService, userId int) []models.ObjUserHistory {
	history := s.GetWechatId(userId)
	var r_list ObjList
	results := make([]models.ObjUserHistory, 0)
	commodityList := commodityAllMap(c.GetAll(0, 0)) //商品列表
	if history.Id > 0 {
		r_historys := HistoryJsonsDate(history.ViewHistory) //排序后的数据
		for _, v := range r_historys {
			v_int, _ := strconv.Atoi(v.Key)
			r_date := comm.FormatFromUnixTimeShort(int64(v.Value))
			userHistory := models.ObjUserHistory{
				CommodityId:       commodityList[v_int].Id,
				CommodityTitle:    commodityList[v_int].Title,
				CommodityNowPrice: commodityList[v_int].PriceNow,
				CommodityPrice:    commodityList[v_int].Price,
				CommodityImage:    commodityList[v_int].Image,
				HTime:             r_date,
				Sort:              v.Value,
			}
			results = append(results, userHistory)
		}

	} else {
		r_list = results
		sort.Sort(r_list)
		if len(results) > 5 {
			return r_list[:5]
		} else {
			return r_list
		}

	}
	r_list = results
	sort.Sort(r_list)
	if len(results) > 5 {
		return r_list[:5]
	} else {
		return r_list
	}
}

func HistoryJsonsDate(historys string) PairList {
	result := make(PairList, 0)

	mapHistory := make(map[string]int, 0)
	err := json.Unmarshal([]byte(historys), &mapHistory) //转map
	if err != nil {
		log.Println("history_api.go historyGetAll err=", err)
	} else {
		result = sortMapByValue(mapHistory)
	}
	return result
}

func commodityAllMap(com []models.MallCommodity) map[int]models.MallCommodity {
	result := make(map[int]models.MallCommodity, 0)
	for _, c := range com {
		result[c.Id] = c
	}
	return result
}

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A slice of Pairs that implements sort.Interface to sort by Value.
type ObjList []models.ObjUserHistory

func (p ObjList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ObjList) Len() int           { return len(p) }
func (p ObjList) Less(i, j int) bool { return p[i].Sort > p[j].Sort }

type ResultH []map[string][]models.ObjUserHistory

func (p ResultH) Len() int      { return len(p) }
func (p ResultH) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ResultH) Less(i, j int) bool {
	var ii string
	var jj string
	for k,_:=range p[i]{
		ii=k
	}
	for k,_:=range p[j]{
		jj=k
	}

	ii=strings.Replace(ii,"-","",2)
	jj=strings.Replace(jj,"-","",2)
	iInt,_:=strconv.Atoi(ii)
	jInt,_:=strconv.Atoi(jj)
	return iInt > jInt
}

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
