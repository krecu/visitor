package model

type Browser struct {
	Name string		`json:"name"`
	Type string		`json:"type"`
	Version string		`json:"version"`
	MajorVer string		`json:"majorver"`
	MinorVer string		`json:"minorver"`
}
