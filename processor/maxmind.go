package processor

import (
	"github.com/oschwald/geoip2-golang"
	"net"
	"visitor/model"
)

const (
	MAXMIND_DB_PATH = "./db/GeoLite2-City.mmdb"
)

type MaxMindProcessor struct {
}

var (
	maxMindClient *geoip2.Reader
	isMmConnect int
)

func (r *MaxMindProcessor) Process(param string) (model.Geo, error) {

	if isMmConnect != 1 {
		conn, err := geoip2.Open(MAXMIND_DB_PATH)
		if err != nil {
			return model.Geo{}, err
		}
		maxMindClient = conn
		isMmConnect = 1
	}


	ip := net.ParseIP(param)
	record, err := maxMindClient.City(ip)

	if err != nil {
		return model.Geo{}, err
	}

	return model.Geo{
		City:model.City{
			Name: record.City.Names["en"],
			Id: uint(record.City.GeoNameID),
		},
		Country:model.Country{
			Name: record.Country.Names["en"],
			Id: uint(record.Country.GeoNameID),
		},
	}, nil
}