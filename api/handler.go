package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tutorme.com/internal/repository"
	"tutorme.com/types"
)

func (s *Server) CreateOrganization(c *fiber.Ctx) error {
	var schoolParams repository.CreateSchoolParams

	if err := json.Unmarshal(c.Body(), &schoolParams); err != nil {
		return err
	}

	ctx := context.Background()
	err := CreateDNSRecord(ctx, schoolParams.Subdomain)

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
