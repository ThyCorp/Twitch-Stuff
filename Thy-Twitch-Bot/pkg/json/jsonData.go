package jsonD

type UserInfo struct {
	ClientID      string `json:"clientId"`
	BotChannel    string `json:"botChannel"`
	StreamChannel string `json:"streamChannel"`
	OAuth         string `json:"oauth"`
}
