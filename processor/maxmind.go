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

func (r *MaxMindProcessor) Process(param string) (model.Geo, error) {

	db, err := geoip2.Open(MAXMIND_DB_PATH)
	if err != nil {
		return model.Geo{}, err
	}
	defer db.Close()

	ip := net.ParseIP(param)
	record, err := db.City(ip)

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