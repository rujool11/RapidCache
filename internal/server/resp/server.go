package respserver

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/rujool11/RapidCache/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	fmt.Println("Running RESP server...")
	ln, err := net.Listen("tcp", ":"+cfg.RESPPort)
	if err != nil {
		return fmt.Errorf("Error listening on port: %s", cfg.RESPPort)
	}
	defer ln.Close()

	go func() {
		<-ctx.Done()
		fmt.Println("Closing RESP server...")
		ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				return fmt.Errorf("Error accepting connection: %s", err)
			}
		}

		go handleConnection(ctx, conn)
	}
}

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close() // close connection when session ends

	reader := bufio.NewReader(conn)

	for {
		// check if server is shutting down
		select {
		case <-ctx.Done():
			return
		default:
		}

		args, err := parseCommand(reader)
		if err != nil {
			// returned when no more inputs - graceful connection close
			if err == io.EOF {
				return
			}
			slog.Error("Error parsing command", "err", err)
			return
		}

		routeRequest(conn, args)
	}
}
