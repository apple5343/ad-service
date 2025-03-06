package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type ServerProvider struct {
	client *http.Client
	path   string
}

func (s *ServerProvider) RestartApp() error {
	println("Restart Application")
	cmd := exec.Command("/bin/sh", "-c", "docker compose -f ../docker-compose.test.yml --env-file ../.env.test restart app-test")
	if err := cmd.Run(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	println("Application restarted")
	return nil
}

func (s *ServerProvider) ResetSuite(name string) error {
	println("Reset Application By", name)
	if err := s.stopApp(); err != nil {
		return err
	}
	if err := s.startApp(name); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	println("Application started")
	return nil
}

func (p *ServerProvider) stopApp() error {
	cmd := exec.Command("docker", "compose", "-f", "../docker-compose.test.yml", "--env-file", "../.env.test", "down")
	return cmd.Run()
}

func (p *ServerProvider) startApp(name string) error {
	cmd := exec.Command("/bin/sh", "-c", "docker compose -f ../docker-compose.test.yml --env-file ../.env.test up -d")
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("/bin/sh", "-c", "docker compose -f ../docker-compose.test.yml --env-file ../.env.test logs -f app-test > logs/"+name+".txt 2>&1 &")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (s *ServerProvider) Do(req *http.Request, expectedStatusCode int, t *testing.T) (*http.Response, error) {
	resp, err := s.client.Do(req)
	require.NoError(t, err)
	if resp.StatusCode != expectedStatusCode {
		body := make(map[string]interface{})
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
		fmt.Printf("Application output: %v\n", body)
		require.Equal(t, expectedStatusCode, resp.StatusCode)
	}
	return resp, nil
}
