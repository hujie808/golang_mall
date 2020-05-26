package utils

//func AddModelsName(models interface{},name string ,service interface{})[]map[string]interface{}{
//	result :=make([]map[string]interface{},0)
//	reflect.Swapper()
//	for _,models :=range models.(reflect.Slice){
//		modelsMap :=structMap(models)
//		for k,v:=range modelsMap{
//			if k==name{
//				modelsMap["name"]=name
//			}
//			fmt.Println(v)
//		}
//		result=append(result,modelsMap)
//	}
//	return result
//}
//
//
//func structMap(obj interface{}) map[string]interface{} {
//	t := reflect.TypeOf(&obj)
//	v := reflect.ValueOf(&obj)
//	var data = make(map[string]interface{})
//	for i := 0; i < t.NumField(); i++ {
//		data[t.Field(i).Name] = v.Field(i).Interface()
//	}
//	return data
//}
