package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/suite"
)

type TimeTestSuite struct {
	suite.Suite
	provider *ServerProvider
}

func (s *TimeTestSuite) SetupSuite() {
	s.provider.ResetSuite("time_test")
}

func (s *TimeTestSuite) TestTime() {
	toSet := gofakeit.Number(50, 100)

	s.T().Run("set time happy path", func(t *testing.T) {
		SetDay(t, s.provider.path, toSet)
		date := GetDay(t, s.provider.path)
		s.Require().Equal(toSet, date)
	})

	s.T().Run("set time negative", func(t *testing.T) {
		type reqBody struct {
			CurrentDay *int `json:"current_date,omitempty"`
		}
		toSet := -1
		reqB, _ := json.Marshal(reqBody{CurrentDay: &toSet})
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/time/advance", bytes.NewBuffer(reqB))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})

	s.T().Run("get time after restart", func(t *testing.T) {
		s.Require().NoError(s.provider.RestartApp())
		date := GetDay(t, s.provider.path)
		s.Require().Equal(toSet, date)
	})
}
