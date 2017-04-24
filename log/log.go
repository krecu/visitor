package log

import (
	"github.com/robertkowalski/graylog-golang"
	"encoding/json"
	"os"
	"time"
	"visitor/conf"
	"fmt"
)

type Message struct {
	Version string		`json:"version"`
	Host string		`json:"host"`
	Timestamp int64		`json:"timestamp"`
	Facility string		`json:"facility"`
	ShortMessage string	`json:"short_message"`
	State string		`json:"state"`
	Duration float64	`json:"duration"`
}

func Notify(m Message) {

	conf := config.New()

	g := gelf.New(gelf.Config{
		GraylogPort:     conf.Logger.Port,
		GraylogHostname: conf.Logger.Host,
	})

	m.Host, _ = os.Hostname()
	m.Version =  "1.0"
	m.Timestamp =  time.Now().Unix()
	m.Facility =  "Visitor"

	message, _ := json.Marshal(m)
	g.Log(string(message))

	fmt.Println(m.ShortMessage)
}