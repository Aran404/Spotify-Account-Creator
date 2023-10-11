package client

import (
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func NewOptions() ClientOptions {
	options := []tls_client.HttpClientOption{
		tls_client.WithClientProfile(profiles.Chrome_112),
		tls_client.WithRandomTLSExtensionOrder(),
		tls_client.WithInsecureSkipVerify(),
	}

	settings := new(Options)
	settings.Settings = options

	return settings
}

func (i *Options) SetTimeout(n int) {
	i.Settings = append(i.Settings, tls_client.WithTimeout(n))
}

func (i *Options) SetNewCookieJar() {
	jar := tls_client.NewHttpCookiejar()

	i.Settings = append(i.Settings, tls_client.WithCookieJar(jar))
}

func (i *Options) SetNotFollowRedirects() {
	i.Settings = append(i.Settings, tls_client.WithNotFollowRedirects())
}

func (i *Options) SetProxy(proxy string) {
	i.Settings = append(i.Settings, tls_client.WithProxyUrl(proxy))
}

func (i *Options) NewClient() (*tls_client.HttpClient, error) {
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), i.Settings...)

	if err != nil {
		return nil, err
	}

	return &client, nil
}
