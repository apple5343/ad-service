package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Client struct {
	ClientID string `json:"client_id"`
	Login    string `json:"login"`
	Age      int    `json:"age"`
	Location string `json:"location"`
	Gender   string `json:"gender"`
}

var (
	genders = []string{"MALE", "FEMALE"}
)

func randomClient() Client {
	return Client{
		ClientID: gofakeit.UUID(),
		Login:    gofakeit.Username(),
		Age:      gofakeit.Number(0, 100),
		Location: gofakeit.City(),
		Gender:   genders[gofakeit.Number(0, 1)],
	}
}

func newRequests(url string) vegeta.Target {
	l := gofakeit.Number(1, 25)
	clients := make([]Client, l)
	for i := 0; i < l; i++ {
		clients[i] = randomClient()
	}
	body, err := json.Marshal(clients)
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
	url := "http://localhost:8080/clients/bulk"
	rate := vegeta.Rate{Freq: 5000, Per: 1 * time.Second}
	duration := 5 * time.Second

	target := vegeta.NewStaticTargeter(newRequests(url))
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for r := range attacker.Attack(target, rate, duration, "Clients load test") {
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
