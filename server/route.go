package server

func (s *Server) routes() {
	s.e.POST("/add-watermark", s.handler.AddWatermark)
}
