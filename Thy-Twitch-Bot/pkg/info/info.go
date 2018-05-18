package info

import (
	"github.com/Techterror12/Twitch-Stuff/Thy-Twitch-Bot/pkg/id-finders"
)

func Info() (string, string) {
	b := id.Callforusrinf()
	return b.BotChannel, b.OAuth
}
