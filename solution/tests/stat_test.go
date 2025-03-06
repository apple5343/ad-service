package tests

import (
	"encoding/json"
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/rand"
)

const (
	clientsCount     = 20
	advertisersCount = 10
	campaignsCount   = 3
)

type StatsExpected map[string]*StatInfo

type StatTestSuite struct {
	suite.Suite
	provider      *ServerProvider
	statsExpected StatsExpected
	statsDaily    map[int]StatsExpected
	campaigns     []*CampaignClear
	clients       []*Client
	adversers     []*Advertiser
}

func (s *StatTestSuite) SetupSuite() {
	s.provider.ResetSuite("stat_test")
	s.statsExpected = make(map[string]*StatInfo)
	s.statsDaily = make(map[int]StatsExpected)
	s.campaigns = []*CampaignClear{}
	s.clients = []*Client{}
	s.adversers = []*Advertiser{}

	for i := 0; i < clientsCount; i++ {
		client := CreateClient(s.T(), s.provider.path, RandomClient())
		s.clients = append(s.clients, client)
	}

	for i := 0; i < advertisersCount; i++ {
		adverister := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
		s.adversers = append(s.adversers, adverister)
		for j := 0; j < campaignsCount; j++ {
			campaign := CreateCampaignNorm()
			campaign.AdvertiserID = adverister.AdvertiserID
			campaign.StartDate = 0
			campaign = CreateCampaignClear(s.T(), s.provider.path, campaign)
			s.campaigns = append(s.campaigns, campaign)
			s.statsExpected[campaign.CampaignID] = &StatInfo{
				Stat:           &Stat{},
				Campaign:       campaign,
				ImpressionsIds: make(map[string]struct{}),
				ClicksIds:      map[string]struct{}{},
			}
		}
	}
}

func (s *StatTestSuite) GetDailyStat() map[string]*StatInfo {
	dailyStat := make(map[string]*StatInfo)
	for k, v := range s.statsExpected {
		dailyStat[k] = &StatInfo{
			Stat:           &Stat{},
			Campaign:       v.Campaign,
			ImpressionsIds: v.ImpressionsIds,
			ClicksIds:      v.ClicksIds,
		}
	}

	return dailyStat
}

func (s *StatTestSuite) JoinStats(stat, dailyStat map[string]*StatInfo) map[string]*StatInfo {
	for k := range stat {
		daily := dailyStat[k]
		stat[k].Stat.ImpressionCount += daily.Stat.ImpressionCount
		stat[k].Stat.ClickCount += daily.Stat.ClickCount
		stat[k].Stat.SpentImpressions += daily.Stat.SpentImpressions
		stat[k].Stat.SpentClicks += daily.Stat.SpentClicks
		stat[k].Stat.Update()
	}
	return stat
}

func (s *StatTestSuite) GetStatByAdvertiser(advertiserID string) *Stat {
	stat := &Stat{}
	for _, v := range s.statsExpected {
		if v.Campaign.AdvertiserID == advertiserID {
			stat.ImpressionCount += v.Stat.ImpressionCount
			stat.ClickCount += v.Stat.ClickCount
			stat.SpentImpressions += v.Stat.SpentImpressions
			stat.SpentImpressions = float32(Round(float64(stat.SpentImpressions)))
			stat.SpentClicks += v.Stat.SpentClicks
			stat.SpentClicks = float32(Round(float64(stat.SpentClicks)))
			stat.Update()
		}
	}
	return stat
}

func (s *StatTestSuite) GetStatByAdvertiserDaily(advertiserID string, day int) *Stat {
	stat := &Stat{}
	for _, v := range s.statsDaily[day] {
		if v.Campaign.AdvertiserID == advertiserID {
			stat.ImpressionCount += v.Stat.ImpressionCount
			stat.ClickCount += v.Stat.ClickCount
			stat.SpentImpressions += v.Stat.SpentImpressions
			stat.SpentImpressions = float32(Round(float64(stat.SpentImpressions)))
			stat.SpentClicks += v.Stat.SpentClicks
			stat.SpentClicks = float32(Round(float64(stat.SpentClicks)))
			stat.Update()
		}
	}
	return stat
}

