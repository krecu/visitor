package processor

import (
	"github.com/mirrr/go-sypexgeo"
	"visitor/model"
	"errors"
)

const (
	SYPEX_DB_PATH = "./db/SxGeoCity.dat"
)

type SyPexGeoProcessor struct {
}

// Вспомогательная функция конвертации, таиственных interface{}
func idConvert(i interface{}) uint {
	switch v := i.(type) {
	case int, int8, int16, int32, int64:
		return uint(v.(int))
	case uint, uint8:
		return uint(v.(uint8))
	case uint16, uint32, uint64:
		return uint(v.(uint32))
	}
	return 0
}

func (r *SyPexGeoProcessor) Process(param string) (model.Geo, error) {

	geo := sypexgeo.New(SYPEX_DB_PATH)

	record, err := geo.GetCityFull(param)

	if err != nil {
		return model.Geo{}, err
	}

	var country = record["country"].(map[string]interface{})
	var city = record["city"].(map[string]interface{})

	// Если не определили страну
	if country["name_en"] == "" {
		return model.Geo{}, errors.New("Could not determine the country by IP: " + param)
	}

	// Если не определили город
	if city["name_en"] == "" {
		return model.Geo{}, errors.New("Could not determine the city by IP: " + param)
	}

	return model.Geo{
		City:model.City{
			Name: city["name_en"].(string),
			Id: idConvert(city["id"]),
		},
		Country:model.Country{
			Name: country["name_en"].(string),
			Id: idConvert(country["id"]),
		},
	}, nil
}