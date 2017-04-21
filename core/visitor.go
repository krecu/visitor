package core

import (
	"visitor/processor"
	"visitor/model"
	"time"
	"visitor/storage"
	"reflect"
	"fmt"
	"visitor/conf"
)

// основная структура
type Visitor struct {
	Id string		`json:"id"`
	Ua string		`json:"-"`
	Ip string		`json:"-"`
	Time string		`json:"gen_time"`
}

// Определяем данные по посетителю
func (v *Visitor) Identify () (model.Visitor, error) {

	total := time.Now()

	var visitor model.Visitor
	var geo model.Geo

	conf := config.New()

	client := storage.AeroSpike{Host:conf.Cache[0].Host, Port:conf.Cache[0].Port, Ns: conf.Ns, Set: conf.Set}
	record, err := client.Get(v.Id)

	// если нашли посетителя в кеше то берем от туда
	// иначе вычисляем его и записываем в кеш
	if record != nil || err != nil {
		visitor = v.UnMarshal(record)
	} else {

		// собираем данные по браузеру
		browser, err := new(processor.BrowsCapProcessor).Process(v.Ua)

		// собираем данные по гео
		geo, err = new(processor.SyPexGeoProcessor).Process(v.Ip)

		// если Сайпекс ничего невернул или произошла ошибка
		if err != nil {
			geo, err = new(processor.MaxMindProcessor).Process(v.Ip)
			if err != nil {
				return visitor, err
			}
		}

		// формируем модель Ip
		ip := model.Ip{V4:v.Ip}

		// формируем персональные данные
		personal, err := new(processor.PersonalProcessor).Process(v.Ua)
		if err != nil {
			return visitor, err
		}

		visitor = model.Visitor{
			Id: v.Id,
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
			City: geo.City,
			Country: geo.Country,
			Location: geo.Location,
			Browser: browser.Browser,
			Device: browser.Device,
			Platform: browser.Platform,
			Personal: personal,
			Ip: ip,
		}

		// формируем массив и записываем в аэроспайк
		err = client.Put(v.Marshal(visitor))

		if err != nil {
			return visitor, err
		}
	}

	// запоминаем потраченное время на определение инфо

	fmt.Printf("Time duration: %v \n", time.Now().Sub(total))

	return visitor, nil
}

// Упаковываем в массив записи аэроспайка
func (v *Visitor) Marshal (visitor model.Visitor) (map[string]interface{}) {

	record := make(map[string]interface{})

	// system
	record["id"] = visitor.Id
	record["created"] = visitor.Created
	record["updated"] = visitor.Updated

	// browser
	record["br_min"] = visitor.Browser.MinorVer
	record["br_maj"] = visitor.Browser.MajorVer
	record["br_type"] = visitor.Browser.Type
	record["br_ver"] = visitor.Browser.Version
	record["br_name"] = visitor.Browser.Name

	// device
	record["dc_name"] = visitor.Device.Name
	record["dc_type"] = visitor.Device.Type
	record["dc_brand"] = visitor.Device.Brand

	// platform
	record["pf_name"] = visitor.Platform.Name
	record["pf_short"] = visitor.Platform.Short
	record["pf_ver"] = visitor.Platform.Version
	record["pf_desc"] = visitor.Platform.Description
	record["pf_maker"] = visitor.Platform.Maker

	// city
	record["ct_name"] = visitor.City.Name
	record["ct_id"] = visitor.City.Id

	// country
	record["cn_name"] = visitor.Country.Name
	record["cn_id"] = visitor.Country.Id
	record["cn_iso"] = visitor.Country.Iso

	// location
	record["lc_lat"] = visitor.Location.Latitude
	record["lc_lon"] = visitor.Location.Longitude
	record["lc_tz"] = visitor.Location.TimeZone

	// personal
	record["pr_ua"] = visitor.Personal.Ua
	record["pr_fn"] = visitor.Personal.FirstName
	record["pr_ln"] = visitor.Personal.LastName
	record["pr_pa"] = visitor.Personal.Patronymic
	record["pr_age"] = visitor.Personal.Age
	record["pr_ge"] = visitor.Personal.Gender

	// ip
	record["ip_v4"] = visitor.Ip.V4
	record["ip_v6"] = visitor.Ip.V6

	return record
}

// Распаковываем запись с аероспайка
func (v *Visitor) UnMarshal (record map[string]interface{}) (model.Visitor) {

	return model.Visitor{
		Id: reflect.ValueOf(record["id"]).String(),
		Created: reflect.ValueOf(record["created"]).Int(),
		Updated: reflect.ValueOf(record["updated"]).Int(),
		Browser: model.Browser{
			MinorVer: reflect.ValueOf(record["br_min"]).String(),
			MajorVer: reflect.ValueOf(record["br_maj"]).String(),
			Type: reflect.ValueOf(record["br_type"]).String(),
			Version: reflect.ValueOf(record["br_ver"]).String(),
			Name: reflect.ValueOf(record["br_name"]).String(),
		},
		Device: model.Device{
			Brand: reflect.ValueOf(record["dc_brand"]).String(),
			Type: reflect.ValueOf(record["dc_type"]).String(),
			Name: reflect.ValueOf(record["dc_name"]).String(),
		},
		Platform: model.Platform{
			Name: reflect.ValueOf(record["pf_name"]).String(),
			Short: reflect.ValueOf(record["pf_short"]).String(),
			Version: reflect.ValueOf(record["pf_ver"]).String(),
			Description: reflect.ValueOf(record["pf_desc"]).String(),
			Maker: reflect.ValueOf(record["pf_maker"]).String(),
		},
		City: model.City{
			Name: reflect.ValueOf(record["ct_name"]).String(),
			Id:   111, //reflect.ValueOf(record["ct_id"]).Uint(),
		},
		Country: model.Country{
			Name: reflect.ValueOf(record["cn_name"]).String(),
			Id:   111, //reflect.ValueOf(record["cn_id"]).Int(),
			Iso:   reflect.ValueOf(record["cn_iso"]).String(),
		},
		Location: model.Location{
			Latitude: reflect.ValueOf(record["lc_lat"]).String(),
			Longitude:  reflect.ValueOf(record["lc_lon"]).String(),
			TimeZone:   reflect.ValueOf(record["lc_tz"]).String(),
		},
		Personal: model.Personal{
			Gender: reflect.ValueOf(record["pr_ge"]).String(),
			Age:  reflect.ValueOf(record["pr_age"]).String(),
			Patronymic:   reflect.ValueOf(record["pr_pa"]).String(),
			LastName:   reflect.ValueOf(record["pr_ln"]).String(),
			FirstName:   reflect.ValueOf(record["pr_fn"]).String(),
			Ua:   reflect.ValueOf(record["pr_ua"]).String(),
		},
		Ip: model.Ip{
			V4: reflect.ValueOf(record["ip_v4"]).String(),
			V6:  reflect.ValueOf(record["ip_v6"]).String(),
		},
	}

}