package tests

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/suite"
)

type AdTestSuite struct {
	suite.Suite
	provider *ServerProvider
}

func (s *AdTestSuite) SetupSuite() {
	s.provider.ResetSuite("ad_test")
}

func (s *AdTestSuite) TestImpression() {
	SetDay(s.T(), s.provider.path, 0)
	advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
	client := CreateClient(s.T(), s.provider.path, RandomClient())
	campaign := RandomCampaign()
	campaign.Target = nil
	start := 0
	end := 20
	campaign.StartDate = &start
	campaign.EndDate = &end
	campaign.AdvertiserID = advertiser.AdvertiserID
	campaign = CreateCampaign(s.T(), s.provider.path, campaign)

	s.T().Run("happy path", func(t *testing.T) {
		_, code := Impression(s.T(), s.provider.path, client.ClientID)
		s.Require().Equal(http.StatusOK, code)
	})

	s.T().Run("inactive campaign", func(t *testing.T) {
		SetDay(t, s.provider.path, 99999)

		_, code := Impression(s.T(), s.provider.path, client.ClientID)
		s.Require().Equal(http.StatusNotFound, code)
		SetDay(t, s.provider.path, 0)
	})
}

func (s *AdTestSuite) TestClick() {
	advertiser := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
	client := CreateClient(s.T(), s.provider.path, RandomClient())
	campaign := RandomCampaign()
	campaign.AdvertiserID = advertiser.AdvertiserID
	campaign.Target = nil
	start := 0
	campaign.StartDate = &start

	campaign = CreateCampaign(s.T(), s.provider.path, campaign)

	s.T().Run("happy path", func(t *testing.T) {
		_, status := Impression(s.T(), s.provider.path, client.ClientID)
		s.Require().Equal(http.StatusOK, status)

		status = Click(s.T(), s.provider.path, campaign.CampaignID, client.ClientID)
		s.Require().Equal(http.StatusNoContent, status)
	})

	s.T().Run("double click", func(t *testing.T) {
		status := Click(s.T(), s.provider.path, campaign.CampaignID, client.ClientID)
		s.Require().Equal(http.StatusNoContent, status)
	})

	s.T().Run("campaign ended", func(t *testing.T) {
		ad, status := Impression(s.T(), s.provider.path, client.ClientID)
		s.Require().Equal(http.StatusOK, status)
		campaign := GetCampaign(s.T(), s.provider.path, ad.AdvertiserId, ad.AdId)
		SetDay(t, s.provider.path, *campaign.EndDate+1)
		status = Click(s.T(), s.provider.path, campaign.CampaignID, client.ClientID)
		s.Require().Equal(http.StatusBadRequest, status)
		SetDay(t, s.provider.path, *campaign.StartDate)
	})

	s.T().Run("client not found", func(t *testing.T) {
		status := Click(s.T(), s.provider.path, campaign.CampaignID, gofakeit.UUID())
		s.Require().Equal(http.StatusBadRequest, status)
	})
}
