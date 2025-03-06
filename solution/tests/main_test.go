package tests

import (
	"net/http"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	path           string
	client         *http.Client
	serverProvider *ServerProvider
}

func (s *APITestSuite) SetupSuite() {
	println("SetupSuite")
	s.initDeps()
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) TestStat() {
	stat := &StatTestSuite{
		provider: s.serverProvider,
	}
	suite.Run(s.T(), stat)
}

func (s *APITestSuite) TestAd() {
	ad := &AdTestSuite{
		provider: s.serverProvider,
	}
	suite.Run(s.T(), ad)
}

func (s *APITestSuite) TestTime() {
	time := &TimeTestSuite{
		provider: s.serverProvider,
	}
	suite.Run(s.T(), time)
}

func (s *APITestSuite) TestAdvertiser() {
	client := &AdvertiserTestSuite{
		provider: s.serverProvider,
	}
	suite.Run(s.T(), client)
}

func (s *APITestSuite) TestCampaign() {
	campaign := &CampaignTestSuite{
		provider: s.serverProvider,
	}
	suite.Run(s.T(), campaign)
}

func (s *APITestSuite) TestClient() {
	client := &ClientTestSuite{
		provider: s.serverProvider,
	}
	suite.Run(s.T(), client)
}

func (s *APITestSuite) initDeps() {
	s.client = &http.Client{}
	s.path = "http://localhost:8081"
	s.serverProvider = &ServerProvider{
		client: s.client,
		path:   s.path,
	}
}

func (s *APITestSuite) stopApp() error {
	cmd := exec.Command("docker", "compose", "-f", "../docker-compose.test.yml", "--env-file", "../.env.test", "down")
	return cmd.Run()
}

func (s *APITestSuite) startApp(name string) error {
	cmd := exec.Command("/bin/sh", "-c", "docker compose -f ../docker-compose.test.yml --env-file ../.env.test up -d")
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("/bin/sh", "-c", "docker compose -f ../docker-compose.test.yml --env-file ../.env.test logs -f > logs/"+name+".txt 2>&1 &")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
