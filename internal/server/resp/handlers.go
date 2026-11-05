package respserver

import (
	"fmt"
	"log/slog"
	"net"
)

type handler func(conn net.Conn, args []string) error

func handlePing(conn net.Conn, args []string) error {
	_, err := conn.Write([]byte("+PONG\r\n"))
	slog.Debug("PING command called")
	return err
}

func handleUnknownCommand(conn net.Conn, cmd string) error {
	response := fmt.Appendf(nil, "-ERR unknown command %s\r\n", cmd)
	_, err := conn.Write((response))
	slog.Debug("Unknown command called", "cmd", cmd)
	return err
}
