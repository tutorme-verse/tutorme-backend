package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gofiber/fiber/v2"
	"tutorme.com/internal/repository"
	"tutorme.com/types"
	"tutorme.com/util"
)

func (s *Server) CreateOrganization(c *fiber.Ctx) error {
	var schoolParams repository.CreateSchoolParams

	if err := json.Unmarshal(c.Body(), &schoolParams); err != nil {
		return err
	}

	ctx := context.Background()
	// err := CreateDNSRecord(ctx, schoolParams.Subdomain)
	apiToken, err := util.ResolveEnv("CF_API_TOKEN")
	if err != nil {
		return err
	}

	s.logger.Info(apiToken)

	apiEmail, err := util.ResolveEnv("CF_API_EMAIL")
	if err != nil {
		return err
	}

	s.logger.Info(apiEmail)
	cfZone, err := util.ResolveEnv("CF_ZONE_ID")
	if err != nil {
		return err
	}

	s.logger.Info(cfZone)
	api, err := cloudflare.New(apiToken, apiEmail)
	if err != nil {
		return err
	}

	zoneId := cloudflare.ZoneIdentifier(cfZone)
	isProxied := true
	recordParams := cloudflare.CreateDNSRecordParams{
		Type:    "A",
		Name:    schoolParams.Subdomain,
		Content: "159.203.82.246",
		TTL:     3600,
		Proxied: &isProxied,
	}

	_, err = api.CreateDNSRecord(ctx, zoneId, recordParams)
	if err != nil {
		return err
	}

	dbName := fmt.Sprintf("tutorme-%s", schoolParams.Subdomain)
	body := []byte(fmt.Sprintf(`{"name": "%s", "group": "default"}`, dbName))
	fmt.Println(string(body))

	reqBody := bytes.NewReader(body)

	client := &http.Client{
		Timeout: time.Second * 6,
	}

	requestUrl := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", os.Getenv("TURSO_ORGANIZATION_SLUG"))
	tursoToken, err := util.ResolveEnv("TURSO_API_TOKEN")
	if err != nil {
		return err
	}
	bearer := "Bearer " + tursoToken

	req, err := http.NewRequest(http.MethodPost, requestUrl, reqBody)

	req.Header.Set("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	jsonResp := make([]byte, 1024)
	n, err := resp.Body.Read(jsonResp)

	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	var tursoDb types.TursoDatabase

	if err := json.Unmarshal(jsonResp[:n], &tursoDb); err != nil {
		return err
	}

	school, err := s.db.CreateSchool(ctx, schoolParams)
	if err != nil {
		return err
	}

	var dbDetailsParams repository.CreateDatabaseDetailParams = repository.CreateDatabaseDetailParams{
		ForeignDatabaseID: tursoDb.Database.DbID,
		SchoolID:          school.SchoolID,
		DatabaseName:      tursoDb.Database.Name,
		ConnectionUri:     tursoDb.Database.Hostname,
	}
	dbDetails, err := s.db.CreateDatabaseDetail(ctx, dbDetailsParams)
	if err != nil {
		return err
	}

	var finalResp types.CreateOrganizationResponse = types.CreateOrganizationResponse{
		School:         school,
		DatabaseDetail: dbDetails,
	}

	return c.JSON(finalResp)
}

func IssueTursoDatabase(dbName string) (types.TursoDatabase, error) {
	body := []byte(fmt.Sprintf(`{"name": "%s", "group": "default"}`, dbName))
	fmt.Println(string(body))

	reqBody := bytes.NewReader(body)

	client := &http.Client{
		Timeout: time.Second * 6,
	}

	requestUrl := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", os.Getenv("TURSO_ORGANIZATION_SLUG"))
	bearer := "Bearer " + os.Getenv("TURSO_API_TOKEN")

	req, err := http.NewRequest(http.MethodPost, requestUrl, reqBody)

	req.Header.Set("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return types.TursoDatabase{}, err
	}

	jsonResp := make([]byte, 1024)
	n, err := resp.Body.Read(jsonResp)

	if err != nil {
		if err != io.EOF {
			return types.TursoDatabase{}, err
		}
	}

	var tursoDbResp types.TursoDatabase

	if err := json.Unmarshal(jsonResp[:n], &tursoDbResp); err != nil {
		return types.TursoDatabase{}, err
	}

	return tursoDbResp, nil
}
