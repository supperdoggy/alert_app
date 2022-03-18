package db

import (
	"github.com/pkg/errors"
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alerts_structs"
	"github.com/u2takey/go-utils/rand"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"gopkg.in/mgo.v2"
	"time"
)

type IDB interface {
	NewAlert(alert alerts_structs.AlertDB) (codes.Code, error)
	GetAllActiveAlerts() (alerts []*alert_proto.Alert, err error)
}

type db struct {
	l *zap.Logger

	// TODO history action records
	dbname           string
	alertsCollection *mgo.Collection
}

// for the queries
type obj map[string]interface{}

var (
	ErrFillAlertFields = errors.New("fill all alert fields")
)

func NewDB(l *zap.Logger, url, dbName, alerts string) (IDB, error) {
	client, err := mgo.Dial(url)
	if err != nil {
		l.Error("error dialing with db", zap.Error(err))
		return nil, err
	}

	database := client.DB(dbName)
	result := db{
		l:                l,
		dbname:           dbName,
		alertsCollection: database.C(alerts),
	}

	return &result, nil
}

func (d *db) NewAlert(alert alerts_structs.AlertDB) (codes.Code, error) {
	if alert.Body == "" || alert.Title == "" ||
		alert.ExpirationDateTime == 0 {
		return codes.InvalidArgument, ErrFillAlertFields
	}

	rand.Seed(time.Now().UnixNano())
	alert.ID = rand.String(24)
	alert.CreationDateTime = time.Now().Unix()

	err := d.alertsCollection.Insert(alert)
	if mgo.IsDup(err) {
		return d.NewAlert(alert)
	}

	if err != nil {
		return codes.Internal, err
	}
	return codes.OK, nil
}

func (d *db) GetAllActiveAlerts() (alerts []*alert_proto.Alert, err error) {
	t := time.Now().Unix()
	// getting alerts which are not expired
	err = d.alertsCollection.Find(obj{"expirationdatetime": obj{"$gt": t}}).All(&alerts)
	return
}
