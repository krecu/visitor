package processor

import (
	"visitor/model"
	"github.com/digitalcrab/browscap_go"
	"path/filepath"
	"os"
)

type BrowsCap struct {
	Browser model.Browser
	Platform model.Platform
	Device model.Device
}

type BrowsCapProcessor struct {
}

func (r *BrowsCapProcessor) Process(param string) (BrowsCap, error) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	if err := browscap_go.InitBrowsCap(dir + "/db/browscap.ini", false); err != nil {
		panic(err)
	}

	browser, ok := browscap_go.GetBrowser(param)
	if !ok || browser == nil {
		panic("Browser not found")
	}

	browscap := BrowsCap{
		Browser: model.Browser{
			Name: browser.Browser,
			MajorVer: browser.BrowserMajorVer,
			MinorVer: browser.BrowserMinorVer,
			Version: browser.BrowserVersion,
			Type: browser.BrowserType,
		},
		Platform: model.Platform{
			Name: browser.Platform,
			Version: browser.PlatformVersion,
			Short: browser.PlatformShort,
		},
		Device: model.Device{
			Name: browser.DeviceName,
			Brand: browser.DeviceBrand,
			Type: browser.DeviceType,
		},
	}

	return browscap, nil
}