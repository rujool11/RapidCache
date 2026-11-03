package respserver

import (
	"context"
	"fmt"
	"net"

	"github.com/rujool11/RapidCache/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	fmt.Println("Running RESP server...")
	ln, err := net.Listen("tcp", ":"+cfg.RESPPort)
	if err != nil {
		return fmt.Errorf("Error listening on port: ", cfg.RESPPort)
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
				return fmt.Errorf("Error accepting connection", err)
			}
		}

		go handleConnection(ctx, conn)
	}
}

func handleConnection(ctx context.Context, conn net.Conn) {
	// stubbed
	defer conn.Close() // close connection when session ends
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		return
	}
}
