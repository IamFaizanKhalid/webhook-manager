package webhook

import (
	"fmt"
	"io"
	"os/exec"
)

type Server struct {
	cmd *exec.Cmd
}

func NewServer(port int, hooksFile string, logFile io.Writer) (*Server, error) {
	p := fmt.Sprintf("%d", port)
	cmd := exec.Command("webhook", "-hooks", hooksFile, "-verbose", "-hotreload", "-port", p)

	if logFile != nil {
		cmd.Stdout = logFile
		cmd.Stderr = logFile
	}

	return &Server{cmd}, nil
}

func (s *Server) Start() error {
	return s.cmd.Start()
}

func (s *Server) Stop() error {
	return s.cmd.Process.Kill()
}
