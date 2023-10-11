package generator

import (
	client "Telegram/Core/Client"
	types "Telegram/Core/Types"
)

type Instance struct {
	*client.Client
	Config         *types.Config
	Splot          string
	CSRFToken      string
	Proxy          string
	CaptchaKey     string
	ApiKey         string
	InstallationId string
	FlowId         string
	SubmissionId   string
	BirthDate      string
	DisplayName    string
	Email          string
	Password       string
	IsRealEmail    bool
	CaptchaRetries int
	HasRetried     bool
}

type RegisterPayload struct {
	AccountDetails struct {
		Birthdate    string `json:"birthdate"`
		ConsentFlags struct {
			EulaAgreed      bool `json:"eula_agreed"`
			SendEmail       bool `json:"send_email"`
			ThirdPartyEmail bool `json:"third_party_email"`
		} `json:"consent_flags"`
		DisplayName                string `json:"display_name"`
		EmailAndPasswordIdentifier struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"email_and_password_identifier"`
		Gender int `json:"gender"`
	} `json:"account_details"`
	CallbackURI string `json:"callback_uri"`
	ClientInfo  struct {
		APIKey         string `json:"api_key"`
		AppVersion     string `json:"app_version"`
		Capabilities   []int  `json:"capabilities"`
		InstallationID string `json:"installation_id"`
		Platform       string `json:"platform"`
	} `json:"client_info"`
	Tracking struct {
		CreationFlow  string `json:"creation_flow"`
		CreationPoint string `json:"creation_point"`
		Referrer      string `json:"referrer"`
	} `json:"tracking"`
	RecaptchaToken string `json:"recaptcha_token"`
	SubmissionID   string `json:"submission_id"`
	FlowID         string `json:"flow_id"`
}

type ChallengePayload struct {
	SessionID            string `json:"session_id"`
	ChallengeID          string `json:"challenge_id"`
	RecaptchaChallengeV1 struct {
		Solve struct {
			RecaptchaToken string `json:"recaptcha_token"`
		} `json:"solve"`
	} `json:"recaptcha_challenge_v1"`
}
