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

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
}

func Run(ctx context.Context, cfg *config.Config) error {
	logger := slog.With("transport", "resp", "port", cfg.RESPPort)
	srv := &Server{
		cfg:    cfg,
		logger: logger,
	}

	return srv.start(ctx)
}

func (s *Server) start(ctx context.Context) error {
	fmt.Println("Running RESP server...")
	ln, err := net.Listen("tcp", ":"+s.cfg.RESPPort)
	if err != nil {
		return fmt.Errorf("Error listening on port: %s", s.cfg.RESPPort)
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

		go s.handleConnection(ctx, conn)
	}
}

func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
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
			s.logger.Error("Error parsing command", "err", err)
			return
		}

		s.routeRequest(conn, args)
	}
}
