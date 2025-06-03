package wago

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/DaviNAzevedo/wago/api/client"
)

type Client struct {
	token  string
	wabaID string
}

func NewClient(token, wabaID string) *Client {
	return &Client{token: token, wabaID: wabaID}
}

func (c *Client) SendMessage(payload []byte) error {
	messageReader := bytes.NewReader(payload)
	messageRequest, err := http.NewRequest("POST", fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", c.wabaID), messageReader)
	if err != nil {
		return err
	}

	messageRequest.Header.Set("Authorization", c.token)
	messageRequest.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(messageRequest)
	if err != nil {
		return err
	}

	i, _ := httputil.DumpRequestOut(res.Request, true)
	log.Println(string(i))

	if res.StatusCode < 200 || res.StatusCode > 299 {
		response, _ := io.ReadAll(res.Body)
		log.Println(string(response))
		return errors.New("Failed to send message")
	}
	return nil
}

func (c *Client) SendText(to, content, originalMessageID string, previewURL bool) error {
	messagePayload := client.SendMessageRequest{
		MessagingProduct: "whatsapp",
		PreviewURL:       &previewURL,
		To:               to,
		Type:             "text",
		Text: &client.TextPayload{
			PreviewURL: previewURL,
			Body:       content,
		},
		Context: &client.MessageContext{
			MessageID: originalMessageID,
		},
	}
	messagePayloadJson, err := json.Marshal(messagePayload)
	if err != nil {
		return err
	}

	return c.SendMessage(messagePayloadJson)
}
