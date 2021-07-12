package main

import (
	// "context"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	_ = purgeURLs //omit 'declared but not used' err
	if present {
		pcr := cloudflare.PurgeCacheRequest{Hosts: strings.Split(purgeURLs, ",")}
		purge, err := api.PurgeCache(ctx, zoneName, pcr)
		ErrCheck(err)
		fmt.Printf("Success: %t", purge.Response.Success)
	} else {
		purge, err := api.PurgeEverything(ctx, zoneName)
		ErrCheck(err)
		fmt.Printf("Success: %t", purge.Response.Success)
	}
}
