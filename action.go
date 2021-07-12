package main

import (
	// "context"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

var (
	CFToken       string = os.Getenv("INPUT_CF_TOKEN")
	CFZoneName    string = os.Getenv("INPUT_CF_ZONE_NAME")
	purgeHosts    string = os.Getenv("INPUT_CF_PURGE_HOSTS")
	purgeURLs     string = os.Getenv("INPUT_CF_PURGE_URLS")
	purgePrefixes string = os.Getenv("INPUT_CF_PURGE_PREFIXES")
)

type CFPurgeWithPrefixes struct {
	Prefixes []string `json:"prefixes,omitempty"`
}

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

	// purgeHosts, present := os.LookupEnv("INPUT_CF_PURGE_HOSTS")
	// purgeURLs, presentURL := os.LookupEnv("INPUT_CF_PURGE_URLS")
	// _ = purgeURLs //omit 'declared but not used' err
	// _ = purgeHosts
	if len(purgeHosts) > 1 {
		pcr := cloudflare.PurgeCacheRequest{Hosts: strings.Split(purgeHosts, ",")}
		purge, err := api.PurgeCache(ctx, zoneName, pcr)
		ErrCheck(err)
		fmt.Printf("Success: %t", purge.Response.Success)
	} else if len(purgeURLs) > 1 {
		pcr := cloudflare.PurgeCacheRequest{Files: strings.Split(purgeURLs, ",")}
		purge, err := api.PurgeCache(ctx, zoneName, pcr)
		ErrCheck(err)
		fmt.Printf("Success: %t", purge.Response.Success)
	} else if len(purgePrefixes) > 1 {
		purgePrefData := CFPurgeWithPrefixes{
			Prefixes: strings.Split(purgePrefixes, ","),
		}
		purgeJSON, errj := json.Marshal(purgePrefData)
		ErrCheck(errj)
		purge, err := http.Post("https://api.cloudflare.com/client/v4/"+zoneName+"/purge_cache",
			"application/json",
			bytes.NewBuffer(purgeJSON))
		ErrCheck(err)
		fmt.Printf("Success status code: %d", purge.StatusCode)
	} else {
		purge, err := api.PurgeEverything(ctx, zoneName)
		ErrCheck(err)
		fmt.Printf("Success: %t", purge.Response.Success)
	}
}
