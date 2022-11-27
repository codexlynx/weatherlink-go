package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"weatherlink-go/pkg/util"
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
	cleanValues := util.RemoveCharacters(values.Encode(), []string{"=", "&"})
	sign.Write([]byte(cleanValues))
	encodedSign := hex.EncodeToString(sign.Sum(nil))
	values.Set("api-signature", encodedSign)
	return values
}

func (w *weatherlink) get(path string, extra map[string]string, includeExtra bool) ([]byte, error) {
	parsed, err := url.Parse(ApiV2Url)
	if err != nil {
		return []byte{}, err
	}
	joined := parsed.JoinPath(path)
	params := parsed.Query()
	params.Set("api-key", w.ApiKey)
	for key, value := range extra {
		params.Set(key, value)
	}
	signedUrl := w.signature(params)
	if !includeExtra {
		keys := util.GetKeys(extra)
		signedUrl = util.RemoveParameters(keys, signedUrl)
	}
	joined.RawQuery = signedUrl.Encode()
	resp, err := w.client.Get(joined.String())
	if err != nil {
		return []byte{}, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}