func (s *StatTestSuite) Impression(userID string) string {
	ad, code := Impression(s.T(), s.provider.path, userID)
	if code != http.StatusOK {
		return ""
	}
	stat := s.statsExpected[ad.AdId]
	stat.Impression(userID)
	return ad.AdId
}

func (s *StatTestSuite) Click(userID, adID string) {
	code := Click(s.T(), s.provider.path, adID, userID)
	if code != http.StatusNoContent {
		return
	}
	stat := s.statsExpected[adID]
	stat.Click(userID)
}

func (s *StatTestSuite) UpdateCampaign(campaign *CampaignClear) {
	newCampaign := CreateCampaignNorm()
	newCampaign.AdvertiserID = campaign.AdvertiserID
	newCampaign.CampaignID = campaign.CampaignID
	newCampaign.StartDate = campaign.StartDate
	newCampaign.EndDate = campaign.EndDate
	newCampaign.ImpressionsLimit = campaign.ImpressionsLimit
	newCampaign.ClicksLimit = campaign.ClicksLimit
	newCampaign, code := UpdateCampaignClear(s.T(), s.provider.path, newCampaign)
	s.Require().Equal(http.StatusOK, code)
	for _, c := range s.campaigns {
		if c.CampaignID == campaign.CampaignID {
			c.CostPerImpression = newCampaign.CostPerImpression
			c.CostPerClick = newCampaign.CostPerClick
			break
		}
	}
	s.statsExpected[campaign.CampaignID].Campaign = newCampaign
}

type StatInfo struct {
	Stat           *Stat
	Campaign       *CampaignClear
	ImpressionsIds map[string]struct{}
	ClicksIds      map[string]struct{}
}

func (i *StatInfo) Impression(userID string) bool {
	if _, ok := i.ImpressionsIds[userID]; ok {
		return false
	}
	i.ImpressionsIds[userID] = struct{}{}
	i.Stat.ImpressionCount++
	i.Stat.SpentImpressions += i.Campaign.CostPerImpression
	i.Stat.SpentImpressions = float32(math.Round(float64(i.Stat.SpentImpressions)*100) / 100)
	i.Stat.Update()
	return true
}

func (i *StatInfo) Click(userID string) bool {
	if _, ok := i.ClicksIds[userID]; ok {
		return false
	}
	i.ClicksIds[userID] = struct{}{}
	i.Stat.ClickCount++
	i.Stat.SpentClicks += i.Campaign.CostPerClick
	i.Stat.SpentClicks = float32(math.Round(float64(i.Stat.SpentClicks)*100) / 100)
	i.Stat.Update()
	return true
}

func (s *StatTestSuite) TestStatClear() {
	SetDay(s.T(), s.provider.path, 0)
	adverister := CreateAdvertiser(s.T(), s.provider.path, RandomAdvertiser())
	campaign := CreateCampaignNorm()
	campaign.AdvertiserID = adverister.AdvertiserID
	campaign = CreateCampaignClear(s.T(), s.provider.path, campaign)

	stat := Stat{}
	s.campaigns = append(s.campaigns, campaign)
	s.statsExpected[campaign.CampaignID] = &StatInfo{
		Stat:     &stat,
		Campaign: campaign,
	}
	s.T().Run("clear stat", func(t *testing.T) {
		req, _ := http.NewRequest("GET", s.provider.path+"/stats/campaigns/"+campaign.CampaignID, nil)
		r, _ := s.provider.Do(req, 200, t)

		var statResponse *Stat
		s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
		s.Require().Equal(stat, *statResponse)
	})

	s.T().Run("clear stat daily", func(t *testing.T) {
		req, _ := http.NewRequest("GET", s.provider.path+"/stats/campaigns/"+campaign.CampaignID+"/daily", nil)
		r, _ := s.provider.Do(req, 200, t)

		var statResponse *Stat
		s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
		s.Require().Equal(stat, *statResponse)
	})

	s.T().Run("clear stat by advertiser", func(t *testing.T) {
		req, _ := http.NewRequest("GET", s.provider.path+"/stats/advertisers/"+adverister.AdvertiserID+"/campaigns", nil)
		r, _ := s.provider.Do(req, 200, t)

		var statResponse *Stat
		s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
		s.Require().Equal(stat, *statResponse)
	})

	s.T().Run("clear stat by advertiser daily", func(t *testing.T) {
		req, _ := http.NewRequest("GET", s.provider.path+"/stats/advertisers/"+adverister.AdvertiserID+"/campaigns/daily", nil)
		r, _ := s.provider.Do(req, 200, t)

		var statResponse *Stat
		s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
		s.Require().Equal(stat, *statResponse)
	})
}

