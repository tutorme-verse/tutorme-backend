package types

import "tutorme.com/internal/repository"

type TursoDatabase struct {
	Database struct {
		DbID     string `json:"DbId"`
		Hostname string `json:"Hostname"`
		Name     string `json:"Name"`
	} `json:"database"`
}

type CreateOrganizationResponse struct {
	School         repository.School         `json:"school"`
	DatabaseDetail repository.CreateDatabaseDetailRow `json:"db_detail"`
}
