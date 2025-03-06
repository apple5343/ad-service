package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Campaign struct {
	ImpressionsLimit  int    `json:"impressions_limit"`
	ClicksLimit       int    `json:"clicks_limit"`
	CostPerClick      int    `json:"cost_per_click"`
	CostPerImpression int    `json:"cost_per_impression"`
	AdTitle           string `json:"ad_title"`
	AdText            string `json:"ad_text"`
	StartDate         int    `json:"start_date"`
	EndDate           int    `json:"end_date"`
	Target            Target `json:"targeting"`
}

type Target struct {
	Gender   string `json:"gender"`
	AgeFrom  int    `json:"age_from"`
	AgeTo    int    `json:"age_to"`
	Location string `json:"location,omitempty"`
}

var (
	genders = []string{"MALE", "FEMALE", "ALL"}
)

func randomCampaign() Campaign {
	from := gofakeit.Number(1, 100)
	return Campaign{
		ImpressionsLimit:  gofakeit.Number(1, 100),
		ClicksLimit:       gofakeit.Number(1, 100),
		CostPerClick:      gofakeit.Number(1, 100),
		CostPerImpression: gofakeit.Number(1, 100),
		AdTitle:           gofakeit.Word(),
		AdText:            gofakeit.Word(),
		StartDate:         0,
		EndDate:           100,
		Target: Target{
			Gender:  genders[gofakeit.Number(0, 2)],
			AgeFrom: from,
			AgeTo:   gofakeit.Number(from, 100),
		},
	}
}

func newRequests(url string) vegeta.Target {
	campaign := randomCampaign()
	body, err := json.Marshal(campaign)
	if err != nil {
		panic(err)
	}
	return vegeta.Target{
		Method: "POST",
		URL:    url,
		Body:   body,
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
}

func main() {
	advertiserID := "3fa85f64-5717-4562-b3fc-2c963f66afa0"
	url := "http://localhost:8080/advertisers/" + advertiserID + "/campaigns"
	rate := vegeta.Rate{Freq: 1500, Per: 1 * time.Second}
	duration := 10 * time.Second
	targets := make([]vegeta.Target, 100)
	for i := 0; i < 100; i++ {
		targets[i] = newRequests(url)
	}
	target := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for r := range attacker.Attack(target, rate, duration, "Campaign create load test") {
		metrics.Add(r)
	}
	metrics.Close()
	fmt.Println("Requests:", metrics.Requests)
	fmt.Println("Success:", metrics.Success*100)
	fmt.Println("Errors:", metrics.Errors)
	fmt.Println("StatusCodes:", metrics.StatusCodes)
	fmt.Println("50 percentile:", metrics.Latencies.P50)
	fmt.Println("90 percentile:", metrics.Latencies.P90)
	fmt.Println("99 percentile:", metrics.Latencies.P99)
}
