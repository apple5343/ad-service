package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	provider *ServerProvider
}

func (s *ClientTestSuite) SetupSuite() {
	s.provider.ResetSuite("client_test")
}

func (s *ClientTestSuite) TestCreateClient() {
	s.T().Run("happy path", func(t *testing.T) {
		l := gofakeit.Number(1, 25)
		clients := make([]*Client, l)
		for i := 0; i < l; i++ {
			clients[i] = RandomClient()
		}

		body, _ := json.Marshal(clients)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/clients/bulk", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusCreated, resp.StatusCode)

		var respClients []*Client
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respClients))
		for i, client := range clients {
			respClient := respClients[i]
			s.Equal(client, respClient)
		}
	})
}

func (s *ClientTestSuite) TestGetClient() {
	s.T().Run("happy path", func(t *testing.T) {
		client := RandomClient()
		body, _ := json.Marshal([]*Client{client})
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/clients/bulk", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().NoError(err)
		s.Require().Equal(http.StatusCreated, resp.StatusCode)

		req, _ = http.NewRequest(http.MethodGet, s.provider.path+"/clients/"+client.ClientID, nil)
		resp, err = s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().NoError(err)
		s.Require().Equal(http.StatusOK, resp.StatusCode)

		var respClient *Client
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respClient))
		s.Equal(client, respClient)
	})

	s.T().Run("not found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/clients/"+gofakeit.UUID(), nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusNotFound, resp.StatusCode)
	})
}
