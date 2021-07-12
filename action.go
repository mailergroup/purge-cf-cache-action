package main

import (
	// "context"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

var (
	CFToken    string = os.Getenv("INPUT_CF_TOKEN")
	CFZoneName string = os.Getenv("INPUT_CF_ZONE_NAME")
)

func ErrCheck(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {
	api, err := cloudflare.NewWithAPIToken(CFToken)
	ErrCheck(err)

	ctx := context.Background()

	zoneName, err := api.ZoneIDByName(CFZoneName)
	ErrCheck(err)

	purgeURLs, present := os.LookupEnv("INPUT_CF_PURGE_HOSTS")
	if present {
		pcr, err := json.Marshal(cloudflare.PurgeCacheRequest{
			Hosts: ["https://blog.aorfanos.com", "https://aorfanos.com"],
		})
		ErrCheck(err)
		api.PurgeCache(ctx, zoneName, pcr)
	} else {
		api.PurgeEverything(ctx, zoneName)
	}

	zoneID, err := api.ZoneIDByName(CFZoneName)
	ErrCheck(err)

	fmt.Printf("%s", zoneID)
}
