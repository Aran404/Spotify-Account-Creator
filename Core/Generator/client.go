package generator

import (
	client "Telegram/Core/Client"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func NewGeneratorInstance(captchaKey, proxy string, captchaRetries int) (*Instance, error) {
	c := client.NewOptions()
	c.SetNewCookieJar()
	c.SetTimeout(60)
	c.SetProxy("http://" + proxy)

	httpClient, err := c.NewClient()

	if err != nil {
		return nil, err
	}

	if strings.Contains(proxy, "@") {
		subs := strings.Split(proxy, "@")
		proxy = fmt.Sprintf("%v:%v", subs[1], subs[0])
	}

	client := &client.Client{HttpClient: *httpClient}

	return &Instance{
		Client:         client,
		CaptchaKey:     captchaKey,
		CaptchaRetries: captchaRetries,
		Proxy:          proxy,
		BirthDate:      GetDob(),
		SubmissionId:   uuid.NewString(),
	}, nil
}

func Suffix(flag string) string {
	for strings.Contains(flag, "%d") {
		flag = strings.Replace(flag, "%d", RandomSuffix(1), 1)
	}

	for strings.Contains(flag, "%s") {
		flag = strings.Replace(flag, "%s", RandomStringSuffix(1), 1)
	}

	return flag
}

func (in *Instance) SetDisplayName(s ...string) {
	var displayName string

	if len(s) >= 1 {
		if len(s[0]) <= 0 {
			in.DisplayName = RandomDisplayName(8)
		}
	}

	switch len(s) {
	case 0:
		displayName = RandomDisplayName(8)
	case 1:
		displayName = s[0]
	case 2:
		flag := s[1]

		suffix := Suffix(flag)

		displayName = s[0] + suffix
	}

	in.DisplayName = displayName
}

func (in *Instance) SetEmail(s ...string) {
	var email string

	if len(s) >= 1 {
		if len(s[0]) <= 0 {
			in.Email = RandomEmail()
		}
	}

	switch len(s) {
	case 0:
		email = RandomEmail()
	case 1:
		email = s[0]
	case 2, 3:
		flag := s[1]

		suffix := Suffix(flag)

		domain := func() string {
			if len(s) == 2 {
				return RandomDomain()
			}

			return s[3]
		}()

		email = s[0] + suffix + domain
	}

	in.Email = email
}

func (in *Instance) SetPassword(s ...string) {
	var password string

	if len(s) >= 1 {
		if len(s[0]) <= 0 {
			in.Password = RandomPassword(8)
		}
	}

	switch len(s) {
	case 0:
		password = RandomPassword(8)
	case 1:
		password = s[0]
	case 2:
		flag := s[1]

		suffix := Suffix(flag)

		password = s[0] + suffix
	}

	in.Password = password + "."
}
