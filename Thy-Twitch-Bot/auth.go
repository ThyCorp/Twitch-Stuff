package main

import (
	"fmt"
	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/twitch"
	"strings"
  "github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/open"
)

var replacerone = strings.NewReplacer("%3A", ",", ":", ",")

func main() {
	out := twitch.New("avq1j7x3f8s9dzesyq67s7nfa8hccm", "a1rrg32a45m8nsx3pfukudymcvgf0x", "http://localhost", "viewing_activity_read")
  open.Open("https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=avq1j7x3f8s9dzesyq67s7nfa8hccm&redirect_uri=http://localhost&scope=viewing_activity_read&state=a1rrg32a45m8nsx3pfukudymcvgf0x")
  out2, err := out.BeginAuth("a1rrg32a45m8nsx3pfukudymcvgf0x")
	if err != nil {
		fmt.Println("oh no")
	}
	open.Open(out2)
}
