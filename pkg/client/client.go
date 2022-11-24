package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (w *weatherlink) signature(values url.Values) url.Values {
	sign := hmac.New(sha256.New, []byte(w.ApiSecret))
	timestamp := time.Now().Unix()
	values.Set("t", strconv.FormatInt(timestamp, 10))
	cleanValues := strings.ReplaceAll(values.Encode(), "=", "")
	cleanValues = strings.ReplaceAll(cleanValues, "&", "")
	sign.Write([]byte(cleanValues))
	encodedSign := hex.EncodeToString(sign.Sum(nil))
	values.Set("api-signature", encodedSign)
	return values
}

func (w *weatherlink) url(path string, extraParams map[string]string) (string, error) {
	parsed, err := url.Parse(ApiV2Url)
	if err != nil {
		return "", err
	}
	joined := parsed.JoinPath(path)
	params := parsed.Query()
	params.Set("api-key", w.ApiKey)
	for key, value := range extraParams {
		params.Set(key, value)
	}
	signedUrl := w.signature(params)
	joined.RawQuery = signedUrl.Encode()
	return joined.String(), err
}
