package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	http "github.com/bogdanfinn/fhttp"
)

func (c *Client) SaveAllCookies(handle *SavedCookies) error {
	cookies := c.GetCookieJar().GetAllCookies()

	handle.JsonPayload.Cookies = base64.StdEncoding.EncodeToString(func() []byte {
		b, _ := json.Marshal(cookies)

		return b
	}())

	handle.Mutex.Lock()
	defer handle.Mutex.Unlock()

	accounts, err := GetAllSessions(handle.JsonPath)

	if err != nil {
		return err
	}

	accounts.Accounts = append(accounts.Accounts, *handle.JsonPayload)

	marshalled, err := json.MarshalIndent(accounts, "", "    ")

	if err != nil {
		return err
	}

	err = DumpJson(marshalled, handle.JsonPath)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) LoadAllCookies(handle *SavedCookies) error {
	sessions, err := GetAllSessions(handle.JsonPath)

	if err != nil {
		return err
	}

	for _, s := range sessions.Accounts {
		valid := (s.Email == handle.JsonPayload.Email &&
			s.EmailPassword == handle.JsonPayload.EmailPassword &&
			s.SpotifyPassword == handle.JsonPayload.SpotifyPassword)

		if valid {
			if c.GetCookieJar() == nil {
				return errors.New("cookie jar is not initialized")
			}

			if s.Proxy != "" {
				c.SetProxy("http://" + s.Proxy)
			}

			decoded, err := base64.StdEncoding.DecodeString(s.Cookies)

			if err != nil {
				return err
			}

			var Cookies map[string][]*http.Cookie

			if err := json.Unmarshal(decoded, &Cookies); err != nil {
				return err
			}

			c.GetCookieJar().SetAllCookies(Cookies)

			return nil
		}
	}

	return ErrorNoSessionSaved
}
