package baidumap

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	API_SERVER = "api.map.baidu.com"
	GEOCONV_F_WGS84  = "1"
	GEOCONV_F_SOGOU  = "2"
	GEOCONV_F_GOOGLE = "3"
	GEOCONV_F_SOSO   = "3"
	GEOCONV_F_ALIYUN = "3"
	GEOCONV_F_MAPABC = "3"
	GEOCONV_F_MET    = "4"
	GEOCONV_F_BD_LL  = "5"
	GEOCONV_F_BD_MET = "6"
	GEOCONV_F_MAPBAR = "7"
	GEOCONV_F_51     = "8"
	GEOCONV_T_BD09ll = "5"
	GEOCONV_T_BD09MC = "6"
)

type BaiduApi struct {
	Ak string
	Version string
}


func (api BaiduApi) call (serviceName string, subService string, params map[string]string, template interface{}) (error){
	params["ak"] = api.Ak
	temp := "?"
	for key, value := range params {
		temp += key + "=" + value + "&"
	}
	queryString := temp[0:len(temp) - 1]
	url := fmt.Sprintf("http://%s/%s/%s/%s%s", 
		API_SERVER, serviceName, api.Version, 
		subService, queryString)
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
	// 解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal(body, &template); err != nil {
		return err	
	}
	return nil
}

// 百度地图坐标转换API是一套以HTTP形式提供的坐标转换接口，用于将常用的非百度坐标 (目前支持GPS设备
// 获取的坐标、 google地图坐标、soso地图坐标、amap地图坐标、mapbar地图坐标）转换成百度地图中使用
// 的坐标，并可将转化后的坐标在百度地图JavaScript API、车联网API、静态图API、web服务API等产品中使用。
// http://lbsyun.baidu.com/index.php?title=webapi/guide/changeposition
func (api BaiduApi) Geoconv (params map[string]string) (*GeoconvRet, error){
	ret := &GeoconvRet{}
	err := api.call("geoconv", "", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
// Geocoding API包括地址解析和逆地址解析功能：
// 地理编码：即地址解析，由详细到街道的结构化地址得到百度经纬度信息
// 逆地理编码：即逆地址解析，由百度经纬度信息得到结构化地址信息
// http://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-geocoding
func (api BaiduApi) Geocoding (params map[string]string) (interface{}, error){
	ret := interface{}(nil)
	if _, ok := params["address"]; ok {
		ret = &GeoEncoderRet{}
	} else {
		ret = &GeocoderRet{}
	}
	params["output"] = "json"
	err := api.call("geocoder", "", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
// 用于返回查询某个区域的某类POI数据，且提供单个POI的详情查询服务
// http://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-placeapi
func (api *BaiduApi) PlaceSearch (params map[string]string) (interface{}, error){
	ret := &PlaceSearchRet{}
	params["output"] = "json"
	err := api.call("place", "search", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
// 用于返回查询某个区域的某类POI数据，且提供单个POI的详情查询服务
// http://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-placeapi
func (api *BaiduApi) PlaceDetail (params map[string]string) (interface{}, error){
	ret := &PlaceDetailRet{}
	params["output"] = "json"
	err := api.call("place", "detail", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
// Place suggestion API 是一套提供匹配用户输入关键字辅助信息、提示服务的接口
// ，可返回json或xml格式的一组建议词条的数据。
func (api *BaiduApi) Suggestion (params map[string]string) (interface{}, error){
	ret := &SuggestionRet{}
	params["output"] = "json"
	err := api.call("place", "suggestion", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}



