package generator

import (
	capsolver "Telegram/Core/Capsolver"
	client "Telegram/Core/Client"
	log "Telegram/Core/Log"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	ParseJsonString = func(b []byte, s string) string {
		return strings.Split(
			strings.Split(
				string(b), fmt.Sprintf(`%v":"`, s),
			)[1], `"`,
		)[0]
	}

	EmptyString = ""
)

func (in *Instance) SolveCaptcha(siteKey, proxy string) (*capsolver.CaptchaResponse, error) {
	task := map[string]interface{}{
		"type":       "ReCaptchaV3EnterpriseTask",
		"websiteURL": "https://www.spotify.com/tr-en/signup",
		"websiteKey": siteKey,
		"pageAction": "website/signup/submit_email",
		"minScore":   0.9,
		"proxy":      proxy,
	}

	solver := capsolver.NewSolver(in.CaptchaKey, in.CaptchaRetries)

	err := solver.CreateTask(task)

	if err != nil {
		return nil, err
	}

	result, err := solver.GetTaskResult()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (in *Instance) SolveVisibleCaptcha(siteKey, proxy, url string) (*capsolver.CaptchaResponse, error) {
	task := map[string]interface{}{
		"type":        "ReCaptchaV2EnterpriseTask",
		"websiteURL":  url,
		"websiteKey":  siteKey,
		"pageAction":  "challenge",
		"isInvisible": false,
		"minScore":    0.9,
		"userAgent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
		"proxy":       proxy,
	}

	solver := capsolver.NewSolver(in.CaptchaKey, in.CaptchaRetries)

	err := solver.CreateTask(task)

	if err != nil {
		return nil, err
	}

	result, err := solver.GetTaskResult()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (in *Instance) GetSignup() error {
	request := client.Request("GET", "https://www.spotify.com/tr-en/signup", nil, map[string]string{}, true, *in.Client)

	if request.Error != nil {
		return request.Error
	}

	if !request.Ok {
		return fmt.Errorf("could not get signup page, status code: %v", request.StatusCode)
	}

	in.ApiKey = ParseJsonString(request.Body, "signupServiceAppKey")
	in.InstallationId = ParseJsonString(request.Body, "spT")
	in.CSRFToken = ParseJsonString(request.Body, "csrfToken")
	in.FlowId = ParseJsonString(request.Body, "flowId")

	return nil
}

func (in *Instance) Register() error {
	in.CheckInstanceFields()

	payload := RegisterPayload{
		AccountDetails: struct {
			Birthdate    string "json:\"birthdate\""
			ConsentFlags struct {
				EulaAgreed      bool "json:\"eula_agreed\""
				SendEmail       bool "json:\"send_email\""
				ThirdPartyEmail bool "json:\"third_party_email\""
			} "json:\"consent_flags\""
			DisplayName                string "json:\"display_name\""
			EmailAndPasswordIdentifier struct {
				Email    string "json:\"email\""
				Password string "json:\"password\""
			} "json:\"email_and_password_identifier\""
			Gender int "json:\"gender\""
		}{
			Birthdate: in.BirthDate,
			ConsentFlags: struct {
				EulaAgreed      bool "json:\"eula_agreed\""
				SendEmail       bool "json:\"send_email\""
				ThirdPartyEmail bool "json:\"third_party_email\""
			}{
				EulaAgreed:      true,
				SendEmail:       true,
				ThirdPartyEmail: false,
			},
			DisplayName: in.DisplayName,
			EmailAndPasswordIdentifier: struct {
				Email    string "json:\"email\""
				Password string "json:\"password\""
			}{
				Email:    in.Email,
				Password: in.Password,
			},
			Gender: 1,
		},
		CallbackURI: "https://www.spotify.com/signup/challenge?locale=tr-en",
		ClientInfo: struct {
			APIKey         string "json:\"api_key\""
			AppVersion     string "json:\"app_version\""
			Capabilities   []int  "json:\"capabilities\""
			InstallationID string "json:\"installation_id\""
			Platform       string "json:\"platform\""
		}{
			APIKey:         in.ApiKey,
			AppVersion:     "v2",
			Capabilities:   []int{1},
			InstallationID: in.InstallationId,
			Platform:       "www",
		},
		Tracking: struct {
			CreationFlow  string "json:\"creation_flow\""
			CreationPoint string "json:\"creation_point\""
			Referrer      string "json:\"referrer\""
		}{
			CreationFlow:  "",
			CreationPoint: "spotify.com",
			Referrer:      "",
		},
		RecaptchaToken: "",
		SubmissionID:   in.SubmissionId,
		FlowID:         in.FlowId,
	}

	token, err := in.SolveCaptcha("6LfCVLAUAAAAALFwwRnnCJ12DalriUGbj8FW_J39", in.Proxy)

	if err != nil {
		return err
	}

	payload.RecaptchaToken = token.CaptchaKey

	encoded, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	request := client.Request("POST", "https://spclient.wg.spotify.com/signup/public/v2/account/create", bytes.NewBuffer(encoded), client.NoHeaders, true, *in.Client)

	if request.Error != nil {
		return request.Error
	}

	if !request.Ok {
		return fmt.Errorf("could not register, status code: %d", request.StatusCode)
	}

	if strings.Contains(string(request.Body), "challenge") {
		if !in.HasRetried && in.Config.Generator.RetryFails {
			in.HasRetried = true

			return in.Register()
		}

		log.LogInfo(log.GetStackTrace(), "Encountered Challenge, Attempting to Solve.")

		return in.HandleChallenge(request.Body)
	}

	in.Splot = ParseJsonString(request.Body, "login_token")

	return nil
}

func (in *Instance) GenerateAccount() error {
	err := in.GetSignup()

	if err != nil {
		return err
	}

	err = in.Register()

	if err != nil {
		return err
	}

	return nil
}
