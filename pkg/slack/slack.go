package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

const (
	slackAPIURL = string("https://api.slack.com/")
	rtmstart    = string("https://slack.com/api/rtm.start")
	messageType = string("message")
)

// Standard slack message format
type messageRx struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      string `json:"ts"`
	// TODO does not include edits or subtypes
}

// Standard slack message format
type messageTx struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// RTMClient is a real time messaging client for slack
type RTMClient struct {
	ws           *websocket.Conn
	id           string
	msg          messageRx
	send         messageTx
	sendSequence uint64
}

// NewRTMClient returns a new real time messaging client for a given token
func NewRTMClient(token string) (*RTMClient, error) {

	// Request an rtm session
	socketURL, id, err := rtmStart(token)
	if err != nil {
		return nil, err
	}

	// Connect to the rtm socket
	ws, err := websocket.Dial(socketURL, "", slackAPIURL)
	if err != nil {
		return nil, err
	}

	return &RTMClient{ws, id, messageRx{}, messageTx{}, 0}, nil
}

func rtmStart(tok string) (socket string, id string, err error) {

	url := fmt.Sprintf("%v?token=%v&no_unreads=true", rtmstart, tok)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("Request Failed: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	// Json format
	type self struct {
		ID string `json:"id"`
	}
	type rtmStartResp struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
		URL   string `json:"url"`
		Self  self   `json:"self"`
	}

	// Parse the response
	var data rtmStartResp
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", "", err
	}

	if !data.Ok {
		return "", "", fmt.Errorf(data.Error)
	}

	return data.URL, data.Self.ID, nil
}

// Receive receives data until a message for the requested id is obtained
func (s *RTMClient) Receive() (string, error) {

	checkMsg := func() error {
		return websocket.JSON.Receive(s.ws, &s.msg)
	}

	var err error
	directMsg := fmt.Sprintf("<@%v>: ", s.id)
	for err = checkMsg(); err != nil || !strings.HasPrefix(s.msg.Text, directMsg) || s.msg.Type != messageType; {
		err = checkMsg()
	}

	text := s.msg.Text[len(s.id)+5:]
	return text, nil
}

// Send a response to the same channel as previous received message
func (s *RTMClient) Send(m string) error {

	// Setup the send message
	s.send.Channel = s.msg.Channel
	s.sendSequence++
	s.send.ID = s.sendSequence
	s.send.Text = m
	s.send.Type = messageType

	return websocket.JSON.Send(s.ws, &s.send)
}
