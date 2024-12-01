package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tutorme.com/repository"
)

func (s *Server) CreateOrganization(c *fiber.Ctx) error {
	// Parse json body into struct
	var request repository.CreateSchoolParams

    fmt.Println(string(c.Body()))

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}

	ctx := context.Background()

	fmt.Println(request.SchoolName)
	fmt.Println(request.Subdomain)
	fmt.Println(request.Status)

	// Use the queries to enter new organization
	if err := s.db.CreateSchool(ctx, request); err != nil {
		return err
	}

	// initiate the db
	// Enter db details into the db details table
	// Return the db details
	return nil
}
