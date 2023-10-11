package client

import (
	types "Telegram/Core/Types"
	"errors"
	"sync"

	http "github.com/bogdanfinn/fhttp"

	tls_client "github.com/bogdanfinn/tls-client"
)

type Options struct {
	Settings []tls_client.HttpClientOption
}

type ClientOptions interface {
	SetTimeout(int)
	SetNewCookieJar()
	SetNotFollowRedirects()
	SetProxy(string)
	NewClient() (*tls_client.HttpClient, error)
}

type Client struct {
	tls_client.HttpClient
}

type SavedCookies struct {
	JsonPath    string
	Mutex       *sync.Mutex
	JsonPayload *types.JsonTags
}

type RequestResponse struct {
	Error                error
	StatusCode           int
	Ok                   bool
	StatusCodeDefinition string
	Body                 []byte
	Json                 map[string]interface{}
	Request              *http.Response
}

var ErrorNoSessionSaved = errors.New("no saved sessions in the database")
