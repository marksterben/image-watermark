package server

func (s *Server) routes() {
	s.e.POST("/add-watermark/:folder", s.handler.AddWatermark)
}
