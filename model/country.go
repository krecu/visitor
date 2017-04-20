package model

type Country struct {
	Name string	`json:"name"`
	Id uint		`json:"geoname_id"`
	Iso string	`json:"iso_code"`
}
