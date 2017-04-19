package main

import (
	"github.com/aerospike/aerospike-client-go"
	"fmt"
	"runtime"
	//"github.com/aerospike/aerospike-client-go/types/particle_type"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	client, err := aerospike.NewClient("192.168.0.5", 3000)

	if (err != nil) {
		panic(err)
	}

	policy := aerospike.NewScanPolicy()
	policy.MaxRetries = 1
	policy.Priority = aerospike.LOW
	policy.ScanPercent = 1

	recordset, err := client.ScanAll(policy, "unstable", "default")

	if (err != nil) {
		panic(err)
	}

	recordCount := 0

	for {
		select {
		case rec := <-recordset.Records:
			if rec == nil {
				break
			}

			if rec.Expiration == 4294967295 {
				recordCount++
				if (recordCount % 10000) == 0 {
					fmt.Println("Records ", recordCount)
				}
			}
		case err := <-recordset.Errors:
			panic(err)
		}
	}

}