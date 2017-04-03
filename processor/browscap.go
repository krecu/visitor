package processor

import (
	"visitor/model"
	"fmt"
	"github.com/digitalcrab/browscap_go"
)

const (
	BROWSCAP_DB_PATH = "./db/browscap.ini"
)

type BrowsCapProcessor struct {
}

func (r *BrowsCapProcessor) Process(param string) (model.Browser, error) {

	if err := browscap_go.InitBrowsCap(BROWSCAP_DB_PATH, false); err != nil {
		panic(err)
	}

	browser, ok := browscap_go.GetBrowser(param)
	if !ok || browser == nil {
		panic("Browser not found")
	} else {
		fmt.Printf("Browser = %s [%s] v%s\n", browser.Browser, browser.BrowserType, browser.BrowserVersion)
		fmt.Printf("Platform = %s v%s\n", browser.Platform, browser.PlatformVersion)
		fmt.Printf("Device = %s [%s] %s\n", browser.DeviceName, browser.DeviceType, browser.DeviceBrand)
		fmt.Printf("IsCrawler = %t\n", browser.IsCrawler())
		fmt.Printf("IsMobile = %t\n", browser.IsMobile())
	}

	return model.Browser{
		Type:"sdfsdf",
	}, nil
}