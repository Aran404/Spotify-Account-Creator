package hotmailbox

type Instance struct {
	ApiKey    string
	Accounts  []string
	ResetLink string
}

type GetEmailResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    struct {
		Emails []struct {
			Email    string `json:"Email"`
			Password string `json:"Password"`
		} `json:"Emails"`
	} `json:"Data"`
}
