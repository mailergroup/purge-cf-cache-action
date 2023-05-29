package main

import (
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

func main() {
	api, err := cloudflare.NewWithAPIToken(CFToken)
	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()

	zoneName, err := api.ZoneIDByName(CFZoneName)
	if err != nil {
		log.Println(err)
	}

	if len(purgeHosts) > 1 {
		pcr := cloudflare.PurgeCacheRequest{Hosts: strings.Split(purgeHosts, ",")}
		purge, err := api.PurgeCache(ctx, zoneName, pcr)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Success: %t", purge.Response.Success)
	} else if len(purgeURLs) > 1 {
		pcr := cloudflare.PurgeCacheRequest{Files: strings.Split(purgeURLs, ",")}
		purge, err := api.PurgeCache(ctx, zoneName, pcr)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Success: %t", purge.Response.Success)
	} else if len(purgePrefixes) > 1 {
		client := &http.Client{}
		purgePrefData := CFPurgeWithPrefixes{
			Prefixes: strings.Split(purgePrefixes, ","),
		}
		purgeJSON, err := json.Marshal(purgePrefData)
		if err != nil {
			log.Println(err)
		}
		purgeReq, err := http.NewRequest("POST", "https://api.cloudflare.com/client/v4/zones/"+zoneName+"/purge_cache",
			bytes.NewBuffer(purgeJSON))
		if err != nil {
			log.Println(err)
		}
		purgeReq.Header.Set("Authorization", "Bearer "+CFToken)
		purgeReq.Header.Set("Content-Type", "application/json")
		purgeResp, err := client.Do(purgeReq)
		if err != nil {
			log.Println(err)
		}
		defer purgeResp.Body.Close()
		fmt.Printf("Success status code: %s", purgeResp.Status)
	} else {
		purge, err := api.PurgeEverything(ctx, zoneName)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Success: %t", purge.Response.Success)
	}
}
