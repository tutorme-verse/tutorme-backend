package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"tutorme.com/types"
	"tutorme.com/util"
)

func CreateDNSRecord(ctx context.Context, subdomain string) error {
	apiToken, err := util.ResolveEnv("CF_API_TOKEN")
	if err != nil {
		return err
	}

	slog.Info(apiToken)

	apiEmail, err := util.ResolveEnv("CF_API_EMAIL")
	if err != nil {
		return err
	}

	slog.Info(apiEmail)
	cfZone, err := util.ResolveEnv("CF_ZONE_ID")
	if err != nil {
		return err
	}

	slog.Info(cfZone)
	api, err := cloudflare.New(apiToken, apiEmail)
	if err != nil {
		return err
	}

	zoneId := cloudflare.ZoneIdentifier(cfZone)
	isProxied := true
	recordParams := cloudflare.CreateDNSRecordParams{
		Type:    "A",
		Name:    subdomain,
		Content: "159.203.82.246",
		TTL:     3600,
		Proxied: &isProxied,
	}

	_, err = api.CreateDNSRecord(ctx, zoneId, recordParams)
	if err != nil {
		return err
	}

	return nil
}

