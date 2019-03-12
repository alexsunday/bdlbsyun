package baidumap

import "strconv"

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (point *Point) ToString() (string) {
	return strconv.FormatFloat(point.X, 'g', 10, 32) + "," +
	strconv.FormatFloat(point.Y, 'g', 10,32) 
}

type GeoconvRet struct {
	Status int `json:"status"`
	Result []Point `json:"result"`
}


type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type AddressComponent struct {
	Country string `json:"country"`
	CountryCode int `json:"country_code"`
	Province string `json:"province"`
	City string `json:"city"`
	District string `json:"district"`
	Adcode string `json:"adcode"`
	Street string `json:"street"`
	Street_number string `json:"street_number"`
	Direction string `json:"direction"`
	Distance string `json:"distance"`
}

type Pois struct {
	Addr string `json:"addr"`
	Cp string `json:"cp"`
	Direction string `json:"direction"`
	Name string `json:"name"`
	PoiType string `json:"poiType"`
	Point Point `json:"point"`
	Tag string `json:"tag"`
	Tel string `json:"tel"`
	Uid string `json:"uid"`
	Zip string `json:"zip"`
}

type PoisRegions struct {
	DirectionDesc string `json:"direction_desc"`
	Name string `json:"name"`
	Tag string `json:"tag"`
}

type GeocoderResult struct {
	Location Location `json:"location"`
	FormattedAddress string `json:"formatted_address"`
	Business string `json:"business"`
	AddressComponent AddressComponent `json:"addressComponent"`
	Pois []Pois `json:"pois"`
	PoisRegions []PoisRegions `json:"poisRegions"`
	SematicDescription string `json:"sematic_description"`
	CityCode int `json:"cityCode"`
}

type GeoEncoderResult struct {
	Location Location `json:"location"`
	Precise int `json:"precise"`
	Confidence int `json:"confidence"`
	Level string `json:"level"`
}

type GeocoderRet struct {
	Status int `json:"status"`
	Result GeocoderResult `json:"result"`
}

type GeoEncoderRet struct {
	Status int `json:"status"`
	Result GeoEncoderResult `json:"result"`
}

type DetailInfo struct {
	Tag string `json:"tag"`
	DetailUrl string `json:"detail_url"`
	Type string `json:"type"`
	Price string `json:"price"`
	OverallRating string `json:"overall_rating"`
	TasteRating string `json:"taste_rating"`
	ServiceRating string `json:"service_rating"`
	EnvionmentRating string `json:"environment_rating"`
	ImageNum string `json:"image_num"`
	CommentNum string `json:"comment_num"`
	ShopHours string `json:"shop_hours"`
	Atmosphere string `json:"atmosphere"`
	FeaturedService string `json:"featured_service"`
	Recommendation string `json:"recommendation"`
	Description string `json:"description"`
	Distance string `json:"distance"`
	FacilityRating  string `json:"facility_rating"`
	HygieneRating string `json:"hygiene_rating"`
	TechnologyRating string `json:"technology_rating"`
	GrouponNum string `json:"groupon_num"`
	DiscountNum string `json:"discount_num"`
	FavoriteNum string `json:"favorite_num"`
	CheckinNum string `json:"checkin_num"`

}

type PlaceResult struct {
	Name string `json:"name"`
	Location Location `json:"location"`
	Address string `json:"address"`
	Telephone string `json:"telephone"`
	Detail int `json:"detail"`
	Uid string `json:"uid"`
	DetailInfo DetailInfo `json:"detail_info"`
	StreetId string `json:"street_id"`
}

type PlaceSearchRet struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Total int `json:"total"`
	Result []PlaceResult `json:"results"`
}

type PlaceDetailRet struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Result PlaceResult `json:"result"`
}

type SuggestionResult struct {
	name string `json:"name"`
	Location Location `json:"location"`
	Uid string `json:"uid"`
	City string `json:"city"`
	District string `json:"district"`
}

type SuggestionRet struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Result []SuggestionResult `json:"result"`
}
