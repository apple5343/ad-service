package tests

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"

	"github.com/brianvoe/gofakeit"
	"golang.org/x/exp/rand"
)

func RandomClient() *Client {
	return &Client{
		ClientID: gofakeit.UUID(),
		Login:    gofakeit.Username(),
		Age:      gofakeit.Number(0, 100),
		Location: gofakeit.City(),
		Gender:   genders[gofakeit.Number(0, 1)],
	}
}

func CreateRandomClient(path string) (*Client, *http.Response, error) {
	client := RandomClient()
	body, _ := json.Marshal([]*Client{client})
	req, _ := http.NewRequest(http.MethodPost, path+"/clients/bulk", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var clients []*Client
	err = json.NewDecoder(resp.Body).Decode(&clients)
	if err != nil {
		return nil, nil, err
	}
	return clients[0], resp, nil
}

func RandomAdvertiser() *Advertiser {
	return &Advertiser{
		AdvertiserID: gofakeit.UUID(),
		Name:         gofakeit.Name(),
	}
}

func CreateRandomAdvertiser(path string) (*Advertiser, *http.Response, error) {
	advertiser := RandomAdvertiser()
	body, _ := json.Marshal([]*Advertiser{advertiser})
	req, _ := http.NewRequest(http.MethodPost, path+"/advertisers/bulk", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var advertisers []*Advertiser
	err = json.NewDecoder(resp.Body).Decode(&advertisers)
	if err != nil {
		return nil, nil, err
	}
	return advertisers[0], resp, nil
}

func RandomCampaign() *Campaign {
	clicksLimit := gofakeit.Number(0, 100)
	impressionsLimit := clicksLimit + gofakeit.Number(0, 50)
	costPerClick := gofakeit.Float32Range(0, 100)
	costPerImpression := gofakeit.Float32Range(0, 100)
	adTitle := gofakeit.Name()
	adText := gofakeit.Name()
	start := gofakeit.Number(0, 100)
	end := start + gofakeit.Number(0, 100)

	gender := gendersAll[gofakeit.Number(0, 2)]
	ageFrom := gofakeit.Number(0, 100)
	ageTo := ageFrom + gofakeit.Number(0, 100)
	location := gofakeit.City()
	return &Campaign{
		ImpressionsLimit:  &impressionsLimit,
		ClicksLimit:       &clicksLimit,
		CostPerClick:      &costPerClick,
		CostPerImpression: &costPerImpression,
		AdTitle:           &adTitle,
		AdText:            &adText,
		StartDate:         &start,
		EndDate:           &end,
		Target: &Target{
			Gender:   &gender,
			AgeFrom:  &ageFrom,
			AgeTo:    &ageTo,
			Location: &location,
		},
	}
}

func CreateRandomCampaign(path, advertiserID string) (*Campaign, *http.Response, error) {
	campaign := RandomCampaign()
	body, _ := json.Marshal(campaign)
	req, _ := http.NewRequest(http.MethodPost, path+"/advertisers/"+advertiserID+"/campaigns", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var campaigns *Campaign
	err = json.NewDecoder(resp.Body).Decode(&campaigns)
	if err != nil {
		return nil, nil, err
	}
	return campaigns, resp, nil
}

func GetNorm(min, max float64) int {
	val := min + rand.Float64()*(max-min)
	return int(math.Round(val))
}

func GetNormFloat(min, max float64) float32 {
	val := min + rand.Float64()*(max-min)
	return float32(math.Round(val*100) / 100)
}

func Round(val float64) float64 {
	return math.Round(val*100) / 100
}

func CreateCampaignNorm() *CampaignClear {
	start := gofakeit.Number(0, 100)
	end := start + gofakeit.Number(0, 20)
	return &CampaignClear{
		ImpressionsLimit:  GetNorm(30, 50),
		ClicksLimit:       GetNorm(15, 25),
		CostPerImpression: GetNormFloat(30, 50),
		CostPerClick:      GetNormFloat(60, 80),
		AdTitle:           gofakeit.Name(),
		AdText:            gofakeit.Name(),
		StartDate:         start,
		EndDate:           end,
	}
}
