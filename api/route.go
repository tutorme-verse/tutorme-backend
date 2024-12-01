package api

func (s *Server) register() {
	v1 := s.server.Group("/v1")

	v1.Post("/org/create", s.CreateOrganization)
}
