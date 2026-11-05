package respserver

import (
	"net"
	"strings"
)

var router = map[string]handler{
	"PING": handlePing,
}

func routeRequest(conn net.Conn, args []string) error {
	command := strings.ToUpper(args[0])

	handler, exists := router[command]
	if !exists {
		// fallback if unkonwn command
		return handleUnknownCommand(conn, command)
	}

	return handler(conn, args)

}
