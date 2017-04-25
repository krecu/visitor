package processor

import (
	"github.com/mirrr/go-sypexgeo"
	"visitor/model"
	"path/filepath"
	"os"
	"github.com/oschwald/geoip2-golang"
	"net"
	"errors"
)

type GeoProcessor struct {
}

var (
	spClient sypexgeo.SxGEO
	spConnect int
	mmClient *geoip2.Reader
	mmConnect int

	spActive int
	mmActive int

	geo model.Geo
	city model.City
	country model.Country
	region model.Region
	location model.Location
	postal model.Postal
)

// Вспомогательная функция конвертации, таиственных interface{}
func numericToUint(i interface{}) uint {
	switch v := i.(type) {
	case int, int8, int16, int32, int64:
		return uint(v.(int))
	case uint, uint8:
		return uint(v.(uint8))
	case uint16, uint32, uint64:
		return uint(v.(uint32))
	case float64:
		return uint(v)
	case float32:
		return uint(v)
	}

	return 0
}

//
func (r *GeoProcessor) Process(param string) (model.Geo, error) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// прогружаем SxGeo базу
	if spConnect != 1 {
		spClient = sypexgeo.New(dir + "/db/SxGeoCity.dat")
		spConnect = 1
	}

	mmClient, err = geoip2.Open(dir + "/db/GeoLite2-City.mmdb")
	if err != nil {
		mmConnect = 0
	} else {
		mmConnect = 1
	}

	spRecord, err := spClient.GetCityFull(param)

	// отмечаем что провайдер Sp загрузился
	if err != nil || spRecord == nil {
		spActive = 0
	} else {
		spActive = 1
	}

	// отмечаем что провайдер MM загрузился
	ip := net.ParseIP(param)
	mmRecord, err := mmClient.City(ip)
	if err != nil || mmRecord == nil {
		mmActive = 0
	} else {
		mmActive = 1
	}

	// приоритет всегда получать ГЕО от SyPex
	if spActive == 1 {

		var countryRaw = spRecord["country"].(map[string]interface{})
		var cityRaw = spRecord["city"].(map[string]interface{})
		var regionRaw = spRecord["region"].(map[string]interface{})

		city = model.City{
			Name: cityRaw["name_en"].(string),
			Id: numericToUint(cityRaw["id"]),
		}

		country = model.Country{
			Name: countryRaw["name_en"].(string),
			Id: numericToUint(countryRaw["id"]),
			Iso: countryRaw["iso"].(string),
		}

		region = model.Region{
			Name: regionRaw["name_en"].(string),
			Id: numericToUint(regionRaw["id"]),
		}

		location = model.Location{
			Latitude: cityRaw["lat"].(float32),
			Longitude: cityRaw["lon"].(float32),
		}
	}

	// если MM прогрузилось и определилось
	if mmActive == 1 {

		// если собрали данные из SP
		if spActive == 1 {
			// проверяем что определился один и тот же город
			// и при успехе дополняем данными из MM
			if city.Name == mmRecord.City.Names["en"] {

				postal = model.Postal{
					Code: mmRecord.Postal.Code,
				}

				// SP дает какуюто хрень но точно не GEONAME_ID
				// поэтому меняем на свои
				city.Id = mmRecord.City.GeoNameID
				country.Id = mmRecord.Country.GeoNameID

				location.TimeZone = mmRecord.Location.TimeZone

			}
		} else {
			city = model.City{
				Name: mmRecord.City.Names["en"],
				Id: mmRecord.City.GeoNameID,
			}

			country = model.Country{
				Name: mmRecord.Country.Names["en"],
				Id: mmRecord.Country.GeoNameID,
			}

			postal = model.Postal{
				Code: mmRecord.Postal.Code,
			}

			location = model.Location{
				Latitude: float32(mmRecord.Location.Latitude),
				Longitude: float32(mmRecord.Location.Longitude),
				TimeZone: mmRecord.Location.TimeZone,
			}
		}

		mmClient.Close()
	}

	// если несняли данные ни с одной системы то геймовер
	if mmActive == 0 && spActive == 0 {
		return model.Geo{}, errors.New("Could not determine the geo by IP: " + param)
	}

	geo = model.Geo{
		City: city,
		Country: country,
		Region: region,
		Location: location,
		Postal:postal,
	}

	return geo, nil
}
