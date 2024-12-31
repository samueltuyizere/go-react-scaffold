package integrations

import (
	"backend/configs"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendEmailWithPlunk(payload, recipient, title, fromEmail string) error {
	type bodyData struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
		Reply   string `json:"reply"`
	}

	url := "https://api.useplunk.com/v1/send"
	data := bodyData{
		To:      recipient,
		Subject: title,
		Body:    payload,
		Reply:   fromEmail,
	}
	var response any
	method := "POST"
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", configs.GetPlunkKey()))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
