package service

import (
	"context"
	db2 "github.com/supperdoggy/alert/alert_db/internal/models/db"
	"github.com/supperdoggy/alert/alert_db/internal/osmapi"
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alerts_structs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	NewAlert(context.Context, *alert_proto.NewAlertRequest) (*alert_proto.NewAlertResponse, error)
	GetAlertStream(*alert_proto.GetAlertStreamRequest, alert_proto.AlertDatabaseService_GetAlertStreamServer) error
	GetAllActiveAlerts(ctx context.Context, in *alert_proto.GetAllActiveAlertsRequest) (*alert_proto.GetAllActiveAlertsResponse, error)
}

type service struct {
	l *zap.Logger
	db db2.IDB
}

var (
	alertChan = make(chan *alert_proto.Alert)
)

func NewService(l *zap.Logger, db db2.IDB) IService {
	return &service{
		l:  l,
		db: db,
	}
}

func (s *service) NewAlert(ctx context.Context, request *alert_proto.NewAlertRequest) (*alert_proto.NewAlertResponse, error) {
	if request.GetTitle() == "" || request.GetBody() == "" {
		return nil, status.Error(codes.InvalidArgument, "no name and title")
	}

	// osm
	osmObj, err := osmapi.GetByCoordinates(request.GetLat(), request.GetLon())
	if err != nil {
		s.l.Error("error osmapi GetByCoordinates", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "error getting address")
	}
	address := osmObj.Address

	a := alerts_structs.AlertDB{
		Title: request.GetTitle(),
		Body: request.GetBody(),
		Lat: request.GetLat(),
		Lon: request.GetLon(),
		Address: address,
		ExpirationDateTime: request.GetExpirationDatetime(),
	}

	code, err := s.db.NewAlert(a)
	if err != nil {
		return nil, status.Error(code, err.Error())
	}

	resp := alert_proto.NewAlertResponse{
		Alert: &alert_proto.Alert{
			Lat:                request.GetLat(),
			Lon:                request.GetLon(),
			Address:            &alert_proto.Address{
				Country:      address.Country,
				CountyCode:   address.CountyCode,
				County:       address.County,
				State:        address.State,
				Town:         address.Town,
				Municipality: address.Municipality,
				District:     address.District,
				Postcode:     address.Postcode,
				Building:     address.Building,
				HouseNumber:  address.HouseNumber,
				Road:         address.Road,
				Suburb:       address.Suburb,
				Borough:      address.Borough,
			},
			Title:              a.Title,
			Body:               a.Body,
			ExpirationDatetime: a.ExpirationDateTime,
		},
	}

	// can block the stream god dammit
	// its just to make it work
	alertChan <- resp.GetAlert()
	return &resp, nil
}

func (s *service) GetAlertStream(req *alert_proto.GetAlertStreamRequest, stream alert_proto.AlertDatabaseService_GetAlertStreamServer) error {
	select {
	case a := <- alertChan:
		resp := alert_proto.GetAlertStreamResponse{Alert: a}
		if err := stream.Send(&resp); err != nil {
			return status.Error(codes.Canceled, "closed stream")
		}
	}
	return nil
}

func (s *service) GetAllActiveAlerts(ctx context.Context, in *alert_proto.GetAllActiveAlertsRequest) (*alert_proto.GetAllActiveAlertsResponse, error) {
	// token and user_id for the future?
	if in == nil || in.GetToken() == "" || in.GetUserId() == "" {
		s.l.Error("error GetAllActiveAlerts did not fill all fields", zap.Any("req", in))
		return nil, status.Error(codes.InvalidArgument, "fill all fields")
	}

	var result alert_proto.GetAllActiveAlertsResponse

	alerts, err := s.db.GetAllActiveAlerts()
	if err != nil {
		s.l.Error("error calling GetAllActiveAlerts", zap.Error(err))
		return nil, status.Error(codes.Internal, "db call error")
	}

	if alerts == nil || len(alerts) == 0 {
		return nil, status.Error(codes.NotFound, "not found any alerts")
	}

	result.Alerts = alerts
	return &result, nil
}

