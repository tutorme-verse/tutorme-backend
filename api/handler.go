package api

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"tutorme.com/repository"
)

func (s *Server) CreateOrganization(c *fiber.Ctx) error {
	// Parse json body into struct
	var request repository.CreateSchoolParams

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}

	ctx := context.Background()

	// Use the queries to enter new organization
	if err := s.db.CreateSchool(ctx, request); err != nil {
		return err
	}

	// initiate the db
	// Enter db details into the db details table
	// Return the db details
	return nil
}
