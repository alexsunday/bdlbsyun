package bdlbsyun

import (
	"fmt"
	"testing"
)

func Test_geoconv(t *testing.T) {
	api := NewBaiduLbsApi("0T3dbQsKaAxfsZfskEDm0OvMVANgTqmf")
	lng, lat, err := api.GeoConvert(113.3532180, 23.1407146, GEOCONV_F_WGS84, GEOCONV_F_BD_LL)
	if err != nil {
		return
	}
	fmt.Println(lng, lat)
}

func Test_loc2addr(t *testing.T) {
	api := NewBaiduLbsApi("0T3dbQsKaAxfsZfskEDm0OvMVANgTqmf")
	result, err := api.Loc2Addr(113.36522341182265, 23.143990588355013)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func Test_addr2loc(t *testing.T) {
	api := NewBaiduLbsApi("0T3dbQsKaAxfsZfskEDm0OvMVANgTqmf")
	lng, lat, err := api.Addr2Loc("", "北京市海淀区上地十街10号")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(lng, lat)
}