func (s *StatTestSuite) TestStat() {
	s.T().Run("stat by campaign", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			for _, client := range s.clients {
				id := s.Impression(client.ClientID)
				if id != "" && rand.Intn(10) < 5 {
					s.Click(client.ClientID, id)
				}
				if id != "" && rand.Intn(10) < 4 {
					s.UpdateCampaign(s.statsExpected[id].Campaign)
				}
			}
		}
		s.provider.RestartApp()
		for _, campaign := range s.campaigns {
			req, _ := http.NewRequest("GET", s.provider.path+"/stats/campaigns/"+campaign.CampaignID, nil)
			r, _ := s.provider.Do(req, 200, t)

			var statResponse *Stat
			s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
			s.Require().Equal(s.statsExpected[campaign.CampaignID].Stat, statResponse)
		}
	})
	day := 5

	s.T().Run("stat by campaign daily", func(t *testing.T) {
		dailyStat := s.GetDailyStat()
		SetDay(t, s.provider.path, day)
		stat := s.statsExpected
		s.statsExpected = dailyStat
		for i := 0; i < 5; i++ {
			for _, client := range s.clients {
				id := s.Impression(client.ClientID)
				if id != "" && rand.Intn(10) < 5 {
					s.Click(client.ClientID, id)
				}
				if id != "" && rand.Intn(10) < 4 {
					s.UpdateCampaign(s.statsExpected[id].Campaign)
				}
			}
		}
		for _, campaign := range s.campaigns {
			req, _ := http.NewRequest("GET", s.provider.path+"/stats/campaigns/"+campaign.CampaignID+"/daily", nil)
			r, _ := s.provider.Do(req, 200, t)

			var statResponse *Stat
			s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
			s.Require().Equal(s.statsExpected[campaign.CampaignID].Stat, statResponse)
		}
		stat = s.JoinStats(stat, dailyStat)
		s.statsExpected = stat
		s.statsDaily[day] = dailyStat
	})

	s.T().Run("stat by advertiser", func(t *testing.T) {
		for _, adverister := range s.adversers {
			req, _ := http.NewRequest("GET", s.provider.path+"/stats/advertisers/"+adverister.AdvertiserID+"/campaigns", nil)
			r, _ := s.provider.Do(req, 200, t)

			var statResponse *Stat
			s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
			expected := s.GetStatByAdvertiser(adverister.AdvertiserID)
			s.Require().Equal(expected, statResponse)
		}
	})

	s.T().Run("stat by advertiser daily", func(t *testing.T) {
		SetDay(t, s.provider.path, day)
		for _, adverister := range s.adversers {
			req, _ := http.NewRequest("GET", s.provider.path+"/stats/advertisers/"+adverister.AdvertiserID+"/campaigns/daily", nil)
			r, _ := s.provider.Do(req, 200, t)

			var statResponse *Stat
			s.Require().NoError(json.NewDecoder(r.Body).Decode(&statResponse))
			expected := s.GetStatByAdvertiserDaily(adverister.AdvertiserID, day)
			s.Require().Equal(expected, statResponse)
		}
	})
}
