package bdlbsyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_SERVER       = "api.map.baidu.com"
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

type BaiduLbsApi struct {
	Ak string
}

func NewBaiduLbsApi(ak string) *BaiduLbsApi {
	return &BaiduLbsApi{
		Ak: ak,
	}
}

func (api *BaiduLbsApi) call(serviceName, subService, version string, params map[string]string, template interface{}) error {
	params["ak"] = api.Ak
	temp := "?"
	for key, value := range params {
		temp += key + "=" + value + "&"
	}
	queryString := temp[0 : len(temp)-1]
	url := fmt.Sprintf("http://%s/%s/%s/%s%s",
		API_SERVER, serviceName, version,
		subService, queryString)
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))
	// 解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal(body, &template); err != nil {
		return err
	}
	return nil
}

func (api *BaiduLbsApi) GeoConvert(lng, lat float64, from, to string) (rlng, rlat float64, err error) {
	params := make(map[string]string)
	params["coords"] = fmt.Sprintf("%v,%v", lng, lat)
	params["from"] = from
	params["to"] = to

	result, err := api.Geoconv(params)
	if err != nil {
		return
	}

	if result.Status != 0 {
		err = errors.New("baidu lbs cloud response error, status not zero.")
		return
	}

	if len(result.Result) != 1 {
		err = errors.New("baidu lbs cloud response error, content empty.")
	}

	rlng = result.Result[0].X
	rlat = result.Result[0].Y
	return
}

// 百度地图坐标转换API是一套以HTTP形式提供的坐标转换接口，用于将常用的非百度坐标 (目前支持GPS设备
// 获取的坐标、 google地图坐标、soso地图坐标、amap地图坐标、mapbar地图坐标）转换成百度地图中使用
// 的坐标，并可将转化后的坐标在百度地图JavaScript API、车联网API、静态图API、web服务API等产品中使用。
// http://lbsyun.baidu.com/index.php?title=webapi/guide/changeposition
func (api *BaiduLbsApi) Geoconv(params map[string]string) (*GeoconvRet, error) {
	ret := &GeoconvRet{}
	err := api.call("geoconv", "", "v1", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (api *BaiduLbsApi) Addr2Loc(city, addr string) (lng, lat float64, err error) {
	params := make(map[string]string)
	params["address"] = addr
	params["city"] = city
	params["output"] = "json"

	res, err := api.Geocoding(params)
	if err != nil {
		return
	}

	result, ok := res.(*GeoEncoderRet)
	if !ok {
		err = errors.New("api return nil pointer.")
		return
	}

	if result.Status != 0 {
		err = errors.New("baidu lbs cloud response error, status not zero.")
		return
	}

	lng = result.Result.Location.Lng
	lat = result.Result.Location.Lat
	return
}

func (api *BaiduLbsApi) Loc2Addr(lng, lat float64) (result *GeocoderRet, err error) {
	params := make(map[string]string)
	params["coordtype"] = "bd09ll"
	params["location"] = fmt.Sprintf("%v,%v", lat, lng)
	params["output"] = "json"
	params["latest_admin"] = "1"

	res, err := api.Geocoding(params)
	if err != nil {
		return
	}

	result, ok := res.(*GeocoderRet)
	if !ok {
		err = errors.New("api return nil pointer.")
		return
	}

	if result.Status != 0 {
		err = errors.New("baidu lbs cloud response error, status not zero.")
		return
	}

	return
}

// Geocoding API包括地址解析和逆地址解析功能：
// 地理编码：即地址解析，由详细到街道的结构化地址得到百度经纬度信息
// 逆地理编码：即逆地址解析，由百度经纬度信息得到结构化地址信息
// http://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-geocoding
func (api *BaiduLbsApi) Geocoding(params map[string]string) (interface{}, error) {
	ret := interface{}(nil)
	if _, ok := params["address"]; ok {
		ret = &GeoEncoderRet{}
	} else {
		ret = &GeocoderRet{}
	}
	params["output"] = "json"
	err := api.call("geocoder", "", "v2", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 用于返回查询某个区域的某类POI数据，且提供单个POI的详情查询服务
// http://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-placeapi
func (api *BaiduLbsApi) PlaceSearch(params map[string]string) (interface{}, error) {
	ret := &PlaceSearchRet{}
	params["output"] = "json"
	err := api.call("place", "search", "v2", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 用于返回查询某个区域的某类POI数据，且提供单个POI的详情查询服务
// http://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-placeapi
func (api *BaiduLbsApi) PlaceDetail(params map[string]string) (interface{}, error) {
	ret := &PlaceDetailRet{}
	params["output"] = "json"
	err := api.call("place", "detail", "v2", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Place suggestion API 是一套提供匹配用户输入关键字辅助信息、提示服务的接口
// ，可返回json或xml格式的一组建议词条的数据。
func (api *BaiduLbsApi) Suggestion(params map[string]string) (interface{}, error) {
	ret := &SuggestionRet{}
	params["output"] = "json"
	err := api.call("place", "suggestion", "v2", params, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
