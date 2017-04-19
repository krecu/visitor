package main

import (
	"fmt"
	"visitor/processor"
	"visitor/model"
	"time"
	"strconv"
	"math/rand"
)

func main() {

	t := time.Now()
	for i := 0; i < 1000; i++ {

		ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36"
		ip := strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222))

		browser, err := new(processor.BrowsCapProcessor).Process(ua)
		geo, err := new(processor.SyPexGeoProcessor).Process(ip)

		if err != nil {
			geo, err = new(processor.MaxMindProcessor).Process(ip)
			if err != nil {
				panic(err)
			}
		}


		visitor := model.Visitor{
			Geo: geo,
			Browser: browser.Browser,
			Device: browser.Device,
			Platform: browser.Platform,
		}

		_ = visitor
		fmt.Println(visitor)
	}

	fmt.Printf("time %v.\n", time.Now().Sub(t))

}