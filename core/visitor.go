package core

import (
	"visitor/processor"
	"visitor/model"
	"time"
)

// основная структура
type Visitor struct {
	Id string		`json:"id"`
	Ua string		`json:"-"`
	Ip string		`json:"-"`
	Time string		`json:"gen_time"`
	Info model.Visitor 	`json:"information"`
}

// Определяем данные по посетителю
func (v *Visitor) Identify () (error) {

	t := time.Now()

	var geo model.Geo

	browser, err := new(processor.BrowsCapProcessor).Process(v.Ua)
	geo, err = new(processor.SyPexGeoProcessor).Process(v.Ip)

	ip := model.Ip{V4:v.Ip}

	if err != nil {
		geo, err = new(processor.MaxMindProcessor).Process(v.Ip)
		if err != nil {
			return err
		}
	}

	personal, err := new(processor.PersonalProcessor).Process(v.Ua)
	if err != nil {
		return err
	}

	v.Info = model.Visitor{
		Geo: geo,
		Browser: browser.Browser,
		Device: browser.Device,
		Platform: browser.Platform,
		Personal: personal,
		Ip: ip,
	}

	v.Time = time.Now().Sub(t).String()

	return nil
}
