package model

type Visitor struct {
	Id string		`json:"id"`
	Created int64		`json:"created"`
	Updated int64		`json:"updated"`
	City City		`json:"city"`
	Country Country		`json:"country"`
	Location Location	`json:"location"`
	Postal Postal		`json:"postal"`
	Browser Browser		`json:"browser"`
	Device Device		`json:"device"`
	Platform Platform	`json:"platform"`
	Personal Personal	`json:"personal"`
	Ip Ip			`json:"ip"`
}