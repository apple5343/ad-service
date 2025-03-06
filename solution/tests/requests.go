package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateCampaign(t *testing.T, path string, campaign *Campaign) *Campaign {
	t.Helper()
	reqBody, _ := json.Marshal(campaign)
	req, _ := http.NewRequest(http.MethodPost, path+"/advertisers/"+campaign.AdvertiserID+"/campaigns", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	var respCampaign *Campaign
	err = json.NewDecoder(resp.Body).Decode(&respCampaign)
	require.NoError(t, err)
	campaign.CampaignID = respCampaign.CampaignID
	return campaign
}

func CreateAdvertiser(t *testing.T, path string, advertiser *Advertiser) *Advertiser {
	t.Helper()
	reqBody, _ := json.Marshal([]*Advertiser{advertiser})
	req, _ := http.NewRequest(http.MethodPost, path+"/advertisers/bulk", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	var respAdvertiser []*Advertiser
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respAdvertiser))
	require.Equal(t, advertiser, respAdvertiser[0])
	return advertiser
}

func CreateClient(t *testing.T, path string, client *Client) *Client {
	t.Helper()
	reqBody, _ := json.Marshal([]*Client{client})
	req, _ := http.NewRequest(http.MethodPost, path+"/clients/bulk", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	var respClient []*Client
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respClient))
	require.Equal(t, client, respClient[0])
	return client
}

func SetDay(t *testing.T, path string, day int) {
	t.Helper()
	type reqBody struct {
		CurrentDay *int `json:"current_date,omitempty"`
	}
	reqB, _ := json.Marshal(reqBody{CurrentDay: &day})
	req, _ := http.NewRequest(http.MethodPost, path+"/time/advance", bytes.NewBuffer(reqB))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func GetDay(t *testing.T, path string) int {
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, path+"/time", nil)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var respDay struct {
		CurrentDay int `json:"current_date"`
	}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respDay))
	return respDay.CurrentDay
}

func Click(t *testing.T, path string, campaignID string, clientID string) int {

	t.Helper()
	type reqBody struct {
		ClientID string `json:"client_id"`
	}
	reqB, _ := json.Marshal(reqBody{ClientID: clientID})
	req, _ := http.NewRequest(http.MethodPost, path+"/ads/"+campaignID+"/click", bytes.NewBuffer(reqB))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	return resp.StatusCode
}

func Impression(t *testing.T, path string, clientID string) (*Ad, int) {
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, path+"/ads?client_id="+clientID, nil)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode
	}
	var respAd *Ad
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respAd))
	return respAd, resp.StatusCode
}

func CreateCampaignClear(t *testing.T, path string, campaign *CampaignClear) *CampaignClear {
	t.Helper()
	reqBody, _ := json.Marshal(campaign)
	req, _ := http.NewRequest(http.MethodPost, path+"/advertisers/"+campaign.AdvertiserID+"/campaigns", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	var respCampaign *Campaign
	err = json.NewDecoder(resp.Body).Decode(&respCampaign)
	require.NoError(t, err)
	campaign.CampaignID = respCampaign.CampaignID
	return campaign
}

func UpdateCampaign(t *testing.T, path string, campaign *Campaign) (*Campaign, int) {
	t.Helper()
	reqBody, _ := json.Marshal(campaign)
	req, _ := http.NewRequest(http.MethodPut, path+"/advertisers/"+campaign.AdvertiserID+"/campaigns/"+campaign.CampaignID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	var respCampaign *Campaign
	err = json.NewDecoder(resp.Body).Decode(&respCampaign)
	require.NoError(t, err)
	return respCampaign, resp.StatusCode
}

func UpdateCampaignClear(t *testing.T, path string, campaign *CampaignClear) (*CampaignClear, int) {
	t.Helper()
	reqBody, _ := json.Marshal(campaign)
	req, _ := http.NewRequest(http.MethodPut, path+"/advertisers/"+campaign.AdvertiserID+"/campaigns/"+campaign.CampaignID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	var respCampaign *CampaignClear
	err = json.NewDecoder(resp.Body).Decode(&respCampaign)
	require.NoError(t, err)
	return respCampaign, resp.StatusCode
}

func GetCampaign(t *testing.T, path string, advertiserID string, campaignID string) *Campaign {
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, path+"/advertisers/"+advertiserID+"/campaigns/"+campaignID, nil)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var respCampaign *Campaign
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respCampaign))
	return respCampaign
}
