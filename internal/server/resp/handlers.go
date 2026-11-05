package respserver

import (
	"fmt"
	"net"
)

type handler func(conn net.Conn, args []string) error

func (s *Server) handlePing(conn net.Conn, args []string) error {
	_, err := conn.Write([]byte("+PONG\r\n"))
	s.logger.Debug("PING command called")
	return err
}

func (s *Server) handleUnknownCommand(conn net.Conn, cmd string) error {
	response := fmt.Appendf(nil, "-ERR unknown command %s\r\n", cmd)
	_, err := conn.Write((response))
	s.logger.Debug("Unknown command called", "cmd", cmd)
	return err
}
