package api

import (
	"context"
	"encoding/json"
	"fmt"

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
	tursoDb, err := IssueTursoDatabase(dbName)

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
