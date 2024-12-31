package integrations

import (
	"backend/configs"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/samueltuyizere/validate_rw_phone_numbers"
)

type PaypackWebhook struct {
	EventID   string      `json:"event_id"`
	EventKind string      `json:"event_kind"`
	CreatedAt time.Time   `json:"created_at"`
	Data      WebhookData `json:"data"`
}

type WebhookData struct {
	Ref         string    `json:"ref"`
	Kind        string    `json:"kind"`
	Fee         float64   `json:"fee"`
	Merchant    string    `json:"merchant"`
	Client      string    `json:"client"`
	Amount      int       `json:"amount"`
	Provider    string    `json:"provider"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	ProcessedAt time.Time `json:"processed_at"`
}

type CashoutResponseParams struct {
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Kind      string    `json:"kind"`
	Ref       string    `json:"ref"`
	Status    string    `json:"status"`
}

type PollResponseParams struct {
	Amount    int       `json:"amount"`
	Client    string    `json:"client"`
	Fee       float64   `json:"fee"`
	Kind      string    `json:"kind"`
	Merchant  string    `json:"merchant"`
	Ref       string    `json:"ref"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type authResp struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	Expires int    `json:"expires"`
}

type cashInRequest struct {
	Amount int    `json:"amount"`
	Number string `json:"number"`
}

type CashinResponseParams struct {
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
	Kind      string `json:"kind"`
	Ref       string `json:"ref"`
	Status    string `json:"status"`
}

func Authenticate() (string, error) {
	url := "https://payments.paypack.rw/api/auth/agents/authorize"
	// authentication logic
	method := "POST"
	var response *authResp
	data := map[string]string{"client_id": configs.GetPaypackId(), "client_secret": configs.GetPaypackSecret()}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return "", err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return response.Access, nil
}

func PaypackCashIn(amount int, number string) (CashinResponseParams, error) {
	var response *CashinResponseParams
	url := "https://payments.paypack.rw/api/transactions/cashin"
	method := "POST"
	token, _ := Authenticate()
	// ensuring the phone number is in the local format
	number = validate_rw_phone_numbers.GetLocalFormat(number)
	data := cashInRequest{
		Amount: amount,
		Number: number,
	}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Webhook-Mode", "production")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}

	return *response, nil
}

func PaypackCashOut(amount int, number string) (CashoutResponseParams, error) {
	var response *CashoutResponseParams
	url := "https://payments.paypack.rw/api/transactions/cashout"
	method := "POST"
	token, _ := Authenticate()

	data := cashInRequest{
		Amount: amount,
		Number: number,
	}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Webhook-Mode", "production")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}

	return *response, nil
}

func PollTransactionStatus(ref string) (PollResponseParams, error) {
	var response *PollResponseParams
	url := fmt.Sprintf("https://payments.paypack.rw/api/transactions/find/%s", ref)
	method := "GET"
	token, _ := Authenticate()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	defer res.Body.Close()
	fmt.Printf("\nresponse: %+v\n", res.Body)
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	fmt.Printf("\nPolling transaction status: %v\n", *response)
	return *response, nil
}

func TestPollTransactionStatus(ref string) (interface{}, error) {
	var response *interface{}
	url := fmt.Sprintf("https://payments.paypack.rw/api/transactions/find/%s", ref)
	method := "GET"
	token, _ := Authenticate()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	defer res.Body.Close()
	fmt.Printf("\nresponse: %+v\n", &res.Body)
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	fmt.Printf("\nPolling transaction status: %v\n", *response)
	return *response, nil
}
