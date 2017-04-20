package model

type Platform struct {
	Name string		`json:"name"`
	Short string		`json:"short"`
	Version string		`json:"version"`
	Description string	`json:"description"`
	Maker string		`json:"maker"`
}
