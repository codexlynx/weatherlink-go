package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ApiV2Url https://weatherlink.github.io/v2-api/api-reference
const ApiV2Url = "https://api.weatherlink.com/v2/"

type weatherlink struct {
	ApiKey    string
	ApiSecret string
	client    *http.Client
}

func New(client *http.Client, apiKey, apiSecret string) *weatherlink {
	if client == nil {
		client = http.DefaultClient
	}
	return &weatherlink{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		client:    client,
	}
}

func (w *weatherlink) signature() (int64, string) {
	timestamp := time.Now().Unix()
	sign := hmac.New(sha256.New, []byte(w.ApiSecret))
	sign.Write([]byte(fmt.Sprintf("api-key%st%d", w.ApiKey, timestamp)))
	return timestamp, hex.EncodeToString(sign.Sum(nil))
}

func (w *weatherlink) url(path string) (string, error) {
	parsed, err := url.Parse(ApiV2Url)
	if err != nil {
		return "", err
	}
	timestamp, sign := w.signature()
	params := parsed.Query()
	params.Set("api-key", w.ApiKey)
	params.Set("api-signature", sign)
	params.Set("t", strconv.FormatInt(timestamp, 10))
	parsed.RawQuery = params.Encode()
	fullUrl := parsed.JoinPath(path)
	return fullUrl.String(), err
}
