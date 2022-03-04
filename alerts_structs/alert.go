package alerts_structs

type AlertDB struct {
	ID                 string `json:"id" mgo:"_id"`
	Title              string `json:"title"`
	Body               string `json:"body"`
	Tag                string `json:"tag"`
	Address            string `json:"address"`
	ExpirationDateTime int64  `json:"expiration_date_time"`
}

type Alert struct {
	Title              string `json:"title"`
	Body               string `json:"body"`
	Tag                string `json:"tag"`
	Address            string `json:"address"`
	ExpirationDateTime int64  `json:"expiration_date_time"`
}
