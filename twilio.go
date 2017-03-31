package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/subosito/twilio"
)

func TwilioAlert(status TargetStatus, config Config) error {

	c := twilio.NewClient(config.Twilio.SID, config.Twilio.Token, nil)
	if config.Twilio.Simple != true {
		statusJson, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			return err
		}
		body := fmt.Sprintf("%s\n\n%s\n", time.Now(), statusJson)

		params := twilio.MessageParams{
			Body: body,
		}
		s, response, err := c.Messages.Send(config.Alert.FromNumber, config.Alert.ToNumber, params)
		if err != nil {
			fmt.Println(s, response, err)
		}
		return nil
	} else {
		var online string
		if status.Online != false {
			online = "Up"
		} else {
			online = "Down"
		}
		body := fmt.Sprintf("Host %s appears to be %s!\n", status.Target.Name, online)

		params := twilio.MessageParams{
			Body: body,
		}
		s, response, err := c.Messages.Send(config.Alert.FromNumber, config.Alert.ToNumber, params)
		if err != nil {
			fmt.Println(s, response, err)
		}
		return nil
	}
}
