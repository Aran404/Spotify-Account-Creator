package generator

import (
	client "Telegram/Core/Client"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func (in *Instance) HandleChallenge(response []byte) error {
	challengeUrl, err := in.GetSession(ParseJsonString(response, "session_id"))

	if err != nil {
		return err
	}

	token, err := in.SolveVisibleCaptcha("6LeO36obAAAAALSBZrY6RYM1hcAY7RLvpDDcJLy3", in.Proxy, challengeUrl)

	if err != nil {
		return err
	}

	err = in.SubmitChallengeCaptcha(token.CaptchaKey, challengeUrl)

	if err != nil {
		return err
	}

	return in.CompleteCreation(ParseJsonString(response, "session_id"))
}

func (in *Instance) GetSession(sessionID string) (string, error) {
	url := `https://challenge.spotify.com/api/v1/get-session`
	payload := strings.NewReader(`{"session_id":"` + sessionID + `"}`)

	request := client.Request("POST", url, payload, client.NoHeaders, true, *in.Client)

	if request.Error != nil {
		return "", request.Error
	}

	if !request.Ok {
		return "", fmt.Errorf("could not get session link, Status Code: %v", request.StatusCode)
	}

	return ParseJsonString(request.Body, "url"), nil
}

func (in *Instance) SubmitChallengeCaptcha(token, url string) error {
	sessionID := strings.Split(strings.Split(url, "c/")[1], "/")[0]
	challengeID := strings.Split(strings.Split(url, sessionID+"/")[1], "/")[0]

	uri := `https://challenge.spotify.com/api/v1/invoke-challenge-command`

	payload := ChallengePayload{
		SessionID:   sessionID,
		ChallengeID: challengeID,
		RecaptchaChallengeV1: struct {
			Solve struct {
				RecaptchaToken string "json:\"recaptcha_token\""
			} "json:\"solve\""
		}{
			Solve: struct {
				RecaptchaToken string "json:\"recaptcha_token\""
			}{
				RecaptchaToken: token,
			},
		},
	}

	encoded, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	headers := map[string]string{
		"X-Cloud-Trace-Context": "000000000000000004ec7cfe60aa92b5/8088460714428896449;o=1",
	}

	request := client.Request("POST", uri, bytes.NewBuffer(encoded), headers, true, *in.Client)

	if request.Error != nil {
		return request.Error
	}

	if !request.Ok {
		return fmt.Errorf("could not invoke challenge command, Status Code: %v, Body: %v", request.StatusCode, string(request.Body))
	}

	return nil
}

func (in *Instance) CompleteCreation(sessionID string) error {
	url := `https://spclient.wg.spotify.com/signup/public/v2/account/complete-creation`
	payload := strings.NewReader(fmt.Sprintf(`{"session_id":"%v"}`, sessionID))

	request := client.Request("POST", url, payload, client.NoHeaders, true, *in.Client)

	if request.Error != nil {
		return request.Error
	}

	if !request.Ok {
		return fmt.Errorf("could not complete creation process, Status Code: %v, Body: %v", request.StatusCode, string(request.Body))
	}

	if !strings.Contains(string(request.Body), "success") {
		return fmt.Errorf("could not complete the creation process, Status Code: %v, Body: %v", request.StatusCode, string(request.Body))
	}

	return nil
}
