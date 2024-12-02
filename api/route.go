package api

import "github.com/gofiber/fiber/v2"

func (s *Server) register() {
	v1 := s.server.Group("/v1")

	v1.Post("/org/create", s.CreateOrganization)

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello World! How are your?"))
	})
}
