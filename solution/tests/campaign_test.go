package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"slices"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/suite"
)

type CampaignTestSuite struct {
	suite.Suite
	provider *ServerProvider
}

func (s *CampaignTestSuite) SetupSuite() {
	s.provider.ResetSuite("campaign_test")
}

func (s *CampaignTestSuite) TestCreateCampaign() {
	s.T().Run("happy path", func(t *testing.T) {
		advertiser, resp, err := CreateRandomAdvertiser(s.provider.path)
		s.Require().NoError(err)
		s.Require().Equal(http.StatusCreated, resp.StatusCode)

		campaign := RandomCampaign()
		campaign.AdvertiserID = advertiser.AdvertiserID
		body, _ := json.Marshal(campaign)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err = s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		var respCampaign *Campaign
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respCampaign))
		campaign.CampaignID = respCampaign.CampaignID
		s.Equal(campaign, respCampaign)
	})

	s.T().Run("unregistered advertiser", func(t *testing.T) {
		campaign := RandomCampaign()
		campaign.AdvertiserID = gofakeit.UUID()
		body, _ := json.Marshal(campaign)
		req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/advertisers/"+gofakeit.UUID()+"/campaigns", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})

	advertiser, resp, err := CreateRandomAdvertiser(s.provider.path)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, resp.StatusCode)

	tests := []struct {
		name       string
		getBody    func() *Campaign
		statusCode int
	}{
		{
			name: "end date less than start date",
			getBody: func() *Campaign {
				campaign := RandomCampaign()
				start := gofakeit.Number(50, 100)
				end := gofakeit.Number(0, start-1)
				campaign.StartDate = &start
				campaign.EndDate = &end
				return campaign
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "start date less than 0",
			getBody: func() *Campaign {
				campaign := RandomCampaign()
				start := gofakeit.Number(-100, -1)
				campaign.StartDate = &start
				return campaign
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "end date less than 0",
			getBody: func() *Campaign {
				campaign := RandomCampaign()
				end := gofakeit.Number(-100, -1)
				campaign.EndDate = &end
				return campaign
			},
			statusCode: http.StatusBadRequest,
		},
		{
			"age from greater than age to",
			func() *Campaign {
				campaign := RandomCampaign()
				ageFrom := gofakeit.Number(50, 100)
				ageTo := gofakeit.Number(0, ageFrom-1)
				campaign.Target.AgeFrom = &ageFrom
				campaign.Target.AgeTo = &ageTo
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"age from less than 0",
			func() *Campaign {
				campaign := RandomCampaign()
				ageFrom := gofakeit.Number(-100, -1)
				campaign.Target.AgeFrom = &ageFrom
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"age to less than 0",
			func() *Campaign {
				campaign := RandomCampaign()
				ageTo := gofakeit.Number(-100, -1)
				campaign.Target.AgeTo = &ageTo
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"unknown gender",
			func() *Campaign {
				campaign := RandomCampaign()
				gender := gofakeit.Name()
				campaign.Target.Gender = &gender
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"empty ad title",
			func() *Campaign {
				campaign := RandomCampaign()
				campaign.AdTitle = nil
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"empty ad text",
			func() *Campaign {
				campaign := RandomCampaign()
				campaign.AdText = nil
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"ad text length 0",
			func() *Campaign {
				campaign := RandomCampaign()
				t := ""
				campaign.AdText = &t
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"empty cost per click",
			func() *Campaign {
				campaign := RandomCampaign()
				campaign.CostPerClick = nil
				return campaign
			},
			http.StatusBadRequest,
		},
		{
			"cost per click less than 0",
			func() *Campaign {
				campaign := RandomCampaign()
				cost := gofakeit.Float32Range(-10, -1)
				campaign.CostPerClick = &cost
				return campaign
			},
			http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			campaign := tt.getBody()
			body, _ := json.Marshal(campaign)
			req, _ := http.NewRequest(http.MethodPost, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := s.provider.client.Do(req)
			s.Require().NoError(err)
			defer resp.Body.Close()
			s.Require().Equal(tt.statusCode, resp.StatusCode)
		})
	}
}

func (s *CampaignTestSuite) TestGetCampaign() {
	s.T().Run("happy path", func(t *testing.T) {
		advertiser, resp, err := CreateRandomAdvertiser(s.provider.path)
		s.Require().NoError(err)
		s.Require().Equal(http.StatusCreated, resp.StatusCode)

		campaign, resp, err := CreateRandomCampaign(s.provider.path, advertiser.AdvertiserID)
		s.Require().NoError(err)
		s.Require().Equal(http.StatusCreated, resp.StatusCode)

		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns/"+campaign.CampaignID, nil)
		resp, err = s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusOK, resp.StatusCode)
	})

	s.T().Run("not found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+gofakeit.UUID()+"/campaigns/"+gofakeit.UUID(), nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})
}

func (s *CampaignTestSuite) TestDeleteCampaign() {
	s.T().Run("happy path", func(t *testing.T) {
		advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
		campaign := RandomCampaign()
		campaign.AdvertiserID = advertiser.AdvertiserID
		campaign = CreateCampaign(s.T(), s.provider.path, campaign)

		req, _ := http.NewRequest(http.MethodDelete, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns/"+campaign.CampaignID, nil)
		_, err := s.provider.Do(req, http.StatusNoContent, t)
		s.Require().NoError(err)
	})

	s.T().Run("campaign not found", func(t *testing.T) {
		advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())

		req, _ := http.NewRequest(http.MethodDelete, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns/"+gofakeit.UUID(), nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})

	s.T().Run("advertiser not found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, s.provider.path+"/advertisers/"+gofakeit.UUID()+"/campaigns/"+gofakeit.UUID(), nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})
}

func (s *CampaignTestSuite) TestListCampaigns() {
	advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())

	campaigns := make([]*Campaign, 10)
	for i := 0; i < 10; i++ {
		campaign := RandomCampaign()
		campaign.AdvertiserID = advertiser.AdvertiserID
		campaign = CreateCampaign(s.T(), s.provider.path, campaign)
		campaigns[i] = campaign
	}

	s.T().Run("happy path", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns?page=2&size=5", nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusOK, resp.StatusCode)
		var respCampaigns []*Campaign
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respCampaigns))
		s.Require().Equal(5, len(respCampaigns))
		slices.Reverse(respCampaigns)
		for i, campaign := range campaigns[:5] {
			campaign.CampaignID = respCampaigns[i].CampaignID
			s.Equal(campaign, respCampaigns[i])
		}
		req, _ = http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+advertiser.AdvertiserID+"/campaigns?size=5", nil)
		resp, err = s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusOK, resp.StatusCode)

		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&respCampaigns))
		s.Require().Equal(5, len(respCampaigns))
		slices.Reverse(respCampaigns)
		for i, campaign := range campaigns[5:] {
			campaign.CampaignID = respCampaigns[i].CampaignID
			s.Equal(campaign, respCampaigns[i])
		}
	})

	s.T().Run("advertiser not found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, s.provider.path+"/advertisers/"+gofakeit.UUID()+"/campaigns", nil)
		resp, err := s.provider.client.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	})
}

func (s *CampaignTestSuite) TestUpdateCampaign() {
	advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
	campaign := RandomCampaign()
	campaign.AdvertiserID = advertiser.AdvertiserID
	start := gofakeit.Number(10, 20)
	end := start + gofakeit.Number(0, 10)
	campaign.StartDate = &start
	campaign.EndDate = &end
	SetDay(s.T(), s.provider.path, 0)
	campaign = CreateCampaign(s.T(), s.provider.path, campaign)
	tests := []struct {
		name    string
		getBody func() *Campaign
		status  int
	}{
		{
			name:   "happy path",
			status: http.StatusOK,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				campaign = newCampaign
				return newCampaign
			},
		},
		{
			name:   "advertiser not found",
			status: http.StatusForbidden,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = gofakeit.UUID()
				newCampaign.CampaignID = campaign.CampaignID
				return newCampaign
			},
		},
		{
			name:   "campaign not found",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = gofakeit.UUID()
				return newCampaign
			},
		},
		{
			name:   "empty start date",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				newCampaign.StartDate = nil
				return newCampaign
			},
		},
		{
			name:   "empty end date",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				newCampaign.EndDate = nil
				return newCampaign
			},
		},
		{
			name:   "end date less than start date",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				strart := gofakeit.Number(10, 20)
				end := strart - 1
				newCampaign.StartDate = &strart
				newCampaign.EndDate = &end
				return newCampaign
			},
		},
		{
			name:   "cost per click less than 0",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				cost := float32(-1.0)
				newCampaign.CostPerClick = &cost
				return newCampaign
			},
		},
		{
			name:   "cost per impression less than 0",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				cost := float32(-1.0)
				newCampaign.CostPerImpression = &cost
				return newCampaign
			},
		},
		{
			name:   "clicks limit less than 0",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				limit := -1
				newCampaign.ClicksLimit = &limit
				return newCampaign
			},
		},
		{
			name:   "impression limit less than 0",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				limit := -1
				newCampaign.ImpressionsLimit = &limit
				return newCampaign
			},
		},
		{
			name:   "impression limit less than clicks limit",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				impressionsLimit := gofakeit.Number(20, 100)
				clicksLimit := impressionsLimit + 1
				newCampaign.ClicksLimit = &clicksLimit
				newCampaign.ImpressionsLimit = &impressionsLimit
				return newCampaign
			},
		},
		{
			name:   "inknown gender",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				invalidGender := "квадробер"
				newCampaign.Target.Gender = &invalidGender
				return newCampaign
			},
		},
		{
			name:   "empty location",
			status: http.StatusOK,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				newCampaign.Target.Location = nil
				campaign = newCampaign
				return newCampaign
			},
		},
		{
			name:   "empty age",
			status: http.StatusOK,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				newCampaign.Target.AgeFrom = nil
				newCampaign.Target.AgeTo = nil
				campaign = newCampaign
				return newCampaign
			},
		},
		{
			name:   "age from greater than age to",
			status: http.StatusBadRequest,
			getBody: func() *Campaign {
				newCampaign := RandomCampaign()
				newCampaign.AdvertiserID = campaign.AdvertiserID
				newCampaign.CampaignID = campaign.CampaignID
				ageFrom := gofakeit.Number(20, 100)
				ageTo := ageFrom - 1
				newCampaign.Target.AgeFrom = &ageFrom
				newCampaign.Target.AgeTo = &ageTo
				return newCampaign
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			newCampaign, code := UpdateCampaign(s.T(), s.provider.path, tt.getBody())
			s.Require().Equal(tt.status, code)
			if code == http.StatusOK {
				s.Require().Equal(campaign, newCampaign)
			}
		})
	}
	end = *campaign.StartDate + 10
	campaign.EndDate = &end
	campaign, code := UpdateCampaign(s.T(), s.provider.path, campaign)
	s.Require().Equal(http.StatusOK, code)
	SetDay(s.T(), s.provider.path, *campaign.StartDate+5)

	s.T().Run("update impressions limit after campaign start", func(t *testing.T) {
		impressionsLimit := *campaign.ImpressionsLimit + 1
		newCampaign := RandomCampaign()
		newCampaign.AdvertiserID = campaign.AdvertiserID
		newCampaign.CampaignID = campaign.CampaignID
		newCampaign.ImpressionsLimit = &impressionsLimit
		_, code := UpdateCampaign(s.T(), s.provider.path, newCampaign)
		s.Require().Equal(http.StatusBadRequest, code)
	})

	s.T().Run("update clicks limit after campaign start", func(t *testing.T) {
		clicksLimit := *campaign.ClicksLimit + 1
		newCampaign := RandomCampaign()
		newCampaign.AdvertiserID = campaign.AdvertiserID
		newCampaign.CampaignID = campaign.CampaignID
		newCampaign.ClicksLimit = &clicksLimit
		_, code := UpdateCampaign(s.T(), s.provider.path, newCampaign)
		s.Require().Equal(http.StatusBadRequest, code)
	})

	s.T().Run("update start date after campaign start", func(t *testing.T) {
		startDate := *campaign.StartDate + 1
		newCampaign := RandomCampaign()
		newCampaign.AdvertiserID = campaign.AdvertiserID
		newCampaign.CampaignID = campaign.CampaignID
		newCampaign.StartDate = &startDate
		_, code := UpdateCampaign(s.T(), s.provider.path, newCampaign)
		s.Require().Equal(http.StatusBadRequest, code)
	})

	s.T().Run("update end date after campaign start", func(t *testing.T) {
		endDate := *campaign.EndDate + 1
		newCampaign := RandomCampaign()
		newCampaign.AdvertiserID = campaign.AdvertiserID
		newCampaign.CampaignID = campaign.CampaignID
		newCampaign.EndDate = &endDate
		_, code := UpdateCampaign(s.T(), s.provider.path, newCampaign)
		s.Require().Equal(http.StatusBadRequest, code)
	})
}
