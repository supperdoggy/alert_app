package alerts_structs

import "github.com/supperdoggy/alert/alert_proto"

type AlertDB struct {
	ID                 string               `json:"id" mgo:"_id"`
	Title              string               `json:"title"`
	Body               string               `json:"body"`
	Lat                float32             `json:"lat"`
	Lon                float32              `json:"lon"`
	Address            OpenStreetMapAddress `json:"address"`
	ExpirationDateTime int64                `json:"expiration_date_time"`
	CreationDateTime   int64                `json:"creation_date_time"`
}

type Alert struct {
	Title              string               `json:"title"`
	Body               string               `json:"body"`
	Lat                float32             `json:"lat"`
	Lon                float32              `json:"lon"`
	Address            OpenStreetMapAddress `json:"address"`
	ExpirationDateTime int64                `json:"expiration_date_time"`
}

type ClientGetAlertStream struct {
	Stream alert_proto.NotificationService_GetAlertsServer
	UserID string
	Token  string

	CloseSignal chan struct{}
}
