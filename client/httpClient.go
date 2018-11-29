package client

import (
	"appbox-launcher/config"
	"fmt"
	"net/http"
	"time"
)

func Init() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return client
}

func CheckUpdateStatus(application *config.Application, client *http.Client) bool {
	serverUrl := application.Server.BaseUrl + "/status/123456"
	resp, err := client.Get(serverUrl)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(resp.StatusCode)
	if resp.StatusCode == 304 {
		return false
	}
	return true
}
