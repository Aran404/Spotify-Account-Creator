package hotmailbox

func NewClient(apiKey string) *Instance {
	return &Instance{
		ApiKey: apiKey,
	}
}
