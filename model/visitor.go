package model

type Visitor struct {
	Geo Geo			`json:"geo"`
	Browser Browser		`json:"browser"`
	Device Device		`json:"device"`
	Platform Platform	`json:"platform"`
	Personal Personal	`json:"personal"`
	Ip Ip			`json:"ip"`
}