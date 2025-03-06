package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/suite"
)

type AdvertiserTestSuite struct {
	suite.Suite
	provider *ServerProvider
}

func (s *AdvertiserTestSuite) SetupSuite() {
	s.provider.ResetSuite("advertiser_test")
}

func (s *AdvertiserTestSuite) TestCreateAdvertiser() {
	s.T().Run("happy path", func(t *testing.T) {
		l := gofakeit.Number(1, 25)
		advertisers := make([]*Advertiser, l)
		for i := 0; i < l; i++ {
			advertisers[i] = RandomAdvertiser()
		}

		body, _ := json.Marshal(advertisers)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/advertisers/bulk", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		s.Require().Equal(http.StatusCreated, resp.StatusCode)

		var respAdvertisers []*Advertiser
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respAdvertisers))
		for i, advertiser := range advertisers {
			respAdvertiser := respAdvertisers[i]
			s.Equal(advertiser, respAdvertiser)
		}
	})
}

func (s *AdvertiserTestSuite) TestGetAdvertiser() {
	s.T().Run("happy path", func(t *testing.T) {

		advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())

		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+advertiser.AdvertiserID, nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusOK, resp.StatusCode)

		var respAdvertiser *Advertiser
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respAdvertiser))
		s.Equal(advertiser, respAdvertiser)
	})

	s.T().Run("not found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+gofakeit.UUID(), nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})
}

func (s *AdvertiserTestSuite) TestMlScore() {
	s.T().Run("happy path", func(t *testing.T) {
		advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
		client := CreateClient(s.T(), s.provider.path, RandomClient())

		score := MlScore{
			ClientID:     client.ClientID,
			AdvertiserID: advertiser.AdvertiserID,
			Score:        gofakeit.Number(0, 100),
		}

		body, _ := json.Marshal(score)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/ml-scores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusOK, resp.StatusCode)
	})

	s.T().Run("client not found", func(t *testing.T) {
		advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())

		score := MlScore{
			ClientID:     gofakeit.UUID(),
			AdvertiserID: advertiser.AdvertiserID,
			Score:        gofakeit.Number(0, 100),
		}
		body, _ := json.Marshal(score)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/ml-scores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})

	s.T().Run("advertiser not found", func(t *testing.T) {
		client := CreateClient(s.T(), s.provider.path, RandomClient())

		score := MlScore{
			ClientID:     client.ClientID,
			AdvertiserID: gofakeit.UUID(),
			Score:        gofakeit.Number(0, 100),
		}
		body, _ := json.Marshal(score)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/ml-scores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})

	s.T().Run("negative score", func(t *testing.T) {
		advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
		client := CreateClient(s.T(), s.provider.path, RandomClient())

		score := MlScore{
			ClientID:     client.ClientID,
			AdvertiserID: advertiser.AdvertiserID,
			Score:        -1,
		}
		body, _ := json.Marshal(score)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/ml-scores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})
}
