package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	botToken := "1077898815:AAEDHyNTrzIeXF9exTFf31makWpGkVAYgcs"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	for {
		updates, err := getUpdates(botUrl)
		if err != nil {
			log.Println("Smth went wrong : ", err.Error())
		}
		fmt.Println(updates)
	}
}

func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond() {

}
