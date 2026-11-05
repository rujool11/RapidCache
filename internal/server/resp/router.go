package respserver

import (
	"net"
	"strings"
)

func (s* Server) getRoutes() map[string]handler{
	return map[string]handler{
		"PING": s.handlePing,
	}
}

func (s *Server) routeRequest(conn net.Conn, args []string) error {
	command := strings.ToUpper(args[0])
	routes := s.getRoutes()

	handler, exists := routes[command]
	if !exists {
		// fallback if unkonwn command
		return s.handleUnknownCommand(conn, command)
	}

	return handler(conn, args)

}
