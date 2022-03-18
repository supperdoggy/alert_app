package alerts_structs

type OpenStreetMapObj struct {
	Address     OpenStreetMapAddress `json:"address"`
	BoundingBox []string             `json:"boundingbox"`
	DisplayName string               `json:"display_name"`
	Lat         string               `json:"lattitute"`
	Licence     string               `json:"licence"`
	Lon         string               `json:"lon"`
	OsmID       int                  `json:"osm_id"`
	OsmType     string               `json:"osm_type"`
	PlaceID     int                  `json:"place_id"`
}

type OpenStreetMapAddress struct {
	Country      string `json:"country"`
	CountyCode   string `json:"county_code"`
	County       string `json:"county"`
	State        string `json:"state"`
	Town         string `json:"town"`
	Municipality string `json:"municipality"`
	District     string `json:"district"`
	Postcode     string `json:"postcode"`
	Building     string `json:"building"`
	HouseNumber  string `json:"house_number"`
	Road         string `json:"road"`
	Suburb       string `json:"suburb"`
	Borough      string `json:"borough"`
}

