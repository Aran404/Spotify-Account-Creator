package types

type Session struct {
	Accounts []JsonTags
}

type JsonTags struct {
	Email           string `json:"email"`
	SpotifyPassword string `json:"spotify_password"`
	Proxy           string `json:"proxy,omitempty"`
	EmailPassword   string `json:"email_password,omitempty"`
	UserID          string `json:"user_id,omitempty"`
	Cookies         string `json:"cookies,omitempty"`
}

type Config struct {
	Generator struct {
		CustomPassword    string `json:"custom_password"`
		PasswordSuffix    string `json:"password_suffix"`
		DisplayName       string `json:"display_name"`
		DisplayNameSuffix string `json:"display_name_suffix"`
		UseHotmailbox     bool   `json:"use_hotmailbox"`
		RetryFails        bool   `json:"retry_fails"`
	} `json:"generator"`
	Threads struct {
		MaxThreads       int `json:"max_threads"`
		AmountOfAccounts int `json:"amount_of_accounts"`
	} `json:"threads"`
	Captcha struct {
		MaxRetries int    `json:"max_retries"`
		APIKey     string `json:"api_key"`
	} `json:"captcha"`
	Hotmailbox struct {
		APIKey string `json:"api_key"`
		Domain string `json:"domain"`
	} `json:"mailbox"`
}
