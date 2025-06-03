package wago

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/DaviNAzevedo/wago/api/webhook"
)

type Webhook struct {
	VerifyToken string
	AppSecret   string
	OnMessage   func(message webhook.Message)
}

func NewWebhook(verifyToken, appSecret string) *Webhook {
	return &Webhook{
		VerifyToken: verifyToken,
		AppSecret:   appSecret,
	}
}

func (wh *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		challenge := wh.validateVerificationToken(r.URL.Query())
		if challenge == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
		return
	}

	var webhookPayload webhook.WebhookPayload
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(requestBody, &webhookPayload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if status := wh.validateSignature(requestBody, r.Header); status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}

	for _, entry := range webhookPayload.Entry {
		for _, change := range entry.Changes {
			switch change.Field {
			case "messages":
				for _, message := range change.Value.Messages {
					wh.OnMessage(message)
				}
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (wh *Webhook) validateSignature(body []byte, h http.Header) int {
	signatureHeader := h.Get("X-Hub-Signature-256")
	if signatureHeader == "" {
		return http.StatusUnauthorized
	}

	parts := strings.SplitN(signatureHeader, "=", 2)
	if strings.ToLower(parts[0]) != "sha256" {
		return http.StatusUnauthorized
	}

	mac := hmac.New(sha256.New, []byte(wh.AppSecret))
	mac.Write(body)
	calculatedSignature := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(calculatedSignature), []byte(parts[1])) {
		log.Println("InvalidSignature")
		log.Println(parts[1])
		log.Println(calculatedSignature)
		return http.StatusUnauthorized
	}

	return http.StatusOK
}

func (wh *Webhook) validateVerificationToken(q url.Values) string {
	mode := q.Get("hub.mode")
	token := q.Get("hub.verify_token")
	log.Println(q.Encode())
	log.Println(token)
	log.Println(wh.VerifyToken)
	if mode == "subscribe" && token == wh.VerifyToken {
		log.Println("TokenVerified")
		return q.Get("hub.challenge")
	}
	log.Println("TokenNotVerified")
	return ""
}
