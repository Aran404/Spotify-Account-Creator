package hotmailbox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (in *Instance) GetNewEmails(amount int, outlook bool) ([]string, error) {
	var Domain string

	if amount <= 9 {
		return []string{}, fmt.Errorf("amount needs to be >= 10")
	}

	if outlook {
		Domain = "OUTLOOK"
	} else {
		Domain = "HOTMAIL"
	}

	var Accs []string
	resp, err := http.Get("https://api.hotmailbox.me/mail/buy?apikey=" + in.ApiKey + "&mailcode=" + Domain + ".TRUSTED&quantity=" + fmt.Sprintf("%v", amount))

	if err != nil {
		return []string{}, err
	}

	var Response *GetEmailResponse

	if err := json.NewDecoder(resp.Body).Decode(&Response); err != nil {
		return []string{}, err
	}

	if Response.Code != 0 {
		return []string{}, fmt.Errorf("could not get new email, Error: %v", Response.Message)
	}

	for _, v := range Response.Data.Emails {
		Accs = append(Accs, v.Email+":"+v.Password)
	}

	in.Accounts = Accs

	return Accs, nil
}
