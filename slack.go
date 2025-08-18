package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Types for fetching only required field from responses in `conversation.create`
type Channel struct {
	Id        string `json:"id"`
	Created   int    `json:"created"`
	Name      string `json:"name"`
	IsChannel bool   `json:"is_channel"`
}

type ChannelResp struct {
	IsOk        bool    `json:"ok"`
	ChannelInfo Channel `json:"channel"`
}

func FetchUserToken() string {
	token, ret := os.LookupEnv("SLACK_USER_TOKEN")
	if ret != true {
		fmt.Println("Failed to fetch Slack User token, please set SLACK_USER_TOKEN first.")
		// Note that os.Exit() will not call the logics of `defer`
		os.Exit(1)
	}
	return token
}

func CreateChannel(token string, name string) string {
	url := "https://slack.com/api/conversations.create"
	body := fmt.Appendf(nil, "name=%s", name)

	fmt.Println("Instantiate http.Request obuject ...")
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	fmt.Println("Instantiate http.Client obuject ...")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	channel := &ChannelResp{}
	if err := json.NewDecoder(res.Body).Decode(channel); err != nil {
		panic("Failed to decode response body to JSON")
	}
	if channel.IsOk != true {
		panic("Got error from Slack API during invoking `conversations.create`")
	}

	// fmt.Println(channel.IsOk)
	// fmt.Println(channel.ChannelInfo.Id)
	// fmt.Println(channel.ChannelInfo.Created)
	// fmt.Println(channel.ChannelInfo.Name)
	// fmt.Println(channel.ChannelInfo.IsChannel)
	return channel.ChannelInfo.Id
}

func AddRepositoryLinkToChannel(token string, id string, repourl string) {
	url := "https://slack.com/api/conversations.setTopic"
	body := fmt.Appendf(nil, "channel=%s&topic=%s", id, repourl)

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

}
