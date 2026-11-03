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
	// stubbed
	defer conn.Close() // close connection when session ends

	reader := bufio.NewReader(conn)

	select {
	case <-ctx.Done():
		return
	default:
	}

	// RESP lines end in \r\n
	_, err := reader.ReadString('\n')
	if err != nil {
		// EOF returned when no more input available
		if err != io.EOF {
			slog.Error("Error reading byte stream: ", "err", err)
			return // conn closed
		}
	}

	// stub PONG
	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		slog.Error("Error writing PONG", "err", err)
	}
}
