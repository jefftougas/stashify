package notifier

type SlackNotifier struct {
	Payload    string `json:"payload"`
	WebHookUrl string `json:"-"`
}
