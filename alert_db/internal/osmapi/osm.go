package osmapi

import (
	"encoding/json"
	"fmt"
	"github.com/supperdoggy/alert/alerts_structs"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)


var logger, _ = zap.NewDevelopment()

func GetByCoordinates(lat, lon float32) (resp alerts_structs.OpenStreetMapObj, err error) {
	var link = fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%v&lon=%v&zoom=10&format=json", lat, lon)
	req, err := http.Get(link)
	if err != nil {
		logger.Error("error making req", zap.Error(err))
		return
	}

	bytesdata, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("error reqding body", zap.Error(err))
		return
	}

	if err = json.Unmarshal(bytesdata, &resp); err != nil {
		logger.Error("error unmarshalling req", zap.Error(err), zap.Any("req", string(bytesdata)))
		return
	}

	return
}

func GetByAddress(address string, limit int) (result []alerts_structs.OpenStreetMapObj, err error) {
	var link = fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%v&format=json&polygon=1&addressdetails=1&limit=%v", address, limit)
	req, err := http.Get(link)
	if err != nil {
		logger.Error("error making req", zap.Error(err))
		return
	}

	bytesdata, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("error reqding body", zap.Error(err))
		return
	}

	if err = json.Unmarshal(bytesdata, &result); err != nil {
		logger.Error("error unmarshalling req", zap.Error(err), zap.Any("req", string(bytesdata)))
		return
	}

	return
}
